package feather_sql_parsing

import (
	"testing"

	feather_sql_reflection "github.com/guidomantilla/go-feather-sql/pkg/feather-sql-reflection"
)

func TestParseColumnAsNameValueSequence(t *testing.T) {
	type args struct {
		value     any
		initChar  string
		endChar   string
		separator string
		cont      int
		fn01      feather_sql_reflection.ColumnFilterFunc
		fn02      EvalColumnFunc
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ParseColumnAsNameValueSequence(tt.args.value, tt.args.initChar, tt.args.endChar, tt.args.separator, tt.args.cont, tt.args.fn01, tt.args.fn02)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseColumnAsNameValueSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseColumnAsNameValueSequence() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ParseColumnAsNameValueSequence() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEvalNameOnly(t *testing.T) {
	type args struct {
		name      string
		in1       int
		separator string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvalNameOnly(tt.args.name, tt.args.in1, tt.args.separator); got != tt.want {
				t.Errorf("EvalNameOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvalNameValueNamed(t *testing.T) {
	type args struct {
		name      string
		in1       int
		separator string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvalNameValueNamed(tt.args.name, tt.args.in1, tt.args.separator); got != tt.want {
				t.Errorf("EvalNameValueNamed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvalValueOnlyNamed(t *testing.T) {
	type args struct {
		name      string
		in1       int
		separator string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvalValueOnlyNamed(tt.args.name, tt.args.in1, tt.args.separator); got != tt.want {
				t.Errorf("EvalValueOnlyNamed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvalNameValueNumbered(t *testing.T) {
	type args struct {
		name      string
		cont      int
		separator string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvalNameValueNumbered(tt.args.name, tt.args.cont, tt.args.separator); got != tt.want {
				t.Errorf("EvalNameValueNumbered() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvalValueOnlyNumbered(t *testing.T) {
	type args struct {
		in0       string
		cont      int
		separator string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvalValueOnlyNumbered(tt.args.in0, tt.args.cont, tt.args.separator); got != tt.want {
				t.Errorf("EvalValueOnlyNumbered() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvalNameValueQuestioned(t *testing.T) {
	type args struct {
		name      string
		in1       int
		separator string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvalNameValueQuestioned(tt.args.name, tt.args.in1, tt.args.separator); got != tt.want {
				t.Errorf("EvalNameValueQuestioned() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvalValueOnlyQuestioned(t *testing.T) {
	type args struct {
		in0       string
		in1       int
		separator string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvalValueOnlyQuestioned(tt.args.in0, tt.args.in1, tt.args.separator); got != tt.want {
				t.Errorf("EvalValueOnlyQuestioned() = %v, want %v", got, tt.want)
			}
		})
	}
}
