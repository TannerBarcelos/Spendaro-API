package utils

func PrepareResponse(data interface{}, status int, message string, err error) CommonResponsePayload {
	return CommonResponsePayload{
		Data:    data,
		Status:  status,
		Message: message,
		Error:   err,
	}
}
