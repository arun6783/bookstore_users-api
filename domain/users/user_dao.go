package users

import (
	"fmt"
	"strings"

	"github.com/arun6783/bookstore_users-api/datasources/mysql/users_db"
	"github.com/arun6783/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	noRowsError           = "no rows in result set"
	queryInsertUser       = "Insert into users(first_name, last_name, email , date_created, password, status) values(?, ?, ?,?,?,?);"
	queryGetUser          = "select id, first_name, last_name, email, date_created, status from users WHERE id = ?;"
	queryUpdateUser       = "update users set first_name=?,last_name=?,email=? WHERE id =?;"
	queryDeleteUser       = "delete from users where id=?;"
	queryFindUserByStatus = "select id , first_name, last_name, email, date_created, status from users where status=?;"
)

func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreted, user.Password, user.Status)
	if saveErr != nil {

		sqlError, ok := saveErr.(*mysql.MySQLError)
		if !ok {
			return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user %s", sqlError.Message))
		}

		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user %s", sqlError.Message))
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

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreted, &user.Status); err != nil {
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
			return errors.NewInternalServerError(fmt.Sprintf("error when trying to update user %s", sqlError.Message))
		}

		return errors.NewInternalServerError(fmt.Sprintf("error when trying to  update %s", sqlError.Message))
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
			return errors.NewInternalServerError(fmt.Sprintf("error when trying to delete user %s", sqlError.Message))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to delete %s", sqlError.Message))
	}

	return nil
}

func (user *User) Search(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreted, &user.Status); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewBadResuestError(fmt.Sprintf("No users matching staus %s", status))
	}

	return results, nil

}
