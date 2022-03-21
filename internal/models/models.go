package models

type User struct {
	// ID       string `json:"_id, omitempty", bson:"_id, omitempty`
	Nickname string `json:"nickname", bson:"nickname,omitempty"`
	Email    string `json:"email", bson:"email,omitempty"`
	Name     string `json:"name" bson:"name,omitempty"`
	Age      int    `json:"age", bson:"age,omitempty"`
}

// type Post struct {
// 	ID string
// 	User
// 	Title     string
// 	Content   string
// 	Timestamp string
// }
