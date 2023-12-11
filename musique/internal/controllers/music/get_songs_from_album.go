package music

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/music"
	"net/http"
	"github.com/go-chi/chi/v5"
)

// GetCollections
// @Tags         albums
// @Summary      Get all albums.
// @Description  Get albums.
// @Param        album           	path      string  album
// @Success      200            {array}  models.Album
// @Failure      500             "Something went wrong"
// @Router       /music [get]

func GetSongsFromAlbum(w http.ResponseWriter, r *http.Request) {
	// calling service

	album := chi.URLParam(r, "album")
	songs, err := music.GetSongsFromAlbum(album)
	if err != nil {
		// logging error
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			// writing http code in header
			w.WriteHeader(customError.Code)
			// writing error message in body
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(songs)
	_, _ = w.Write(body)
	return
}