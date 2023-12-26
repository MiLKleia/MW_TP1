package users

import (
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/users"
	"net/http"
)

func GetAllUsers() ([]models.User, error) {
	var err error
	// calling repository
	users, err := repository.GetAllUsers()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving users : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return users, nil
}

func GetUserByUid(id uuid.UUID) (*models.User, error) {
	var err error
	
	user , err := repository.GetUserByUid(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "user not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving users : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return user, err
}


func DeleteUserByUid(id uuid.UUID) (error) {
	err := repository.DeleteUserByUid(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return &models.CustomError{
				Message: "user not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving users : %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}
	return &models.CustomError{
			Message: "No Content",
			Code:    204,
		}
	
	
}


func CreateUser(name string, username string) (*models.User, error) {
	user, err := repository.CreateUser(name, username)

	if err != nil {
		logrus.Errorf("error adding user : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}
	return user, &models.CustomError{
			Message: "Created",
			Code:    201,
		}
	
}

func UpdateUserByUid(id uuid.UUID, name string, username string) (*models.User, error) {
	var err error
	
	user , err := repository.UpdateUserByUid(id, name, username)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "user not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving users : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return user, err
}



