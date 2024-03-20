package communitymember

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/domain/auth"
	"github.com/mhmdiamd/go-social-service/external/database"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/mhmdiamd/go-social-service/internal/config"
	tempdata "github.com/mhmdiamd/go-social-service/temp_data"
	"github.com/stretchr/testify/require"
)

var svc service
var communityService service
var authService auth.Service

var tempEmail string
var tempPassword string

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.Db)
	if err != nil {
		panic(err)
	}

	authRepo := auth.NewRepository(db)
	authService = auth.NewService(authRepo)

	tempEmail = "milham0140@gmail.com"
	tempPassword = "password"

	repo := newRepository(db)
	svc = newService(repo)

  fmt.Println(tempdata.TempLastUserPublicId)
}

func Test_SendOtp(t *testing.T) {
	models, _ := authService.Repo.GetOtpByEmail(context.Background(), tempEmail)
	if len(models) >= 3 {
		t.Run("fail to much send email", func(t *testing.T) {

			req := auth.SendOtpRequestPayload{
				Email: tempEmail,
			}

			err := authService.SendOtp(context.Background(), req)
			require.NotNil(t, err)
			require.Equal(t, response.ErrToMuchSendEmail, err)
		})
	} else {

		t.Run("success", func(t *testing.T) {

			req := auth.SendOtpRequestPayload{
				Email: tempEmail,
			}

		   // Then Delete all the otp
      t.Run("success, delete otp", func(t *testing.T) {
			  err := authService.Repo.DeleteOtpByEmail(context.Background(), tempEmail)
        require.Nil(t, err)
      })

		  t.Run("success, delete user auth", func(t *testing.T) {
			  // Delete temp account in the database first
			  err := authService.DeleteAuth(context.Background(), req.Email)
			  require.Nil(t, err)

			  err = authService.SendOtp(context.Background(), req)
			  require.Nil(t, err)
      })


			t.Run("success verify otp", func(t *testing.T) {
				req := auth.VerifyOtpRequestPayload{
					Email: tempEmail,
					Otp:   tempdata.TempRegisterOtp,
				}

				_, err := authService.VerifyOtp(context.Background(), req)
				require.Nil(t, err)
			})

			t.Run("success register user", func(t *testing.T) {
				req := auth.RegisterRequestPayload{
					Name:                 "Member Muhamad Ilham",
					Password:             tempPassword,
					PasswordConfirmation: tempPassword,
					PublicIdUserOtp:      uuid.MustParse(tempdata.TempPublicIdUserOtp),
				}

				err := authService.Register(context.Background(), req)
				require.Nil(t, err)
			})
		})
	}
}

func Test_VerifyOtp(t *testing.T) {
	t.Run("fail, otp invalid", func(t *testing.T) {
		req := auth.VerifyOtpRequestPayload{
			Email: tempEmail,
			Otp:   "2131",
		}

		_, err := authService.VerifyOtp(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrOtpInvalid, err)
	})
}

func Test_AddNewMember(t *testing.T) {

	t.Run("success", func(t *testing.T) {
    req := AddCommunityMemberRequestPayload{
			Role:         CommunityMemberRole_member,
			UserPublicId: tempdata.TempLastUserPublicId,
      CommunityId: 57,
		}

    err := svc.AddMember(context.Background(), req)
    require.Nil(t, err)
	})

	t.Run("fail, member not found", func(t *testing.T) {

	})
}


