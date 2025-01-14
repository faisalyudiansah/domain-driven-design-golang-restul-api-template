package apperror

import (
	"errors"
	"fmt"

	"server/pkg/constant"
)

func NewEntityNotFoundError(entity string) *AppError {
	msg := fmt.Sprintf(constant.EntityNotFoundErrorMessage, entity)

	err := errors.New(msg)

	return NewAppError(err, DefaultClientErrorCode, msg)
}
