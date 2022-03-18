package models

type User struct {
	ID    string `json:"id,omitempty", bson:"_id,omitempty"`
	Name  string `json:"name,omitempty", bson:"name,omitempty"`
	Email string `json:"email,omitempty", bson:"email,omitempty"`
	Age   int    `json:"age,omitempty", bson:"age,omitempty"`
}

// type Post struct {
// 	ID string
// 	User
// 	Title     string
// 	Content   string
// 	Timestamp string
// }
