package users

import (
	"strings"

	"github.com/arun6783/bookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id         int64  `json:"user_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	DateCreted string `json:"date_created"`
	Status     string `json:"status"`
	Password   string `json:password"`
}

func (user *User) Validate() *errors.RestErr {

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return errors.NewBadResuestError("invalid email address")
	}
	user.Password = strings.TrimSpace(user.Password)

	if user.Password == "" {
		return errors.NewBadResuestError("invalid password")
	}
	return nil
}
