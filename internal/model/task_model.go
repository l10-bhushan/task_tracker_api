package model

type Task struct {
	Id                  string `json:"id"`
	Title               string `json:"title"`
	Status              bool   `json:"status"`
	Timestamp           string `json:"timestamp"`
	CompletionTimestamp string `json:"completiontime"`
	// Todo : We can make this better by introducing enums
	Category string `json:"category"`
}
