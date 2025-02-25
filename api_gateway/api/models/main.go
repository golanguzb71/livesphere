package models

type Response struct {
	Status  string `json:"statusCode"`
	Message string `json:"message"`
}
