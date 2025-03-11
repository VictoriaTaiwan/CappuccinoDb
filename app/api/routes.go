package handlers

type Response struct {
	Status string
	Data   []*Info
}

type Info struct {
	Content string `json:"content"`
	ID      string `json:"id"`
}