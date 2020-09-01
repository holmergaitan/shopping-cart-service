package domain

type Cart struct {
	ID    string `json:"id"`
	Category  string `json:"category,omitempty"`
}
