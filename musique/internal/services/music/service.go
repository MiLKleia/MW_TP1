package music

import (
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/music"
	"net/http"
)

func GetAllSongs() ([]models.Song, error) {
	var err error
	// calling repository
	songs, err := repository.GetAllSongs()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving songs : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return songs, nil
}

func GetSongById(id uuid.UUID) (*models.Song, error) {
	var err error
	
	song , err := repository.GetSongById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "song not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving song : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return song, err
}


func DeleteSongById(id uuid.UUID) (error) {
	err := repository.DeleteSongById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return &models.CustomError{
				Message: "song not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error deleting song : %s", err.Error())
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


func AddSong(name string, artist string, album string) (*models.Song, error) {
	song, err := repository.AddSong(name, artist, album)

	if err != nil {
		logrus.Errorf("error adding song : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}
	return song, &models.CustomError{
			Message: "Created",
			Code:    201,
		}
	
}

func UpdateSongById(id uuid.UUID, name string, artist string, album string) (*models.Song, error) {
	var err error
	
	song , err := repository.UpdateSongById(id, name, artist, album)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "song not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error updating song : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return song, err
}



func GetAllAlbums() ([]models.Album, error) {
	var err error
	// calling repository
	albums, err := repository.GetAllAlbums()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving albums : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return albums, nil
}


func GetSongsFromAlbum(album string) ([]models.Song, error) {
	var err error
	// calling repository
	songs, err := repository.GetSongsFromAlbum(album)
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving albums : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return songs, nil
}

func GetAllArtists() ([]models.Artist, error) {
	var err error
	// calling repository
	artists, err := repository.GetAllArtists()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving albums : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return artists, nil
}


func GetSongsFromArtist(artist string) ([]models.Song, error) {
	var err error
	// calling repository
	songs, err := repository.GetSongsFromArtist(artist)
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving artists : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return songs, nil
}
