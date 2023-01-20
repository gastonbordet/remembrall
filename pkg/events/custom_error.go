package events

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	web "github.com/gastonbordet/remembrall/pkg/web/error"
)

func handleError(ctx context.Context, eventID string, err error) *web.WebError {
	code := http.StatusInternalServerError
	msg := fmt.Sprintf("Events Internal error: %s", err.Error())

	if errors.Is(err, EventNotFoundError) {
		code = http.StatusNotFound
		msg = fmt.Sprintf("Event %s not found.", eventID)
	}

	if errors.Is(err, EventInternalError) {
		code = http.StatusInternalServerError
		msg = fmt.Sprintf("Event %s internal error.", eventID)
	}

	return web.NewWebError(code, msg)
}
