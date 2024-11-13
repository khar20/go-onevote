package models

import (
	"context"
	"fmt"
	"onevote/database"
	"time"

	"github.com/jackc/pgx/v5"
)

type Candidate struct {
	ID                 int     `db:"id"`                  // SERIAL, PRIMARY KEY
	CIP                string  `db:"cip"`                 // TEXT, NOT NULL
	DNI                string  `db:"dni"`                 // TEXT, NOT NULL
	Name               string  `db:"name"`                // TEXT, NOT NULL
	FirstSurname       string  `db:"first_surname"`       // TEXT, NOT NULL
	SecondSurname      string  `db:"second_surname"`      // TEXT, NOT NULL
	PartyID            int     `db:"party_id"`            // INTEGER, FOREIGN KEY to parties table
	BranchID           int     `db:"branch_id"`           // INTEGER, FOREIGN KEY to branches table
	PositionApplied    string  `db:"position_applied"`    // TEXT, NOT NULL
	PreviousExperience *string `db:"previous_experience"` // TEXT, nullable
	AcademicBackground *string `db:"academic_background"` // TEXT, nullable
	AdditionalInfo     *string `db:"additional_info"`     // TEXT, nullable

	CreatedAt time.Time `db:"created_at"` // TIMESTAMP, NOT NULL, DEFAULT CURRENT_TIMESTAMP
	UpdatedAt time.Time `db:"updated_at"` // TIMESTAMP, NOT NULL, DEFAULT CURRENT_TIMESTAMP
}

func GetCandidates() ([]Candidate, error) {
	conn, err := database.ConnectCoreDB()
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), `
		SELECT
			id, cip, dni, name, first_surname, second_surname, party_id, branch_id, position_applied
		FROM candidates`)
	if err != nil {
		return nil, fmt.Errorf("failed to query candidates: %v", err)
	}
	defer rows.Close()

	var candidates []Candidate

	for rows.Next() {
		var candidate Candidate
		err := rows.Scan(
			&candidate.ID,
			&candidate.CIP,
			&candidate.DNI,
			&candidate.Name,
			&candidate.FirstSurname,
			&candidate.SecondSurname,
			&candidate.PartyID,
			&candidate.BranchID,
			&candidate.PositionApplied,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		candidates = append(candidates, candidate)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows iteration error: %v", rows.Err())
	}

	return candidates, nil
}

func GetCandidateByCIP(cip string) (*Candidate, error) {
	conn, err := database.ConnectCoreDB()
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	row := conn.QueryRow(context.Background(), `
		SELECT 
			id, cip, dni, name, first_surname, second_surname, 
			party_id, branch_id, position_applied, previous_experience, 
			academic_background, additional_info, created_at, updated_at 
		FROM candidates 
		WHERE cip = $1`, cip)

	var candidate Candidate

	err = row.Scan(
		&candidate.ID,
		&candidate.CIP,
		&candidate.DNI,
		&candidate.Name,
		&candidate.FirstSurname,
		&candidate.SecondSurname,
		&candidate.PartyID,
		&candidate.BranchID,
		&candidate.PositionApplied,
		&candidate.PreviousExperience,
		&candidate.AcademicBackground,
		&candidate.AdditionalInfo,
		&candidate.CreatedAt,
		&candidate.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan candidate: %v", err)
	}

	return &candidate, nil
}

func GetCandidateByID(id string) (*Candidate, error) {
	conn, err := database.ConnectCoreDB()
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	row := conn.QueryRow(context.Background(), `
		SELECT 
			id, cip, dni, name, first_surname, second_surname, 
			party_id, branch_id, position_applied, previous_experience, 
			academic_background, additional_info, created_at, updated_at 
		FROM candidates 
		WHERE id = $1`, id)

	var candidate Candidate

	err = row.Scan(
		&candidate.ID,
		&candidate.CIP,
		&candidate.DNI,
		&candidate.Name,
		&candidate.FirstSurname,
		&candidate.SecondSurname,
		&candidate.PartyID,
		&candidate.BranchID,
		&candidate.PositionApplied,
		&candidate.PreviousExperience,
		&candidate.AcademicBackground,
		&candidate.AdditionalInfo,
		&candidate.CreatedAt,
		&candidate.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan candidate: %v", err)
	}

	return &candidate, nil
}
