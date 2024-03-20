package communitymember

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/stretchr/testify/require"
)

func Test_CommunityMemberValidation_Success(t *testing.T)  {
  req := AddCommunityMemberRequestPayload{
    UserPublicId : uuid.NewString(),
    CommunityId : 123,
  }

  cm := NewCommunityMemberFromAdd(req);
  
  err := cm.Validate();
  require.Nil(t, err)
}

func Test_CommunityMemberValidation_Failed(t *testing.T) {
  t.Run("failed, user id is required", func (t *testing.T) {
    req := AddCommunityMemberRequestPayload{
      UserPublicId : "",
    }
    cm := NewCommunityMemberFromAdd(req);

    err := cm.Validate();
    require.NotNil(t, err)
    require.Equal(t, response.ErrUserPublicIdRequired, err)
  }) 

  t.Run("failed, community id is required", func (t *testing.T) {
    req := AddCommunityMemberRequestPayload{
      UserPublicId : uuid.NewString(),
    }
    cm := NewCommunityMemberFromAdd(req);

    err := cm.Validate();
    require.NotNil(t, err)
    require.Equal(t, response.ErrCommunityIdRequired, err)
  }) 
}
