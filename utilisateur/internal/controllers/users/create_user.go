package users

import (

	"github.com/go-chi/chi/v5"
	"encoding/json"
	//"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/users"
	"net/http"
)

// GetUser
// @Tags         users
// @Summary      Get a user.
// @Description  Get a user.
// @Param        uid           	path      string  true  "user UUID formatted ID"
// @Success      200            {object}  models.user
// @Failure      422            "Cannot parse uid"
// @Failure      500            "Something went wrong"
// @Router       /users/{uid} [get]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	alias := chi.URLParam(r, "alias")

	err := users.CreateUser(name, surname, alias)
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
