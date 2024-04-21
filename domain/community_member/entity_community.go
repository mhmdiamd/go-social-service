package communitymember

import "time"

type Community struct {
	Id                  int    `db:"id"`
	Name                string `db:"name"`
	Description         string `db:"description"`
	Logo                string `db:"logo"`
	ExternalCategories  string `db:"external_categories"`
	CategoryCommunityID int    `db:"category_community_id"`
	FileIdGdrive        string `db:"file_id_gdrive"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
