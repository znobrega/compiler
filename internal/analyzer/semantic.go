package analyzer

import (
	"github.com/znobrega/compiler/internal/entities"
	"github.com/znobrega/compiler/internal/infra"
	"log"
)

var (
	scope = entities.Identifier{
		Name: "$",
		Type: "$",
	}
)

type Semantic struct {
	identifiersStack []entities.Identifier
	depth            int64
	commandType      string
}

func NewSemantic() Semantic {
	return Semantic{
		identifiersStack: make([]entities.Identifier, 0),
		depth:            0,
		commandType:      "",
	}
}

func (s *Semantic) checkVariable(identifier entities.Identifier) bool {
	if s.depth > 0 {
		return s.processIdentifier(identifier.Name)
	} else {
		return s.pushVariable(identifier)
	}
}

func (s *Semantic) checkProcedure(identifier entities.Identifier) bool {
	if s.depth > 0 {
		return s.processIdentifier(identifier.Name)
	} else {
		return s.pushProcedure(identifier)
	}
}

func (s *Semantic) processIdentifier(identifierName string) bool {
	stackTop := len(s.identifiersStack) - 1
	for i := stackTop; i >= 0; i-- {
		currentIdentifier := s.identifiersStack[i]
		if currentIdentifier.Name == identifierName {
			if currentIdentifier.Type == "procedure" {
				return true
			} else {
				return s.typeAnalyses(currentIdentifier.Type)
			}
		}
	}
	log.Fatal("identifier not found ", identifierName)
	return false
}

func (s *Semantic) typeAnalyses(identifierType string) bool {
	if s.commandType == "" {
		s.commandType = identifierType
		return true
	} else if s.commandType == "integer" {
		if identifierType == "integer" {
			return true
		} else if ok := infra.MatchString(IS_MULTIPLIER_OPERATOR, identifierType); ok {
			return true
		} else if ok := infra.MatchString(IS_RELACIONAL_OPERATOR, identifierType); ok {
			s.commandType += "_relacional"
			return true
		} else {
			log.Fatal("incompatible integer type")
			return false
		}
	} else if s.commandType == "real" {
		if identifierType == "integer" {
			return true
		} else if ok := infra.MatchString(IS_MULTIPLIER_OPERATOR, identifierType); ok {
			return true
		} else if identifierType == "real" {
			return true
		} else if ok := infra.MatchString(IS_RELACIONAL_OPERATOR, identifierType); ok {
			s.commandType += "_relacional"
			return true
		} else {
			log.Fatal("incompatible real type")
			return false
		}
	} else if s.commandType == "boolean" {
		if identifierType == "boolean" {
			return true
		} else if ok := infra.MatchString(IS_RELACIONAL_OPERATOR, identifierType); ok {
			return true
		} else {
			log.Fatal("incompatible boolean type")
			return false
		}
	} else if ok := infra.MatchString("integer_relacional|real_relacional", s.commandType); ok {
		if ok := infra.MatchString("integer|real", identifierType); ok {
			return true
		} else {
			log.Fatal("quando integer|real && operador elacional, o próximo deve ser integer|real")
			return false
		}
	} else {
		return false
	}
	return false
}

func (s *Semantic) findIdentifierOnScope(identifier entities.Identifier) bool {
	stackTop := len(s.identifiersStack) - 1
	for i := stackTop; i >= 0; i-- {
		currentIdentifier := s.identifiersStack[i]

		if currentIdentifier.Name == "$" {
			break
		}

		// TODO REVIEW && currentIdentifier.Type == identifier.Type
		if currentIdentifier.Name == identifier.Name {
			return true
		}
	}
	return false
}

func (s *Semantic) pushVariable(identifier entities.Identifier) bool {
	if s.findIdentifierOnScope(identifier) {
		//log.Fatal("variable already declareD")
		return false
	}

	s.pushIdentifier(identifier)
	return true
}

func (s *Semantic) pushProcedure(identifier entities.Identifier) bool {
	if identifier.Type == "procedure" {
		if s.findIdentifierOnScope(identifier) {
			log.Fatal("identifier already declare")
			return false
		}

		if identifier.Name == s.identifiersStack[0].Name {
			log.Fatal("identifier used as program name")
			return false
		}

		s.pushIdentifier(identifier)
		s.pushIdentifier(scope)

		return true
	} else if identifier.Type == "program" {
		s.pushIdentifier(identifier)
		s.pushIdentifier(scope)
		return true
	} else {
		log.Fatal("error with procedure")
		return false
	}
}

func (s *Semantic) clearProcedure() {
	// TODO CHECH
	stackTop := len(s.identifiersStack) - 1
	for i := stackTop; i >= 0; i-- {
		if s.identifiersStack[i].Name == "$" {
			s.identifiersStack = s.identifiersStack[:i+1]
			break
		}
	}
}

func (s *Semantic) clearCommand() {
	s.commandType = ""
}

func (s *Semantic) initScope() {
	s.depth++
}

func (s *Semantic) endScope() {
	// TODO CHECH
	s.depth--
	if s.depth == 0 {
		s.clearProcedure()
	}
}

func (s *Semantic) defineVariableType(variableType string) {
	stackTop := len(s.identifiersStack) - 1
	for i := stackTop; i >= 0; i-- {
		if s.identifiersStack[i].Type != "" {
			break
		}

		s.identifiersStack[i].Type = variableType
	}
}

func (s *Semantic) isEmptyStack() bool {
	isEmpty := len(s.identifiersStack) == 0
	return isEmpty
}

func (s *Semantic) pushIdentifier(identifier entities.Identifier) {
	s.identifiersStack = append(s.identifiersStack, identifier)
}

func (s *Semantic) checkFinalType(required string, delivered string) bool {
	if required == "integer" || required == "real" {
		if delivered == "integer" || delivered == "real" {
			return true
		}
	} else if required == "boolean" {
		if infra.MatchString(BOOLEAN, delivered) {
			return true
		}
	} else if required == "procedure" && delivered == "procedure" {
		return true
	}
	return false
}
