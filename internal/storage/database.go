package storage

import (
	_ "github.com/mattn/go-sqlite3"
)

func RunDb() {
	// fmt.Println("aaa")
	// file, err := os.Create("database.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// file.Close()
	// --------- Trying-----

	// db, err := sql.Open("sqlite3", "database.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// err = createTables(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// user := NewUser("askhajan666@gmail.com", "darzox", "123")

	// err = createUser(db, user)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// post := NewPost("Salam alekum", 1)
	// err = createPost(db, post)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// post = NewPost("Salam alekum2", 2)
	// err = createPost(db, post)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// comment := NewComment("ualekum assalam", 1, 1)

	// err = createComment(db, comment)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = PostLike(db, 1, 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = PostLike(db, 1, 2)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = PostLike(db, 2, 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = PostDisLike(db, 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = CommentLike(db, 1, 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = CommentDisLike(db, 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = GetAllPosts(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = GetMostLikedPosts(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// --------- Trying-----
}
