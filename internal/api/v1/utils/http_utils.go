package utils

import "spendaro-api/internal/api/v1/models"

func PrepareResponse(data interface{}, status int, message string, err error) models.CommonResponsePayload {
	return models.CommonResponsePayload{
		Data:    data,
		Status:  status,
		Message: message,
		Error:   err,
	}
}
