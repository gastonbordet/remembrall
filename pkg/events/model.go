package events

import "time"

const (
	title_max_length       = 30
	title_min_length       = 1
	description_max_length = 200
)

type Event struct {
	ID          uint   `json:"id"`
	Title       string `json:"Title"`
	Date        string `json:"Date"`
	Status      bool   `json:"Status"`
	Description string `json:"Description"`
}

func (event *Event) ValidateEvent() bool {
	result := true

	if event.ID < 1 {
		result = false
	}

	if len(event.Title) > title_max_length || len(event.Title) < title_min_length {
		result = false
	}

	if len(event.Description) > description_max_length {
		result = false
	}

	if _, err := time.Parse("01/02/2006", event.Date); err != nil {
		result = false
	}

	return result
}
