package music

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/music"
	"net/http"
)

// Getsong
// @Tags         song
// @Summary      Get a song.
// @Description  Get a song.
// @Param        uid           	path      string  true  "song UUID formatted ID"
// @Success      200            {object}  models.Song
// @Failure      422            "Cannot parse uid"
// @Failure      500            "Something went wrong"
// @Router       /music/{uid} [get]
func GetSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	songId, _ := ctx.Value("music_id").(uuid.UUID)

	song, err := music.GetSongById(songId)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(song)
	_, _ = w.Write(body)
	return
}
