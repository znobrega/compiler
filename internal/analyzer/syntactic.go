package analyzer

import (
	"errors"
	"fmt"
	"github.com/znobrega/compiler/internal/entities"
	"github.com/znobrega/compiler/internal/infra"
	"log"
)

var (
	ErrSymbolsIsOver          = errors.New("symbols is over  ")
	ErrIsNotDot               = errors.New("expecting a dot")
	ErrIsNotSemiColen         = errors.New("expecting a ; ")
	ErrIsNotIdentifier        = errors.New("expecting a identifier")
	ErrIsNotProgram           = errors.New("expecting a key word program")
	ErrVariableDeclaration    = errors.New("variable declaration")
	ErrSubProgramDeclaration  = errors.New("sub program declaration")
	ErrCompostCommand         = errors.New("compost command")
	ErrVariableAlrealdyExists = errors.New("variable alrealdy exists")
)

const (
	PROGRAM             = "program"
	SEMICOLEN           = ";"
	COLEN               = ":"
	COMMAN              = ","
	DOT                 = "\\."
	DOT_VANILLA         = "."
	PROCEDURE           = "procedure"
	BEGIN               = "begin"
	END                 = "end"
	VAR                 = "var"
	DO                  = "do"
	WHILE               = "while"
	IF                  = "if"
	THEN                = "then"
	ELSE                = "else"
	TRUE                = "true"
	FALSE               = "false"
	NOT                 = "not"
	TYPES               = "integer|real|boolean|char"
	OPEN_PARENTHESIS    = "\\("
	CLOSE_PARENTHESIS   = "\\)"
	IDENTIFIER          = "IDENTIFICADOR"
	MULTIPLIER_OPERATOR = "OPERADORES MULTIPLICATIVOS"
	ADDITION_OPERATOR   = "OPERADORES ADITIVOS"
	RELATIONAL_OPERATOR = "OPERADORES RELACIONAIS"
	ASSIGNMENT          = "ATRIBUICAO"
	SINE                = "[+|-]"
	BOOLEAN             = "boolean|integer_relacional|real_relacional"
)

type Syntactic struct {
	table             []entities.Symbol
	variables         []entities.Variable
	procedures        []entities.Procedure
	index             int64
	currentSymbol     entities.Symbol
	beforeSymbol      entities.Symbol
	currentIdentifier entities.Identifier
	semantic          Semantic
}

func NewSyntactic() Syntactic {
	return Syntactic{
		table:             nil,
		index:             -1,
		currentSymbol:     entities.Symbol{},
		beforeSymbol:      entities.Symbol{},
		variables:         make([]entities.Variable, 0),
		procedures:        make([]entities.Procedure, 0),
		currentIdentifier: entities.Identifier{},
		semantic:          NewSemantic(),
	}
}

func (s *Syntactic) Analyze(table []entities.Symbol) error {
	s.table = table
	s.getNextSymbol()
	return s.program()
}

