package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/controllers/music"
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Route("/music", func(r chi.Router) {
		r.Get("/", music.GetAllSongs)
		r.Post("/", music.AddSong)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(music.Ctx)
			r.Get("/", music.GetSong)
			r.Delete("/", music.DeleteSong)
			r.Put("/", music.UpdateSong)
		})
		r.Route("/album", func(r chi.Router) {
			r.Get("/", music.GetAllAlbums)
			r.Route("/{album}", func(r chi.Router) {
				r.Get("/", music.GetSongsFromAlbum)
			})

		})
		r.Route("/artist", func(r chi.Router) {
			r.Get("/", music.GetAllArtists)
			r.Route("/{artist}", func(r chi.Router) {
				r.Get("/", music.GetSongsFromArtist)
			})

		})
	})


	logrus.Info("[INFO] Web server started. Now listening on *:8083")
	logrus.Fatalln(http.ListenAndServe(":8083", r))



}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS songs (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			name VARCHAR(255) NOT NULL,
			artist VARCHAR(255) NOT NULL,
			album VARCHAR(255) NOT NULL
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
