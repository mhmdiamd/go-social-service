package auth

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/infra/response"
)

type OtpEntity struct {
  Id int `db:"id"`
  PublicId uuid.UUID `db:"public_id"`
  Otp string `db:"otp"`
  Email string `db:"email"`
  IsActive int `db:"is_active"`
  Password string `db:"password"`
  ExpiredAt time.Time `db:"expired_at"`

  CreatedAt time.Time `db:"created_at"`
  UpdatedAt time.Time `db:"updated_at"`
}

func NewOtpEntity(otp, email string) OtpEntity {

  expiredTime := time.Now().Add(time.Minute * time.Duration(30))

  entity := OtpEntity{
    Otp: otp,
    PublicId: uuid.New(),
    Email: email,
    IsActive : 1,
    ExpiredAt: expiredTime,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }

  return entity
}

func (a *OtpEntity) Validate() (err error) {

  if err = a.ValidateEmail(); err != nil {
    return 
  }

 if err = a.ValidateOtp(); err != nil {
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

func (a *OtpEntity) ValidateOtp() (err error) {

  if a.Otp == "" {
    return response.ErrOtpRequired
  }

  if len(a.Otp) != 4 {
    return response.ErrOtpInvalid
  }

  return
}

func (a *OtpEntity) IsExpired() bool {
  return a.ExpiredAt.Unix() < time.Now().Unix()
}


