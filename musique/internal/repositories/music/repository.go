package music

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)


func UpdateSongById(id uuid.UUID, name string, artist string, album string) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	_ , err2 := db.Exec("UPDATE songs SET name = ?, artist = ?, album = ? WHERE id = ?", name, artist, album, id.String())
	
	
	if err2 != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM songs WHERE id=?", id.String())
	helpers.CloseDB(db)

	var data models.Song
	err = row.Scan(&data.Id, &data.Name, &data.Artist, &data.Album)
	if err != nil {
		return nil, err
	}
	return &data, err
	
}

func AddSong(name string, artist string, album string) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	// TODO boucle tant que pas nouveau UID

	new_id, _ := uuid.NewV4()

	
	_ , err2 := db.Exec("INSERT INTO songs (id, name, artist, album) VALUES(? ,? , ?, ?)", new_id.String(), name, artist, album)

	if err2 != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM songs WHERE id=?", new_id.String())
	helpers.CloseDB(db)

	var data models.Song
	err = row.Scan(&data.Id, &data.Name, &data.Artist, &data.Album)
	if err != nil {
		return nil, err
	}
	return &data, err
	
}
	

func DeleteSongById(id uuid.UUID) (error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	_ , err2 := db.Exec("DELETE FROM songs WHERE id=?", id.String())
	helpers.CloseDB(db)

	return err2
}

func GetAllSongs() ([]models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM songs")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	songs := []models.Song{}
	for rows.Next() {
		var data models.Song
		err = rows.Scan(&data.Id, &data.Name, &data.Artist, &data.Album)
		if err != nil {
			return nil, err
		}
		songs = append(songs, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return songs, err
}

func GetSongById(id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM songs WHERE id=?", id.String())
	helpers.CloseDB(db)

	var data models.Song
	err = row.Scan(&data.Id, &data.Name, &data.Artist, &data.Album)
	if err != nil {
		return nil, err
	}
	return &data, err
}




