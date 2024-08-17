package entity

import (
	"reflect"
	"testing"
	"time"

	"github.com/ecintiawan/loan-service/internal/constant"
)

func TestLoan_IsValid(t *testing.T) {
	type fields struct {
		ID                 int64
		BorrowerID         int64
		Amount             float64
		Rate               float64
		ApprovalProofURL   string
		AgreementLetterURL string
		Status             constant.LoanStatus
		CreatedBy          int64
		ApprovedBy         int64
		DisbursedBy        int64
		CreatedAt          time.Time
		UpdatedAt          time.Time
		ApprovedAt         time.Time
		InvestedAt         time.Time
		DisbursedAt        time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "valid",
			fields: fields{
				BorrowerID: 1,
				Amount:     10000,
				Rate:       10,
			},
			want: true,
		},
		{
			name: "invalid",
			fields: fields{
				BorrowerID: 1,
				Amount:     0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &Loan{
				ID:                 tt.fields.ID,
				BorrowerID:         tt.fields.BorrowerID,
				Amount:             tt.fields.Amount,
				Rate:               tt.fields.Rate,
				ApprovalProofURL:   tt.fields.ApprovalProofURL,
				AgreementLetterURL: tt.fields.AgreementLetterURL,
				Status:             tt.fields.Status,
				CreatedBy:          tt.fields.CreatedBy,
				ApprovedBy:         tt.fields.ApprovedBy,
				DisbursedBy:        tt.fields.DisbursedBy,
				CreatedAt:          tt.fields.CreatedAt,
				UpdatedAt:          tt.fields.UpdatedAt,
				ApprovedAt:         tt.fields.ApprovedAt,
				InvestedAt:         tt.fields.InvestedAt,
				DisbursedAt:        tt.fields.DisbursedAt,
			}
			if got := data.IsValid(); got != tt.want {
				t.Errorf("Loan.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoanProceed_IsValid(t *testing.T) {
	type fields struct {
		Action          constant.LoanAction
		ApprovalProof   File
		AgreementLetter File
		Data            *Loan
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "valid",
			fields: fields{
				Action: 1,
				Data: &Loan{
					ID: 1,
				},
			},
			want: true,
		},
		{
			name: "invalid",
			fields: fields{
				Action: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &LoanProceed{
				Action:          tt.fields.Action,
				ApprovalProof:   tt.fields.ApprovalProof,
				AgreementLetter: tt.fields.AgreementLetter,
				Data:            tt.fields.Data,
			}
			if got := req.IsValid(); got != tt.want {
				t.Errorf("LoanProceed.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoanProceed_IsValidExt(t *testing.T) {
	type fields struct {
		Action          constant.LoanAction
		ApprovalProof   File
		AgreementLetter File
		Data            *Loan
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "valid",
			fields: fields{
				ApprovalProof: File{
					File:    []byte{123},
					FileExt: ".png",
				},
			},
			want: true,
		},
		{
			name: "invalid",
			fields: fields{
				AgreementLetter: File{
					File:    []byte{123},
					FileExt: ".png",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &LoanProceed{
				Action:          tt.fields.Action,
				ApprovalProof:   tt.fields.ApprovalProof,
				AgreementLetter: tt.fields.AgreementLetter,
				Data:            tt.fields.Data,
			}
			if got := req.IsValidExt(); got != tt.want {
				t.Errorf("LoanProceed.IsValidExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoanFilter_Validate(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name   string
		fields fields
		want   *LoanFilter
	}{
		{
			name:   "success",
			fields: fields{},
			want: &LoanFilter{
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
			filter := &LoanFilter{
				DataTable:      tt.fields.DataTable,
				ID:             tt.fields.ID,
				BorrowerID:     tt.fields.BorrowerID,
				Status:         tt.fields.Status,
				CreatedAtStart: tt.fields.CreatedAtStart,
				CreatedAtEnd:   tt.fields.CreatedAtEnd,
				UpdatedAtStart: tt.fields.UpdatedAtStart,
				UpdatedAtEnd:   tt.fields.UpdatedAtEnd,
				ApprovedBy:     tt.fields.ApprovedBy,
				DisbursedBy:    tt.fields.DisbursedBy,
			}
			filter.Validate()
			if !reflect.DeepEqual(filter, tt.want) {
				t.Errorf("LoanFilter.Validate() = %v, want %v", filter, tt.want)
			}
		})
	}
}
