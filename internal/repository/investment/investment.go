package investment

import (
	"context"
	"fmt"

	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/internal/repository"
	"github.com/ecintiawan/loan-service/pkg/database"
	"github.com/ecintiawan/loan-service/pkg/errorwrapper"
	"github.com/ecintiawan/loan-service/pkg/sqlbuilder"
	"github.com/jackc/pgx/v5"
)

type (
	// repoImpl implements Investment interface
	repoImpl struct {
		client database.DB
	}
)

// New creates a new instance of repoImpl
func New(client database.DB) repository.Investment {
	return &repoImpl{
		client: client,
	}
}

// Get will return investment data based on filter
func (r *repoImpl) Get(
	ctx context.Context,
	filter *entity.InvestmentFilter,
) (entity.InvestmentResult, error) {
	var (
		result = entity.InvestmentResult{
			List: []*entity.Investment{},
			Pagination: entity.Pagination{
				Page: filter.DataTable.Pagination.Page,
				Row:  filter.DataTable.Pagination.Limit,
			},
		}
		builder = sqlbuilder.NewBuilder()
		err     error
	)

	if filter.ID > 0 {
		builder.AddWhereClause("id", "=", filter.ID)
	}

	if filter.InvestorID > 0 {
		builder.AddWhereClause("investor_id", "=", filter.InvestorID)
	}

	if filter.LoanID > 0 {
		builder.AddWhereClause("loan_id", "=", filter.LoanID)
	}

	if filter.Status > 0 {
		builder.AddWhereClause("status", "=", filter.Status)
	}

	if !filter.CreatedAtStart.IsZero() {
		builder.AddWhereClause("created_at", ">=", filter.CreatedAtStart)
	}
	if !filter.CreatedAtEnd.IsZero() {
		builder.AddWhereClause("created_at", "<=", filter.CreatedAtEnd)
	}

	if !filter.UpdatedAtStart.IsZero() {
		builder.AddWhereClause("updated_at", ">=", filter.UpdatedAtStart)
	}
	if !filter.UpdatedAtEnd.IsZero() {
		builder.AddWhereClause("updated_at", "<=", filter.UpdatedAtEnd)
	}

	if filter.DataTable.IsPaginated() {
		countQuery := fmt.Sprintf(
			`SELECT COUNT(*) FROM investment WHERE 1 = 1 %s`,
			builder.WhereClause(),
		)
		err = r.client.QueryRow(ctx, countQuery, builder.Args()...).Scan(&result.Count)
		if err != nil {
			return result, errorwrapper.E(err, errorwrapper.CodeInternal)
		}
	}

	query := fmt.Sprintf(`
		SELECT
			id,
			investor_id,
			loan_id,
			amount,
			roi,
			status,
			created_at,
			COALESCE(updated_at, '0001-01-01 00:00:00'::timestamp)
		FROM
			investment
		WHERE
			1 = 1
			%s`,
		builder.WhereClause(),
	)

	if filter.DataTable.IsPaginated() {
		query = fmt.Sprintf(
			"%s ORDER BY %s %s LIMIT %d offset %d",
			query,
			filter.DataTable.Sort.Field,
			filter.DataTable.Sort.Direction,
			filter.DataTable.Pagination.Limit,
			filter.DataTable.Pagination.Offset,
		)
	}

	var rows pgx.Rows
	rows, err = r.client.Query(ctx, query, builder.Args()...)
	if err != nil {
		return result, errorwrapper.E(err, errorwrapper.CodeInternal)
	}
	defer rows.Close()

	for rows.Next() {
		var investment = &entity.Investment{}
		err = rows.Scan(
			&investment.ID,
			&investment.InvestorID,
			&investment.LoanID,
			&investment.Amount,
			&investment.ROI,
			&investment.Status,
			&investment.CreatedAt,
			&investment.UpdatedAt,
		)
		if err != nil {
			return result, errorwrapper.E(err, errorwrapper.CodeInternal)
		}

		result.List = append(result.List, investment)
	}
	err = rows.Err()
	if err != nil {
		return result, errorwrapper.E(err, errorwrapper.CodeInternal)
	}

	return result, nil
}

// GetAmountSum will return investment amount sum data based on filter
func (r *repoImpl) GetAmountSum(
	ctx context.Context,
	filter *entity.InvestmentFilter,
) (float64, error) {
	var (
		result  float64
		builder = sqlbuilder.NewBuilder()
		err     error
	)

	if filter.ID > 0 {
		builder.AddWhereClause("id", "=", filter.ID)
	}

	if filter.InvestorID > 0 {
		builder.AddWhereClause("investor_id", "=", filter.InvestorID)
	}

	if filter.LoanID > 0 {
		builder.AddWhereClause("loan_id", "=", filter.LoanID)
	}

	if filter.Status > 0 {
		builder.AddWhereClause("status", "=", filter.Status)
	}

	if !filter.CreatedAtStart.IsZero() {
		builder.AddWhereClause("created_at", ">=", filter.CreatedAtStart)
	}
	if !filter.CreatedAtEnd.IsZero() {
		builder.AddWhereClause("created_at", "<=", filter.CreatedAtEnd)
	}

	if !filter.UpdatedAtStart.IsZero() {
		builder.AddWhereClause("updated_at", ">=", filter.UpdatedAtStart)
	}
	if !filter.UpdatedAtEnd.IsZero() {
		builder.AddWhereClause("updated_at", "<=", filter.UpdatedAtEnd)
	}

	query := fmt.Sprintf(`
		SELECT
			COALESCE(SUM(amount), 0)
		FROM
			investment
		WHERE
			1 = 1
			%s`,
		builder.WhereClause(),
	)

	err = r.client.QueryRow(ctx, query, builder.Args()...).Scan(&result)
	if err != nil {
		return result, errorwrapper.E(err, errorwrapper.CodeInternal)
	}

	return result, nil
}

// Create will insert initial investment data
func (r *repoImpl) Create(
	ctx context.Context,
	model *entity.Investment,
) error {
	var (
		err error
	)

	tx, err := r.client.Begin(ctx)
	if err != nil {
		return errorwrapper.E(err, errorwrapper.CodeInternal)
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	query := `
		INSERT INTO investment (
			investor_id,
			loan_id,
			amount,
			roi,
			status,
			created_at
		)
		VALUES (
			$1,
			$2,
			$3,
		    $4,
		    $5,
		    NOW()
		)
	`

	_, err = tx.Exec(
		ctx,
		query,
		model.InvestorID,
		model.LoanID,
		model.Amount,
		model.ROI,
		model.Status,
	)
	if err != nil {
		return errorwrapper.E(err, errorwrapper.CodeInternal)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return errorwrapper.E(err, errorwrapper.CodeInternal)
	}

	return nil
}
