package users

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/users"
	"net/http"
)

// delete user
// @Tags         users
// @Summary      delete a user.
// @Description  delete a user.
// @Param        uid           	path      string  true  "user UUID formatted ID"
// @Success      
// @Failure      
// @Failure      500            "Something went wrong"
// @Router       /users/{uid} [delete]

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	userUid, _ := r.Context().Value("userUid").(uuid.UUID)
	
	err := users.DeleteUserByUid(userUid)

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