package models

import (
	"errors"
	"github.com/upper/db/v4"
	"time"
)

var (
	ErrDuplicateTitle = errors.New("title already exist in database")
	ErrDuplicateVotes = errors.New("you already voted")

	queryTemplate = `
	SELECT COUNT(*) OVER() AS total_records, pq.*, u.name as uname FROM (
	    SELECT p.id, p.title, p.url, p.created_at, p.user_id as uid, COUNT(c.post_id) as comment_count, count(v.post_id) as votes
		FROM posts p 
		LEFT JOIN comments c ON p.id = c.post_id 
	    LEFT JOIN votes v ON p.id = v.post_id
	 	#where#
		GROUP BY p.id
		#orderby#
		) AS pq
	LEFT JOIN users u ON u.id = uid
	#limit#
	`
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
