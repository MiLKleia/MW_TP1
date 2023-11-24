package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
	Id      *uuid.UUID `json:"id"`
	Name string     `json:"content"`
	Surname string     `json:"content"`
	Alias string     `json:"content"`
}
