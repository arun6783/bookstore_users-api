package users

import (
	"fmt"
	"strings"

	"github.com/arun6783/bookstore_users-api/datasources/mysql/users_db"
	"github.com/arun6783/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	noRowsError     = "no rows in result set"
	queryInsertUser = "Insert into users(first_name, last_name, email , date_created) values(?, ?, ?,?);"
	queryGetUser    = "select id, first_name, last_name, email, date_created from users WHERE id = ?;"
	queryUpdateUser = "update users set first_name=?,last_name=?,email=? WHERE id =?;"
	queryDeleteUser = "delete from users where id=?;"
)

func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreted)
	if saveErr != nil {

		sqlError, ok := saveErr.(*mysql.MySQLError)
		if !ok {
			return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user %s", sqlError.Error))
		}

		fmt.Println(sqlError.Error)
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user %s", sqlError.Error))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user %s", err.Error()))
	}
	user.Id = userId

	return nil

}

func (user *User) Get() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreted); err != nil {
		println(err.Error())
		if strings.Contains(err.Error(), noRowsError) {
			return errors.NewNotFoundError(fmt.Sprintf("User id %d does not exist in database ", user.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error occured when trying to get data from table %s", err.Error()))
	}
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()
	_, updateError := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if updateError != nil {

		sqlError, ok := updateError.(*mysql.MySQLError)
		if !ok {
			return errors.NewInternalServerError(fmt.Sprintf("error when trying to update user %s", sqlError.Error))
		}

		fmt.Println(sqlError.Error)
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to  update %s", sqlError.Error))
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()
	_, updateError := stmt.Exec(user.Id)
	if updateError != nil {

		sqlError, ok := updateError.(*mysql.MySQLError)
		if !ok {
			return errors.NewInternalServerError(fmt.Sprintf("error when trying to delete user %s", sqlError.Error))
		}

		fmt.Println(sqlError.Error)
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to delete %s", sqlError.Error))
	}

	return nil
}
