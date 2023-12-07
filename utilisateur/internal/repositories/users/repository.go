package users

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)


func UpdateUserByUid(uid uuid.UUID, name string, surname string, alias string) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	_ , err2 := db.Exec("UPDATE users SET name = ?, surname = ?, alias = ? WHERE uid = ?", name, surname, alias, uid.String())
	
	
	if err2 != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM users WHERE uid=?", uid.String())
	helpers.CloseDB(db)

	var data models.User
	err = row.Scan(&data.Uid, &data.Name, &data.Surname, &data.Alias)
	if err != nil {
		return nil, err
	}
	return &data, err
	
}

func CreateUser(name string, surname string, alias string) (error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	// TODO boucle tant que pas nouveau UID

	new_uid, _ := uuid.NewV4()

	
	_ , err2 := db.Exec("INSERT INTO users (uid, name, surname, alias) VALUES(? ,? , ?, ?)", new_uid.String(), name, surname, alias)

	helpers.CloseDB(db)

	
	return err2
	
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
		err = rows.Scan(&data.Uid, &data.Name, &data.Surname, &data.Alias)
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
	err = row.Scan(&data.Uid, &data.Name, &data.Surname, &data.Alias)
	if err != nil {
		return nil, err
	}
	return &data, err
}




