package sqlbuilder

import "fmt"

func NewBuilder() Builder {
	return &builderImpl{
		counter:         0,
		whereClause:     "",
		updateSetClause: "",
		args:            []interface{}{},
	}
}

func (b *builderImpl) WhereClause() string {
	return b.whereClause
}

func (b *builderImpl) UpdateSetClause() string {
	return b.updateSetClause
}

func (b *builderImpl) Args() []interface{} {
	return b.args
}

func (b *builderImpl) AddWhereClause(column, sign string, value interface{}) {
	b.counter++
	b.whereClause += fmt.Sprintf(" AND %s %s $%d", column, sign, b.counter)
	b.args = append(b.args, value)
}

func (b *builderImpl) AddUpdateSetClause(column, value interface{}) {
	b.counter++
	b.updateSetClause += fmt.Sprintf(", %s = $%d", column, b.counter)
	b.args = append(b.args, value)
}
