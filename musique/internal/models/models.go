package models

import (
	"github.com/gofrs/uuid"
)

type Song struct {
	Id      *uuid.UUID `json:"id"`
	Name   string `json:"name"`
    Artist string `json:"artist"`
	Album string `json:"album"`
}

type Album struct {
	Album string `json:"album"`
}

type Song_no_id struct {
	Name   string `json:"name"`
    Artist string `json:"artist"`
	Album string `json:"album"`
}