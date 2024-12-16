package usecase

import (
	"context"
	"path/filepath"
	"strings"

	appErrorAuth "server/internal/auth/apperror"
	dtoAuth "server/internal/auth/dto"
	authRepo "server/internal/auth/repository"
	utilsAuth "server/internal/auth/utils"
	appErrorProfile "server/internal/profile/apperror"
	"server/internal/profile/constant"
	dtoProfile "server/internal/profile/dto"
	profileRepo "server/internal/profile/repository"
	"server/internal/profile/utils"
	appErrorPkg "server/pkg/apperror"
	"server/pkg/database/transactor"
	"server/pkg/utils/cloudinaryutils"

	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type ProfileUseCase interface {
	GetProfile(ctx context.Context, userId int64, serve string) (*dtoProfile.ResponseProfile, error)
	PutMyProfile(ctx context.Context, reqBody *dtoProfile.RequestPutMyProfile, userId int64, roleId int64) (*dtoProfile.ResponseProfile, error)
}

type profileUseCaseImpl struct {
	profileRepo    profileRepo.ProfileRepository
	addressRepo    profileRepo.AddressRepository
	userRepo       authRepo.UserRepository
	transactor     transactor.Transactor
	cloudinaryUtil cloudinaryutils.CloudinaryUtil
}

func NewProfileUseCase(
	profileRepo profileRepo.ProfileRepository,
	addressRepo profileRepo.AddressRepository,
	userRepo authRepo.UserRepository,
	transactor transactor.Transactor,
	cloudinaryUtil cloudinaryutils.CloudinaryUtil,
) *profileUseCaseImpl {
	return &profileUseCaseImpl{
		profileRepo:    profileRepo,
		addressRepo:    addressRepo,
		userRepo:       userRepo,
		transactor:     transactor,
		cloudinaryUtil: cloudinaryUtil,
	}
}

func (pu *profileUseCaseImpl) GetProfile(ctx context.Context, userId int64, serve string) (*dtoProfile.ResponseProfile, error) {
	var (
		addresses  []*dtoProfile.ResponseAddress
		user       dtoAuth.ResponseUser
		userDetail *dtoAuth.ResponseUserDetail
	)
	err := pu.transactor.Atomic(ctx, func(cForTx context.Context) error {
		userDb, err := pu.userRepo.FindByIDWithCompleteData(cForTx, userId)
		if err != nil || userDb.ID == 0 {
			if serve == "all" {
				return appErrorPkg.NewForbiddenAccessError()
			}
			return appErrorProfile.NewInvalidIdUserProfileNotExistsError()
		}
		user = dtoAuth.ResponseUser{
			ID:         userDb.ID,
			RoleId:     userDb.Role,
			Role:       utilsAuth.SpecifyRole(userDb.Role),
			Email:      userDb.Email,
			IsVerified: userDb.IsVerified,
			IsOauth:    userDb.IsOauth,
			CreatedAt:  userDb.CreatedAt,
			UpdatedAt:  userDb.UpdatedAt,
			DeletedAt:  userDb.DeletedAt,
		}
		userDetailDb, err := pu.userRepo.GetUserDetailByUserID(cForTx, userId)
		if err != nil {
			return appErrorPkg.NewServerError(err)
		}
		if userDetailDb != nil {
			userDetail = &dtoAuth.ResponseUserDetail{
				Id:             *userDetailDb.ID,
				UserId:         *userDetailDb.UserId,
				Fullname:       *userDetailDb.Fullname,
				WhatsappNumber: userDetailDb.WhatsappNumber,
				ImageUrl:       *userDetailDb.ImageUrl,
				CreatedAt:      *userDetailDb.CreatedAt,
				UpdatedAt:      *userDetailDb.UpdatedAt,
				DeletedAt:      userDetailDb.DeletedAt,
			}
		}
		dataAddresses, err := pu.addressRepo.GetAllAddressesByUserId(cForTx, userId)
		if err != nil {
			return appErrorPkg.NewServerError(err)
		}
		addresses = dtoProfile.ConvertToAddressResponses(dataAddresses)
		return nil
	})
	if err != nil {
		return nil, err
	}
	profile := dtoProfile.ResponseProfile{
		ID:         user.ID,
		RoleId:     user.RoleId,
		Role:       user.Role,
		Email:      user.Email,
		IsVerified: user.IsVerified,
		UserDetail: userDetail,
		Address:    addresses,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		DeletedAt:  user.DeletedAt,
	}
	return &profile, nil
}

func (pu *profileUseCaseImpl) PutMyProfile(ctx context.Context, reqBody *dtoProfile.RequestPutMyProfile, userId int64, roleId int64) (*dtoProfile.ResponseProfile, error) {
	var (
		addresses  []*dtoProfile.ResponseAddress
		user       dtoAuth.ResponseUser
		userDetail *dtoAuth.ResponseUserDetail
	)
	if reqBody.ProfileImage != nil {
		if filepath.Ext(reqBody.ProfileImage.Filename) != ".png" && filepath.Ext(reqBody.ProfileImage.Filename) != ".jpg" && filepath.Ext(reqBody.ProfileImage.Filename) != ".jpeg" {
			return nil, appErrorProfile.NewInvalidImageFormatErrorMessageError()
		}
		if reqBody.ProfileImage.Size > constant.MAX_IMAGE_SIZE {
			return nil, appErrorProfile.NewInvalidPhotoMaxSizeError()
		}
	}
	err := pu.transactor.Atomic(ctx, func(cForTx context.Context) error {
		userDb, err := pu.userRepo.FindByIDWithCompleteData(cForTx, userId)
		if err != nil || userDb.ID == 0 {
			return appErrorPkg.NewForbiddenAccessError()
		}
		user = dtoAuth.ResponseUser{
			ID:         userDb.ID,
			RoleId:     userDb.Role,
			Role:       utilsAuth.SpecifyRole(userDb.Role),
			Email:      userDb.Email,
			IsVerified: userDb.IsVerified,
			CreatedAt:  userDb.CreatedAt,
			UpdatedAt:  userDb.UpdatedAt,
			DeletedAt:  userDb.DeletedAt,
		}
		checkUserDetail, err := pu.userRepo.GetUserDetailByUserID(cForTx, userId)
		if err != nil {
			return appErrorPkg.NewServerError(err)
		}
		if checkUserDetail != nil && checkUserDetail.WhatsappNumber != nil && *checkUserDetail.WhatsappNumber != reqBody.WhatsappNumber {
			checkWa, err := pu.userRepo.FindByWhatsappNumber(cForTx, reqBody.WhatsappNumber)
			if err != nil || checkWa != 0 {
				return appErrorAuth.NewInvalidWaNumberAlreadyExists()
			}
		}
		if checkUserDetail == nil {
			checkWa, err := pu.userRepo.FindByWhatsappNumber(cForTx, reqBody.WhatsappNumber)
			if err != nil || checkWa != 0 {
				return appErrorAuth.NewInvalidWaNumberAlreadyExists()
			}
			_, err = pu.userRepo.SaveUserDetail(cForTx, user.ID, reqBody.Fullname)
			if err != nil {
				return appErrorPkg.NewServerError(err)
			}
		}
		var imgUrl string
		if reqBody.ProfileImage != nil {
			f, err := reqBody.ProfileImage.Open()
			if err != nil {
				return err
			}
			imgUrl, err = pu.cloudinaryUtil.UploadImage(cForTx, f, uploader.UploadParams{
				PublicID:       strings.ToLower(strings.ReplaceAll(utils.GeneratePhotoProfileTitle(userDb.ID, *checkUserDetail.Fullname), " ", "-")),
				UniqueFilename: api.Bool(true),
				Overwrite:      api.Bool(true),
				Invalidate:     api.Bool(true),
			})
			if err != nil {
				return appErrorPkg.NewServerError(err)
			}
		}
		userDetailDb, err := pu.profileRepo.PutMyProfile(cForTx, reqBody, userId, roleId, &imgUrl)
		if err != nil {
			return appErrorPkg.NewServerError(err)
		}
		userDetail = &dtoAuth.ResponseUserDetail{
			Id:             *userDetailDb.ID,
			UserId:         *userDetailDb.UserId,
			Fullname:       *userDetailDb.Fullname,
			WhatsappNumber: userDetailDb.WhatsappNumber,
			ImageUrl:       *userDetailDb.ImageUrl,
			CreatedAt:      *userDetailDb.CreatedAt,
			UpdatedAt:      *userDetailDb.UpdatedAt,
			DeletedAt:      userDetailDb.DeletedAt,
		}
		dataAddresses, err := pu.addressRepo.GetAllAddressesByUserId(cForTx, userId)
		if err != nil {
			return appErrorPkg.NewServerError(err)
		}

		addresses = dtoProfile.ConvertToAddressResponses(dataAddresses)
		return nil
	})
	if err != nil {
		return nil, err
	}
	profile := dtoProfile.ResponseProfile{
		ID:         user.ID,
		RoleId:     user.RoleId,
		Role:       user.Role,
		Email:      user.Email,
		IsVerified: user.IsVerified,
		UserDetail: userDetail,
		Address:    addresses,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		DeletedAt:  user.DeletedAt,
	}
	return &profile, nil
}
