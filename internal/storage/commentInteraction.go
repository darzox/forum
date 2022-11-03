package storage

import "database/sql"

func CommentLike(db *sql.DB, user_id uint32, comment_id uint32) error {
	records := `INSERT INTO comment_like(user_id, comment_id)
				VALUES (?,?)`
	query, err := db.Prepare(records)
	if err != nil {
		return err
	}
	_, err = query.Exec(user_id, comment_id)
	if err != nil {
		return err
	}
	return nil
}

func CommentDisLike(db *sql.DB, commentLikeId uint32) error {
	records := `DELETE FROM comment_like WHERE comment_like_id=?`
	query, err := db.Prepare(records)
	if err != nil {
		return err
	}
	_, err = query.Exec(commentLikeId)
	if err != nil {
		return err
	}
	return nil
}
