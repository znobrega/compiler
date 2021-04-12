package analyzer

import (
	"github.com/znobrega/compiler/internal/entities"
	"reflect"
	"testing"
)

func TestNewSyntactic(t *testing.T) {
	tests := []struct {
		name string
		want Syntactic
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSyntactic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSyntactic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSyntactic_Analyze(t *testing.T) {
	type fields struct {
		table         []entities.Symbol
		index         int64
		currentSymbol entities.Symbol
	}
	type args struct {
		table []entities.Symbol
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "basic",
			fields: fields{
				table:         nil,
				index:         -1,
				currentSymbol: entities.Symbol{},
			},
			args: args{
				table: []entities.Symbol{
					{"program", "PALAVRA CHAVE", 1},
					{"teste", "IDENTIFICADOR", 1},
					{";", "DELIMITADOR", 1},
					{".", "DELIMITADOR", 2}},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Syntactic{
				table:         tt.fields.table,
				index:         tt.fields.index,
				currentSymbol: tt.fields.currentSymbol,
			}
			if err := s.Analyze(tt.args.table); (err != nil) != tt.wantErr {
				t.Errorf("Analyze() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSyntactic_CompostCommand(t *testing.T) {
	type fields struct {
		table         []entities.Symbol
		index         int64
		currentSymbol entities.Symbol
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Syntactic{
				table:         tt.fields.table,
				index:         tt.fields.index,
				currentSymbol: tt.fields.currentSymbol,
			}
			if err := s.CompostCommand(); (err != nil) != tt.wantErr {
				t.Errorf("CompostCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSyntactic_Program(t *testing.T) {
	type fields struct {
		table         []entities.Symbol
		index         int64
		currentSymbol entities.Symbol
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Syntactic{
				table:         tt.fields.table,
				index:         tt.fields.index,
				currentSymbol: tt.fields.currentSymbol,
			}
			if err := s.Program(); (err != nil) != tt.wantErr {
				t.Errorf("Program() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSyntactic_SubProgramDeclaration(t *testing.T) {
	type fields struct {
		table         []entities.Symbol
		index         int64
		currentSymbol entities.Symbol
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Syntactic{
				table:         tt.fields.table,
				index:         tt.fields.index,
				currentSymbol: tt.fields.currentSymbol,
			}
			if err := s.SubProgramDeclaration(); (err != nil) != tt.wantErr {
				t.Errorf("SubProgramDeclaration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSyntactic_VariableDeclaration(t *testing.T) {
	type fields struct {
		table         []entities.Symbol
		index         int64
		currentSymbol entities.Symbol
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Syntactic{
				table:         tt.fields.table,
				index:         tt.fields.index,
				currentSymbol: tt.fields.currentSymbol,
			}
			if err := s.VariableDeclaration(); (err != nil) != tt.wantErr {
				t.Errorf("VariableDeclaration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSyntactic_getNextSymbol(t *testing.T) {
	type fields struct {
		table         []entities.Symbol
		index         int64
		currentSymbol entities.Symbol
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Syntactic{
				table:         tt.fields.table,
				index:         tt.fields.index,
				currentSymbol: tt.fields.currentSymbol,
			}

			s.getNextSymbol()
		})
	}
}
