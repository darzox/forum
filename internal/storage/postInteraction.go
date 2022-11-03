package storage

import "database/sql"

func PostLike(db *sql.DB, user_id uint32, post_id uint32) error {
	records := `INSERT INTO post_like(user_id, post_id)
				VALUES (?,?)`
	query, err := db.Prepare(records)
	if err != nil {
		return err
	}
	_, err = query.Exec(user_id, post_id)
	if err != nil {
		return err
	}
	return nil
}

func PostDisLike(db *sql.DB, postLikeId uint32) error {
	records := `DELETE FROM post_like WHERE post_like_id=?`
	query, err := db.Prepare(records)
	if err != nil {
		return err
	}
	_, err = query.Exec(postLikeId)
	if err != nil {
		return err
	}
	return nil
}
