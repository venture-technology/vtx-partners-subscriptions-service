package models

type Driver struct {
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	CNH        string  `json:"cnh"`
	Street     string  `json:"street"`
	Number     string  `json:"number"`
	Complement string  `json:"complement,omitempty"`
	ZIP        string  `json:"zip"`
	Price      float64 `json:"price"`
}
