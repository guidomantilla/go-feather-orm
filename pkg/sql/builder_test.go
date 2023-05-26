package sql

import (
	"testing"
)

func TestCreateSelectSQL(t *testing.T) {
	type model struct {
		Id   string `db_column:"id,pk,generated"`
		Name string `db_column:"name"`
	}
	type args struct {
		table       string
		value       any
		in0         DriverName
		paramHolder ParamHolder
		fn01        ColumnFilterFunc
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Err sequence Path",
			args: args{
				value: nil,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Happy ColumnFilterFunc is Nil Path",
			args: args{
				in0:   0,
				table: "some",
				value: model{},
				fn01:  nil,
			},
			want:    "SELECT id, name FROM some",
			wantErr: false,
		},
		{
			name: "Err whereSequence Path",
			args: args{
				in0:   0,
				table: "some",
				value: model{},
				fn01:  NoneColumnFilter,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Happy Path",
			args: args{
				in0:         0,
				table:       "some",
				value:       model{},
				paramHolder: NamedParamHolder,
				fn01:        ColumnFilter,
			},
			want:    "SELECT id, name FROM some WHERE id = :id AND name = :name",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateSelectSQL(tt.args.table, tt.args.value, tt.args.in0, tt.args.paramHolder, tt.args.fn01)
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
	type model struct {
		Id string `db_column:"id,pk"`
	}
	type model2 struct {
		Id   string `db_column:"id,pk,generated"`
		Name string `db_column:"name"`
	}
	type args struct {
		table       string
		value       any
		driverName  DriverName
		paramHolder ParamHolder
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Err nameSequence Path",
			args: args{
				value: nil,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Err valueSequence Path",
			args: args{
				driverName: 0,
				table:      "some",
				value:      model{},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Happy Non Oracle Path",
			args: args{
				driverName:  MysqlDriverName,
				table:       "some",
				paramHolder: NamedParamHolder,
				value:       model{},
			},
			want:    "INSERT INTO some (id) VALUES (:id)",
			wantErr: false,
		},
		{
			name: "Err pkNameSequence Path",
			args: args{
				driverName: OracleDriverName,
				table:      "some",
				value:      model{},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Happy Oracle Path",
			args: args{
				driverName:  OracleDriverName,
				paramHolder: NamedParamHolder,
				table:       "some",
				value:       model2{},
			},
			want:    "INSERT INTO some (name) VALUES (:name) RETURNING id INTO :id",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateInsertSQL(tt.args.table, tt.args.value, tt.args.driverName, tt.args.paramHolder)
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
	type model struct {
		Id string `db_column:"id"`
	}
	type args struct {
		table       string
		value       any
		in0         DriverName
		paramHolder ParamHolder
		fn01        ColumnFilterFunc
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Err nameSequence Path",
			args: args{
				in0:   0,
				value: nil,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Err valueSequence Path",
			args: args{
				in0:   0,
				table: "some",
				value: model{},
				fn01:  nil,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Happy Path",
			args: args{
				in0:         0,
				table:       "some",
				paramHolder: NamedParamHolder,
				value:       model{},
				fn01:        ColumnFilter,
			},
			want:    "UPDATE some SET id = :id WHERE id = :id",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateUpdateSQL(tt.args.table, tt.args.value, tt.args.in0, tt.args.paramHolder, tt.args.fn01)
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
	type model struct {
		Id string `db_column:"id"`
	}
	type args struct {
		table       string
		value       any
		in0         DriverName
		paramHolder ParamHolder
		fn01        ColumnFilterFunc
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Err ParseColumnAsNameValueSequence Path",
			args: args{
				in0:   0,
				value: nil,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Happy Path",
			args: args{
				in0:         0,
				table:       "some",
				paramHolder: NamedParamHolder,
				value:       model{},
				fn01:        ColumnFilter,
			},
			want:    "DELETE FROM some WHERE id = :id",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateDeleteSQL(tt.args.table, tt.args.value, tt.args.in0, tt.args.paramHolder, tt.args.fn01)
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
