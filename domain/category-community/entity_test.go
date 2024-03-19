package categorycomunity

import (
	"testing"

	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/stretchr/testify/require"
)


func Test_ValidateCategoryCommunity(t *testing.T) {
  t.Run("success" , func(t *testing.T) {
    entity := NewCategoryCommunityEntity("financial");

    err := entity.Validate();
    require.Nil(t, err)
  })

  t.Run("fail, name is required", func(t *testing.T) {
    entity := NewCategoryCommunityEntity("");

    err := entity.Validate();
    require.NotNil(t, err)
    require.Equal(t, response.ErrNameRequired, err)
  })
}
