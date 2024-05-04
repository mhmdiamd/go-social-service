package communitymember

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
  repo := newRepository(db)
  svc := newService(repo)
  handler := newHandler(svc)

  communityMemberRouter := router.Group("/community-member")
  {
    communityMemberRouter.Get("/:community_id/members", handler.GetAllMemberByCommunityId)
    communityMemberRouter.Get("/:community_id/members/:member_id", handler.GetDetailMemberById)

    communityMemberRouter.Post("", handler.CreateNewMember)
    communityMemberRouter.Put("/:community_id/members/:member_id", handler.UpdateMemberById)
    communityMemberRouter.Delete("/:community_id/members/:member_id", handler.DeleteMemberById)
    
    communityMemberRouter.Post("/:community_id/members/:member_id/kick", handler.KickMemberById)
  }
}
