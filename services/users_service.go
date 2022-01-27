package services

import (
	"github.com/arun6783/bookstore_users-api/domain/users"
	date_utils "github.com/arun6783/bookstore_users-api/utils"
	"github.com/arun6783/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	if usrErr := user.Validate(); usrErr != nil {
		return nil, usrErr
	}

	user.DateCreted = date_utils.GetNowDBFormat()

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

func UpdateUser(isPartialUpdate bool, user users.User) (*users.User, *errors.RestErr) {

	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if isPartialUpdate {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	if updateErr := current.Update(); updateErr != nil {
		return nil, updateErr
	}
	return current, nil
}

func Delete(userId int64) *errors.RestErr {
	if userId <= 0 {
		return errors.NewBadResuestError("User id should be greater than 0")
	}

	current := &users.User{Id: userId}
	return current.Delete()
}
