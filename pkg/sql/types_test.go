package sql

import (
	"reflect"
	"testing"
)

func TestDriverName_String(t *testing.T) {
	tests := []struct {
		name string
		enum DriverName
		want string
	}{
		{
			name: "UndefinedDriverName Path",
			enum: UndefinedDriverName,
			want: UndefinedDriverName.String(),
		},
		{
			name: "OracleDriverName Path",
			enum: OracleDriverName,
			want: OracleDriverName.String(),
		},
		{
			name: "MysqlDriverName Path",
			enum: MysqlDriverName,
			want: MysqlDriverName.String(),
		},
		{
			name: "PostgresDriverName Path",
			enum: PostgresDriverName,
			want: PostgresDriverName.String(),
		},
		{
			name: "Nil Path",
			enum: -2,
			want: UndefinedDriverName.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.String(); got != tt.want {
				t.Errorf("DriverName.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriverName_ValueOf(t *testing.T) {
	type args struct {
		driverName string
	}
	tests := []struct {
		name string
		enum DriverName
		args args
		want DriverName
	}{
		{
			name: "UndefinedDriverName Path",
			enum: UndefinedDriverName,
			args: args{driverName: UndefinedDriverName.String()},
			want: UndefinedDriverName,
		},
		{
			name: "OracleDriverName Path",
			enum: OracleDriverName,
			args: args{driverName: OracleDriverName.String()},
			want: OracleDriverName,
		},
		{
			name: "MysqlDriverName Path",
			enum: MysqlDriverName,
			args: args{driverName: MysqlDriverName.String()},
			want: MysqlDriverName,
		},
		{
			name: "PostgresDriverName Path",
			enum: PostgresDriverName,
			args: args{driverName: PostgresDriverName.String()},
			want: PostgresDriverName,
		},

		{
			name: "Nil Path",
			enum: -2,
			args: args{driverName: ""},
			want: UndefinedDriverName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.ValueFromName(tt.args.driverName); got != tt.want {
				t.Errorf("DriverName.ValueOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParamHolder_EvalNameValue(t *testing.T) {
	tests := []struct {
		name string
		enum ParamHolder
		want EvalColumnFunc
	}{
		{
			name: "NamedParamHolder Path",
			enum: NamedParamHolder,
			want: EvalNameValueNamed,
		},
		{
			name: "NumberedParamHolder Path",
			enum: NumberedParamHolder,
			want: EvalNameValueNumbered,
		},
		{
			name: "QuestionedParamHolder Path",
			enum: QuestionedParamHolder,
			want: EvalNameValueQuestioned,
		},
		{
			name: "Nil Path",
			enum: -1,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.EvalNameValue(); !reflect.DeepEqual(reflect.ValueOf(got).Pointer(), reflect.ValueOf(tt.want).Pointer()) {
				t.Errorf("ParamHolder.EvalNameValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParamHolder_EvalValueOnly(t *testing.T) {
	tests := []struct {
		name string
		enum ParamHolder
		want EvalColumnFunc
	}{
		{
			name: "NamedParamHolder Path",
			enum: NamedParamHolder,
			want: EvalValueOnlyNamed,
		},
		{
			name: "NumberedParamHolder Path",
			enum: NumberedParamHolder,
			want: EvalValueOnlyNumbered,
		},
		{
			name: "QuestionedParamHolder Path",
			enum: QuestionedParamHolder,
			want: EvalValueOnlyQuestioned,
		},
		{
			name: "Nil Path",
			enum: -1,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.EvalValueOnly(); !reflect.DeepEqual(reflect.ValueOf(got).Pointer(), reflect.ValueOf(tt.want).Pointer()) {
				t.Errorf("ParamHolder.EvalValueOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}
