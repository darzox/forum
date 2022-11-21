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
	records := `SELECT t4.comment_id,
t4.text,
t4.user_id,
t4.username,
COALESCE(t5.likes, 0) as comment_likes,
COALESCE(t5.dislikes, 0) as comment_dislikes
FROM 
(SELECT t1.comment_id as comment_id,
t1.text as text,
t1.user_id as user_id,
t2.username as username
FROM
          (SELECT comment_id, text, post_id, user_id 
          FROM comment) as t1
LEFT JOIN (SELECT user_id, username 
          FROM user) AS t2 ON t1.user_id = t2.user_id) AS t4
LEFT JOIN (SELECT comment_id,
SUM(CASE WHEN positive = true THEN 1 ELSE 0 END) AS likes,
SUM(CASE WHEN positive = false THEN 1 ELSE 0 END) AS dislikes
FROM comment_like) AS t5 ON t4.comment_id = t5.comment_id`
}
