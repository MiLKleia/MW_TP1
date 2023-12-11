package users

import (

	"io/ioutil"
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
//		  name, surname, alias  body .json
// @Success      200            {object}  models.user
// @Failure      422            "Cannot parse uid"
// @Failure      500            "Something went wrong"
// @Router       /users/{uid} [put]




func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userUid, _ := r.Context().Value("userUid").(uuid.UUID)



	body_in, _ := ioutil.ReadAll(r.Body)
	bodyString := string(body_in)
	var user_in models.User_no_id
	json.Unmarshal([]byte(bodyString), &user_in)

	
	name := user_in.Name
	surname := user_in.Surname
	alias := user_in.Alias
	

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
