package response

import "errors"

var (
  // Auth Error
  ErrEmailRequired = errors.New("email is required")
  ErrEmailInvalid = errors.New("email is invalid")
  ErrEmailAlreadyUsed = errors.New("email already used")
  ErrPasswordInvalid = errors.New("password is invalid, character min 6")
  ErrPasswordRequired = errors.New("password is required")
)
