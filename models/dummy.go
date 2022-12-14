package models

import (
	"context"
	"github.com/hiroBzinga/bun"
	"time"
)

type Dummy struct {
	bun.BaseModel `bun:"table:dummy,select:v_dummy"`
	ID            int `bun:"id"`

	CreatedAt time.Time `bun:"created_at,nullzero,default:current_timestamp" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bun:"updated_at,nullzero" json:"updatedAt,omitempty"`
	DeletedAt time.Time `bun:"deleted_at,nullzero,soft_delete" json:"deletedAt,omitempty"`
}

func (d *Dummy) Fetch(db *bun.DB, ctx context.Context) (err error) {
	query := db.NewSelect().Model(d)
	query.Where("id=?", d.ID)
	err = query.Scan(ctx)
	return
}

func (d *Dummy) FetchAll(db *bun.DB, ctx context.Context) (dummies []Dummy, err error) {
	query := db.NewSelect().Model(&dummies)
	err = query.Scan(ctx)
	return
}
