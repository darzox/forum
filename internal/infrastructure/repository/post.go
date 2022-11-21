package repository

import (
	"database/sql"

	"forum/internal/model"
)

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}

func (pr *postRepository) GetAllPosts() ([]model.PostRepresentation, error) {
	records := `SELECT t4.post_id,
						t4.text,
						t4.username,
						t4.heading,
						t4.comments,
						t5.likes,
						t5.dislikes
				FROM (SELECT t1.post_id AS post_id, 
									t1.text as text, 
									t1.username as username,
									t1.heading as heading, 
									coalesce(t2.amount_comments, 0) AS comments
							FROM 
							(SELECT post_id, text, username, heading 
							FROM post INNER JOIN user
							ON post.user_id = user.user_id) AS t1
								LEFT JOIN 
							(SELECT post_id, COUNT(comment_id) AS amount_comments
							FROM comment
							GROUP BY post_id) AS t2 
								ON t1.post_id = t2.post_id) AS t4
								LEFT JOIN (SELECT post_id,
				SUM(CASE WHEN positive = true THEN 1 ELSE 0 END) AS likes,
				SUM(CASE WHEN positive = false THEN 1 ELSE 0 END) AS dislikes
				FROM post_like
				GROUP BY post_id) AS t5 ON t4.post_id = t5.post_id`
	rows, err := pr.db.Query(records)
	if err != nil {
		return nil, err
	}
	var tempPost model.PostRepresentation
	var allPosts []model.PostRepresentation
	for rows.Next() {
		rows.Scan(&tempPost.PostId, &tempPost.Text, &tempPost.Username, &tempPost.Heading, &tempPost.AmountComments, &tempPost.AmountLikes, &tempPost.AmountDisLikes)
		allPosts = append(allPosts, tempPost)
	}
	return allPosts, nil
}

func (pr *postRepository) CreatePost(heading string, text string, userId uint) (uint, error) {
	records := `INSERT INTO post(heading, text, user_id)
				VALUES (?,?,?)`
	query, err := pr.db.Prepare(records)
	if err != nil {
		return 0, err
	}
	result, err := query.Exec(heading, text, userId)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return uint(id), err
}

func (pr *postRepository) GetPostById(postId uint) (*model.PostRepresentation, error) {
	records := `SELECT t4.post_id,
	t4.text,
	t4.username,
	t4.heading,
	t4.comments,
	t5.likes,
	t5.dislikes
FROM (SELECT t1.post_id AS post_id, 
				t1.text as text, 
				t1.username as username,
				t1.heading as heading, 
				coalesce(t2.amount_comments, 0) AS comments
		FROM 
		(SELECT post_id, text, username, heading 
		FROM post INNER JOIN user
		ON post.user_id = user.user_id) AS t1
			LEFT JOIN 
		(SELECT post_id, COUNT(comment_id) AS amount_comments
		FROM comment
		GROUP BY post_id) AS t2 
			ON t1.post_id = t2.post_id) AS t4
			LEFT JOIN (SELECT post_id,
SUM(CASE WHEN positive = true THEN 1 ELSE 0 END) AS likes,
SUM(CASE WHEN positive = false THEN 1 ELSE 0 END) AS dislikes
FROM post_like
GROUP BY post_id) AS t5 ON t4.post_id = t5.post_id
				WHERE t4.post_id = ?`
	rows, err := pr.db.Query(records, postId)
	if err != nil {
		return nil, err
	}
	var tempPost model.PostRepresentation
	for rows.Next() {
		rows.Scan(&tempPost.PostId, &tempPost.Text, &tempPost.Username, &tempPost.Heading, &tempPost.AmountComments, &tempPost.AmountLikes, &tempPost.AmountDisLikes)
	}
	return &tempPost, nil
}
