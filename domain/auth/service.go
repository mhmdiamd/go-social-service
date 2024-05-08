package auth

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/external/google"
	"github.com/mhmdiamd/go-social-service/helper"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/mhmdiamd/go-social-service/internal/config"
	"github.com/mhmdiamd/go-social-service/internal/log"
	tempdata "github.com/mhmdiamd/go-social-service/temp_data"
)

type Repository interface {
	GetOtpByEmail(ctx context.Context, email string) (models []OtpEntity, err error)
	GetDetailOtp(ctx context.Context, publid_id_user_otp uuid.UUID) (model OtpEntity, err error)
	GetDetailOtpByEmailAndOtp(ctx context.Context, req VerifyOtpRequestPayload) (model OtpEntity, err error)
	CreateOTP(ctx context.Context, model OtpEntity) (err error)
	UnactiveOtp(ctx context.Context, email string) (err error)

	DeleteAuthByEmail(ctx context.Context, email string) (err error)

	GetAuthByEmail(ctx context.Context, email string) (model AuthEntity, err error)
	CreateAuth(ctx context.Context, model AuthEntity) (err error)
  DeleteOtpByEmail(ctx context.Context, email string) (err error)
}

type Service struct {
	Repo Repository
}

func NewService(repo Repository) Service {
	return Service{
		Repo: repo,
	}
}

func (s Service) Register(ctx context.Context, req RegisterRequestPayload) (err error) {

	authEntity := NewAuthEntityFromRegister(req)

	// Validation auth register payload
	if err = authEntity.Validate(); err != nil {
		return
	}

	// Get data auth by OTP from DB
	userOtp, err := s.Repo.GetDetailOtp(ctx, req.PublicIdUserOtp)
	if err != nil {
		if err == response.ErrNotFound {
			err = response.ErrOtpInvalid
		}

		return
	}

	// check is otp expired or not
	if userOtp.IsExpired() {
		return response.ErrOtpExpired
	}

	// Encrypt Password
	if err = authEntity.EncryptPassword(req.Password, int(config.Cfg.App.Encryption.Salt)); err != nil {
		return
	}

	// Add email to entity
	authEntity.Email = userOtp.Email
	// Add last public id to Tempdata
	tempdata.TempLastUserPublicId = authEntity.PublicId.String()

	// Execute Serivce
	return s.Repo.CreateAuth(ctx, authEntity)
}

func (s Service) login(ctx context.Context, req LoginRequestPayload) (token string, err error) {
	authEntity := NewAuthEntityFromLogin(req)

	// Validation
	if err = authEntity.ValidateLoginPayload(); err != nil {
		return
	}

	// Check is account exists
	model, err := s.Repo.GetAuthByEmail(ctx, req.Email)
	if err != nil {
		if err == response.ErrNotFound {
			err = response.ErrUnauthorized
		}

		return
	}

	// Validation Password
	if err = authEntity.VerifyPasswordFromPlain(model.Password); err != nil {
		return "", response.ErrPasswordNotMatch
	}

	token, err = model.GenerateToken(config.Cfg.App.Encryption.JWTSecret)

	return
}

func (s Service) SendOtp(ctx context.Context, req SendOtpRequestPayload) (err error) {

	// Generate OTP
	otp := helper.GenerateOTP()

	// than create the entity for insert into the database
	otpEntity := NewOtpEntity(otp, req.Email)

	// Validating the payload
	err = otpEntity.ValidateEmail()
	if err != nil {
    log.Log.Errorf(ctx, "[SendOtp, ValidateEmail] with error detail %v", err.Error())
		return
  }

	// Check is otp already exists or not, if yes then regenerate the otp itself
	userOtps, err := s.Repo.GetOtpByEmail(ctx, req.Email)
	maxSendEmail := 3

	// Check is user to much send email
	if len(userOtps) >= maxSendEmail {
		return response.ErrToMuchSendEmail
	}

	for _, userOtp := range userOtps {
		for userOtp.Otp == otp && userOtp.IsActive == 1 {
			otp = helper.GenerateOTP()
		}
	}

	// Send OTP to Gmail
	header := google.SendEmailHeaderParams{
		Sender_email:   config.Cfg.App.External.Google.Smtp_sender_email,
		Password:       config.Cfg.App.External.Google.Smtp_password,
		Recipent_email: req.Email,
		Mail_subject:   "Sign up OTP",
		Mail_body:      otp,
	}

	// Save OTP to temp data for unitest
	tempdata.TempRegisterOtp = otp

	newSmtpHeader := google.NewSendEmailHeader(header)
	go newSmtpHeader.SendGmailWithSMTP()

	// Deactivated existing OTP
	err = s.Repo.UnactiveOtp(ctx, req.Email)
	if err != nil {
		return
	}

	// Create user otp
	err = s.Repo.CreateOTP(ctx, otpEntity)
	if err != nil {
		return
	}

	return
}

func (s Service) VerifyOtp(ctx context.Context, req VerifyOtpRequestPayload) (otp_id string, err error) {

	otpEntity := NewOtpEntity(req.Otp, req.Email)

	if err = otpEntity.Validate(); err != nil {
		return
	}

	model, err := s.Repo.GetDetailOtpByEmailAndOtp(ctx, req)
	if err != nil {
		if err == response.ErrNotFound {
			err = response.ErrOtpInvalid
			return
		}
	}

	otp_id = model.PublicId.String()

	// set otp id for unitesting
	tempdata.TempPublicIdUserOtp = otp_id

	return
}

func (s Service) DeleteAuth(ctx context.Context, email string) (err error) {
	err = s.Repo.DeleteAuthByEmail(ctx, email)
	if err != nil {
    fmt.Println(err)
		return
	}

	return
}
