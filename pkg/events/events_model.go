package events

type Event struct {
	ID          uint   `json:"id"`
	Title       string `json:"Title"`
	Date        string `json:"Date"`
	Status      bool   `json:"Status"`
	Description string `json:"Description"`
}
