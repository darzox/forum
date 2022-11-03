package storage

import (
	"database/sql"
)

type Comment struct {
	Id     uint32
	Text   string
	PostId uint32
	UserId uint32
}

func NewComment(text string, postId uint32, userId uint32) *Comment {
	return &Comment{
		Text:   text,
		PostId: postId,
		UserId: userId,
	}
}

func createComment(db *sql.DB, comment *Comment) error {
	records := `INSERT INTO comment(text, post_id, user_id)
				VALUES (?,?,?)`
	query, err := db.Prepare(records)
	if err != nil {
		return err
	}
	_, err = query.Exec(comment.Text, comment.PostId, comment.UserId)
	if err != nil {
		return err
	}
	return nil
}
