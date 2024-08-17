package sqlbuilder

type (
	Builder interface {
		WhereClause() string
		UpdateSetClause() string
		Args() []interface{}
		AddWhereClause(column, sign string, value interface{})
		AddUpdateSetClause(column, value interface{})
	}
	builderImpl struct {
		counter         int
		whereClause     string
		updateSetClause string
		args            []interface{}
	}
)
