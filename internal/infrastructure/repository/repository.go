package repository

import (
	"database/sql"
	"errors"
	"os"

	"forum/internal/storage"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	userRepository
	sessionRepository
	postRepository
	commentRepository
	reactionRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		*NewUserRepository(db),
		*NewSessionRepository(db),
		*NewPostRepository(db),
		*NewCommentRepository(db),
		*NewReactionRepository(db),
	}
}

func RunDb() (*sql.DB, error) {
	if _, err := os.Stat("database.db"); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create("database.db")
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}
	err = storage.CreateTables(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}
