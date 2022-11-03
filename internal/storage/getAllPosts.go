package storage

import (
	"database/sql"
	"fmt"
)

func GetAllPosts(db *sql.DB) error {
	records := `SELECT post_id, text, username
				FROM post INNER JOIN user
				ON post.user_id = user.user_id;`
	rows, err := db.Query(records)
	if err != nil {
		return err
	}
	temp := struct {
		postId   uint32
		text     string
		username string
	}{}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&temp.postId, &temp.text, &temp.username)
	}
	fmt.Println(temp)

	records = `SELECT post_id, COUNT(comment_id) AS amount_comments
				FROM comment
				GROUP BY post_id;`

	temp1 := struct {
		postId         uint32
		amountComments uint32
	}{}
	rows, err = db.Query(records)
	if err != nil {
		return err
	}
	for rows.Next() {
		rows.Scan(&temp1.postId, &temp1.amountComments)
	}
	fmt.Println(temp1)

	records = `SELECT post_id, COUNT(post_like_id) AS amount_likes
				FROM post_like
				GROUP BY post_id;`

	temp2 := struct {
		postId      uint32
		amountLikes uint32
	}{}
	rows, err = db.Query(records)
	if err != nil {
		return err
	}
	for rows.Next() {
		rows.Scan(&temp2.postId, &temp2.amountLikes)
	}
	fmt.Println(temp2)

	records = `SELECT t1.text, 
					  t1.username, 
					  t2.amount_comments, 
					  t3.amount_likes
				FROM 
					(SELECT post_id, text, username
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
					  `
	rows, err = db.Query(records)
	if err != nil {
		return err
	}
	temp3 := struct {
		text           string
		amountLikes    uint32
		amountComments uint32
		username       string
	}{}
	for rows.Next() {
		rows.Scan(&temp3.text, &temp3.username, &temp3.amountComments, &temp3.amountLikes)
	}
	fmt.Println(temp3)
	return nil
}
