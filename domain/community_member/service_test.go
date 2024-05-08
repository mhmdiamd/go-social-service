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

      t.Run("success, send otp to email", func (t *testing.T) {
        err := authService.SendOtp(context.Background(), req)
        require.Nil(t, err)
      })

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
      CommunityId: 65,
		}

    err := svc.AddMember(context.Background(), req)
    require.Nil(t, err)
	})

	t.Run("fail, user not found", func(t *testing.T) {
    req := AddCommunityMemberRequestPayload{
			Role:         CommunityMemberRole_member,
			UserPublicId: uuid.NewString(),
      CommunityId: 65,
		}

    err := svc.AddMember(context.Background(), req)
    require.NotNil(t, err)
    require.Equal(t, response.ErrNotFound, err)
	})
}

func Test_GetAllMemberByCommunityId(t *testing.T) {
	t.Run("success", func(t *testing.T) {
    pagination := CommunityMemberListRequestPayload{}
    members, err := svc.GetAllMemberByCommunityId(context.Background(), 65 , pagination.GenerateDefaultValue())
    require.Nil(t, err)
    require.NotNil(t, members)
	})
}

func Test_UpdateMember(t *testing.T) {
  type Mock struct {
    Status string
    Payload UpdateCommunityMemberRequestPayload
    Err error
  }

  // Get detail user 
  member, err := svc.repo.GetDetailMember(context.Background(), tempdata.TempLastUserPublicId, 65)
  if err != nil {
    panic(err)
  }

  var data = map[string]Mock{
    "update community member" : {
      Status: "success",
      Payload: UpdateCommunityMemberRequestPayload {
        CommunityId: 65,
        UserPublicId: member.UserPublicId,
        Nik: member.Nik,
        Role : member.Role,
        IsActive: member.IsActive,
        PhotoKTP: member.PhotoKTP,
      },
      Err : nil,
    },

    "member not found" : {
      Status: "fail",
      Payload: UpdateCommunityMemberRequestPayload {
        CommunityId: 65,
        UserPublicId: uuid.NewString(),
        Nik: member.Nik,
        Role : member.Role,
        IsActive: member.IsActive,
        PhotoKTP: member.PhotoKTP,
      },
      Err : response.ErrNotFound,
    },

    "community not found" : {
      Status: "fail",
      Payload: UpdateCommunityMemberRequestPayload {
        CommunityId: 9999,
        UserPublicId: uuid.NewString(),
        Nik: member.Nik,
        Role : member.Role,
        IsActive: member.IsActive,
        PhotoKTP: member.PhotoKTP,
      },
      Err : response.ErrNotFound,
    },
  }

  for i, m := range data {
    t.Run(fmt.Sprintf("%s, %s", m.Status, i), func (t *testing.T) {
      err := svc.Update(context.Background(), m.Payload);
      
      if m.Err != nil {
        require.NotNil(t, err)
        require.Equal(t, m.Err , err)
      }else {
        require.Nil(t, err)
      }
    })
  }
}

func Test_KickMember(t *testing.T) {

  type Req struct {
    EditorId string 
    MemberId string 
    CommunityId int
  }

  type Mock struct {
    Status string
    Payload Req
    Err error
  }

  // Get detail user 
  editor, err := svc.repo.GetDetailMember(context.Background(), "4631d7f4-ebe6-4065-9d42-a6b89aa639ad", 65)
  if err != nil {
    panic(err)
  }

  var data = map[string]Mock{
    "kick community member" : {
      Status: "success",
      Payload: Req {
        CommunityId: 65,
        MemberId : tempdata.TempLastUserPublicId,
        EditorId : editor.UserPublicId,
      },
      Err : nil,
    },

    "not permitted" : {
      Status: "fail",
      Payload: Req {
        CommunityId: 65,
        MemberId : uuid.NewString(),
        EditorId : tempdata.TempLastUserPublicId,
      },
      Err : response.ErrCommunityMemberRoleNotPermitted,
    },

    "member not found" : {
      Status: "fail",
      Payload: Req {
        CommunityId: 65,
        MemberId : uuid.NewString(),
        EditorId : editor.UserPublicId,
      },
      Err : response.ErrNotFound,
    },

    "community not found" : {
      Status: "fail",
      Payload: Req {
        CommunityId: 999,
        MemberId : tempdata.TempLastUserPublicId,
        EditorId : editor.UserPublicId,
      },
      Err : response.ErrNotFound,
    },
  }

  for i, m := range data {
    t.Run(fmt.Sprintf("%s, %s", m.Status, i), func (t *testing.T) {
      err := svc.KickMember(context.Background(), m.Payload.EditorId, m.Payload.MemberId, m.Payload.CommunityId);
      
      if m.Err != nil {
        require.NotNil(t, err)
        require.Equal(t, m.Err , err)
      }else {
        require.Nil(t, err)
      }
    })
  }
}

func Test_DeleteMember(t *testing.T) {
  t.Run("success, delete member", func (t *testing.T) {
    err := svc.DeleteCommunityMember(context.Background(), tempdata.TempLastUserPublicId, 65)
    require.Nil(t, err)
  })

  t.Run("fail, member not found", func(t *testing.T) {
    err := svc.DeleteCommunityMember(context.Background(), tempdata.TempLastUserPublicId, 65)
    require.NotNil(t, err)
    require.Equal(t, response.ErrNotFound, err)
  })
}

