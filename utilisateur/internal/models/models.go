package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
	Uid      *uuid.UUID `json:"uid"`
	Name string     `json:"name"`
	Username string     `json:"username"`
	Date string     `json:"inscription_date"`
}

type User_no_id struct {
    Name   string `json:"name"`
    Username string `json:"username"`
	Date string `json:"inscription_date"`
}


