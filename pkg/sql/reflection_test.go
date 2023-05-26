package sql

import (
	"reflect"
	"testing"
)

func TestRetrieveColumnNames(t *testing.T) {

	type modelOK01 struct {
		Id   string
		Name string
	}
	type modelOK02 struct {
		Id   string `db_column:"-"`
		Name string `db_column:"-"`
	}

	type modelOK03 struct {
		Id   string `db_column:""`
		Name string `db_column:""`
	}

	type modelOK04 struct {
		Id   string `db_column:","`
		Name string `db_column:","`
	}

	type modelOK05 struct {
		Id   string `db_column:"id"`
		Name string `db_column:"name"`
	}

	type args struct {
		value            any
		columnFilterFunc ColumnFilterFunc
	}
	tests := []struct {
		name      string
		args      args
		want      []string
		wantErr   bool
		errWanted error
	}{
		{
			name:      "Err retrieveReflectedStruct Path",
			args:      args{},
			wantErr:   true,
			errWanted: ErrAnyIsNil,
		},

		{
			name: "Err columnFilterFunc Path",
			args: args{
				value:            modelOK01{},
				columnFilterFunc: nil,
			},
			wantErr:   true,
			errWanted: ErrColumnFilterFuncIsNil,
		},
		{
			name: "Ignore No Tag reflectedTagKey Path",
			args: args{
				value:            modelOK01{},
				columnFilterFunc: ColumnFilter,
			},
			want:    []string{},
			wantErr: false,
		},
		{
			name: "Ignore Dash reflectedTagKey Path",
			args: args{
				value:            modelOK02{},
				columnFilterFunc: ColumnFilter,
			},
			want:    []string{},
			wantErr: false,
		},
		{
			name: "Ignore Empty reflectedTagKey Path",
			args: args{
				value:            modelOK03{},
				columnFilterFunc: ColumnFilter,
			},
			want:    []string{},
			wantErr: false,
		},
		{
			name: "Empty GeneratedColumnFilter Path",
			args: args{
				value:            modelOK04{},
				columnFilterFunc: NoneColumnFilter,
			},
			want:    []string{},
			wantErr: false,
		},
		{
			name: "Not Empty GeneratedColumnFilter Path",
			args: args{
				value:            modelOK05{},
				columnFilterFunc: ColumnFilter,
			},
			want:    []string{"id", "name"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RetrieveColumnNames(tt.args.value, tt.args.columnFilterFunc)
			if tt.wantErr {
				if err != nil {
					if !reflect.DeepEqual(err, tt.errWanted) {
						t.Errorf("RetrieveColumnNames() error = %v, wantErr %v", err, tt.wantErr)
						return
					}
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("RetrieveColumnNames() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestKeyColumnFilter(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Happy Path",
			args: args{values: []string{"hello", TagPkValue}},
			want: true,
		},
		{
			name: "Unhappy Path",
			args: args{values: []string{TagPkValue, TagPkValue}},
			want: false,
		},
		{
			name: "Happy Path",
			args: args{values: []string{"hello", TagUqValue}},
			want: true,
		},
		{
			name: "Unhappy Path",
			args: args{values: []string{TagPkValue, TagUqValue}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeyColumnFilter(tt.args.values); got != tt.want {
				t.Errorf("KeyColumnFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNonePkGeneratedColumnFilter(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Unhappy Path",
			args: args{values: []string{TagPkValue, TagUqValue}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NonePkGeneratedColumnFilter(tt.args.values); got != tt.want {
				t.Errorf("NonePkGeneratedColumnFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoneUqGeneratedColumnFilter(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Unhappy Path",
			args: args{values: []string{TagPkValue, TagUqValue}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoneUqGeneratedColumnFilter(tt.args.values); got != tt.want {
				t.Errorf("NoneUqGeneratedColumnFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPkGeneratedColumnFilter(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Happy Path",
			args: args{values: []string{"hello", TagPkValue, TagGeneratedValue}},
			want: true,
		},
		{
			name: "Unhappy Path",
			args: args{values: []string{TagPkValue, TagUqValue}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PkGeneratedColumnFilter(tt.args.values); got != tt.want {
				t.Errorf("PkGeneratedColumnFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUqGeneratedColumnFilter(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Happy Path",
			args: args{values: []string{"hello", TagUqValue, TagGeneratedValue}},
			want: true,
		},
		{
			name: "Unhappy Path",
			args: args{values: []string{TagPkValue, TagUqValue}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UqGeneratedColumnFilter(tt.args.values); got != tt.want {
				t.Errorf("UqGeneratedColumnFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGeneratedColumnFilter(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Happy Path",
			args: args{values: []string{"hello", TagGeneratedValue}},
			want: true,
		},
		{
			name: "Happy Path",
			args: args{values: []string{"hello", TagPkValue, TagGeneratedValue}},
			want: true,
		},
		{
			name: "Unhappy Path",
			args: args{values: []string{TagPkValue}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GeneratedColumnFilter(tt.args.values); got != tt.want {
				t.Errorf("GeneratedColumnFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPkColumnFilter(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Happy Path",
			args: args{values: []string{"hello", TagPkValue}},
			want: true,
		},
		{
			name: "Unhappy Path",
			args: args{values: []string{TagPkValue, TagPkValue}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PkColumnFilter(tt.args.values); got != tt.want {
				t.Errorf("PkColumnFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUqColumnFilter(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Happy Path",
			args: args{values: []string{"hello", TagUqValue}},
			want: true,
		},
		{
			name: "Unhappy Path",
			args: args{values: []string{TagPkValue, TagUqValue}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UqColumnFilter(tt.args.values); got != tt.want {
				t.Errorf("UqColumnFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColumnFilter(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Happy Path",
			args: args{values: []string{"hello"}},
			want: true,
		},
		{
			name: "Unhappy Path",
			args: args{values: []string{TagPkValue}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ColumnFilter(tt.args.values); got != tt.want {
				t.Errorf("ColumnFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoneColumnFilter(t *testing.T) {
	type args struct {
		in0 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Nil Path",
			args: args{in0: nil},
			want: false,
		},
		{
			name: "Empty Array Path",
			args: args{in0: []string{}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoneColumnFilter(tt.args.in0); got != tt.want {
				t.Errorf("NoneColumnFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_retrieveFields(t *testing.T) {
	type modelOK struct {
		Id   string
		Name string
	}
	reflectedValueOK := reflect.ValueOf(modelOK{})
	reflectedType := reflectedValueOK.Type()
	idField, _ := reflectedType.FieldByName("Id")
	nameField, _ := reflectedType.FieldByName("Name")
	valuesOK := []reflect.StructField{idField, nameField}

	type modelNoOk struct {
		id string
	}
	reflectedValueNoOk := reflect.ValueOf(modelNoOk{})
	valuesNoOK := make([]reflect.StructField, 0)

	type args struct {
		reflectedValue reflect.Value
	}

	//

	tests := []struct {
		name string
		args args
		want []reflect.StructField
	}{
		{
			name: "HappyPath",
			args: args{reflectedValue: reflectedValueOK},
			want: valuesOK,
		},

		{
			name: "UnhappyPath",
			args: args{reflectedValue: reflectedValueNoOk},
			want: valuesNoOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RetrieveFields(tt.args.reflectedValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("retrieveFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_retrieveReflectedStruct(t *testing.T) {
	type model struct {
	}
	reflectedValue := reflect.ValueOf(model{})
	type args struct {
		value any
	}

	//

	tests := []struct {
		name      string
		args      args
		want      *reflect.Value
		wantErr   bool
		errWanted error
	}{
		{
			name:    "HappyPath As Struct",
			args:    args{value: model{}},
			want:    &reflectedValue,
			wantErr: false,
		},
		{
			name:    "HappyPath As Pointer",
			args:    args{value: &model{}},
			want:    &reflectedValue,
			wantErr: false,
		},
		{
			name:      "Error As nil",
			args:      args{value: nil},
			wantErr:   true,
			errWanted: ErrAnyIsNil,
		},
		{
			name:      "Error As Not a Pointer or Struct",
			args:      args{value: "hello"},
			wantErr:   true,
			errWanted: ErrAnyNotPointerOrStruct,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RetrieveReflectedStruct(tt.args.value)
			if tt.wantErr {
				if err != nil {
					if !reflect.DeepEqual(err, tt.errWanted) {
						t.Errorf("retrieveReflectedStruct() error = %v, wantErr %v", err, tt.wantErr)
						return
					}
				}
			} else {
				if !reflect.DeepEqual(got.Kind(), tt.want.Kind()) {
					t.Errorf("retrieveReflectedStruct() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
