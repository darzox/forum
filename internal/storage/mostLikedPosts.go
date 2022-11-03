package storage

import (
	"database/sql"
	"fmt"
)

func GetMostLikedPosts(db *sql.DB) error {
	records := `SELECT post_id, COUNT(post_like_id) AS c
				FROM post_like
				GROUP BY post_id
				ORDER BY c DESC;`
	rows, err := db.Query(records)
	if err != nil {
		return err
	}

	type t1 struct {
		postId uint32
		count  uint32
	}

	var posts []t1

	for rows.Next() {
		var temp1 t1
		rows.Scan(&temp1.postId, &temp1.count)
		posts = append(posts, temp1)
	}
	fmt.Println(posts)
	return nil
}
