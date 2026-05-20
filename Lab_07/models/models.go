package models

//easyjson:json
type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

//easyjson:json
type NoteList []Note

//easyjson:json
type ErrorResponse struct {
	Error string `json:"error"`
}
