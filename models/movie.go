package models

type Movie struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Duration    string `json:"duration"`
	Description string `json:"description"`
}