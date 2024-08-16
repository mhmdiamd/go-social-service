package tempdata

import "github.com/google/uuid"

var (
	TempRegisterOtp      string
	TempPublicIdUserOtp  uuid.UUID
	TempLastUserPublicId uuid.UUID
	LastCommunityID      int

	// Event
	TempCurrentEventPublicId uuid.UUID
)
