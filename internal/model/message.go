package model

type Message struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
