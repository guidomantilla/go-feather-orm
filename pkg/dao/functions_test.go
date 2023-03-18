package dao

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
)

func TestWriteContext(t *testing.T) {
	type args struct {
		ctx          context.Context
		sqlStatement string
		args         []any
	}
	tests := []struct {
		name    string
		args    args
		want    *int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WriteContext(tt.args.ctx, tt.args.sqlStatement, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WriteContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadContext(t *testing.T) {
	type args struct {
		ctx          context.Context
		sqlStatement string
		fn           ReadFunction
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadContext(tt.args.ctx, tt.args.sqlStatement, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("ReadContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadRowContext(t *testing.T) {
	type args struct {
		ctx          context.Context
		sqlStatement string
		key          any
		dest         []any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadRowContext(tt.args.ctx, tt.args.sqlStatement, tt.args.key, tt.args.dest...); (err != nil) != tt.wantErr {
				t.Errorf("ReadRowContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContext(t *testing.T) {
	type args struct {
		ctx          context.Context
		sqlStatement string
		fn           Function
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Context(tt.args.ctx, tt.args.sqlStatement, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("Context() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_closeStatement(t *testing.T) {
	type args struct {
		statement *sql.Stmt
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			closeStatement(tt.args.statement)
		})
	}
}

func Test_closeResultSet(t *testing.T) {
	type args struct {
		rows *sql.Rows
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			closeResultSet(tt.args.rows)
		})
	}
}
