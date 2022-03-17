package models

type User struct {
	ID       string
	Username string
	Email    string
	Password string
}

type Post struct {
	ID string
	User
	Title     string
	Content   string
	Timestamp string
}
