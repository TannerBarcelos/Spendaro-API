package utils

import (
	"errors"
	"net/http"
	"spendaro-api/internal/api/v1/models"
	"testing"
)

func TestPrepareResponse(t *testing.T) {

	t.Run("Test PrepareResponse", func(t *testing.T) {
		data := "some data"
		status := http.StatusOK
		message := "message"
		want := models.CommonResponsePayload{
			Data:    data,
			Status:  status,
			Message: message,
			Error:   nil,
		}
		got := PrepareResponse(data, status, message, nil)
		if got != want {
			t.Errorf("PrepareResponse() = %v, want %v", got, want)
		}
	})

	t.Run("Test PrepareResponse with error", func(t *testing.T) {
		data := "some data"
		status := http.StatusOK
		message := "message"
		err := errors.New("error")
		want := models.CommonResponsePayload{
			Data:    data,
			Status:  status,
			Message: message,
			Error:   err,
		}
		got := PrepareResponse(data, status, message, err)
		if got != want {
			t.Errorf("PrepareResponse() = %v, want %v", got, want)
		}
	})
}
