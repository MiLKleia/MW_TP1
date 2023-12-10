package music

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/music"
	"net/http"
)

// delete song
// @Tags         song
// @Summary      delete a song.
// @Description  delete a song.
// @Param        uid           	path      string  true  "song UUID formatted ID"
// @Success      
// @Failure      
// @Failure      500            "Something went wrong"
// @Router       /music/{uid} [delete]

func DeleteSong(w http.ResponseWriter, r *http.Request) {

	songId, _ := r.Context().Value("music_id").(uuid.UUID)
	
	err := music.DeleteSongById(songId)

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
	return

}