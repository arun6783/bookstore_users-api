package users

type User struct {
	Id         int64  `json:"user_id"`
	FirstName  string `json: "first_name"`
	LastName   string `json: "last_name"`
	Email      string `json: "email"`
	DateCreted string `json: "date_created"`
}
