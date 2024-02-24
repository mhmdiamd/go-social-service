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
  ErrEmailRequired = errors.New("email is required")
  ErrEmailInvalid = errors.New("email is invalid")
  ErrEmailAlreadyUsed = errors.New("email already used")
  ErrPasswordInvalid = errors.New("password is invalid, character min 6")
  ErrPasswordNotMatch = errors.New("password is not match")
  ErrPasswordRequired = errors.New("password is required")
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

  ErrorPasswordRequired = NewError(ErrPasswordInvalid.Error(), "40003", http.StatusBadRequest)
  ErrorPasswordInvalid = NewError(ErrPasswordInvalid.Error(), "40004", http.StatusBadRequest)
  ErrorPasswordNotMatch = NewError(ErrPasswordNotMatch.Error(), "40005", http.StatusUnauthorized)
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
}