func (s *Syntactic) program() error {
	if infra.MatchString(s.currentSymbol.Token, PROGRAM) {

		s.getNextSymbol()
		if infra.MatchString(s.currentSymbol.Classification, IDENTIFIER) {
			s.currentIdentifier.Name = s.currentSymbol.Token
			s.currentIdentifier.Type = "program"
			s.semantic.checkProcedure(s.currentIdentifier)
			s.getNextSymbol()
			if s.currentSymbol.Token == SEMICOLEN {
				s.getNextSymbol()

				err := s.VariableDeclaration()
				if err != nil {
					log.Printf("%s: ", ErrVariableDeclaration.Error())
					return err
				}
				err = s.SubProgramDeclaration()
				if err != nil {
					log.Printf("%s: ", ErrSubProgramDeclaration.Error())
					return err
				}
				err = s.CompostCommand()
				if err != nil {
					log.Printf("%s: ", ErrCompostCommand.Error())
					return err
				}

				if s.currentSymbol.Token != DOT_VANILLA {
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
	if infra.MatchString(VAR, s.currentSymbol.Token) {
		s.getNextSymbol()
		err := s.getVariables()
		if err != nil {
			return err
		}

		// TODO CHECK NECESSITY
		//if s.semantic.isEmptyStack() {
		//	if infra.MatchString(BEGIN, s.currentSymbol.Token) {
		//	 // log.Println("begin found")
		//		return nil
		//} else if infra.MatchString(PROCEDURE, s.currentSymbol.Token) {
		// // log.Println("procedure found")
		// 	return nil
		//} else {
		//	return fmt.Errorf("after VAR is a variable is expected")
		//}
		//}
	} else {
		log.Println("there is not a variable declaration")
	}
	return nil
}

func (s *Syntactic) getVariables() error {
	err := s.getVariablesNames()
	if infra.MatchString(BEGIN, s.currentSymbol.Token) {
		// log.Println("begin found")
		return nil
	} else if infra.MatchString(PROCEDURE, s.currentSymbol.Token) {
		// log.Println("procedure found")
		return nil
	}
	if err != nil {
		// TODO STRANGE
		//if len(s.variables) == 0 {
		//	return err
		//} else

		if err != nil {
			return err
		}
	}

	s.getNextSymbol()

	err = s.getVariableType()
	if err != nil {
		return err
	}
	return nil
}

func (s *Syntactic) getVariableType() error {
	if infra.MatchString(TYPES, s.currentSymbol.Token) {

		s.addTypeToVariableList()
		s.semantic.defineVariableType(s.currentSymbol.Token)

		s.getNextSymbol()
		if infra.MatchString(SEMICOLEN, s.currentSymbol.Token) {
			s.getNextSymbol()
			return s.getVariables()
		} else {
			return fmt.Errorf("expecting a ; after type declaration")
		}
	} else {
		return fmt.Errorf("expectiong a variable type line %d", s.currentSymbol.Line)
	}

}

func (s *Syntactic) getVariablesNames() error {
	if infra.MatchString(IDENTIFIER, s.currentSymbol.Classification) {
		s.currentIdentifier = entities.Identifier{Name: s.currentSymbol.Token}
		if !s.semantic.checkVariable(s.currentIdentifier) {
			return fmt.Errorf("error varible already exists")
		}
		//err := s.addVariableToList(s.currentSymbol.Token)
		//if err != nil {
		//	return err
		//}
		s.getNextSymbol()
		if infra.MatchString(COMMAN, s.currentSymbol.Token) {
			s.getNextSymbol()
			return s.getVariables()
		} else if infra.MatchString(COLEN, s.currentSymbol.Token) {
			return nil
		} else {
			return fmt.Errorf("Expecting a : or a , on line %d", s.currentSymbol.Line)
		}

	} else {
		return fmt.Errorf("expecting a identifier, receive a %s line %d", s.currentSymbol.Token, s.currentSymbol.Line)
	}
	return nil
}

func (s *Syntactic) SubProgramDeclaration() error {
	err := s.isSubProgram()
	if err != nil {
		return err
	}

	return nil
}

func (s *Syntactic) isSubProgram() error {
	if !infra.MatchString(PROCEDURE, s.currentSymbol.Token) {
		return nil
	}

	err := s.getProcedureId()
	if err != nil {
		return err
	}

	s.getNextSymbol()

	err = s.VariableDeclaration()
	if err != nil {
		return err
	}

	err = s.SubProgramDeclaration()
	if err != nil {
		return err
	}

	err = s.CompostCommand()
	if err != nil {
		return err
	}

	if infra.MatchString(SEMICOLEN, s.currentSymbol.Token) {
		s.getNextSymbol()
	}

	return nil
}

func (s *Syntactic) getProcedureId() error {
	if infra.MatchString(PROCEDURE, s.currentSymbol.Token) {
		s.getNextSymbol()
		if infra.MatchString(IDENTIFIER, s.currentSymbol.Classification) {

			if !s.semantic.checkProcedure(entities.Identifier{Name: s.currentSymbol.Token, Type: s.beforeSymbol.Token}) {
				return fmt.Errorf("error procedure")
			}

			procedure := entities.Procedure{Name: s.currentSymbol.Token, Arguments: make([]entities.Variable, 0)}
			s.getNextSymbol()

			err := s.getProcedureArguments(&procedure)
			if err != nil {
				return err
			}

			err = s.checkIfProcedureAlreadyExists(&procedure)
			if err != nil {
				return err
			}

			s.procedures = append(s.procedures, procedure)
		} else {
			fmt.Errorf("expecting a identifier to initialize the procedure arguments, line: %d", s.currentSymbol.Line)
		}
	}
	// TODO CHECK THIS RETURN, THERE IS NO PROCEDURE
	return nil
}

func (s *Syntactic) checkIfProcedureAlreadyExists(newProcedure *entities.Procedure) error {
	for _, procedure := range s.procedures {
		if procedure.Name == newProcedure.Name {
			return fmt.Errorf("procedure %s already exists", procedure.Name)
		}
	}
	return nil
}

func (s *Syntactic) getProcedureArguments(procedure *entities.Procedure) error {
	if infra.MatchString(OPEN_PARENTHESIS, s.currentSymbol.Token) {
		return s.getArguments(procedure)
	} else if infra.MatchString(SEMICOLEN, s.currentSymbol.Token) {
		return nil
	} else {
		return fmt.Errorf("expecting a ; or a argument list")
	}
}

func (s *Syntactic) getArguments(procedure *entities.Procedure) error {
	s.getNextSymbol()
	if infra.MatchString(VAR, s.currentSymbol.Token) {
		s.getNextSymbol()
	}

	if infra.MatchString(IDENTIFIER, s.currentSymbol.Classification) {
		argumentName := s.currentSymbol.Token
		s.semantic.checkVariable(entities.Identifier{Name: argumentName})
		s.getNextSymbol()
		if infra.MatchString(COLEN, s.currentSymbol.Token) {
			s.getNextSymbol()
			if infra.MatchString(TYPES, s.currentSymbol.Token) {
				argumentType := s.currentSymbol.Token
				procedure.Arguments = append(procedure.Arguments, entities.Variable{
					Name: argumentName,
					Type: argumentType,
				})
				s.semantic.defineVariableType(argumentType)
				s.getNextSymbol()
				if infra.MatchString(CLOSE_PARENTHESIS, s.currentSymbol.Token) {
					s.getNextSymbol()
					if infra.MatchString(SEMICOLEN, s.currentSymbol.Token) {
						return nil
					} else {
						return fmt.Errorf("expecting a ; to complete the procedure")
					}
				} else if infra.MatchString(SEMICOLEN, s.currentSymbol.Token) {
					return s.getArguments(procedure)
				} else {
					fmt.Errorf("expecting a ) to complete argument list ou a ; to another arguments line: %d", s.currentSymbol.Line)
				}
			} else {
				return fmt.Errorf("expecting a variable type to argument line: %d", s.currentSymbol.Line)
			}
		} else {
			return fmt.Errorf("expectiog a colen to after argument name line: %d", s.currentSymbol.Line)
		}
	} else {
		return fmt.Errorf("expecting arument name to a procedure line: %d", s.currentSymbol.Line)
	}

	// TODO CHECK THIS ERROR
	return fmt.Errorf("ERROR UNIDENTIGIER inside getArguments")
}

func (s *Syntactic) CompostCommand() error {
	if infra.MatchString(DOT, s.currentSymbol.Token) {
		return formatError(SEMICOLEN, s.currentSymbol)
	} else if !infra.MatchString(BEGIN, s.currentSymbol.Token) {
		return formatError(BEGIN, s.currentSymbol)
	}
	s.semantic.initScope()
	s.getNextSymbol()

	err := s.optionalsCommands()
	if err != nil {
		return err
	}

	if infra.MatchString(END, s.currentSymbol.Token) {
		s.semantic.endScope()
		s.getNextSymbol()
		return nil
	}

	return formatError("end", s.currentSymbol)
}

func (s *Syntactic) optionalsCommands() error {
	err := s.checkCommands()
	if err != nil {
		return err
	}

	return nil
}

func (s *Syntactic) checkCommands() error {
	err := s.isCommand()
	if err != nil {
		return err
	}

	if infra.MatchString(SEMICOLEN, s.currentSymbol.Token) {
		s.semantic.clearCommand()
		s.getNextSymbol()
		if infra.MatchString(END, s.currentSymbol.Token) {
			return nil
		} else {
			return s.checkCommands()
		}
	} else {
		return formatError(SEMICOLEN, s.currentSymbol)
	}

	return fmt.Errorf("ERRO NO CHECK COMMANDS")
}

func (s *Syntactic) isCommand() error {
	s.currentIdentifier.Name = s.currentSymbol.Token
	s.currentIdentifier.Type = ""
	if infra.MatchString(IDENTIFIER, s.currentSymbol.Classification) {
		if s.isProcedureActivation() {
			return nil
		} else {
			if infra.MatchString(ASSIGNMENT, s.currentSymbol.Classification) {
				s.semantic.clearCommand()
				s.semantic.checkVariable(s.currentIdentifier)
				finalExpressionType := s.semantic.commandType
				s.semantic.clearCommand()
				s.getNextSymbol()
				if s.isExpression() {
					if !s.semantic.checkFinalType(finalExpressionType, s.semantic.commandType) {
						return fmt.Errorf("wrong expression")
					}
					return nil
				} else {
					return fmt.Errorf("expecting a expression")
				}
			} else {
				return formatError(ASSIGNMENT, s.currentSymbol)
			}
		}
		// TODO TEST cases with comandoCompos()
	} else if infra.MatchString(IF, s.currentSymbol.Token) {
		s.getNextSymbol()
		if s.isExpression() {
			if !infra.MatchString(BOOLEAN, s.semantic.commandType) {
				return fmt.Errorf("command if must to be a boolean")
			}

			if infra.MatchString(THEN, s.currentSymbol.Token) {
				s.semantic.clearCommand()
				s.getNextSymbol()
				if err := s.isCommand(); err == nil {
					//s.getNextSymbol()
					if infra.MatchString(ELSE, s.currentSymbol.Token) {
						s.semantic.clearCommand()
						s.getNextSymbol()
						return s.isCommand()
					} else {
						return nil
					}
				}
			}
		}
	} else if infra.MatchString(WHILE, s.currentSymbol.Token) {
		s.getNextSymbol()
		if s.isExpression() {
			if !infra.MatchString(BOOLEAN, s.semantic.commandType) {
				return fmt.Errorf("command while must to be a boolean")
			}

			if infra.MatchString(DO, s.currentSymbol.Token) {
				s.getNextSymbol()
				return s.isCommand()
			}
		}
	}
	return fmt.Errorf("is not command")
}

func (s *Syntactic) isProcedureActivation() bool {
	s.currentIdentifier = entities.Identifier{
		Name: s.currentSymbol.Token,
		Type: "procedure",
	}
	s.getNextSymbol()
	if infra.MatchString(OPEN_PARENTHESIS, s.currentSymbol.Token) {
		s.getNextSymbol()
		if !s.semantic.checkProcedure(s.currentIdentifier) {
			log.Println("procedure error checkProcedure")
			return false
		}

		if s.expressionList() {
			if infra.MatchString(CLOSE_PARENTHESIS, s.currentSymbol.Token) {
				s.getNextSymbol()
				return true
			} else {
				log.Print("Error procedimento linha 379")
				return false
			}
		}
	}
	return false
}

func (s *Syntactic) expressionList() bool {
	if s.isExpression() {
		s.getNextSymbol()
		if infra.MatchString(COMMAN, s.currentSymbol.Token) {
			return s.expressionList()
		}
		s.bootstrapSymbol()
		return true
	}

	log.Println("problema na lista de expressões.")
	return false
}
func (s *Syntactic) factor() bool {
	s.currentIdentifier = entities.Identifier{
		Name: s.currentSymbol.Token,
		Type: "",
	}
	if infra.MatchString(IDENTIFIER, s.currentSymbol.Classification) {
		s.semantic.checkVariable(s.currentIdentifier)
		s.getNextSymbol()
		if infra.MatchString(OPEN_PARENTHESIS, s.currentSymbol.Token) {
			s.getNextSymbol()
			if s.expressionList() {
				s.getNextSymbol()
				if infra.MatchString(CLOSE_PARENTHESIS, s.currentSymbol.Token) {
					return true
				}
			}
		}
		s.bootstrapSymbol()
		return true
	} else if infra.MatchString(IS_DIGIT, s.currentSymbol.Token) {
		s.semantic.typeAnalyses("integer")
		return true
	} else if infra.MatchString(IS_NUMBER, s.currentSymbol.Token) {
		s.semantic.typeAnalyses("real")
		return true
	} else if infra.MatchString(TRUE, s.currentSymbol.Token) {
		s.semantic.typeAnalyses("boolean")
		return true
	} else if infra.MatchString(FALSE, s.currentSymbol.Token) {
		s.semantic.typeAnalyses("boolean")
		return true
	} else if infra.MatchString(OPEN_PARENTHESIS, s.currentSymbol.Token) {
		s.getNextSymbol()
		if s.isExpression() {
			if infra.MatchString(CLOSE_PARENTHESIS, s.currentSymbol.Token) {
				return true
			} else {
				return false
			}
		} else {
			log.Print(formatError("factor expression", s.currentSymbol).Error())
			return false
		}
		return true
	} else if infra.MatchString(NOT, s.currentSymbol.Token) {
		s.getNextSymbol()
	}

	log.Print(formatError("a factor error happens", s.currentSymbol).Error())
	return false

}

func (s *Syntactic) isExpression() bool {
	if s.isSimpleExpression() {
		if infra.MatchString(RELATIONAL_OPERATOR, s.currentSymbol.Classification) {
			s.semantic.typeAnalyses(s.currentSymbol.Token)
			s.getNextSymbol()
			return s.isSimpleExpression()
		}
		return true
	}
	return false
}

func (s *Syntactic) isSimpleExpression() bool {
	if s.isTerm() {
		if infra.MatchString(ADDITION_OPERATOR, s.currentSymbol.Classification) {
			s.getNextSymbol()
			return s.isSimpleExpression()
		}
		return true
	} else if infra.MatchString(SINE, s.currentSymbol.Token) {
		// TODO o isTerm pode ser: um número que recebe um sinal negativo antes. exemplo: -4
		return s.isTerm()
	}
	println("ERRO: é esperado pelo menos um isTerm")
	return false
}

func (s *Syntactic) isTerm() bool {
	if s.factor() {
		s.getNextSymbol()
		if infra.MatchString(MULTIPLIER_OPERATOR, s.currentSymbol.Classification) {
			s.semantic.typeAnalyses(s.currentSymbol.Token)
			s.getNextSymbol()
			return s.isTerm()
		}
		return true
	}
	return false
}

func formatError(expected string, currentSymbol entities.Symbol) error {
	return fmt.Errorf("expecting a %s but receive a %s line: %d", expected, currentSymbol.Token, currentSymbol.Line)
}

func (s *Syntactic) getNextSymbol() {
	s.index = s.index + 1

	//TODO DEAL BETTER
	if s.index > int64(len(s.table)) {
		s.currentSymbol = entities.Symbol{}
	}
	s.beforeSymbol = s.currentSymbol
	s.currentSymbol = s.table[s.index]

	if s.currentSymbol.Classification == "COMMENT" {
		log.Printf("comment: %s", s.currentSymbol.Token)
		s.getNextSymbol()
	}
}

func (s *Syntactic) bootstrapSymbol() {
	s.currentSymbol = s.beforeSymbol
	s.index = s.index - 1
}

func (s *Syntactic) addVariableToList(name string) error {
	// TODO THIS MUST BE A SET
	for _, variable := range s.variables {
		if variable.Name == name {
			return ErrVariableAlrealdyExists
		}
	}

	s.variables = append(s.variables, entities.Variable{
		Name: name,
		Type: "",
	})

	return nil
}

func (s *Syntactic) addTypeToVariableList() {
	for i, variable := range s.variables {
		if variable.Type == "" {
			s.variables[i].Type = s.currentSymbol.Token
		}
	}
}
