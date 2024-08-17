package entity

import (
	"time"

	"github.com/ecintiawan/loan-service/internal/constant"
)

type (
	// Loan reflects loan table
	// contains loan data by certain borrower
	Loan struct {
		ID                 int64               `json:"id"                             db:"id"`
		BorrowerID         int64               `json:"borrower_id"                    db:"borrower_id"`
		Amount             float64             `json:"amount"                          db:"amount"`
		Rate               float64             `json:"rate"                            db:"rate"`
		ApprovalProofURL   string              `json:"approval_proof_url,omitempty"   db:"approval_proof_url"`
		AgreementLetterURL string              `json:"agreement_letter_url,omitempty" db:"agreement_letter_url"`
		Status             constant.LoanStatus `json:"status"                         db:"status"`
		CreatedBy          int64               `json:"created_by"                     db:"created_by"`
		ApprovedBy         int64               `json:"approved_by,omitempty"          db:"approved_by"`
		DisbursedBy        int64               `json:"disbursed_by,omitempty"         db:"disbursed_by"`
		CreatedAt          time.Time           `json:"created_at"                     db:"created_at"`
		UpdatedAt          time.Time           `json:"updated_at,omitempty"           db:"updated_at"`
		ApprovedAt         time.Time           `json:"approved_at,omitempty"          db:"approved_at"`
		InvestedAt         time.Time           `json:"invested_at,omitempty"          db:"invested_at"`
		DisbursedAt        time.Time           `json:"disbursed_at,omitempty"         db:"disbursed_at"`
	}

	// LoanFilter stores pagination and filter used in get loan request
	LoanFilter struct {
		DataTable      DataTableFilter
		ID             int64
		BorrowerID     int64
		Status         constant.LoanStatus
		CreatedAtStart time.Time
		CreatedAtEnd   time.Time
		UpdatedAtStart time.Time
		UpdatedAtEnd   time.Time
		ApprovedBy     int64
		DisbursedBy    int64
	}

	// LoanResult for API fetch response with pagination
	LoanResult struct {
		List []*Loan
		Pagination
	}

	// LoanProceed for API fetch request with additional parameters
	LoanProceed struct {
		Action          constant.LoanAction `json:"action"`
		ApprovalProof   File                `json:"-"`
		AgreementLetter File                `json:"-"`
		Data            *Loan               `json:"data"`
	}
)

func (data *Loan) IsValid() bool {
	return data.BorrowerID > 0 && data.Amount > 0 && data.Rate > 0
}

func (req *LoanProceed) IsValid() bool {
	return req.Action > 0 && req.Data != nil && req.Data.ID > 0
}

func (req *LoanProceed) IsValidExt() bool {
	if req.ApprovalProof.File != nil {
		return req.ApprovalProof.FileExt == ".jpg" ||
			req.ApprovalProof.FileExt == ".jpeg" ||
			req.ApprovalProof.FileExt == ".png"
	}

	if req.AgreementLetter.File != nil {
		return req.AgreementLetter.FileExt == ".pdf"
	}

	return true
}

// Validate self corrects IP Whitelist filter
func (filter *LoanFilter) Validate() {
	filter.DataTable.Validate()

	columnSortability := map[string]bool{
		"id":          true,
		"borrower_id": true,
		"amount":      true,
		"rate":        true,
		"status":      true,
		"created_at":  true,
		"updated_at":  true,
	}

	sortable, valid := columnSortability[filter.DataTable.Sort.Field]
	if !(valid && sortable) {
		filter.DataTable.Sort.Field = "id"
		filter.DataTable.Sort.Direction = "desc"
	}
}
