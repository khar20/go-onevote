package models

import (
	"time"
)

type User struct {
	ID             string `json:"id" db:"id"`
	Cip            string `json:"cip" db:"cip"`
	Name           string
	FatherLastName string
	MotherLastName string
	Password       string
	PasswordHash   string    `json:"password_hash" db:"password_hash"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// todo
func GetUsers() []User {
	return []User{
		{
			ID:             "1",
			Cip:            "12345",
			Password:       "hashedpassword1",
			Name:           "Alex",
			FatherLastName: "Rodríguez",
			MotherLastName: "Díaz",
			CreatedAt:      time.Now().AddDate(0, -1, 0),
			UpdatedAt:      time.Now(),
		},
		{
			ID:             "2",
			Cip:            "67890",
			Password:       "hashedpassword2",
			Name:           "Alex",
			FatherLastName: "Rodríguez",
			MotherLastName: "Díaz",
			CreatedAt:      time.Now().AddDate(0, -2, 0),
			UpdatedAt:      time.Now(),
		},
		{
			ID:             "3",
			Cip:            "24680",
			Password:       "hashedpassword3",
			Name:           "Alex",
			FatherLastName: "Rodríguez",
			MotherLastName: "Díaz",
			CreatedAt:      time.Now().AddDate(0, -3, 0),
			UpdatedAt:      time.Now(),
		},
	}
}
