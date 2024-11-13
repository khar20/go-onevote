package models

import (
	"context"
	"fmt"
	"onevote/database"
	"time"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID            int    `db:"id"`             // SERIAL, PRIMARY KEY
	CIP           string `db:"cip"`            // TEXT, NOT NULL
	DNI           string `db:"dni"`            // TEXT, NOT NULL
	Name          string `db:"name"`           // TEXT, NOT NULL
	FirstSurname  string `db:"first_surname"`  // TEXT, NOT NULL
	SecondSurname string `db:"second_surname"` // TEXT, NOT NULL
	Email         string `db:"email"`          // TEXT, NOT NULL
	BranchID      int    `db:"branch_id"`      // INTEGER, FOREIGN KEY to branches table
	Role          string `db:"role"`           // TEXT, CHECK(role IN ('ADMIN', 'VOTER', 'MONITOR')), NOT NULL
	Attended      int    `db:"attended"`       // INTEGER, NOT NULL, DEFAULT 0

	CreatedAt time.Time `db:"created_at"` // TIMESTAMP, NOT NULL, DEFAULT CURRENT_TIMESTAMP
	UpdatedAt time.Time `db:"updated_at"` // TIMESTAMP, NOT NULL, DEFAULT CURRENT_TIMESTAMP
}

func GetUsers() ([]User, error) {
	conn, err := database.ConnectCoreDB()
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), `
		SELECT
			u.id, u.cip, u.dni, u.name, u.first_surname, u.second_surname, u.email, u.branch_id, b.branch_name, u.role, u.attended, u.created_at, u.updated_at
		FROM users u
		JOIN branches b ON u.branch_id = b.id`)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %v", err)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.CIP,
			&user.DNI,
			&user.Name,
			&user.FirstSurname,
			&user.SecondSurname,
			&user.Email,
			&user.BranchID,
			&user.Role,
			&user.Attended,
			&user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows iteration error: %v", rows.Err())
	}

	return users, nil
}

func GetUserByCIP(cip string) (*User, error) {
	conn, err := database.ConnectCoreDB()
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	row := conn.QueryRow(context.Background(), `
		SELECT
			id, cip, dni, name, first_surname, second_surname, email, branch_id, role, attended, created_at, updated_at
		FROM users
		WHERE cip=$1`, cip)

	var user User

	err = row.Scan(
		&user.ID,
		&user.CIP,
		&user.DNI,
		&user.Name,
		&user.FirstSurname,
		&user.SecondSurname,
		&user.Email,
		&user.BranchID,
		&user.Role,
		&user.Attended,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan user: %v", err)
	}

	return &user, nil
}

func GetUserByID(id string) (*User, error) {
	conn, err := database.ConnectCoreDB()
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	row := conn.QueryRow(context.Background(), `
		SELECT
			id, cip, dni, name, first_surname, second_surname, email, branch_id, role, attended, created_at, updated_at
		FROM users
		WHERE id=$1`, id)

	var user User

	err = row.Scan(
		&user.ID,
		&user.CIP,
		&user.DNI,
		&user.Name,
		&user.FirstSurname,
		&user.SecondSurname,
		&user.Email,
		&user.BranchID,
		&user.Role,
		&user.Attended,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan user: %v", err)
	}

	return &user, nil
}
