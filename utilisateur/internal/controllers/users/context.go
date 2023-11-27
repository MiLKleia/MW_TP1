package users

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"net/http"
)

func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userUid, err := uuid.FromString(chi.URLParam(r, "uid"))
		if err != nil {
			logrus.Errorf("parsing error : %s", err.Error())
			customError := &models.CustomError{
				Message: fmt.Sprintf("cannot parse uid (%s) as UUID", chi.URLParam(r, "uid")),
				Code:    http.StatusUnprocessableEntity,
			}
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
			return
		}

		ctx := context.WithValue(r.Context(), "userUid", userUid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
