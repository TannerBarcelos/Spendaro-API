package models

// CommonResponsePayload is a struct that represents the common response payload for all API responses.
type CommonResponsePayload struct {
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   error       `json:"error"`
}
