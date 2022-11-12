package repository

import (
	"database/sql"
	"fmt"

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
	records := `SELECT t1.post_id  
					  t1.text, 
					  t1.username,
					  t1.heading 
					  t2.amount_comments, 
					  t3.amount_likes
				FROM 
					(SELECT post_id, text, username, heading 
					 FROM post INNER JOIN user
					 ON post.user_id = user.user_id) AS t1
					 	INNER JOIN 
					(SELECT post_id, COUNT(comment_id) AS amount_comments
					 FROM comment
					 GROUP BY post_id) AS t2 
					 	ON t1.post_id = t2.post_id
						INNER JOIN
					(SELECT post_id, COUNT(post_like_id) AS amount_likes
					 FROM post_like
					 GROUP BY post_id) AS t3 
					 	ON t2.post_id = t3.post_id;
				ORDER BY t1.post_id DESC`
	rows, err := pr.db.Query(records)
	if err != nil {
		return nil, err
	}
	var tempPost model.PostRepresentation
	var allPosts []model.PostRepresentation
	for rows.Next() {
		rows.Scan(&tempPost.PostId, &tempPost.Text, &tempPost.Username, &tempPost.Heading, &tempPost.AmountComments, &tempPost.AmountLikes)
		allPosts = append(allPosts, tempPost)
	}
	fmt.Println(allPosts)
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
	records := `SELECT post_id, text, username, heading 
				FROM post INNER JOIN user
				ON post.user_id = user.user_id
				WHERE post_id = ?`
	rows, err := pr.db.Query(records, postId)
	if err != nil {
		return nil, err
	}
	var tempPost model.PostRepresentation
	for rows.Next() {
		rows.Scan(&tempPost.PostId, &tempPost.Text, &tempPost.Username, &tempPost.Heading)
	}
	fmt.Println(tempPost)
	records = `SELECT (SELECT count(positive)
				FROM post_like
				WHERE positive = true AND post_id= ?) - (SELECT count(positive) FROM post_like WHERE positive = false AND post_id= ?)`
	rows, err = pr.db.Query(records, postId, postId)
	if err != nil {
		return nil, err
	}
	var CountLikes int
	for rows.Next() {
		rows.Scan(&CountLikes)
	}
	tempPost.AmountLikes = CountLikes
	return &tempPost, nil
}
