package communitymember

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-social-service/domain/auth"
	"github.com/mhmdiamd/go-social-service/infra/response"
)

type repository struct {
  Db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
  return repository {
    Db : db,
  }
}

func (r repository) Begin(ctx context.Context) (tx *sqlx.Tx, err error) {
  return r.Db.BeginTxx(ctx, &sql.TxOptions{})
}

func (r repository) Commit(ctx context.Context, tx *sqlx.Tx) (err error) {
  return tx.Commit()
}

func (r repository) Rollback(ctx context.Context, tx *sqlx.Tx) (err error) {
  return tx.Rollback()
}

func (r repository) Create(ctx context.Context, cm CommunityMember) (err error) {
  query := `
    INSERT INTO community_members (
      role, is_active, nik, photoktp, user_public_id, community_id, created_at, updated_at 
    ) VALUES (
      :role, :is_active, :nik, :photoktp, :user_public_id, :community_id, :created_at, :updated_at 
    )
  `

  stmt, err  := r.Db.PrepareNamedContext(ctx, query)
  if err != nil {
    return
  }

  _, err = stmt.ExecContext(ctx, cm)
  if err != nil {
    return
  }

  defer stmt.Close()

  return 
}

func (r repository) Update(ctx context.Context, cm CommunityMember) (err error) {
  query := `
    UPDATE community_members
    SET role=:role, is_active=:is_active, nik=:nik, photo_ktp=:photoktp, created_at=:created_at, updated_at=:updated_at 
    WHERE community_id=:community_id AND user_public_id=:user_public_id 
  `

  stmt, err  := r.Db.PrepareNamedContext(ctx, query)
  if err != nil {
    return
  }

  _, err = stmt.ExecContext(ctx, cm)
  if err != nil {
    return
  }

  defer stmt.Close()

  return
}
  
func (r repository) Delete(ctx context.Context, userId string, communityId int) (err error) {
  query := `
    DELETE FROM community_members
    WHERE community_id=:community_id AND user_public_id=:user_public_id 
  `

  _, err = r.Db.QueryContext(ctx, query, userId, communityId)
  if err != nil {
    return
  }

  return
}

func (r repository) UpdateRole(ctx context.Context, memberId string, role CommunityMemberRole) (err error) {
  query := `
    UPDATE community_members
    SET role=$1
    WHERE id=$2
  `

  _, err = r.Db.QueryContext(ctx, query, role, memberId)
  if err != nil {
    return
  }

  return
}

func (r repository) GetAllByCommunityId(ctx context.Context, communityId int, pagination CommunityMemberPagination) (members CommunityMembers, err error)  {
  query := `
    SELECT id, role, is_active, nik, photoktp, user_public_id, community_id, created_at, updated_at
    FROM community_members
    WHERE community_id=$1
    ORDER BY id ASC
    LIMIT $2
  `

  err = r.Db.SelectContext(ctx, &members, query, communityId, pagination.Size)
  if err != nil {
    if err != sql.ErrNoRows {
      return
    }
  }

  return
}

func (r repository) GetDetailMember(ctx context.Context, userId string, communityId int) (member CommunityMember, err error) {
  query := `
    SELECT id, role, is_active, nik, photoktp, user_public_id, community_id, created_at, updated_at
    FROM community_members
    WHERE community_id=$1 AND user_public_id=$2
  `

  err = r.Db.GetContext(ctx, &member, query, communityId, userId)
  if err != nil {
    if err == sql.ErrNoRows {
      err = response.ErrNotFound
    }
  }

  return
}

func (r repository) UpdateMemberStatus(ctx context.Context, communityId int, userId string, status int) (err error) {
  query := `
    UPDATE community_members
    SET is_active=$1
    WHERE community_id=$2 AND user_public_id=$3
  `

  _, err = r.Db.QueryContext(ctx, query, status, communityId, userId)
  if err != nil {
    return
  }

  return
}
  
func (r repository) GetCommunityById(ctx context.Context, communityId int) (community Community, err error) {
  query := `
    SELECT id, name, description, logo, category_community_id, external_categories, created_at, updated_at, file_id_gdrive
    FROM community
    WHERE id=$1
  `

  err = r.Db.GetContext(ctx, &community, query, communityId)
  if err != nil {
    if err == sql.ErrNoRows {
      err = response.ErrNotFound
    }
  }

  return
}

func (r repository) IsUserAuthExist(ctx context.Context, userId string) (err error) {

  fmt.Println(userId)
  query := `
    SELECT 
      id, public_id, email, password, name, gender, no_tlp, address
    FROM auth 
    WHERE public_id=$1
  `

 var auth auth.AuthEntity

  err = r.Db.GetContext(ctx, &auth, query, userId)
  if err != nil {
    if err == sql.ErrNoRows {
      err = response.ErrNotFound
    }
  }

  return
}

