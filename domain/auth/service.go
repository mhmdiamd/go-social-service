package auth

import (
	"context"
	"fmt"
	"log"

	"github.com/mhmdiamd/go-social-service/external/google"
	"github.com/mhmdiamd/go-social-service/helper"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/mhmdiamd/go-social-service/internal/config"
)

type Repository interface {
  GetOtpByEmail(ctx context.Context, email string) (models []OtpEntity, err error)
  CreateOTP(ctx context.Context, model OtpEntity) (err error)
  GetAuthByEmail(ctx context.Context, email string) (model AuthEntity, err error)
  CreateAuth(ctx context.Context, model AuthEntity) (err error) 
}

type service struct {
  repo Repository
}

func newService(repo Repository) service {
  return service{
    repo: repo,
  }
}

func (s service) register(ctx context.Context, req RegisterRequestPayload) (err error) {

  authEntity := NewAuthEntityFromRegister(req)

  // Validation
  if err = authEntity.Validate(); err != nil {
    return 
  }

  // Password encryption 
  if err = authEntity.EncryptPassword(int(config.Cfg.App.Encryption.Salt)); err != nil {
    return
  }

  // Check is account exists
  model, err := s.repo.GetAuthByEmail(ctx, authEntity.Email)
  if err != nil {
    if err != response.ErrNotFound {
      return
    }
  }

  if model.IsExists() {
    return response.ErrEmailAlreadyUsed
  }

  // Execute Serivce
  return s.repo.CreateAuth(ctx, authEntity)
}

func (s service) login(ctx context.Context, req LoginRequestPayload) (token string, err error) {
  authEntity := NewAuthEntityFromLogin(req)

  // Validation
  if err = authEntity.Validate(); err != nil {
    return 
  }

  // Check is account exists
  model, err := s.repo.GetAuthByEmail(ctx, req.Email)
  if err != nil {
    return
  }

  // Validation Password
  if err = authEntity.VerifyPasswordFromPlain(model.Password); err != nil {
    err = response.ErrPasswordNotMatch
    return
  }

  token, err = model.GenerateToken(config.Cfg.App.Encryption.JWTSecret)

  return 
}

func (s service) sendOtp(ctx context.Context, req SendOtpRequestPayload) (err error) {

  // first is Check otp is already exists or not

  // Generate OTP
  otp := helper.GenerateOTP()
  log.Println(otp)
  userOtps, err := s.repo.GetOtpByEmail(ctx, req.Email)

  for _, userOtp := range userOtps {
    if userOtp.Otp == otp && userOtp.IsActive != 0 {
      otp = helper.GenerateOTP()
    }
  }

  // Send OTP to Gmail
  header := google.SendEmailHeaderParams{
    Sender_email: config.Cfg.App.External.Google.Smtp_sender_email,
    Password: config.Cfg.App.External.Google.Smtp_password,
    Recipent_email: req.Email,
    Mail_subject: "Sign up OTP",
    Mail_body: otp,
  }

  newSmtpHeader := google.NewSendEmailHeader(header)
  err = newSmtpHeader.SendGmailWithSMTP()

  if err != nil {
    err = response.ErrBadRequest
  }

  // than create the entity into the database
  otpEntity := NewOtpEntity(otp, req)
  fmt.Println(otpEntity)
  err = s.repo.CreateOTP(ctx, otpEntity)
  if err != nil {
    return 
  }

  return
}


