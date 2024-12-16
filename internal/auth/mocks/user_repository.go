// Code generated by mockery v2.47.0. DO NOT EDIT.

package mocks

import (
	context "context"
	dto "server/internal/auth/dto"
	entity "server/internal/auth/entity"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// FindByEmail provides a mock function with given fields: ctx, email
func (_m *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for FindByEmail")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, id
func (_m *UserRepository) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*entity.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entity.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByIDWithCompleteData provides a mock function with given fields: ctx, id
func (_m *UserRepository) FindByIDWithCompleteData(ctx context.Context, id int64) (*entity.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindByIDWithCompleteData")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*entity.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entity.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByWhatsappNumber provides a mock function with given fields: ctx, waNumber
func (_m *UserRepository) FindByWhatsappNumber(ctx context.Context, waNumber string) (int64, error) {
	ret := _m.Called(ctx, waNumber)

	if len(ret) == 0 {
		panic("no return value specified for FindByWhatsappNumber")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int64, error)); ok {
		return rf(ctx, waNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, waNumber)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, waNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserDetailByUserID provides a mock function with given fields: ctx, userId
func (_m *UserRepository) GetUserDetailByUserID(ctx context.Context, userId int64) (*entity.UserDetail, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for GetUserDetailByUserID")
	}

	var r0 *entity.UserDetail
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*entity.UserDetail, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entity.UserDetail); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UserDetail)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, reqBody, hashPassword, from
func (_m *UserRepository) Save(ctx context.Context, reqBody dto.RequestUserRegister, hashPassword string, from string) (*entity.User, error) {
	ret := _m.Called(ctx, reqBody, hashPassword, from)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.RequestUserRegister, string, string) (*entity.User, error)); ok {
		return rf(ctx, reqBody, hashPassword, from)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.RequestUserRegister, string, string) *entity.User); ok {
		r0 = rf(ctx, reqBody, hashPassword, from)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.RequestUserRegister, string, string) error); ok {
		r1 = rf(ctx, reqBody, hashPassword, from)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveOauth provides a mock function with given fields: ctx, _a1
func (_m *UserRepository) SaveOauth(ctx context.Context, _a1 *entity.User) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SaveOauth")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveUserDetail provides a mock function with given fields: ctx, userId, fullname
func (_m *UserRepository) SaveUserDetail(ctx context.Context, userId int64, fullname string) (*entity.UserDetail, error) {
	ret := _m.Called(ctx, userId, fullname)

	if len(ret) == 0 {
		panic("no return value specified for SaveUserDetail")
	}

	var r0 *entity.UserDetail
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) (*entity.UserDetail, error)); ok {
		return rf(ctx, userId, fullname)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) *entity.UserDetail); ok {
		r0 = rf(ctx, userId, fullname)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UserDetail)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, userId, fullname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateIsVerified provides a mock function with given fields: ctx, user
func (_m *UserRepository) UpdateIsVerified(ctx context.Context, user *entity.User) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateIsVerified")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePassword provides a mock function with given fields: ctx, user
func (_m *UserRepository) UpdatePassword(ctx context.Context, user *entity.User) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
