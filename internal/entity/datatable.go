package entity

import "strings"

type (
	// Pagination used on API response struct
	Pagination struct {
		Count int64 `json:"count"`
		Row   int64 `json:"row"`

		// offset pagination
		Page int64 `json:"page"`
	}

	// DataTableFilter stores sorting and pagination options
	DataTableFilter struct {
		Sort       DataTableSort
		Pagination DataTablePagination
	}

	// DataTableSort stores field and direction for sql sorting
	DataTableSort struct {
		Field     string
		Direction string
	}

	// DataTablePagination caters for both cursor and offset pagination
	DataTablePagination struct {
		DisablePagination bool
		Limit             int64
		Cursor            int64 // cursor pagination, for big table
		Page, Offset      int64 // offset pagination
	}
)

var (
	// GetByIDFilter is default datatable filter on fetching row by id
	GetByIDFilter = DataTableFilter{
		Pagination: DataTablePagination{
			DisablePagination: true,
			Limit:             1,
		},
	}
)

// Validate self corrects both general sorting and pagination filter
func (filter *DataTableFilter) Validate() {
	filter.Sort.validate()
	filter.Pagination.validate()
}

// IsPaginated returns true if pagination is used
func (filter *DataTableFilter) IsPaginated() bool {
	return !filter.Pagination.isDisabled()
}

// Validate self corrects direction
func (sort *DataTableSort) validate() {
	// set default direction to ascending
	if !strings.EqualFold(sort.Direction, "desc") {
		sort.Direction = "asc"
	}
}

// Validate self corrects limit and sets offset
func (pagination *DataTablePagination) validate() {
	// no need to validate limit and offset if pagination is not used
	if pagination.isDisabled() {
		return
	}

	// set default page to 1
	if !(pagination.Page > 0) {
		pagination.Page = 1
	}

	// set default limit to 10
	if !(pagination.Limit > 0) {
		pagination.Limit = 10
	}

	// set upper limit to 300 only
	if pagination.Limit > 300 {
		pagination.Limit = 300
	}

	// calculate offset value
	pagination.setOffset()
}

// setOffset calculates offset from given page and limit
func (pagination *DataTablePagination) setOffset() {
	// only calculate when page is greater than zero to avoid negative offset
	if pagination.Page > 0 {
		pagination.Offset = (pagination.Page - 1) * pagination.Limit
	}
}

// IsDisabled determines whether pagination is used or not
func (pagination *DataTablePagination) isDisabled() bool {
	return pagination.DisablePagination
}
