package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataTableFilter_Validate(t *testing.T) {
	tests := []struct {
		name          string
		filter        DataTableFilter
		wantDirection string
		wantLimit     int64
		wantOffset    int64
	}{
		{
			name: "success",
			filter: DataTableFilter{
				Pagination: DataTablePagination{
					Page: 9,
				},
			},
			wantDirection: "asc",
			wantLimit:     10,
			wantOffset:    80,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.filter.Validate()
			assert.Equal(t, tt.wantDirection, tt.filter.Sort.Direction)
			assert.Equal(t, tt.wantLimit, tt.filter.Pagination.Limit)
			assert.Equal(t, tt.wantOffset, tt.filter.Pagination.Offset)
		})
	}
}

func TestDataTableFilter_IsPaginated(t *testing.T) {
	tests := []struct {
		name   string
		filter DataTableFilter
		want   bool
	}{
		{
			name: "not paginated",
			filter: DataTableFilter{
				Pagination: DataTablePagination{
					DisablePagination: true,
				},
			},
			want: false,
		},
		{
			name: "paginated",
			filter: DataTableFilter{
				Pagination: DataTablePagination{
					DisablePagination: false,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.filter.IsPaginated()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDataTableSort_Validate(t *testing.T) {
	tests := []struct {
		name string
		sort DataTableSort
		want string
	}{
		{
			name: "empty direction",
			sort: DataTableSort{
				Direction: "",
			},
			want: "asc",
		},
		{
			name: "success",
			sort: DataTableSort{
				Direction: "DESC",
			},
			want: "DESC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sort.validate()
			assert.Equal(t, tt.want, tt.sort.Direction)
		})
	}
}

func TestDataTablePagination_Validate(t *testing.T) {
	tests := []struct {
		name       string
		pagination DataTablePagination
		wantPage   int64
		wantLimit  int64
		wantOffset int64
	}{
		{
			name: "pagination disabled",
			pagination: DataTablePagination{
				DisablePagination: true,
			},
			wantPage:   0,
			wantLimit:  0,
			wantOffset: 0,
		},
		{
			name: "empty page",
			pagination: DataTablePagination{
				Page:  0,
				Limit: 0,
			},
			wantPage:   1,
			wantLimit:  10,
			wantOffset: 0,
		},
		{
			name: "empty limit",
			pagination: DataTablePagination{
				Page:  6,
				Limit: 0,
			},
			wantPage:   6,
			wantLimit:  10,
			wantOffset: 50,
		},
		{
			name: "limit is too big",
			pagination: DataTablePagination{
				Page:  1,
				Limit: 1500,
			},
			wantPage:   1,
			wantLimit:  300,
			wantOffset: 0,
		},
		{
			name: "success",
			pagination: DataTablePagination{
				Page:  2,
				Limit: 25,
			},
			wantPage:   2,
			wantLimit:  25,
			wantOffset: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pagination.validate()
			assert.Equal(t, tt.wantPage, tt.pagination.Page)
			assert.Equal(t, tt.wantLimit, tt.pagination.Limit)
			assert.Equal(t, tt.wantOffset, tt.pagination.Offset)
		})
	}
}

func TestDataTablePagination_SetOffset(t *testing.T) {
	tests := []struct {
		name       string
		pagination DataTablePagination
		want       int64
	}{
		{
			name: "empty page",
			pagination: DataTablePagination{
				Page:  0,
				Limit: 10,
			},
			want: 0,
		},
		{
			name: "success",
			pagination: DataTablePagination{
				Page:  3,
				Limit: 30,
			},
			want: 60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pagination.setOffset()
			assert.Equal(t, tt.want, tt.pagination.Offset)
		})
	}
}

func TestDataTablePagination_IsDisabled(t *testing.T) {
	tests := []struct {
		name       string
		pagination DataTablePagination
		want       bool
	}{
		{
			name: "pagination disabled",
			pagination: DataTablePagination{
				DisablePagination: false,
			},
			want: false,
		},
		{
			name: "pagination enabled",
			pagination: DataTablePagination{
				DisablePagination: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.pagination.isDisabled()
			assert.Equal(t, tt.want, got)
		})
	}
}
