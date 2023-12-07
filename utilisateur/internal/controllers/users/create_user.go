package users

import (

	"io/ioutil"
	"encoding/json"
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
	
	body_in, _ := ioutil.ReadAll(r.Body)
	bodyString := string(body_in)
	user_in := user{}
	json.Unmarshal([]byte(bodyString), &user_in)

	name := user_in.Name
	surname := user_in.Surname
	alias := user_in.Alias

	user, err := users.CreateUser(name, surname, alias)
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
