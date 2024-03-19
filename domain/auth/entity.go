package auth

import (
	"database/sql"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/external/jwt"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"golang.org/x/crypto/bcrypt"
)

type Gender string

const (
	GENDER_Male   Gender = "male"
	GENDER_Female Gender = "female"
)

type AuthEntity struct {
	Id                   int            `db:"id"`
	PublicId             uuid.UUID      `db:"public_id"`
	Name                 string         `db:"name"`
	Email                string         `db:"email"`
	Password             string         `db:"password"`
	PasswordConfirmation string         `db:"-"`
	Gender               sql.NullString `db:"gender"`
	No_tlp               sql.NullString `db:"no_tlp"`
	Address              sql.NullString `db:"address"`

	PublicIdUserOtp uuid.UUID `db:"user_otp_public_id"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewAuthEntityFromRegister(req RegisterRequestPayload) AuthEntity {
	return AuthEntity{
		PublicId:             uuid.New(),
		Name:                 req.Name,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
		PublicIdUserOtp:      req.PublicIdUserOtp,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}
}

func NewAuthEntityFromLogin(req LoginRequestPayload) AuthEntity {
	return AuthEntity{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (r AuthEntity) Validate() (err error) {

	if err = r.ValidateName(); err != nil {
		return
	}

	if err = r.ValidatePassword(); err != nil {
		return
	}

	if err = r.ValidatePasswordConfirmation(); err != nil {
		return
	}

	if err = r.ValidateUserOtp(); err != nil {
		return
	}

	return
}

func (r AuthEntity) ValidateLoginPayload() (err error) {

	if err = r.ValidateEmail(); err != nil {
		return
	}

	if err = r.ValidatePassword(); err != nil {
		return
	}

	return
}

func (r AuthEntity) ValidateEmail() (err error) {
	if r.Email == "" {
		return response.ErrEmailRequired
	}

	if len(strings.Split(r.Email, "@")) != 2 {
		return response.ErrEmailInvalid
	}

	return
}

func (r AuthEntity) ValidateName() (err error) {
	if r.Name == "" {
		return response.ErrNameRequired
	}

	if len(r.Name) < 4 {
		return response.ErrNameInvalid
	}

	return
}

func (r AuthEntity) ValidatePassword() (err error) {
	if r.Password == "" {
		return response.ErrPasswordRequired
	}

	if len(r.Password) < 6 {
		return response.ErrPasswordInvalid
	}

	return
}

func (r AuthEntity) ValidatePasswordConfirmation() (err error) {
	if r.PasswordConfirmation == "" {
		return response.ErrPasswordConfirmationRequired
	}

	if len(r.PasswordConfirmation) < 6 {
		return response.ErrPasswordConfirmationInvalid
	}

	if r.Password != r.PasswordConfirmation {
		return response.ErrPasswordConfirmationNotMatch
	}

	return
}

func (r AuthEntity) ValidateUserOtp() (err error) {

	if r.PublicIdUserOtp.String() == "" {
		return response.ErrPublicIdUserOtpRequired
	}

	return
}

func (a AuthEntity) IsExists() bool {
	return a.Id != 0
}

func (a *AuthEntity) EncryptPassword(password string, salt int) (err error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	if err != nil {
		return
	}

	a.Password = string(encryptedPassword)

	return
}

func (a AuthEntity) VerifyPasswordFromPlain(encrypted string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(a.Password))
}

func (a AuthEntity) GenerateToken(secret string) (tokenString string, err error) {
	return jwt.GenerateToken(a.PublicId.String(), secret)
}
