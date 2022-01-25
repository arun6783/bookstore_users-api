package services

import (
	"github.com/arun6783/bookstore_users-api/domain/users"
	"github.com/arun6783/bookstore_users-api/utils/errors"

	"github.com/arun6783/bookstore_users-api/utils/dateutils"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	if usrErr := user.Validate(); usrErr != nil {
		return nil, usrErr
	}

	user.DateCreted = dateutils.GetNowString()

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	if userId <= 0 {
		return nil, errors.NewBadResuestError("User id should be greater than 0")
	}

	result := &users.User{Id: userId}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}
