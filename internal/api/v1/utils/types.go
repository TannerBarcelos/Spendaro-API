package utils

type CommonResponsePayload struct {
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   error       `json:"error"`
}
