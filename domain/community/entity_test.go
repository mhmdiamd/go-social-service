package community

import (
	"testing"

	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/stretchr/testify/require"
)

type MockPaylod struct {
  Title string
  Payload CreateCommunityRequestPayload
}

func Test_EntityCommunityValidation_Success(t *testing.T) {
  conditionMock := []MockPaylod {
    {
      Title : "success_with_category_id",
      Payload : CreateCommunityRequestPayload {
        Name: "example community",
        CategoryCommunityID: 0,
        ExternalCategories : []string{"Organization"},
      },
    },
    {
      Title : "success_without_category_id",
      Payload : CreateCommunityRequestPayload{
        Name: "example community",
        CategoryCommunityID: 0,
        ExternalCategories : []string{"Organization"},
      },
    },
  }

  for _, dataMock := range conditionMock {
    t.Run(dataMock.Title, func(t *testing.T) {
      entity := NewCommunityFromCreate(dataMock.Payload)

      err := entity.Validate()
      require.Nil(t, err)
    })
  }

}

func Test_EntityCommunityValidation_Fail(t *testing.T) {
  t.Run("fail, name required", func(t *testing.T) {
    req := CreateCommunityRequestPayload{
      Name: "",
      CategoryCommunityID: 1,
    }

    entity := NewCommunityFromCreate(req)

    err := entity.Validate()
    require.NotNil(t, err)
    require.Equal(t, response.ErrNameRequired, err)
  })
  
  t.Run("fail, category community id required", func(t *testing.T) {
    req := CreateCommunityRequestPayload{
      Name: "example name",
      CategoryCommunityID: 0,
      ExternalCategories: []string{},
    }

    entity := NewCommunityFromCreate(req)

    err := entity.Validate()
    require.NotNil(t, err)
    require.Equal(t, response.ErrCommunityCategoryIdRequired, err)
  })
}
