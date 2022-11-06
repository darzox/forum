package repository

import "database/sql"

type Repository struct {
	userRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		*NewUserRepository(db),
	}
}
