package loan

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
	// repoImpl implements Loan interface
	repoImpl struct {
		client database.DB
	}
)

// New creates a new instance of repoImpl
func New(client database.DB) repository.Loan {
	return &repoImpl{
		client: client,
	}
}

// Get will return loan data based on filter
func (r *repoImpl) Get(
	ctx context.Context,
	filter *entity.LoanFilter,
) (entity.LoanResult, error) {
	var (
		result = entity.LoanResult{
			List: []*entity.Loan{},
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

	if filter.BorrowerID > 0 {
		builder.AddWhereClause("borrower_id", "=", filter.BorrowerID)
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

	if filter.ApprovedBy > 0 {
		builder.AddWhereClause("approved_by", "=", filter.ApprovedBy)
	}

	if filter.DisbursedBy > 0 {
		builder.AddWhereClause("disbursed_by", "=", filter.DisbursedBy)
	}

	if filter.DataTable.IsPaginated() {
		countQuery := fmt.Sprintf(
			`SELECT COUNT(*) FROM loan WHERE 1 = 1 %s`,
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
			borrower_id,
			amount,
			rate,
			approval_proof_url,
			agreement_letter_url,
			status,
			created_by,
			approved_by,
			disbursed_by,
			created_at,
			COALESCE(updated_at, '0001-01-01 00:00:00'::timestamp),
			COALESCE(approved_at, '0001-01-01 00:00:00'::timestamp),
			COALESCE(invested_at, '0001-01-01 00:00:00'::timestamp),
			COALESCE(disbursed_at, '0001-01-01 00:00:00'::timestamp)
		FROM
			loan
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
		var loan = &entity.Loan{}
		err = rows.Scan(
			&loan.ID,
			&loan.BorrowerID,
			&loan.Amount,
			&loan.Rate,
			&loan.ApprovalProofURL,
			&loan.AgreementLetterURL,
			&loan.Status,
			&loan.CreatedBy,
			&loan.ApprovedBy,
			&loan.DisbursedBy,
			&loan.CreatedAt,
			&loan.UpdatedAt,
			&loan.ApprovedAt,
			&loan.InvestedAt,
			&loan.DisbursedAt,
		)
		if err != nil {
			return result, errorwrapper.E(err, errorwrapper.CodeInternal)
		}

		result.List = append(result.List, loan)
	}
	err = rows.Err()
	if err != nil {
		return result, errorwrapper.E(err, errorwrapper.CodeInternal)
	}

	return result, nil
}

// GetDetail will return loan data based on filter
func (r *repoImpl) GetDetail(
	ctx context.Context,
	id int64,
) (*entity.Loan, error) {
	var (
		err error
	)

	list, err := r.Get(ctx, &entity.LoanFilter{
		DataTable: entity.DataTableFilter{
			Pagination: entity.DataTablePagination{
				DisablePagination: true,
			},
		},
		ID: id,
	})
	if err != nil {
		return nil, errorwrapper.E(err, errorwrapper.CodeInternal)
	}
	if len(list.List) <= 0 {
		return nil, errorwrapper.E("data does not exist", errorwrapper.CodeNotFound)
	}

	return list.List[0], nil
}

// Create will insert initial loan data
func (r *repoImpl) Create(
	ctx context.Context,
	model *entity.Loan,
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
		INSERT INTO loan (
			borrower_id,
			amount,
			rate,
			approval_proof_url,
			agreement_letter_url,
			status,
			created_by,
			approved_by,
			disbursed_by,
			created_at,
			approved_at,
			invested_at,
			disbursed_at
		)
		VALUES (
			$1,
			$2,
			$3,
		    $4,
		    $5,
		    $6,
		    $7,
		    $8,
		    $9,
		    NOW(),
		    $10,
		    $11,
			$12
		)
	`

	_, err = tx.Exec(
		ctx,
		query,
		model.BorrowerID,
		model.Amount,
		model.Rate,
		model.ApprovalProofURL,
		model.AgreementLetterURL,
		model.Status,
		model.CreatedBy,
		model.ApprovedBy,
		model.DisbursedBy,
		model.ApprovedAt,
		model.InvestedAt,
		model.DisbursedAt,
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

// Proceed is an action to go through all loan states for certain loan data
func (r *repoImpl) Update(
	ctx context.Context,
	model *entity.Loan,
) error {
	var (
		err     error
		builder = sqlbuilder.NewBuilder()
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

	if model.ID > 0 {
		builder.AddWhereClause("id", "=", model.ID)
	}

	if model.ApprovalProofURL != "" {
		builder.AddUpdateSetClause("approval_proof_url", model.ApprovalProofURL)
	}

	if model.AgreementLetterURL != "" {
		builder.AddUpdateSetClause("agreement_letter_url", model.AgreementLetterURL)
	}

	if model.Status > 0 {
		builder.AddUpdateSetClause("status", model.Status)
	}

	if model.ApprovedBy > 0 {
		builder.AddUpdateSetClause("approved_by", model.ApprovedBy)
	}

	if model.DisbursedBy > 0 {
		builder.AddUpdateSetClause("disbursed_by", model.DisbursedBy)
	}

	if !model.ApprovedAt.IsZero() {
		builder.AddUpdateSetClause("approved_at", model.ApprovedAt)
	}

	if !model.InvestedAt.IsZero() {
		builder.AddUpdateSetClause("invested_at", model.InvestedAt)
	}

	if !model.DisbursedAt.IsZero() {
		builder.AddUpdateSetClause("disbursed_at", model.DisbursedAt)
	}

	query := fmt.Sprintf(`
		UPDATE
			loan
		SET
			updated_at = NOW()
			%s
		WHERE
			1 = 1
			%s
	`, builder.UpdateSetClause(), builder.WhereClause())
	_, err = tx.Exec(ctx, query, builder.Args()...)
	if err != nil {
		return errorwrapper.E(err, errorwrapper.CodeInternal)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return errorwrapper.E(err, errorwrapper.CodeInternal)
	}

	return nil
}
