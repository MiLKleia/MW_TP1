package users

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"time"
)


func UpdateUserByUid(uid uuid.UUID, name string, username string) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	_ , err2 := db.Exec("UPDATE users SET name = ?, username = ? WHERE uid = ?", name, username, uid.String())
	
	
	if err2 != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM users WHERE uid=?", uid.String())
	helpers.CloseDB(db)

	var data models.User
	err = row.Scan(&data.Uid, &data.Name, &data.Username, &data.Date)
	if err != nil {
		return nil, err
	}
	return &data, err
	
}

func CreateUser(name string, username string) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}


	// does the username already exist
	row_test := db.QueryRow("SELECT * FROM users WHERE username=?", username)
	var data_test_exist models.User
	err = row_test.Scan(&data_test_exist.Uid, &data_test_exist.Name, &data_test_exist.Username, &data_test_exist.Date)
	if err == nil {
		return nil, &models.CustomError{
			Message: "Username exists already",
			Code:    409,
		}
	}

	// TODO boucle tant que pas nouveau UID

	new_uid, _ := uuid.NewV4()

	dt := time.Now()
	_ , err2 := db.Exec("INSERT INTO users (uid, name, username, inscription_date) VALUES(? ,? , ?, ?)", new_uid.String(), name, username, dt.Format("2006-1-2"))

	if err2 != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM users WHERE uid=?", new_uid.String())
	helpers.CloseDB(db)

	var data models.User
	err = row.Scan(&data.Uid, &data.Name, &data.Username, &data.Date)
	if err != nil {
		return nil, err
	}
	return &data, err
	
}
	

func DeleteUserByUid(uid uuid.UUID) (error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	_ , err2 := db.Exec("DELETE FROM users WHERE uid=?", uid.String())
	helpers.CloseDB(db)

	return err2
}

func GetAllUsers() ([]models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM users")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	users := []models.User{}
	for rows.Next() {
		var data models.User
		err = rows.Scan(&data.Uid, &data.Name, &data.Username, &data.Date)
		if err != nil {
			return nil, err
		}
		users = append(users, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return users, err
}

func GetUserByUid(uid uuid.UUID) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM users WHERE uid=?", uid.String())
	helpers.CloseDB(db)

	var data models.User
	err = row.Scan(&data.Uid, &data.Name, &data.Username, &data.Date)
	if err != nil {
		return nil, err
	}
	return &data, err
}




