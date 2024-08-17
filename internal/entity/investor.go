package entity

import "time"

type (
	// Investor reflects investor table
	// contains investor data that invests on certain loans
	Investor struct {
		ID                   int64     `json:"id"                    db:"id"`
		IdentificationNumber string    `json:"identification_number"   db:"identification_number"`
		Name                 string    `json:"name"                  db:"name"`
		Email                string    `json:"email"                  db:"email"`
		Status               int       `json:"status"                db:"status"`
		CreatedAt            time.Time `json:"created_at"            db:"created_at"`
		UpdatedAt            time.Time `json:"updated_at"            db:"updated_at"`
	}
)
