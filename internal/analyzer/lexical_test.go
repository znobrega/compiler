package analyzer

import (
	"fmt"
	"github.com/znobrega/compiler/internal/entities"
	"github.com/znobrega/compiler/internal/infra"
	"log"
	"reflect"
	"testing"
)

func TestLexical_Analyze(t *testing.T) {
	code, err := infra.ReadFile("code")
	if err != nil {
		log.Fatal(err)
	}

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
		want    []entities.Symbol
		wantErr bool
	}{
		{
			name: "sucess",
			fields: fields{
				table: make([]entities.Symbol, 0),
			},
			args: args{
				code: code,
			},
			want: []entities.Symbol{
				{"program", "PALAVRA CHAVE", 1},
				{"teste", "IDENTIFICADOR", 1},
				{";", "DELIMITADOR", 1},
				{"{programa exemplo}", "COMMENT", 1},
				{"var", "PALAVRA CHAVE", 2},
				{"valor1", "IDENTIFICADOR", 3},
				{":", "DELIMITADOR", 3},
				{"integer", "PALAVRA CHAVE", 3},
				{";", "DELIMITADOR", 3},
				{"valor2", "IDENTIFICADOR", 4},
				{":", "DELIMITADOR", 4},
				{"real", "PALAVRA CHAVE", 4},
				{";", "DELIMITADOR", 4},
				{"begin", "PALAVRA CHAVE", 5},
				{"valor1", "IDENTIFICADOR", 6},
				{":=", "ATRIBUICAO", 6},
				{"10", "INTEGER", 6},
				{";", "DELIMITADOR", 6},
				{"end", "PALAVRA CHAVE", 7},
				{".", "DELIMITADOR", 7},
			},
			wantErr: false,
		},
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
			want: []entities.Symbol{
				{"10", "INTEGER", 1},
				{";", "DELIMITADOR", 1},
				{"10.44444", "FLOAT", 2},
				{"99999.99999", "FLOAT", 3},
				{".", "DELIMITADOR", 3},
				{"10234", "INTEGER", 4},
				{"end", "PALAVRA CHAVE", 5},
				{".", "DELIMITADOR", 5},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Lexical{
				table: tt.fields.table,
			}
			got, err := l.Analyze(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("Analyze() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Analyze() got = %v, want %v", got, tt.want)
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
			if got := infra.MatchString(tt.args.expression, tt.args.letter); got != tt.want {
				t.Errorf("MatchString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexical_addSymbolToTable(t *testing.T) {
	type fields struct {
		table []entities.Symbol
	}
	type args struct {
		word           string
		classification string
		i              int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexical{
				table: tt.fields.table,
			}

			fmt.Println(l)
		})
	}
}

func TestLexical_buildMultilineComment(t *testing.T) {
	type fields struct {
		table []entities.Symbol
	}
	type args struct {
		i       *int
		line    string
		pattern string
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
			if got := l.buildMultilineComment(tt.args.i, tt.args.line, tt.args.pattern); got != tt.want {
				t.Errorf("buildMultilineComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexical_buildWord(t *testing.T) {
	type fields struct {
		table []entities.Symbol
	}
	type args struct {
		i       *int
		line    string
		pattern string
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
			if got := l.buildWord(tt.args.i, tt.args.line, tt.args.pattern); got != tt.want {
				t.Errorf("buildWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexical_isAdditionOperator(t *testing.T) {
	type fields struct {
		table []entities.Symbol
	}
	type args struct {
		letter      string
		line        string
		letterIndex *int
		lineNumber  int
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
			if got := l.isAdditionOperator(tt.args.letter, tt.args.line, tt.args.letterIndex, tt.args.lineNumber); got != tt.want {
				t.Errorf("isAdditionOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexical_isComment(t *testing.T) {
	type fields struct {
		table []entities.Symbol
	}
	type args struct {
		letter      string
		line        string
		letterIndex *int
		lineNumber  int
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
			if got := l.isComment(tt.args.letter, tt.args.line, tt.args.letterIndex, tt.args.lineNumber); got != tt.want {
				t.Errorf("isComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexical_isDelimiter(t *testing.T) {
	type fields struct {
		table []entities.Symbol
	}
	type args struct {
		letter      string
		line        string
		letterIndex *int
		lineNumber  int
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
			if got := l.isDelimiter(tt.args.letter, tt.args.line, tt.args.letterIndex, tt.args.lineNumber); got != tt.want {
				t.Errorf("isDelimiter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexical_isKeyWordOrIdentifierOrAndOr(t *testing.T) {
	type fields struct {
		table []entities.Symbol
	}
	type args struct {
		letter      string
		line        string
		letterIndex *int
		lineNumber  int
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
			if got := l.isKeyWordOrIdentifierOrAndOr(tt.args.letter, tt.args.line, tt.args.letterIndex, tt.args.lineNumber); got != tt.want {
				t.Errorf("isKeyWordOrIdentifierOrAndOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexical_isMultiplierOperator(t *testing.T) {
	type fields struct {
		table []entities.Symbol
	}
	type args struct {
		letter      string
		line        string
		letterIndex *int
		lineNumber  int
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
			if got := l.isMultiplierOperator(tt.args.letter, tt.args.line, tt.args.letterIndex, tt.args.lineNumber); got != tt.want {
				t.Errorf("isMultiplierOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexical_isNumber(t *testing.T) {
	type fields struct {
		table []entities.Symbol
	}
	type args struct {
		letter      string
		line        string
		letterIndex *int
		lineNumber  int
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
			if got := l.isNumber(tt.args.letter, tt.args.line, tt.args.letterIndex, tt.args.lineNumber); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexical_isRelacionalOrAssignmentOperator(t *testing.T) {
	type fields struct {
		table []entities.Symbol
	}
	type args struct {
		letter      string
		line        string
		letterIndex *int
		lineNumber  int
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
			if got := l.isRelacionalOrAssignmentOperator(tt.args.letter, tt.args.line, tt.args.letterIndex, tt.args.lineNumber); got != tt.want {
				t.Errorf("isRelacionalOrAssignmentOperator() = %v, want %v", got, tt.want)
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
