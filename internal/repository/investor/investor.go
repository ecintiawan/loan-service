package investor

import (
	"context"
	"fmt"

	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/internal/repository"
	"github.com/ecintiawan/loan-service/pkg/database"
	"github.com/ecintiawan/loan-service/pkg/errorwrapper"
	"github.com/ecintiawan/loan-service/pkg/sqlbuilder"
)

type (
	// repoImpl implements Investor interface
	repoImpl struct {
		client database.DB
	}
)

// New creates a new instance of repoImpl
func New(client database.DB) repository.Investor {
	return &repoImpl{
		client: client,
	}
}

// GetDetail will return investor data based on filter
func (r *repoImpl) GetDetail(
	ctx context.Context,
	id int64,
) (*entity.Investor, error) {
	var (
		result  = &entity.Investor{}
		builder = sqlbuilder.NewBuilder()
		err     error
	)

	builder.AddWhereClause("id", "=", id)
	query := fmt.Sprintf(`
		SELECT
			id,
			identification_number,
			name,
			email,
			status,
			created_at,
			COALESCE(updated_at, '0001-01-01 00:00:00'::timestamp)
		FROM
			investor
		WHERE
			1 = 1
			%s`,
		builder.WhereClause(),
	)

	err = r.client.QueryRow(ctx, query, builder.Args()...).
		Scan(
			&result.ID,
			&result.IdentificationNumber,
			&result.Name,
			&result.Email,
			&result.Status,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	if err != nil {
		return result, errorwrapper.E(err, errorwrapper.CodeInternal)
	}

	return result, nil
}
