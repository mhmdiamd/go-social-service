package community

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-social-service/infra/response"
)

type repository struct {
	Db *sqlx.DB
}

func NewRepository(db *sqlx.DB) repository {
	return repository{
		Db: db,
	}
}

func (r repository) Begin(ctx context.Context) (tx *sqlx.Tx, err error) {
	tx, err = r.Db.BeginTxx(ctx, &sql.TxOptions{})
	return
}

func (r repository) Commit(ctx context.Context, tx *sqlx.Tx) (err error) {
	return tx.Commit()
}

func (r repository) Rollback(ctx context.Context, tx *sqlx.Tx) (err error) {
	return tx.Rollback()
}

func (r repository) Create(ctx context.Context, entity Community) (id int, err error) {
	query := `
    INSERT INTO community (
      name, description, logo, external_categories, category_community_id, created_at, updated_at, file_id_gdrive
    ) VALUES (
      $1, $2, $3, $4, $5, $6, $7, $8
    ) RETURNING id
  `

	tx := r.Db.MustBegin()

	var newCommunity Community

	row := tx.QueryRowxContext(ctx, query, entity.Name, entity.Description, entity.Logo, entity.ExternalCategories, entity.CategoryCommunityID, entity.CreatedAt, entity.UpdatedAt, entity.FileIdGdrive)
	if err = row.Scan(&newCommunity.Id); err != nil {
		return
	}

	id = newCommunity.Id

	defer tx.Commit()

	return
}

func (r repository) UpdateById(ctx context.Context, entity Community) (err error) {
	query := `
    UPDATE community 
    SET 
      name=:name, 
      description=:description, 
      logo=:logo, 
      external_categories=:external_categories, 
      category_community_id=:category_community_id, 
      file_id_gdrive=:file_id_gdrive
    
    WHERE id=:id
  `

	stmt, err := r.Db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, entity)
	if err != nil {
		return
	}

	return
}

func (r repository) DeleteById(ctx context.Context, id int) (err error) {
	query := `DELETE FROM community WHERE id=$1`

	_, err = r.Db.QueryContext(ctx, query, id)
	if err != nil {
		return
	}

	return
}

func (r repository) DeleteCommunityMemberByIdCommunity(ctx context.Context, communityId int) (err error) {
	query := `DELETE FROM community_members WHERE community_id=$1`

	_, err = r.Db.QueryContext(ctx, query, communityId)
	if err != nil {
		return
	}

	return
}

func (r repository) GetAll(ctx context.Context, communityPagination CommunityPagination) (communities []Community, err error) {
	query := `
    SELECT 
      id, name, description, logo, external_categories, category_community_id 
    FROM community 
    WHERE id>$1
    ORDER BY id ASC
    LIMIT $2
  `

	err = r.Db.SelectContext(ctx, &communities, query, communityPagination.Cursor, communityPagination.Size)
	if err != nil {
		if err == sql.ErrNoRows {
			return []Community{}, nil
		}

		return
	}

	return
}

func (r repository) GetById(ctx context.Context, id int) (model Community, err error) {
	query := `SELECT * FROM community WHERE id=$1`

	err = r.Db.GetContext(ctx, &model, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
		}
	}

	return
}

// Required by admin
func (r repository) GetAllByUserPublicId(ctx context.Context, publicId string) (communities []Community, err error) {
	return
}

func (r repository) UpdateActivatedCommunityMember(ctx context.Context, is_active int, community_id int, publicUserId uuid.UUID) (communityMembers []CommunityMember, err error) {
	query := `
    UPDATE community_member 
    SET is_acitve=$1 
    WHERE community_id=$2 AND user_public_id=$3
  `

	err = r.Db.SelectContext(ctx, &communityMembers, query, is_active, community_id, publicUserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return []CommunityMember{}, nil
		}
	}

	return
}

func (r repository) GetAllCommunityMemberByIdCommunity(ctx context.Context, community_id int) (communityMembers []CommunityMember, err error) {
	query := `
    SELECT 
      user_public_id, community_id, role, nik, photoktp, is_active
    FROM community_members
    WHERE community_id=:community_id
  `

	err = r.Db.SelectContext(ctx, &communityMembers, query, community_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return []CommunityMember{}, nil
		}
	}

	return
}

func (r repository) CreateCommunityMember(ctx context.Context, entity CommunityMember) (err error) {
	query := `
    INSERT INTO community_members (
      community_id, user_public_id, role, photoktp, nik, is_active
    ) VALUES (
      :community_id, :user_public_id, :role, :photoktp, :nik, :is_active
    )
  `

	stmt, err := r.Db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, entity)
	if err != nil {
		return
	}

	defer stmt.Close()

	return
}
