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
// @Tags         songs
// @Summary      Get all songs from an artist.
// @Description  Get songs.
// @Param        artist           	path      string  artist
// @Success      200            {array}  models.Alsongsbum
// @Failure      500             "Something went wrong"
// @Router       /music/artist/{artist} [get]

func GetSongsFromArtist(w http.ResponseWriter, r *http.Request) {
	// calling service

	artist := chi.URLParam(r, "artist")

	songs, err := music.GetSongsFromArtist(artist)
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