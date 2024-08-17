package entity

import (
	"reflect"
	"testing"
	"time"
)

func TestInvestment_IsValid(t *testing.T) {
	type fields struct {
		ID         int64
		InvestorID int64
		LoanID     int64
		Amount     float64
		ROI        float64
		Status     int
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "valid",
			fields: fields{
				InvestorID: 1,
				LoanID:     3,
				Amount:     10000,
			},
			want: true,
		},
		{
			name: "invalid",
			fields: fields{
				InvestorID: 1,
				LoanID:     3,
				Amount:     0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &Investment{
				ID:         tt.fields.ID,
				InvestorID: tt.fields.InvestorID,
				LoanID:     tt.fields.LoanID,
				Amount:     tt.fields.Amount,
				ROI:        tt.fields.ROI,
				Status:     tt.fields.Status,
				CreatedAt:  tt.fields.CreatedAt,
				UpdatedAt:  tt.fields.UpdatedAt,
			}
			if got := data.IsValid(); got != tt.want {
				t.Errorf("Investment.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvestmentFilter_Validate(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name   string
		fields fields
		want   *InvestmentFilter
	}{
		{
			name:   "success",
			fields: fields{},
			want: &InvestmentFilter{
				DataTable: DataTableFilter{
					Sort: DataTableSort{
						Field:     "id",
						Direction: "desc",
					},
					Pagination: DataTablePagination{
						Limit: 10,
						Page:  1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &InvestmentFilter{
				DataTable:      tt.fields.DataTable,
				ID:             tt.fields.ID,
				InvestorID:     tt.fields.InvestorID,
				LoanID:         tt.fields.LoanID,
				Status:         tt.fields.Status,
				CreatedAtStart: tt.fields.CreatedAtStart,
				CreatedAtEnd:   tt.fields.CreatedAtEnd,
				UpdatedAtStart: tt.fields.UpdatedAtStart,
				UpdatedAtEnd:   tt.fields.UpdatedAtEnd,
			}
			filter.Validate()
			if !reflect.DeepEqual(filter, tt.want) {
				t.Errorf("InvestmentFilter.Validate() = %v, want %v", filter, tt.want)
			}
		})
	}
}
