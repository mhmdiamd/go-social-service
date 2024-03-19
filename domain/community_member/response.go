package communitymember

import "time"

type CommunityMemberResponse struct {
  Id int `json:"id"`
  Role CommunityMemberRole `json:"role"`
  Nik string `json:"nik"`
  IsActive int `json:"is_active"`
  PhotoKTP int `json:"photo_ktp"`

  // user public id foreign key
  UserPublicId string `json:"user_public_id"`

  // community id 
  CommunityId int `json:"community_id"`

  JoinAt int `json:"join_at"`

  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  
}

func (c *CommunityMemberResponse) SetJoinAt() {
  c.JoinAt = 10
}

func NewCommunityMemberListFromEntity(communities CommunityMembers) []CommunityMemberResponse {
  var communitiesResponse []CommunityMemberResponse

  for _ , community := range communities {
    communityMember := community.ToCommunityMemberResponse()

    // Set Join At
    communityMember.SetJoinAt()
    communitiesResponse = append(communitiesResponse, communityMember)
  }

  return communitiesResponse
}


