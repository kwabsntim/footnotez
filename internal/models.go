package internal

import (
	"time"
)

type Article struct {
	ID        int       `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Content   string    `db:"content" json:"content"`
	Slug      string    `db:"slug" json:"slug"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
