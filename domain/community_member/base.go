package communitymember

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	infrafiber "github.com/mhmdiamd/go-social-service/infra/fiber"
)

func Init(router fiber.Router, db *sqlx.DB) {
  repo := newRepository(db)
  svc := newService(repo)
  handler := newHandler(svc)

  communityMemberRouter := router.Group("community-member")
  {
    communityMemberRouter.Get("/:community_id/members", handler.GetAllMemberByCommunityId)
    communityMemberRouter.Post("", handler.CreateNewMember)
    communityMemberRouter.Put("/:community_id/members/:member_id", infrafiber.CheckAuth(), handler.UpdateMemberById)
    communityMemberRouter.Delete("/:community_id/members/:member_id", handler.DeleteMemberById)
    
  // Authorization middleware
    communityMemberRouter.Put("/:community_id/members/:member_id/kick", infrafiber.CheckAuth(), handler.KickMemberById)
  }
}
