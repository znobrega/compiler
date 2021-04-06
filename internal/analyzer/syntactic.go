package analyzer

import (
	"errors"
	"github.com/znobrega/compiler/internal/entities"
	"github.com/znobrega/compiler/internal/infra"
)

var (
	ErrSymbolsIsOver = errors.New("symbols is over  ")

	ErrIsNotDot              = errors.New("expecting a dot")
	ErrIsNotSemiColen        = errors.New("expecting a ; ")
	ErrIsNotIdentifier       = errors.New("expecting a identifier")
	ErrIsNotProgram          = errors.New("expecting a key word program")
	ErrVariableDeclaration   = errors.New("variable declaration")
	ErrSubProgramDeclaration = errors.New("sub program declaration")
	ErrCompostCommand        = errors.New("compost commanf")
)

const (
	program       = "program|PROGRAM"
	identificador = "identificador|IDENTIFICADOR"
	semiColen     = ";"
	dot           = "."
)

type Syntactic struct {
	table         []entities.Symbol
	index         int64
	currentSymbal entities.Symbol
}

func NewSyntactic() Syntactic {
	return Syntactic{
		table:         nil,
		index:         -1,
		currentSymbal: entities.Symbol{},
	}
}

func (s *Syntactic) Analyze(table []entities.Symbol) error {
	s.table = table
	s.getNextSymbol()
	return s.Program()
}

func (s *Syntactic) Program() error {
	if infra.MatchString(s.currentSymbal.Token, program) {
		s.getNextSymbol()
		if infra.MatchString(s.currentSymbal.Classification, identificador) {
			s.getNextSymbol()
			if s.currentSymbal.Token == semiColen {
				s.getNextSymbol()

				err := s.VariableDeclaration()
				if err != nil {
					return ErrVariableDeclaration
				}
				err = s.SubProgramDeclaration()
				if err != nil {
					return ErrSubProgramDeclaration
				}
				err = s.CompostCommand()
				if err != nil {
					return ErrCompostCommand
				}

				if s.currentSymbal.Token != dot {
					return ErrIsNotDot
				} else {
					// SUCESS
					return nil
				}
			} else {
				return ErrIsNotSemiColen
			}
		} else {
			return ErrIsNotIdentifier
		}
	} else {
		return ErrIsNotProgram
	}
}

func (s *Syntactic) VariableDeclaration() error {
	return nil
}

func (s *Syntactic) SubProgramDeclaration() error {
	return nil
}

func (s *Syntactic) CompostCommand() error {
	return nil
}

func (s *Syntactic) getNextSymbol() {
	s.index = s.index + 1
	//TODO DEAL BETTER
	if s.index > int64(len(s.table)) {
		s.currentSymbal = entities.Symbol{}
	}

	s.currentSymbal = s.table[s.index]
}