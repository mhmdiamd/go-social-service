package auth

import (
	"testing"

	"github.com/mhmdiamd/go-social-service/helper"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/stretchr/testify/require"
)


func Test_ValidationSendOTP(t *testing.T) {

  t.Run("success", func(t *testing.T) {

    randOtp := helper.GenerateOTP()

    email := "am@gmail.com"
    
    payload := NewOtpEntity(randOtp, email)

    err := payload.Validate()
    require.Nil(t, err)
  })

  t.Run("fail email invalid", func(t *testing.T){
    randOtp := helper.GenerateOTP()

    email := "amgmail.com"
    
    payload := NewOtpEntity(randOtp, email)

    err := payload.Validate()
    require.NotNil(t, err)
    require.Equal(t, response.ErrEmailInvalid, err)
  })

  t.Run("fail email is required", func(t *testing.T) {
    randOtp := helper.GenerateOTP()

    email := ""

    payload := NewOtpEntity(randOtp, email)

    err := payload.Validate()
    require.NotNil(t, err)
    require.Equal(t, response.ErrEmailRequired, err)
  })
}

func Test_ValidationVerifyOTP(t *testing.T) {

  t.Run("success", func(t *testing.T) {

    randOtp := helper.GenerateOTP()

    email := "am@gmail.com"
    
    payload := NewOtpEntity(randOtp, email)

    err := payload.ValidateEmail()
    require.Nil(t, err)
  })

  t.Run("fail email invalid", func(t *testing.T){
    randOtp := helper.GenerateOTP()

    email := "amgmail.com"
    
    payload := NewOtpEntity(randOtp, email)

    err := payload.ValidateEmail()
    require.NotNil(t, err)
    require.Equal(t, response.ErrEmailInvalid, err)
  })

  t.Run("fail email is required", func(t *testing.T) {
    randOtp := helper.GenerateOTP()

    email := ""

    payload := NewOtpEntity(randOtp, email)

    err := payload.ValidateEmail()
    require.NotNil(t, err)
    require.Equal(t, response.ErrEmailRequired, err)
  })

  t.Run("fail otp invalid", func(t *testing.T){
    randOtp := helper.GenerateOTP()

    email := "amgmail.com"
    
    payload := NewOtpEntity(randOtp, email)

    err := payload.Validate()
    require.NotNil(t, err)
    require.Equal(t, response.ErrEmailInvalid, err)
  })

  t.Run("fail otp is required", func(t *testing.T) {
    // randOtp := helper.GenerateOTP()

    email := "am@gmail.com"

    payload := NewOtpEntity("", email)

    err := payload.Validate()
    require.NotNil(t, err)
    require.Equal(t, response.ErrOtpRequired, err)
  })
}
