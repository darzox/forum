package repository

import (
	"database/sql"

	"forum/internal/storage"
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
	// file, err := os.Create("database.db")
	// if err != nil {
	// 	return nil, err
	// }
	// defer file.Close()
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
