package users

import "encoding/json"

type PublicUser struct {
	Id         int64  `json:"id"`
	DateCreted string `json:"date_created"`
	Status     string `json:"status"`
}

type PrivateUser struct {
	Id         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	DateCreted string `json:"date_created"`
	Status     string `json:"status"`
}

func (users Users) Marshal(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshal(isPublic)
	}
	return result
}

func (user *User) Marshal(isPublic bool) interface{} {

	if isPublic {
		return PublicUser{

			Id:         user.Id,
			DateCreted: user.DateCreted,
			Status:     user.Status,
		}
	}

	jsonUser, _ := json.Marshal((user))
	var privateUser PrivateUser
	json.Unmarshal(jsonUser, &privateUser)
	return privateUser
}
