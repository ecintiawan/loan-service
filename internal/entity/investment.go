package entity

import "time"

type (
	// Investment reflects investment table
	// contains investment data by certain investor/user
	Investment struct {
		ID         int64     `json:"id"          db:"id"`
		InvestorID int64     `json:"investor_id" db:"investor_id"`
		LoanID     int64     `json:"loan_id"     db:"loan_id"`
		Amount     float64   `json:"amount"       db:"amount"`
		ROI        float64   `json:"roi"          db:"roi"`
		Status     int       `json:"status"      db:"status"`
		CreatedAt  time.Time `json:"created_at"  db:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"  db:"updated_at"`
	}

	// InvestmentFilter stores pagination and filter used in get investment request
	InvestmentFilter struct {
		DataTable      DataTableFilter
		ID             int64
		InvestorID     int64
		LoanID         int64
		Status         int
		CreatedAtStart time.Time
		CreatedAtEnd   time.Time
		UpdatedAtStart time.Time
		UpdatedAtEnd   time.Time
	}

	// InvestmentResult for API fetch response with pagination
	InvestmentResult struct {
		List []*Investment
		Pagination
	}
)

func (data *Investment) IsValid() bool {
	return data.InvestorID > 0 && data.LoanID > 0 && data.Amount > 0
}

// Validate self corrects IP Whitelist filter
func (filter *InvestmentFilter) Validate() {
	filter.DataTable.Validate()

	columnSortability := map[string]bool{
		"id":          true,
		"investor_id": true,
		"loan_id":     true,
		"amount":      true,
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
