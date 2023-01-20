package events_test

import (
	"fmt"
	"testing"

	"github.com/gastonbordet/remembrall/pkg/events"
	"github.com/stretchr/testify/assert"
)

func TestValidateEvent(t *testing.T) {
	type Test struct {
		Title          string
		Input          *events.Event
		ExpectedResult bool
	}

	tests := []Test{{
		Title: "Event invalid title too long.",
		Input: &events.Event{
			ID:          1,
			Title:       "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaas",
			Description: "a",
			Date:        "01/02/2006",
		},
		ExpectedResult: false,
	}, {
		Title: "Event invalid title too short.",
		Input: &events.Event{
			ID:          1,
			Title:       "",
			Description: "a",
			Date:        "01/02/2006",
		},
		ExpectedResult: false,
	}, {
		Title: "Event invalid description too long.",
		Input: &events.Event{
			ID:          1,
			Title:       "test",
			Description: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			Date:        "01/02/2006",
		},
		ExpectedResult: false,
	}, {
		Title: "Event invalid ID 0.",
		Input: &events.Event{
			ID:          0,
			Title:       "test",
			Description: "hehe",
			Date:        "01/02/2006",
		},
		ExpectedResult: false,
	}, {
		Title: "Event invvalid date is wrong.",
		Input: &events.Event{
			ID:          1,
			Title:       "test",
			Description: "hehe",
			Date:        "brr",
		},
		ExpectedResult: false,
	}, {
		Title: "Event valid.",
		Input: &events.Event{
			ID:          1,
			Title:       "test",
			Description: "hehe",
			Date:        "01/02/2006",
		},
		ExpectedResult: true,
	}}

	for _, test := range tests {
		assert.Equal(t, test.ExpectedResult, test.Input.ValidateEvent(), fmt.Sprintf("Failed test: %s", test.Title))
	}
}
