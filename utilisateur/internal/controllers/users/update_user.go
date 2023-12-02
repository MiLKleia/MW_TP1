package users

import (
	"github.com/go-chi/chi/v5"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/users"
	"net/http"
)

// UpdateUser
// @Tags         users
// @Summary      Update an user and return it.
// @Description  Update an user.
// @Param        uid           	path      string  true  "user UUID formatted ID"
// @Success      200            {object}  models.user
// @Failure      422            "Cannot parse uid"
// @Failure      500            "Something went wrong"
// @Router       /users/{uid} [get]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userUid, _ := ctx.Value("userUid").(uuid.UUID)
	name := chi.URLParam(r, "name")
	surname := chi.URLParam(r, "surname")
	alias := chi.URLParam(r, "alias")

	user, err := users.UpdateUserByUid(userUid, name, surname, alias)
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
	body, _ := json.Marshal(user)
	_, _ = w.Write(body)
	return
}
