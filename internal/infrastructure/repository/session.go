package repository

import "database/sql"

type sessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) *sessionRepository {
	return &sessionRepository{
		db: db,
	}
}

func (sr *sessionRepository) CreateSession(cookie string, userId uint) error {
	records := `INSERT INTO session(user_id, cookie)
				VALUES(?,?)`
	query, err := sr.db.Prepare(records)
	if err != nil {
		return err
	}
	_, err = query.Exec(cookie, userId)
	if err != nil {
		return err
	}
	return nil
}
