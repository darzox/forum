package storage

import (
	"database/sql"
	"fmt"
)

func createTables(db *sql.DB) error {
	usersTable := `
	CREATE TABLE user (
	    user_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	    email TEXT,
		username TEXT,
	    password TEXT
	);`
	postTable := `
	CREATE TABLE post (
		post_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		text TEXT,
		user_id INTEGER REFERENCES user(user_id)
	);`
	commentTable := `
	CREATE TABLE comment (
		comment_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		text TEXT,
		post_id INTEGER REFERENCES post(post_id),
		user_id INTEGER REFERENCES user(user_id)
	);`
	postLikeTable := `
	CREATE TABLE post_like (
		post_like_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		user_id INTEGER REFERENCES user(user_id),
		post_id INTEGER REFERENCES post(post_id)
	);`
	commentLikeTable := `
	CREATE TABLE comment_like (
		comment_like_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		user_id INTEGER REFERENCES user(user_id),
		comment_id INTEGER REFERENCES comment(comment_id)
	);`
	categoryTable := `
		CREATE TABLE category (
			category_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			category_name TEXT
		);`
	postCategoryTable := `
	CREATE TABLE post_category (
		post_category_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		post_id INTEGER REFERENCES post(post_id),
		category_id INTEGER REFERENCES category(category_id)

	);`
	sessionTable := `
	CREATE TABLE session (
		session_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		user_id INTEGER REFERENCES user(user_id),
		cookie TEXT
	);`
	allTables := []string{usersTable, postTable, commentTable, postLikeTable, commentLikeTable, categoryTable, postCategoryTable, sessionTable}
	for _, table := range allTables {
		query, err := db.Prepare(table)
		if err != nil {
			fmt.Println("aaa")
			return err
		}
		_, err = query.Exec()
		if err != nil {
			return err
		}
	}
	fmt.Println("Tables created successfully!")
	return nil
}
