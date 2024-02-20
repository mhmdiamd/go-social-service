package auth

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"golang.org/x/crypto/bcrypt"
)

type Gender string

const (
  GENDER_Male Gender = "male"
  GENDER_Female Gender = "female"
)

type AuthEntity struct {
  Id int `db:"id"`
  PublicId uuid.UUID `db:"public_id"`
  Nik string `db:"nik"`
  Name string `db:"name"`
  Email string `db:"email"`
  Password string `db:"password"`
  Gender Gender `db:"gender"`

  CreatedAt time.Time `db:"created_at"`
  UpdatedAt time.Time `db:"updated_at"`
}

func NewAuthEntityFromRegister(req RegisterRequestPayload) AuthEntity {
  return AuthEntity{
    PublicId: uuid.New(),
    Email: req.Email,
    Password: req.Password,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }
}

func NewAuthEntityFromLogin(req RegisterRequestPayload) AuthEntity {
  return AuthEntity{
    Email: req.Email,
    Password: req.Password,
  }
}

func (a *AuthEntity) Validate() (err error) {

  if err = a.ValidateEmail(); err != nil {
    return 
  }

  if err = a.ValidatePassword(); err != nil {
    return
  }

  return
}

func (a *AuthEntity) ValidateEmail() (err error) {

  if a.Email == "" {
    return response.ErrEmailRequired
  }

  if len(strings.Split(a.Email, "@")) != 2 {
    return response.ErrEmailInvalid
  }

  return
}

func (a *AuthEntity) ValidatePassword() (err error) {

  if (a.Password == ""){
    return response.ErrPasswordRequired
  }

  if len(a.Password) < 6 {
   return response.ErrPasswordInvalid
  }

  return
}

func (a AuthEntity) IsExists() bool {
  return a.Id != 0
}

func (a *AuthEntity) EncryptedPassword(salt int) (err error) {
  encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), salt)
  if err != nil {
    return err
  }

  a.Password = string(encryptedPassword)
  return
}


