package auth

import "github.com/google/uuid"

type RegisterRequestPayload struct {
  Name string `json:"name"`
  Password string `json:"password"`
  PasswordConfirmation string `json:"password_confirmation"`
  PublicIdUserOtp uuid.UUID `json:"otp_id"`
}

type LoginRequestPayload struct {
  Email string `json:"email"`
  Password string `json:"password"`
}

type SendOtpRequestPayload struct {
  Email string `json:"email"`
}

type VerifyOtpRequestPayload struct {
  Email string `json:"email"`
  Otp string `json:"otp"`
}


