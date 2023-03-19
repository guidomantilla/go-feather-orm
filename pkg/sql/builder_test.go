package sql

import (
	"testing"
)

func TestCreateSelectSQL(t *testing.T) {
	type model struct {
		Id   string `sql:"id,pk,generated"`
		Name string `sql:"name"`
	}
	type args struct {
		in0   DriverName
		table string
		value any
		fn01  ColumnFilterFunc
		fn02  EvalColumnFunc
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
				fn01:  ColumnFilter,
				fn02:  nil,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Happy Path",
			args: args{
				in0:   0,
				table: "some",
				value: model{},
				fn01:  ColumnFilter,
				fn02:  EvalValueOnlyNamed,
			},
			want:    "SELECT id, name FROM some WHERE :id AND :name",
			wantErr: false,
		},
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
	type model struct {
		Id string `sql:"id,pk"`
	}
	type model2 struct {
		Id   string `sql:"id,pk,generated"`
		Name string `sql:"name"`
	}
	type args struct {
		driverName DriverName
		table      string
		value      any
		fn02       EvalColumnFunc
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
				fn02:       nil,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Happy Non Oracle Path",
			args: args{
				driverName: MysqlDriverName,
				table:      "some",
				value:      model{},
				fn02:       EvalNameValueNumbered,
			},
			want:    "INSERT INTO some (id) VALUES (id = :1)",
			wantErr: false,
		},
		{
			name: "Err pkNameSequence Path",
			args: args{
				driverName: OracleDriverName,
				table:      "some",
				value:      model{},
				fn02:       EvalNameValueNumbered,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Happy Oracle Path",
			args: args{
				driverName: OracleDriverName,
				table:      "some",
				value:      model2{},
				fn02:       EvalValueOnlyNumbered,
			},
			want:    "INSERT INTO some (name) VALUES (:1) RETURNING id INTO :2",
			wantErr: false,
		},
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
	type model struct {
		Id string `sql:"id"`
	}
	type args struct {
		in0   DriverName
		table string
		value any
		fn01  ColumnFilterFunc
		fn02  EvalColumnFunc
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
				fn02:  EvalNameValueNumbered,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Happy Path",
			args: args{
				in0:   0,
				table: "some",
				value: model{},
				fn01:  ColumnFilter,
				fn02:  EvalNameValueNumbered,
			},
			want:    "UPDATE some SET id = :1 WHERE id = :2",
			wantErr: false,
		},
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
	type model struct {
		Id string `sql:"id"`
	}
	type args struct {
		in0   DriverName
		table string
		value any
		fn01  ColumnFilterFunc
		fn02  EvalColumnFunc
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
				in0:   0,
				table: "some",
				value: model{},
				fn01:  ColumnFilter,
				fn02:  EvalNameValueNumbered,
			},
			want:    "DELETE FROM some WHERE id = :1",
			wantErr: false,
		},
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
