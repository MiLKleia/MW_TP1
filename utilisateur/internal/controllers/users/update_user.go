package users

import (
	"fmt"

	"io/ioutil"

	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/users"
	"net/http"
)

type user struct {
    Name   []string `json:"name"`
    Surname []string `json:"surname"`
	Alias []string `json:"alias"`
}

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
	userUid, _ := r.Context().Value("userUid").(uuid.UUID)

	//r.Body 
	//json.Unmarshall 


	body_in, _ := ioutil.ReadAll(r.Body)
	bodyString := string(body_in)
    fmt.Println(bodyString)

	user_in := user{}
	json.Unmarshal([]byte(bodyString), &user_in)
	fmt.Println(user_in)


	
	name := "john"
	surname := "doe"
	alias := "JD"
	

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
