package repository

import (
	"database/sql"
	"fmt"
)

type reactionRepository struct {
	db *sql.DB
}

func NewReactionRepository(db *sql.DB) *reactionRepository {
	return &reactionRepository{
		db: db,
	}
}

func (rr *reactionRepository) React(postOrComment string, userId uint, postId uint, positive bool) (uint, error) {
	records := fmt.Sprintf(`REPLACE INTO %s_like(%s_like_id, user_id, post_id, positive)
	VALUES((SELECT post_like_id FROM post_like WHERE user_id = ? AND post_id = ?), ?, ?, ?);`, postOrComment, postOrComment)
	query, err := rr.db.Prepare(records)
	if err != nil {
		return 0, err
	}
	result, err := query.Exec(userId, postId, userId, postId, positive)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return uint(id), err
}
