package model

type Comment struct {
	Text   string
	PostId uint
	UserUd uint
}

type CommentRepresentation struct {
	CommentId uint
	Text      string
	Like      int
	Username  string
}
