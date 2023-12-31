package music

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/music"
	"net/http"
)

// GetCollections
// @Tags         songs
// @Summary      Get songs.
// @Description  Get song.
// @Success      200            {array}  models.Song
// @Failure      500             "Something went wrong"
// @Router       /music [get]

func GetAllSongs(w http.ResponseWriter, _ *http.Request) {
	// calling service
	songs, err := music.GetAllSongs()
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
