package users

import (
	"fmt"

	"github.com/arun6783/bookstore_users-api/datasources/mysql/users_db"
	"github.com/arun6783/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = "Insert into users(first_name, last_name, email , date_created) values(?, ?, ?,?);"
)

func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreted)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user %s", err.Error()))
	}
	user.Id = userId

	return nil

}

func (user *User) Get() *errors.RestErr {

	println(fmt.Sprintf("number of connections is %v", users_db.Client.Stats()))

	return nil
}
