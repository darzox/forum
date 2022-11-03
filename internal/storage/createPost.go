package storage

import (
	"database/sql"
)

type Post struct {
	Id       uint32
	TextPost string
	UserId   uint32
}

func NewPost(textPost string, userId uint32) *Post {
	return &Post{
		TextPost: textPost,
		UserId:   userId,
	}
}

func createPost(db *sql.DB, post *Post) error {
	records := `INSERT INTO post(text, user_id)
				VALUES (?,?)`
	query, err := db.Prepare(records)
	if err != nil {
		return err
	}
	_, err = query.Exec(post.TextPost, post.UserId)
	if err != nil {
		return err
	}
	return nil
}
