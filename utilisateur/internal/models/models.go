package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
	Uid      *uuid.UUID `json:"uid"`
	Name string     `json:"name"`
	Surname string     `json:"surname"`
	Alias string     `json:"alias"`
}