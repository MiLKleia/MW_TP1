package music

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/music"
	"net/http"
)

// GetCollections
// @Tags         artist
// @Summary      Get all artists.
// @Description  Get artists.
// @Success      200            {array}  models.Album
// @Failure      500             "Something went wrong"
// @Router       /music/artists [get]

func GetAllArtists(w http.ResponseWriter, _ *http.Request) {
	// calling service
	artists, err := music.GetAllArtists()
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
	body, _ := json.Marshal(artists)
	_, _ = w.Write(body)
	return
}
