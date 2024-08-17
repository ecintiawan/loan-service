package sqlbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBuilder(t *testing.T) {
	tests := []struct {
		name string
		want *builderImpl
	}{
		{
			name: "success",
			want: &builderImpl{
				counter:     0,
				whereClause: "",
				args:        []interface{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBuilder()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBuilderImpl_WhereClause(t *testing.T) {
	tests := []struct {
		name string
		b    *builderImpl
		want string
	}{
		{
			name: "success",
			b: &builderImpl{
				whereClause: "where clause",
			},
			want: "where clause",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.b.WhereClause()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBuilderImpl_UpdateSetClause(t *testing.T) {
	tests := []struct {
		name string
		b    *builderImpl
		want string
	}{
		{
			name: "success",
			b: &builderImpl{
				updateSetClause: "update set clause",
			},
			want: "update set clause",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.b.UpdateSetClause()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBuilderImpl_Args(t *testing.T) {
	tests := []struct {
		name string
		b    *builderImpl
		want []interface{}
	}{
		{
			name: "success",
			b: &builderImpl{
				args: []interface{}{1, 2, 3, "test1", "test2"},
			},
			want: []interface{}{1, 2, 3, "test1", "test2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.b.Args()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBuilderImpl_AddWhereClause(t *testing.T) {
	type args struct {
		column string
		sign   string
		value  interface{}
	}
	tests := []struct {
		name      string
		b         *builderImpl
		args      args
		wantWhere string
		wantArgs  []interface{}
	}{
		{
			name: "success",
			b: &builderImpl{
				counter:     0,
				whereClause: "",
				args:        []interface{}{},
			},
			args: args{
				column: "id",
				sign:   "=",
				value:  5,
			},
			wantWhere: " AND id = $1",
			wantArgs:  []interface{}{5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.AddWhereClause(tt.args.column, tt.args.sign, tt.args.value)
			assert.Equal(t, tt.wantWhere, tt.b.WhereClause())
			assert.Equal(t, tt.wantArgs, tt.b.Args())
		})
	}
}

func TestBuilderImpl_AddUpdateSetClause(t *testing.T) {
	type args struct {
		column string
		value  interface{}
	}
	tests := []struct {
		name          string
		b             *builderImpl
		args          args
		wantUpdateSet string
		wantArgs      []interface{}
	}{
		{
			name: "success",
			b: &builderImpl{
				counter:         0,
				updateSetClause: "",
				args:            []interface{}{},
			},
			args: args{
				column: "id",
				value:  5,
			},
			wantUpdateSet: ", id = $1",
			wantArgs:      []interface{}{5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.AddUpdateSetClause(tt.args.column, tt.args.value)
			assert.Equal(t, tt.wantUpdateSet, tt.b.UpdateSetClause())
			assert.Equal(t, tt.wantArgs, tt.b.Args())
		})
	}
}
