package model

type Response struct {
	Message string      `json:"message" example:"operasi berhasil"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty" example:"detail error"`
}
