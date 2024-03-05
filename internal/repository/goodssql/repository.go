package goodssql

import (
	"context"
	"database/sql"

	"github.com/itc1205/little-crud/internal/repository"
)

type GoodsSql struct {
	db *sql.DB
}

func New(db *sql.DB) GoodsSql {
	return GoodsSql{db: db}
}

func (r *GoodsSql) Create(ctx context.Context, e *repository.Goods) (*repository.Goods, error) {
	const q = `
		INSERT INTO "GOODS" (project_id, name, description, priority, removed) 
			VALUES ($1, $2, $3, $4, $5)
		RETURNING id, project_id, name, description, priority, removed, created_at
	`

	goods := new(repository.Goods)
	err := r.db.QueryRowContext(ctx, q, e.ProjectID, e.Name, e.Description, e.Priority, e.Removed).Scan(
		&goods.ID,
		&goods.ProjectID,
		&goods.Name,
		&goods.Description,
		&goods.Priority,
		&goods.Removed,
		&goods.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return goods, nil
}

func (r *GoodsSql) Delete(ctx context.Context, id int32) error {
	const q = `
		DELETE FROM "GOODS"
			WHERE id=$1
	`
	_, err := r.db.ExecContext(ctx, q, id)
	return err
}

func (r *GoodsSql) Get(ctx context.Context, id int32) (*repository.Goods, error) {
	const q = `
		SELECT (id, project_id, name, description, priority, removed, created_at) FROM "GOODS"
			WHERE id=$1
	`
	goods := new(repository.Goods)
	err := r.db.QueryRowContext(ctx, q, id).Scan(
		&goods.ID,
		&goods.ProjectID,
		&goods.Name,
		&goods.Description,
		&goods.Priority,
		&goods.Removed,
		&goods.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return goods, nil
}
func (r *GoodsSql) List(ctx context.Context, offset, limit int32) ([]*repository.Goods, error) {
	const q = `
		SELECT (id, project_id, name, description, priority, removed, created_at) 
			FROM "GOODS" 
			OFFSET $1 LIMIT $2
	`
	var list []*repository.Goods
	rows, err := r.db.QueryContext(ctx, q, offset, limit)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		goods := new(repository.Goods)
		err := rows.Scan(
			&goods.ID,
			&goods.ProjectID,
			&goods.Name,
			&goods.Description,
			&goods.Priority,
			&goods.Removed,
			&goods.Created_at,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, goods)
	}
	return list, nil
}
func (r *GoodsSql) Update(ctx context.Context, e *repository.Goods) error {
	const q = `
		UPDATE "GOODS" SET project_id=$2, name=$3, description=$3, priority=$4, removed=$5
			WHERE id=$1
	`
	_, err := r.db.ExecContext(ctx, q, e.ProjectID, e.Name, e.Description, e.Priority, e.Removed)
	return err
}
