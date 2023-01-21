package events_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/gastonbordet/remembrall/pkg/events"
	web "github.com/gastonbordet/remembrall/pkg/web/error"
	"github.com/stretchr/testify/assert"
)

func TestHandleError(t *testing.T) {
	type TestHandleError struct {
		Title          string
		ExpectedResult *web.WebError
		Input          map[string]interface{}
	}

	tests := []TestHandleError{{
		Title: "HandleError return WebError when not custom error.",
		ExpectedResult: &web.WebError{
			Status:  http.StatusInternalServerError,
			Message: "Events Internal error: error",
		},
		Input: map[string]interface{}{
			"eventID": "1",
			"err":     fmt.Errorf("error"),
		},
	}, {
		Title: "HandlerError return WebError when EventNotFoundError.",
		ExpectedResult: &web.WebError{
			Status:  http.StatusNotFound,
			Message: "Event 1 not found.",
		},
		Input: map[string]interface{}{
			"eventID": "1",
			"err":     events.EventNotFoundError,
		},
	}, {
		Title: "HandleError return WebError when EventInternalError.",
		ExpectedResult: &web.WebError{
			Status:  http.StatusInternalServerError,
			Message: "Event 1 generated internal error.",
		},
		Input: map[string]interface{}{
			"eventID": "1",
			"err":     events.EventInternalError,
		},
	}}

	for _, test := range tests {
		webErr := events.HandleError(context.Background(), test.Input["eventID"].(string), test.Input["err"].(error))
		assert.Equal(t, test.ExpectedResult.Status, webErr.Status, fmt.Sprintf("Failed test: %s", test.Title))
		assert.Equal(t, test.ExpectedResult.Message, webErr.Message)
	}
}
