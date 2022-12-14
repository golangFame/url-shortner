package models

import (
	"github.com/hiroBzinga/bun"
	"time"
)

type Url struct {
	bun.BaseModel `bun:"table:urls"`
	ID            int    `bun:"id"`
	Url           string `bun:"urls"`
	ShortenedUrl  string `bun:"short_urls"`

	CreatedAt time.Time `bun:"created_at,nullzero,default:current_timestamp" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bun:"updated_at,nullzero" json:"updatedAt,omitempty"`
	DeletedAt time.Time `bun:"deleted_at,nullzero,soft_delete" json:"deletedAt,omitempty"`
}
