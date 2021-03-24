package analyzer

import (
	"github.com/znobrega/compiler/internal/entities"
	"reflect"
	"testing"
)

func TestLexical_Analyze(t *testing.T) {
	type fields struct {
		table []entities.Symbol
	}
	type args struct {
		code []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "numbers",
			fields: fields{
				table: make([]entities.Symbol, 0),
			},
			args: args{
				code: []string{"10;",
					"10.44444",
					"99999.99999.",
					"10234",
					"end."},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Lexical{
				table: tt.fields.table,
			}
			if err := l.Analyze(tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("Analyze() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLexical_MatchString(t *testing.T) {
	type fields struct {
		table []entities.Symbol
	}
	type args struct {
		expression string
		letter     string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexical{
				table: tt.fields.table,
			}
			if got := l.MatchString(tt.args.expression, tt.args.letter); got != tt.want {
				t.Errorf("MatchString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexical_buildNumber(t *testing.T) {
	type fields struct {
		table []entities.Symbol
	}
	type args struct {
		i    *int
		line string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexical{
				table: tt.fields.table,
			}
			if got := l.buildNumber(tt.args.i, tt.args.line); got != tt.want {
				t.Errorf("buildNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexical_buildWord(t *testing.T) {
	type fields struct {
		table []entities.Symbol
	}
	type args struct {
		i    *int
		line string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexical{
				table: tt.fields.table,
			}
			if got := l.buildWord(tt.args.i, tt.args.line); got != tt.want {
				t.Errorf("buildWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLexical(t *testing.T) {
	tests := []struct {
		name string
		want Lexical
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLexical(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLexical() = %v, want %v", got, tt.want)
			}
		})
	}
}
