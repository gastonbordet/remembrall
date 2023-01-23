package events_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gastonbordet/remembrall/pkg/events"
	web "github.com/gastonbordet/remembrall/pkg/web/error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEventsService struct {
	mock.Mock
}

func (mock *MockEventsService) GetAll(ctx context.Context) ([]events.Event, error) {
	args := mock.Called()
	mock_events := args.Get(0).([]events.Event)

	return mock_events, args.Error(1)
}

func (mock *MockEventsService) GetByEventID(ctx context.Context, eventID uint) (*events.Event, error) {
	args := mock.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	mock_event := args.Get(0).(events.Event)

	return &mock_event, args.Error(1)
}

type MockResponseWriter struct {
	mock.Mock
}

func (mock *MockResponseWriter) Header() http.Header {
	x := &http.Header{}
	return *x
}

func (mock *MockResponseWriter) WriteHeader() {
	return
}

func (mock *MockResponseWriter) Write([]byte) (int, error) {
	return 1, nil
}

func TestGetAllOk(t *testing.T) {
	// Setup
	var response []events.Event
	var error_response error
	events_response := []events.Event{{
		ID:          1,
		Title:       "test",
		Description: "test",
		Date:        "",
		Status:      true,
	}}

	eventService := &MockEventsService{}
	eventService.On("GetAll").Return(events_response, error_response)

	eventHandler := events.BuildEventsHandler(eventService)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/events/", nil)

	// Act
	eventHandler.GetAll(w, r)
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "test", response[0].Title)
}

func TestGetAllInternalError(t *testing.T) {
	// Setup
	var response []events.Event
	eventsService := &MockEventsService{}
	eventsService.On("GetAll").Return(response, events.EventInternalError)
	eventsHandler := events.BuildEventsHandler(eventsService)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/events/", nil)

	// Act
	eventsHandler.GetAll(w, r)
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert
	assert.Nil(t, err)
	assert.Empty(t, response)
	assert.Equal(t, w.Code, http.StatusInternalServerError)
}

func TestGetByIDOk(t *testing.T) {
	// Setup
	var response *events.Event
	event_response := &events.Event{
		ID:          1,
		Title:       "test",
		Description: "test",
		Date:        "",
		Status:      true,
	}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/events/1", nil)
	eventsService := &MockEventsService{}
	eventsService.On("GetByEventID").Return(*event_response, nil)
	eventsHandler := events.BuildEventsHandler(eventsService)

	// Act
	eventsHandler.GetByEventID(w, r)
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "test", response.Title)
}

func TestGetByIDNotFound(t *testing.T) {
	// Setup
	var response *events.Event
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/events/1", nil)
	eventsService := &MockEventsService{}
	eventsService.On("GetByEventID").Return(nil, events.EventNotFoundError)
	eventsHandler := events.BuildEventsHandler(eventsService)

	// Act
	eventsHandler.GetByEventID(w, r)
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert
	assert.Nil(t, err)
	assert.Nil(t, response)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

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
