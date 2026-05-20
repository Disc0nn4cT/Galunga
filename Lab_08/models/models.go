package models

//easyjson:json
type Contact struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

//easyjson:json
type ContactList []Contact

//easyjson:json
type ErrorResponse struct {
	Error string `json:"error"`
}
