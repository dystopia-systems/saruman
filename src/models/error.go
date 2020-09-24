package models

type Error struct {
	Code int `json:"core"`
	Message string `json:"message"`
}