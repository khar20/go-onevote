package models

import (
	"time"
)

type Candidate struct {
	ID             string `json:"id" db:"id"`
	Cip            string
	Name           string `json:"name" db:"name"`
	FatherLastName string
	MotherLastName string
	Description    string    `json:"description" db:"description"`
	Votes          int       `json:"votes" db:"votes"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// todo
func GetCandidates() ([]Candidate, error) {
	return []Candidate{
		{
			ID:        "1",
			Cip:       "12345",
			Name:      "Alex",
			CreatedAt: time.Now().AddDate(0, -1, 0),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "2",
			Cip:       "67890",
			Name:      "Juan",
			CreatedAt: time.Now().AddDate(0, -2, 0),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "3",
			Cip:       "24680",
			Name:      "Pedro",
			CreatedAt: time.Now().AddDate(0, -3, 0),
			UpdatedAt: time.Now(),
		},
	}, nil
}
