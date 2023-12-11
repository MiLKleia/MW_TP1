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


	// does the song already exist
	row_test := db.QueryRow("SELECT * FROM songs WHERE name=? AND artist=? AND album=?",name, artist, album)
	var data_test_exist models.Song
	err = row_test.Scan(&data_test_exist.Id, &data_test_exist.Name, &data_test_exist.Artist, &data_test_exist.Album)
	if err == nil {
		return nil, &models.CustomError{
			Message: "Song exists already",
			Code:    409,
		}
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


func GetAllAlbums() ([]models.Album, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT DISTINCT album FROM songs")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	albums := []models.Album{}
	for rows.Next() {
		var data models.Album
		err = rows.Scan(&data.Album)
		if err != nil {
			return nil, err
		}
		albums = append(albums, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return albums, err
}


func GetSongsFromAlbum(album string) ([]models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM songs WHERE album=?", album)
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


func GetAllArtists() ([]models.Artist, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT DISTINCT artist FROM songs")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	artists := []models.Artist{}
	for rows.Next() {
		var data models.Artist
		err = rows.Scan(&data.Artist)
		if err != nil {
			return nil, err
		}
		artists = append(artists, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return artists, err
}


func GetSongsFromArtist(artist string) ([]models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	
	rows, err := db.Query("SELECT * FROM songs WHERE artist=?", artist)
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



