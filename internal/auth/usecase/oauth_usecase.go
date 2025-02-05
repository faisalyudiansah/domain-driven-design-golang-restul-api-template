package usecase

import (
	"context"
	"fmt"

	apperrorAuth "server/internal/auth/apperror"
	"server/internal/auth/dto"
	"server/internal/auth/entity"
	"server/internal/auth/repository"
	apperrorPkg "server/pkg/apperror"
	"server/pkg/database/transactor"
	"server/pkg/utils/jwtutils"

	"github.com/google/uuid"
	"github.com/markbates/goth"
)

type OauthUseCase interface {
	Login(ctx context.Context, request *goth.User) (*dto.ResponseLogin, error)
}

type oauthUseCaseImpl struct {
	jwtUtil                jwtutils.JwtUtil
	userRepository         repository.UserRepository
	refreshTokenRepository repository.RefreshTokenRepository
	transactor             transactor.Transactor
}

func NewOauthUseCase(
	jwtUtil jwtutils.JwtUtil,
	userRepository repository.UserRepository,
	refreshTokenRepository repository.RefreshTokenRepository,
	transactor transactor.Transactor,
) *oauthUseCaseImpl {
	return &oauthUseCaseImpl{
		jwtUtil:                jwtUtil,
		userRepository:         userRepository,
		refreshTokenRepository: refreshTokenRepository,
		transactor:             transactor,
	}
}

func (u *oauthUseCaseImpl) Login(ctx context.Context, request *goth.User) (*dto.ResponseLogin, error) {
	user, err := u.userRepository.FindByEmail(ctx, request.Email)
	if err != nil {
		return nil, apperrorPkg.NewServerError(err)
	}
	if user != nil && !user.IsOauth {
		return nil, apperrorAuth.NewInvalidEmailAlreadyExists(err)
	}
	if user == nil {
		user = &entity.User{Email: request.Email}
		if err := u.userRepository.SaveOauth(ctx, user); err != nil {
			return nil, apperrorPkg.NewServerError(err)
		}

		if _, err := u.userRepository.SaveUserDetail(ctx, user.ID, fmt.Sprintf("%v %v", request.FirstName, request.LastName)); err != nil {
			return nil, apperrorPkg.NewServerError(err)
		}
	}

	jti := uuid.NewString()
	token, err := u.jwtUtil.Sign(user.ID, user.Role, jti)
	if err != nil {
		return nil, apperrorAuth.NewInvalidLoginCredentials(err)
	}

	refreshToken, err := u.jwtUtil.SignRefresh()
	if err != nil {
		return nil, apperrorAuth.NewInvalidLoginCredentials(err)
	}

	if err := u.refreshTokenRepository.Save(
		ctx, &entity.RefreshToken{UserID: user.ID, RefreshToken: refreshToken, JTI: jti},
	); err != nil {
		return nil, apperrorPkg.NewServerError(err)
	}

	return &dto.ResponseLogin{AccessToken: token}, nil
}
