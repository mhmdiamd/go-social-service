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

    OtpEntity :=  SendOtpRequestPayload{
      Email :"am@gmail.com",
      Password : "password",
    }
    
    payload := NewOtpEntity(randOtp, OtpEntity)

    err := payload.Validate()
    require.Nil(t, err)
  })

  t.Run("fail email invalid", func(t *testing.T){
    randOtp := helper.GenerateOTP()

    OtpEntity :=  SendOtpRequestPayload{
      Email :"amgmail.com",
      Password : "password",
    }
    
    payload := NewOtpEntity(randOtp, OtpEntity)

    err := payload.Validate()
    require.NotNil(t, err)
    require.Equal(t, response.ErrEmailInvalid, err)
  })

  t.Run("fail email is required", func(t *testing.T) {
    randOtp := helper.GenerateOTP()

    OtpEntity :=  SendOtpRequestPayload{
      Email :"",
      Password : "password",
    }
    
    payload := NewOtpEntity(randOtp, OtpEntity)

    err := payload.Validate()
    require.NotNil(t, err)
    require.Equal(t, response.ErrEmailRequired, err)
  })

  t.Run("fail password invalid", func(t *testing.T){
    randOtp := helper.GenerateOTP()

    OtpEntity :=  SendOtpRequestPayload{
      Email :"am@gmail.com",
      Password : "pass",
    }
    
    payload := NewOtpEntity(randOtp, OtpEntity)

    err := payload.Validate()
    require.NotNil(t, err)
    require.Equal(t, response.ErrPasswordInvalid, err)
  })

  t.Run("fail password is required", func(t *testing.T) {
    randOtp := helper.GenerateOTP()

    OtpEntity :=  SendOtpRequestPayload{
      Email :"am@gmail.com",
      Password : "",
    }
    
    payload := NewOtpEntity(randOtp, OtpEntity)

    err := payload.Validate()
    require.NotNil(t, err)
    require.Equal(t, response.ErrPasswordRequired, err)
  })
}

