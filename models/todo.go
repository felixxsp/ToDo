package models

type Todo struct {
	UUID      string `json:"uuid"` // looks like this: 6ba7b810-9dad-11d1-80b4-00c04fd430c8
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Index     int    `json:"index"`
}
