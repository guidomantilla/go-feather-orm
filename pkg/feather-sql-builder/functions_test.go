package feather_sql_builder

import (
	"testing"

	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/feather-sql"
	feather_sql_parsing "github.com/guidomantilla/go-feather-sql/pkg/feather-sql-parsing"
	feather_sql_reflection "github.com/guidomantilla/go-feather-sql/pkg/feather-sql-reflection"
)

func TestCreateSelectSQL(t *testing.T) {
	type args struct {
		in0   feather_sql.DriverName
		table string
		value any
		fn01  feather_sql_reflection.ColumnFilterFunc
		fn02  feather_sql_parsing.EvalColumnFunc
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateSelectSQL(tt.args.in0, tt.args.table, tt.args.value, tt.args.fn01, tt.args.fn02)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSelectSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateSelectSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateInsertSQL(t *testing.T) {
	type args struct {
		driverName feather_sql.DriverName
		table      string
		value      any
		fn02       feather_sql_parsing.EvalColumnFunc
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateInsertSQL(tt.args.driverName, tt.args.table, tt.args.value, tt.args.fn02)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateInsertSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateInsertSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateUpdateSQL(t *testing.T) {
	type args struct {
		in0   feather_sql.DriverName
		table string
		value any
		fn01  feather_sql_reflection.ColumnFilterFunc
		fn02  feather_sql_parsing.EvalColumnFunc
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateUpdateSQL(tt.args.in0, tt.args.table, tt.args.value, tt.args.fn01, tt.args.fn02)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUpdateSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateUpdateSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateDeleteSQL(t *testing.T) {
	type args struct {
		in0   feather_sql.DriverName
		table string
		value any
		fn01  feather_sql_reflection.ColumnFilterFunc
		fn02  feather_sql_parsing.EvalColumnFunc
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateDeleteSQL(tt.args.in0, tt.args.table, tt.args.value, tt.args.fn01, tt.args.fn02)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDeleteSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateDeleteSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
