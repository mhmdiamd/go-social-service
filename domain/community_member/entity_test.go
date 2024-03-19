package communitymember

import (
	"testing"

	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/stretchr/testify/require"
)

func Test_CommunityMemberValidation_Success(t *testing.T)  {
  req := AddCommunityMemberRequestPayload{
    UserPublicId : "7eaf37cb-1408-459c-ba13-94d3dbcff0aa",
  }

  cm := NewCommunityMemberFromAdd(req);
  
  err := cm.Validate();
  require.Nil(t, err)
}

func Test_CommunityMemberValidation_Failed(t *testing.T) {
  req := AddCommunityMemberRequestPayload{
    UserPublicId : "",
  }
  cm := NewCommunityMemberFromAdd(req);


  err := cm.Validate();
  require.NotNil(t, err)
  require.Equal(t, response.ErrIdRequired, err)
}
