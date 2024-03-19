package community

import "time"

type CommunityListResponse struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Description string `json:"description"`
  Logo string `json:"logo"`
  ExternalCategories []string `json:"external_categories"`
  CategoryCommunityId int `json:"category_community_id"`
}


type CommunityDetailResponse struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Description string `json:"description"`
  Logo string `json:"logo"`
  ExternalCategories []string `json:"external_categories"`
  CategoryCommunityId int `json:"category_community_id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}

func NewCommunityListResponseFromEntity(communities []Community) []CommunityListResponse {
  var communityList []CommunityListResponse

  for _, community := range communities {
    communityList = append(communityList, community.ToCommunityListResponse())
  }

  return communityList
}
