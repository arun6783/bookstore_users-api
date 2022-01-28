package users

import (
	"fmt"
	"strings"

	"github.com/arun6783/bookstore_users-api/datasources/mysql/users_db"
	"github.com/arun6783/bookstore_users-api/logger"
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
		logger.Error("Error occured when trying to prepare insert user statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreted, user.Password, user.Status)
	if saveErr != nil {

		sqlError, ok := saveErr.(*mysql.MySQLError)
		if !ok {
			logger.Error("Error occured when trying to parse to sqlError object", nil)
			return errors.NewInternalServerError("database error")
		}

		logger.Error("Error occured when trying to execute insert statement", sqlError)
		return errors.NewInternalServerError("database error")
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("Error occured when trying to save and get last inserted id", err)
		return errors.NewInternalServerError("database error")
	}
	user.Id = userId

	return nil

}

func (user *User) Get() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("Error occured when trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreted, &user.Status); err != nil {

		logger.Error("Error occured when trying to scan user", err)

		if strings.Contains(err.Error(), noRowsError) {
			return errors.NewNotFoundError(fmt.Sprintf("User id %d does not exist in database ", user.Id))
		}
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("Error occured when trying to prepare update user statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()
	_, updateError := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if updateError != nil {

		sqlError, ok := updateError.(*mysql.MySQLError)

		if !ok {
			logger.Error("Error occured when trying to parse to sqlError object - update user", nil)
			return errors.NewInternalServerError("database error")
		}

		logger.Error("Error occured when trying to execute update statement - update user", sqlError)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("Error occured when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()
	_, updateError := stmt.Exec(user.Id)
	if updateError != nil {

		sqlError, ok := updateError.(*mysql.MySQLError)
		if !ok {
			logger.Error("Error occured when trying to parse to sqlError object - delete user", nil)
			return errors.NewInternalServerError("database error")
		}

		logger.Error("Error occured when trying to execute delete statement - delete user", sqlError)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) Search(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("Error occured when trying to prepare search user statement", err)
		return nil, errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("Error occured when trying to execute search statement ", err)
		return nil, errors.NewInternalServerError("database error")
	}

	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreted, &user.Status); err != nil {
			logger.Error("Error occured when trying to execute scan row - search user", err)
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewBadResuestError(fmt.Sprintf("No users matching staus %s", status))
	}

	return results, nil

}
