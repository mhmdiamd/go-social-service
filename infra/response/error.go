package response

import (
	"errors"
	"net/http"
)

var (
  ErrNotFound = errors.New("not found")
  ErrUnauthorized = errors.New("unauthorized")
  ErrBadRequest = errors.New("bad request")
)

var (
  // Auth Error
  ErrPublicIdUserOtpRequired = errors.New("otp id is required")

  ErrEmailRequired = errors.New("email is required")
  ErrEmailInvalid = errors.New("email is invalid")
  ErrEmailAlreadyUsed = errors.New("email already used")

  ErrNameRequired = errors.New("name is required")
  ErrNameInvalid = errors.New("name is invalid")

  ErrPasswordInvalid = errors.New("password is invalid, character min 6")
  ErrPasswordRequired = errors.New("password is required")
  ErrPasswordNotMatch = errors.New("password is not match")

  ErrPasswordConfirmationInvalid = errors.New("password confirmation is invalid, character min 6")
  ErrPasswordConfirmationRequired = errors.New("password confirmation is required")
  ErrPasswordConfirmationNotMatch = errors.New("password confirmation is not match")

  ErrOtpRequired = errors.New("otp is required")
  ErrOtpInvalid = errors.New("otp is invalid")
  ErrOtpExpired = errors.New("otp is expired")
  ErrToMuchSendEmail = errors.New("to much send email (more than 3x), your account get blocked")
)

type Error struct {
  Message string
  Code string
  HttpCode int
}

func (e Error) Error() string {
  return e.Message
}

func NewError(message, code string, httpCode int) Error {
  return Error {
    Message: message,
    Code: code,
    HttpCode: httpCode,
  }
}

var (
  // Global error
  ErrorGeneral = NewError("internal server error", "99999", http.StatusInternalServerError)
  ErrorBadRequest = NewError(ErrBadRequest.Error(), "40000", http.StatusBadRequest)
  ErrorNotFound = NewError(ErrNotFound.Error(), "40400", http.StatusNotFound)
  ErrorUnauthorized = NewError(ErrUnauthorized.Error(), "40900", http.StatusBadRequest)
)

var (
  // entity error
  ErrorEmailRequired = NewError(ErrEmailRequired.Error(), "40001", http.StatusBadRequest)
  ErrorEmailInvalid = NewError(ErrEmailInvalid.Error(), "40002", http.StatusBadRequest)
  ErrorEmailAlreadyUsed = NewError(ErrEmailAlreadyUsed.Error(), "40901", http.StatusConflict)

  ErrorNameRequired = NewError(ErrNameRequired.Error(), "40003", http.StatusBadRequest)
  ErrorNameInvalid = NewError(ErrNameInvalid.Error(), "40004", http.StatusBadRequest)

  ErrorPasswordRequired = NewError(ErrPasswordInvalid.Error(), "40005", http.StatusBadRequest)
  ErrorPasswordInvalid = NewError(ErrPasswordInvalid.Error(), "40006", http.StatusBadRequest)
  ErrorPasswordNotMatch = NewError(ErrPasswordNotMatch.Error(), "40007", http.StatusUnauthorized)

  ErrorPasswordConfirmationRequired = NewError(ErrPasswordConfirmationRequired.Error(), "40008", http.StatusBadRequest)
  ErrorPasswordConfirmationInvalid = NewError(ErrPasswordConfirmationInvalid.Error(), "40009", http.StatusBadRequest)
  ErrorPasswordConfirmationNotMatch = NewError(ErrPasswordConfirmationNotMatch.Error(), "40010", http.StatusUnauthorized)

  ErrorPublicIdUserOtpRequired = NewError(ErrPublicIdUserOtpRequired.Error(), "40011", http.StatusUnauthorized)

  ErrorOtpRequired = NewError(ErrOtpRequired.Error(), "40012", http.StatusBadRequest)
  ErrorOtpInvalid = NewError(ErrOtpInvalid.Error(), "400013", http.StatusBadRequest)
  ErrorOtpExpired = NewError(ErrOtpExpired.Error(), "40901", http.StatusUnauthorized)
  ErrorToMuchSendEmail = NewError(ErrToMuchSendEmail.Error(), "40014", http.StatusBadRequest)
)

var ErrorMapping = map[string]Error{
  ErrorBadRequest.Error() : ErrorBadRequest,
  ErrorGeneral.Error() : ErrorGeneral,
  ErrorNotFound.Error() : ErrorNotFound,
  ErrorUnauthorized.Error() : ErrorUnauthorized,

  // Auth
  ErrorEmailRequired.Error() : ErrorEmailRequired ,
  ErrorEmailInvalid.Error() : ErrorEmailInvalid,
  ErrorEmailAlreadyUsed.Error() : ErrorEmailAlreadyUsed,
  ErrorPasswordInvalid.Error() : ErrorPasswordRequired,
  ErrorPasswordInvalid.Error() : ErrorPasswordInvalid, 
  ErrorPasswordNotMatch.Error() : ErrorPasswordNotMatch,

  ErrorOtpRequired.Error() : ErrorOtpRequired,
  ErrorOtpInvalid.Error() : ErrorOtpInvalid,
  ErrorOtpExpired.Error() : ErrorOtpExpired,
  ErrorToMuchSendEmail.Error() : ErrorToMuchSendEmail,
}



