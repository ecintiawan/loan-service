package entity

import "time"

type (
	// Borrower reflects borrower table
	// contains borrower data that submit certain loans
	Borrower struct {
		ID                   int64     `json:"id"                    db:"id"`
		IdentificationNumber string    `json:"identification_number"   db:"identification_number"`
		Name                 string    `json:"name"                  db:"name"`
		Status               int       `json:"status"                db:"status"`
		CreatedAt            time.Time `json:"created_at"            db:"created_at"`
		UpdatedAt            time.Time `json:"updated_at"            db:"updated_at"`
	}
)
