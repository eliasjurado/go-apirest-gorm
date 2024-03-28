package models

type Response struct {
	IsSuccess  bool        `json:"isSuccess"`
	Status     int         `json:"status"`
	StatusCode string      `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Metadata   []string    `json:"metadata"`
}
