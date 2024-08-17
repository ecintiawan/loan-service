package database

import (
	"context"
	"errors"
	"reflect"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// MockPgxTx represents a mock implementation of pgx.Tx
type MockPgxTx struct {
	ExecFunc   func(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	CommitFunc func(ctx context.Context) error
}

func (m *MockPgxTx) Begin(ctx context.Context) (pgx.Tx, error) {
	return nil, nil
}
func (m *MockPgxTx) BeginFunc(ctx context.Context, f func(pgx.Tx) error) (err error) {
	return nil
}
func (m *MockPgxTx) Commit(ctx context.Context) error {
	if m.CommitFunc != nil {
		return m.CommitFunc(ctx)
	}
	return nil
}
func (m *MockPgxTx) Rollback(ctx context.Context) error {
	return nil
}

func (m *MockPgxTx) CopyFrom(
	ctx context.Context,
	tableName pgx.Identifier,
	columnNames []string,
	rowSrc pgx.CopyFromSource,
) (int64, error) {
	return 0, nil
}
func (m *MockPgxTx) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return nil
}
func (m *MockPgxTx) LargeObjects() pgx.LargeObjects {
	return pgx.LargeObjects{}
}

func (m *MockPgxTx) Prepare(
	ctx context.Context,
	name, sql string,
) (*pgconn.StatementDescription, error) {
	return nil, nil
}
func (m *MockPgxTx) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return nil, nil
}
func (m *MockPgxTx) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return nil
}
func (m *MockPgxTx) Conn() *pgx.Conn {
	return nil
}

func (m *MockPgxTx) Exec(
	ctx context.Context,
	sql string,
	args ...interface{},
) (pgconn.CommandTag, error) {
	if m.ExecFunc != nil {
		return m.ExecFunc(ctx, sql, args...)
	}
	return pgconn.CommandTag{}, nil
}

// MockPgxRows represents a mock implementation of pgx.Rows
type MockPgxRows struct {
	data     [][]interface{}
	columns  []string
	rowIndex int
	closed   bool
	errValue error
}

// NewMockPgxRows creates a new MockPgxRows instance
func NewMockPgxRows(columns []string, data [][]interface{}) *MockPgxRows {
	return &MockPgxRows{
		columns:  columns,
		data:     data,
		rowIndex: -1,
	}
}

func (m *MockPgxRows) Next() bool {
	m.rowIndex++
	return m.rowIndex < len(m.data)
}

func (m *MockPgxRows) Scan(dest ...interface{}) error {
	if m.rowIndex >= len(m.data) {
		return pgx.ErrNoRows
	}

	rowData := m.data[m.rowIndex]

	if len(rowData) != len(dest) {
		return errors.New("number of columns doesn't match number of destinations")
	}

	for i, val := range rowData {
		destVal := reflect.ValueOf(dest[i])
		if destVal.Kind() != reflect.Ptr {
			return errors.New("destination is not a pointer")
		}

		if reflect.TypeOf(val).AssignableTo(destVal.Elem().Type()) {
			destVal.Elem().Set(reflect.ValueOf(val))
		} else {
			return errors.New("value type does not match destination type: " + reflect.TypeOf(val).String() + " vs " + destVal.Elem().Type().String())
		}
	}

	return nil
}

func (m *MockPgxRows) Close() {
	m.closed = true
}

func (m *MockPgxRows) Err() error {
	return m.errValue
}

func (m *MockPgxRows) CommandTag() pgconn.CommandTag {
	// Implement if needed
	return pgconn.CommandTag{}
}

func (m *MockPgxRows) FieldDescriptions() []pgconn.FieldDescription {
	// Implement if needed
	return nil
}

func (m *MockPgxRows) RawValues() [][]byte {
	// Implement if needed
	return nil
}

func (m *MockPgxRows) Values() ([]interface{}, error) {
	// Implement if needed
	return nil, nil
}

func (m *MockPgxRows) Conn() *pgx.Conn {
	return nil
}

// MockPgxRow represents a mock implementation of pgx.Row
type MockPgxRow struct {
	columns []string
	values  []interface{}
	index   int
}

// NewMockPgxRow creates a new MockPgxRow instance
func NewMockPgxRow(columns []string, values []interface{}) *MockPgxRow {
	return &MockPgxRow{
		columns: columns,
		values:  values,
		index:   -1,
	}
}

func (m *MockPgxRow) Scan(dest ...interface{}) error {
	if m.index >= len(m.values)-1 {
		return pgx.ErrNoRows
	}
	m.index++

	if len(m.values) != len(dest) {
		return errors.New("number of columns doesn't match number of destinations")
	}

	for i, val := range m.values {
		destVal := reflect.ValueOf(dest[i])
		if destVal.Kind() != reflect.Ptr {
			return errors.New("destination is not a pointer")
		}

		if reflect.TypeOf(val).AssignableTo(destVal.Elem().Type()) {
			destVal.Elem().Set(reflect.ValueOf(val))
		} else {
			return errors.New("value type does not match destination type: " + reflect.TypeOf(val).String() + " vs " + destVal.Elem().Type().String())
		}
	}
	return nil
}

func (m *MockPgxRow) Close() error {
	return nil
}
