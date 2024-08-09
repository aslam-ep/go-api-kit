package utils

// MessageRes struct for default response success status and message
type MessageRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
