package categorycomunity

import "context"

type Repository interface {
  CreateCategoryCommunity(ctx context.Context, model CategoryCommunityEntity) (err error)
}

type service struct {
  repo Repository
}

func newService(repo Repository) service {
  return service{
    repo: repo,
  }
}

func (s service) AddCategoryCommunity(ctx context.Context, req CreateCategoryCommunityRequestPayload) (err error) {
  entity := NewCategoryCommunityEntity(req.Name)

  if err = entity.Validate(); err != nil {
    return
  }

  return s.repo.CreateCategoryCommunity(ctx, entity);
}
