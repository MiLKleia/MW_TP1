package main

import (

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/controllers/users"
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Get("/", users.GetAllUsers)
		r.Route("/{uid}", func(r chi.Router) {
			r.Use(users.Ctx)
			r.Get("/", users.GetUserByUid)
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8082")
	logrus.Fatalln(http.ListenAndServe(":8082", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS users (
			uid VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			name VARCHAR(255) NOT NULL,
			surname VARCHAR(255) NOT NULL,
			alias VARCHAR(255) NOT NULL
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
