package auth

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/external/database"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/mhmdiamd/go-social-service/internal/config"
	tempdata "github.com/mhmdiamd/go-social-service/temp_data"
	"github.com/stretchr/testify/require"
)

var svc service
var tempEmail string 
var tempPassword string 

func init() {
  filename := "../../cmd/api/config.yaml"
  err := config.LoadConfig(filename)

  db, err := database.ConnectPostgres(config.Cfg.Db)
  if err != nil {
    panic(err)
  }

  tempEmail = "milham0141@gmail.com"
  tempPassword = "password" 

  repo := newRepository(db)
  svc = newService(repo)
}



func Test_SendOTP(t *testing.T) {

  models, _ := svc.repo.GetOtpByEmail(context.Background(), tempEmail)
  if len(models) >= 3 {
    t.Run("fail to much send email", func(t *testing.T) {

      req := SendOtpRequestPayload{
        Email: tempEmail,
      }

      err := svc.sendOtp(context.Background(), req)
      require.NotNil(t, err)
      require.Equal(t, response.ErrToMuchSendEmail, err)
    })
  } else {

    t.Run("success", func(t *testing.T) {
      req := SendOtpRequestPayload{
        Email: tempEmail,
      }

      // Delete temp account in the database first
      err := svc.deleteAuth(context.Background(), req.Email)
      require.Nil(t, err)

      err = svc.sendOtp(context.Background(), req)
      require.Nil(t, err)

        t.Run("success verify otp", func(t *testing.T) {
          req := VerifyOtpRequestPayload{
            Email: tempEmail,
            Otp: tempdata.TempRegisterOtp,
          }  

          _, err := svc.verifyOtp(context.Background(), req)
          require.Nil(t, err)
        })


        t.Run("success", func(t *testing.T) {
          req := RegisterRequestPayload{
            Name : "Muhamad Ilham",
            Password : tempPassword,
            PasswordConfirmation : tempPassword,
            PublicIdUserOtp : uuid.MustParse(tempdata.TempPublicIdUserOtp),
          }
            
          err := svc.register(context.Background(), req)
          require.Nil(t, err)
        })
    })
  }


  // // error otp get blocked
  // t.Run("fail, otp get blocked", func(t *testing.T) {
  //
  //   req := SendOtpRequestPayload{
  //     Email: tempEmail,
  //   }
  //
  //   // Delete temp account in the database first
  //   err := svc.deleteAuth(context.Background(), req.Email)
  //   require.Nil(t, err)
  //
  //   err = svc.sendOtp(context.Background(), req)
  //   require.Nil(t, err)
  // })
}

func Test_VerifyOtp(t *testing.T) {
  // t.Run("success", func(t *testing.T) {
  //   req := VerifyOtpRequestPayload{
  //     Email: tempEmail,
  //     Otp: tempdata.TempRegisterOtp,
  //   }
  //
  //   _, err := svc.verifyOtp(context.Background(), req)
  //   require.Nil(t, err)
  // })

  t.Run("fail, otp invalid", func(t *testing.T) {
    req := VerifyOtpRequestPayload{
      Email: tempEmail,
      Otp: "2131",
    }

    _, err := svc.verifyOtp(context.Background(), req)
    require.NotNil(t, err)
    require.Equal(t, response.ErrOtpInvalid, err)
  })
}

func Test_Register(t *testing.T) {
  t.Run("fail, name required", func(t *testing.T) {
    req := RegisterRequestPayload{
      Name : "",
      Password : tempPassword,
      PasswordConfirmation : tempPassword,
      PublicIdUserOtp : uuid.New(),
    }

    err := svc.register(context.Background(), req)
    require.NotNil(t, err)
    require.Equal(t, response.ErrNameRequired, err)
  })

  t.Run("fail, password required", func(t *testing.T) {
    req := RegisterRequestPayload{
      Name : "Muhamad Ilham Darmawan",
      Password :"" ,
      PasswordConfirmation : tempPassword,
      PublicIdUserOtp : uuid.New(),
    }

    err := svc.register(context.Background(), req)
    require.NotNil(t, err)
    require.Equal(t, response.ErrPasswordRequired, err)
  })

  t.Run("fail, password confirmation required", func(t *testing.T) {
    req := RegisterRequestPayload{
      Name : "Muhamad Ilham Darmawan",
      Password :  tempPassword, 
      PasswordConfirmation : "",
      PublicIdUserOtp : uuid.New(),
    }

    err := svc.register(context.Background(), req)
    require.NotNil(t, err)
    require.Equal(t, response.ErrPasswordConfirmationRequired, err)
  })

  t.Run("fail, password confirmation required", func(t *testing.T) {
    req := RegisterRequestPayload{
      Name : "Muhamad Ilham Darmawan",
      Password :  tempPassword, 
      PasswordConfirmation : "",
      PublicIdUserOtp : uuid.New(),
    }

    err := svc.register(context.Background(), req)
    require.NotNil(t, err)
    require.Equal(t, response.ErrPasswordConfirmationRequired, err)
  })
}

func Test_Login(t *testing.T) {
  t.Run("success", func(t *testing.T) {
    reqLogin := LoginRequestPayload{
      Email: tempEmail,
      Password : tempPassword,
    }

    _, err := svc.login(context.Background(), reqLogin)
    require.Nil(t, err)
  })

  t.Run("error unauthorized", func(t *testing.T) {
    reqLogin := LoginRequestPayload{
      Email: "wrong@gmail.com",
      Password : tempPassword,
    }

    _, err := svc.login(context.Background(), reqLogin)
    require.NotNil(t, err)
    require.Equal(t, response.ErrUnauthorized, err)
  })

  t.Run("error password not match", func(t *testing.T) {
    reqLogin := LoginRequestPayload{
      Email: tempEmail,
      Password : "password123",
    }

    _, err := svc.login(context.Background(), reqLogin)
    require.NotNil(t, err)
    require.Equal(t, response.ErrPasswordNotMatch, err)
  })
}

