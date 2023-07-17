package models

import (
	"github.com/upper/db/v4"
	"time"
)

type Post struct {
	Id           int       `db:"id,omitempty"`
	Title        string    `db:"title"`
	Url          string    `db:"url"`
	CreatedAt    time.Time `db:"created_at"`
	UserID       string    `db:"user_id"`
	Votes        int       `db:"votes,omitempty"`
	CommentCount int       `db:"comment_count,omitempty"`
	TotalRecords int       `db:"total_records,omitempty"`
}

type PostModel struct {
	db db.Session
}
