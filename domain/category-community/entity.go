package categorycomunity

import (
	"time"

	"github.com/mhmdiamd/go-social-service/infra/response"
)

type CategoryCommunityEntity struct {
  Name string `db:"name"`
  CreatedAt time.Time `db:"created_at"`
  UpdatedAt time.Time `db:"updated_at"`
}

func NewCategoryCommunityEntity(name string) CategoryCommunityEntity {
  return CategoryCommunityEntity {
    Name : name,
    CreatedAt : time.Now(),
    UpdatedAt : time.Now(),
  }
}

func (c *CategoryCommunityEntity) Validate() (err error) {

  if err = c.ValidateName(); err != nil {
    return
  }

  return 
}

func (c *CategoryCommunityEntity) ValidateName() (err error) {
  if c.Name == "" {
    return response.ErrNameRequired
  }

  return
}
