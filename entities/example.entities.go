package entities

type Example struct {
	Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
