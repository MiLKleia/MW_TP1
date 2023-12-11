package music

import (

	"io/ioutil"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/music"
	"net/http"
)


// AddSong
// @Tags         songs
// @Summary      Get a song.
// @Description  Get a song.
// @Param        id           	path      string  true  "song UUID formatted ID"
//				 name, artist, album      as .json in body
// @Success      200            {object}  models.music
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /music   [post]


func AddSong(w http.ResponseWriter, r *http.Request) {
	
	body_in, _ := ioutil.ReadAll(r.Body)
	bodyString := string(body_in)
	var song_in models.Song_no_id
	json.Unmarshal([]byte(bodyString), &song_in)

	name := song_in.Name
	artist := song_in.Artist
	album := song_in.Album

	song, err := music.AddSong(name, artist, album)
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
