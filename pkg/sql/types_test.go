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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.String(); got != tt.want {
				t.Errorf("DriverName.String() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.EvalNameValue(); !reflect.DeepEqual(got, tt.want) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.EvalValueOnly(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParamHolder.EvalValueOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}
