package auth

import (
	"strings"
	"time"

	"github.com/mhmdiamd/go-social-service/infra/response"
)

type OtpEntity struct {
  Id int `db:"id"`
  Otp string `db:"otp"`
  Email string `db:"email"`
  Password string `db:"password"`
  IsActive int `db:"is_active"`
  ExpiredAt time.Time `db:"expired_at"`

  CreatedAt time.Time `db:"created_at"`
  UpdatedAt time.Time `db:"updated_at"`
}

func NewOtpEntity(otp string, req SendOtpRequestPayload) OtpEntity {
  entity := OtpEntity{
    Otp: otp,
    Email: req.Email,
    Password: req.Password,
    IsActive: 1,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }

  entity.SetExpiredTime()

  return entity
}

func (a *OtpEntity) Validate() (err error) {

  if err = a.ValidateEmail(); err != nil {
    return 
  }

  if err = a.ValidatePassword(); err != nil {
    return 
  }

  return
}

func (a *OtpEntity) ValidateEmail() (err error) {

  if a.Email == "" {
    return response.ErrEmailRequired
  }

  if len(strings.Split(a.Email, "@")) != 2 {
    return response.ErrEmailInvalid
  }

  return
}

func (a *OtpEntity) ValidatePassword() (err error) {

  if (a.Password == ""){
    return response.ErrPasswordRequired
  }

  if len(a.Password) < 6 {
   return response.ErrPasswordInvalid
  }

  return
}

func (o *OtpEntity) SetExpiredTime() {
  expiredTime := time.Now().Add(time.Minute * time.Duration(30))
  o.ExpiredAt = expiredTime
}

