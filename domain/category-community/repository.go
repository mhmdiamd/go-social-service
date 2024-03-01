package categorycomunity

import ( 
	"context"

	"github.com/jmoiron/sqlx"
)

type repository struct {
  db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
  return repository{
    db: db,
  }
}

func (r repository) CreateCategoryCommunity(ctx context.Context, model CategoryCommunityEntity) (err error) {

  query := `
    INSERT INTO category_community (
      name, created_at, updated_at 
    )
    VALUES (
      :name, :created_at, :updated_at
    )
  `

  stmt, err := r.db.PrepareNamedContext(ctx, query)
  
  if err != nil {
    return 
  }

  _, err = stmt.ExecContext(ctx, model)

  if err != nil {
    return
  }

  defer stmt.Close()

  return
}
