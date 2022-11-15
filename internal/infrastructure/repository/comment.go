package repository

import "database/sql"

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *commentRepository {
	return &commentRepository{
		db: db,
	}
}

func (cr *commentRepository) CreateComment(userId, postId uint, text string) (uint, error) {
	records := `INSERT INTO comment(text, post_id, user_id)
				VALUES (?,?,?)`
	query, err := cr.db.Prepare(records)
	if err != nil {
		return 0, err
	}
	result, err := query.Exec(text, postId, userId)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return uint(id), err
}

func (cr *commentRepository) GetAllCommentsByPostId(postId uint) {
	
}
