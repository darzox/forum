package repository

import (
	"database/sql"

	"forum/internal/model"
)

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
	_, err = query.Exec(userId, cookie)
	if err != nil {
		return err
	}
	return nil
}

func (sr *sessionRepository) RetrieveSession(cookie string) (string, error) {
	records := `SELECT cookie
				FROM session
				WHERE cookie = ?
				LIMIT 1`
	query, err := sr.db.Prepare(records)
	if err != nil {
		return "", err
	}
	rows, err := query.Query(cookie)
	if err != nil {
		return "", err
	}
	var tempCookie string
	for rows.Next() {
		rows.Scan(&tempCookie)
	}
	return tempCookie, nil
}

func (sr *sessionRepository) RetrieveUserBySession(cookie string) (*model.User, error) {
	records := `SELECT user_id, email, username
				FROM user
				WHERE user_id = (
					SELECT user_id
					FROM session
					WHERE cookie = ?)`
	query, err := sr.db.Prepare(records)
	if err != nil {
		return nil, err
	}

	rows, err := query.Query(cookie)
	if err != nil {
		return nil, err
	}
	var tempUser model.User
	for rows.Next() {
		rows.Scan(&tempUser.ID, &tempUser.Email, &tempUser.Username)
	}
	return &tempUser, nil
}

func (sr *sessionRepository) DeleteSessionByUserId(userId uint) error {
	records := `DELETE FROM session
				WHERE user_id = ?`
	query, err := sr.db.Prepare(records)
	if err != nil {
		return err
	}
	
	_, err = query.Exec(userId)
	if err != nil {
		return err
	}
	return nil
}
