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

  ErrIdRequired = errors.New("id is required")
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

  // Event Demographics
  ErrStartAgeRequired = errors.New("start age is required")
  ErrStartAgeToMax = errors.New("start age is to much bigger, must be lower than 99")
  ErrStartAgeToMin = errors.New("start age is to much lower, must be greater then 0")
  ErrStartAgeGreaterThenEndAge = errors.New("start age is greater than end age")
  ErrEndAgeRequired = errors.New("end age is required")
  ErrEndAgeToMax = errors.New("end age is to much bigger, must be lower than 99")
  ErrEndAgeToMin = errors.New("end age is to much lower, must be greater then 0")
  ErrEndAgeLowerThenStartAge = errors.New("end age is lower than start age")

  ErrGenderInvalid = errors.New("gender invalid, should be male or female")

  ErrDescriptionRequired = errors.New("description is required")
  ErrAddressRequired = errors.New("address is required")
  ErrThumbnailRequired = errors.New("thumbnail is required")
  ErrStartDateRequired = errors.New("start date is required")
  ErrEndDateRequired = errors.New("end date is required")
  ErrStartDateGreaterThanEndDate = errors.New("start date greater than end date")

  ErrCommunityCategoryIdRequired = errors.New("community category id is required")

  ErrFileTypeNotCompatible = errors.New("community category id is required")
  ErrImageTypeNotCompatible = errors.New("image type not compatible, only png, jpg, jpeg")
  ErrImageOversize = errors.New("image size is oversized, max size 1mb")

  ErrCommunityMemberRoleNotPermitted = errors.New("role is not permitted")
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

  ErrorStartAgeRequired = NewError(ErrStartAgeRequired.Error(), "40015", http.StatusBadRequest)
  ErrorStartAgeToMax = NewError(ErrStartAgeToMax.Error(), "40016", http.StatusBadRequest)
  ErrorStartAgeToMin = NewError(ErrStartAgeToMin.Error(), "40017", http.StatusBadRequest)
  ErrorStartAgeGreaterThenEndAge = NewError(ErrStartAgeGreaterThenEndAge.Error(), "40018", http.StatusBadRequest)
  ErrorEndAgeRequired = NewError(ErrEndAgeRequired.Error(), "40019", http.StatusBadRequest)
  ErrorEndAgeToMax = NewError(ErrEndAgeToMax.Error(), "40020", http.StatusBadRequest)
  ErrorEndAgeToMin = NewError(ErrEndAgeToMin.Error(), "40021", http.StatusBadRequest)
  ErrorEndAgeLowerThenStartAge = NewError(ErrEndAgeLowerThenStartAge.Error(), "40022", http.StatusBadRequest)

  ErrorIdRequired = NewError(ErrIdRequired.Error(), "40023", http.StatusBadRequest)

  ErrorGenderInvalid = NewError(ErrGenderInvalid.Error(), "40024", http.StatusBadRequest)

  ErrorCommunityCategoryIdRequired = NewError(ErrCommunityCategoryIdRequired.Error(), "40025", http.StatusBadRequest)

  ErrorFileTypeNotCompatible = NewError(ErrFileTypeNotCompatible.Error(), "40026", http.StatusBadRequest)
  ErrorImageTypeNotCompatible = NewError(ErrImageTypeNotCompatible.Error(), "40027", http.StatusBadRequest)
  ErrorImageOversize = NewError(ErrImageOversize.Error(), "40028", http.StatusBadRequest)

  // Community Member
  ErrorCommunityMemberRoleNotPermitted =  NewError(ErrCommunityMemberRoleNotPermitted.Error(), "40029", http.StatusBadRequest)
)

var ErrorMapping = map[string]Error{
  ErrorBadRequest.Error() : ErrorBadRequest,
  ErrorGeneral.Error() : ErrorGeneral,
  ErrorNotFound.Error() : ErrorNotFound,
  ErrorUnauthorized.Error() : ErrorUnauthorized,

  ErrorIdRequired.Error() : ErrorIdRequired,

  // Auth
  ErrorEmailRequired.Error() : ErrorEmailRequired ,
  ErrorEmailInvalid.Error() : ErrorEmailInvalid,
  ErrorEmailAlreadyUsed.Error() : ErrorEmailAlreadyUsed,
  ErrorPasswordRequired.Error() : ErrorPasswordRequired,
  ErrorPasswordInvalid.Error() : ErrorPasswordInvalid, 
  ErrorPasswordNotMatch.Error() : ErrorPasswordNotMatch,

  // Auth - Register
  ErrorNameRequired.Error() : ErrorNameRequired ,
  ErrorNameInvalid.Error() : ErrorNameInvalid,
  ErrorPasswordConfirmationRequired.Error() : ErrorPasswordConfirmationRequired,
  ErrorPasswordConfirmationInvalid.Error() : ErrorPasswordConfirmationInvalid, 
  ErrorPasswordConfirmationNotMatch.Error() : ErrorPasswordConfirmationNotMatch,
  ErrorOtpRequired.Error() : ErrorOtpRequired,
  ErrorOtpInvalid.Error() : ErrorOtpInvalid,
  ErrorOtpExpired.Error() : ErrorOtpExpired,
  ErrorToMuchSendEmail.Error() : ErrorToMuchSendEmail,

  // Event Demographies
  ErrStartAgeRequired.Error() : ErrorStartAgeRequired,
  ErrStartAgeToMax.Error() :  ErrorStartAgeToMax,
  ErrStartAgeToMin.Error() : ErrorStartAgeToMin,
  ErrStartAgeGreaterThenEndAge.Error() : ErrorStartAgeGreaterThenEndAge, 
  ErrEndAgeRequired.Error() : ErrorEndAgeRequired,
  ErrEndAgeToMax.Error() : ErrorEndAgeToMax,
  ErrEndAgeToMin.Error() : ErrorEndAgeToMin,  
  ErrorEndAgeLowerThenStartAge.Error() : ErrorEndAgeLowerThenStartAge,

  ErrorGenderInvalid.Error() : ErrorGenderInvalid,

  ErrCommunityCategoryIdRequired.Error() : ErrorCommunityCategoryIdRequired,

  ErrFileTypeNotCompatible.Error() : ErrorFileTypeNotCompatible,
  ErrImageTypeNotCompatible.Error() : ErrorImageTypeNotCompatible,
  ErrImageOversize.Error() : ErrorImageOversize,

  // Community member
  ErrCommunityMemberRoleNotPermitted.Error() : ErrorCommunityMemberRoleNotPermitted,
}



