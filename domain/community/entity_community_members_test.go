package community

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_CommunityMembersValidation(t *testing.T) {

  t.Run("success", func(t *testing.T) {
    req := CreateCommunityMembersRequestPayload{
      CommunityId: 1,
      UserPublicId: uuid.New(),
    }

    communityMemberEntity := NewCommunityMembersFromCreate(req)
    err := communityMemberEntity.Validate()

    require.Nil(t, err)
  })

  t.Run("fail, id required", func(t *testing.T) {
    req := CreateCommunityMembersRequestPayload{
      UserPublicId: uuid.New(),
    }

    communityMemberEntity := NewCommunityMembersFromCreate(req)
    err := communityMemberEntity.Validate()

    require.NotNil(t, err)
  })
}

