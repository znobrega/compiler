package analyzer

import (
	"fmt"
	"github.com/znobrega/compiler/internal/entities"
	"github.com/znobrega/compiler/internal/infra"
	"log"
	"regexp"
)

const (
	IS_KEY_WORD   = "^(program|var|integer|real|boolean|procedure|begin|end|if|then|else|while|do|not|case|true|false)(;|\\.)?$"
	IS_IDENTIFIER = "\\_*\\w+[\\_\\w+]*"
	IS_DELIMITER  = "(;|\\.|:|,|\\(|\\))"
	//  [^=] ===> deixar claro que o '=' de ':=' nao eh operador relacional
	IS_RELACIONAL_OPERATOR = "=|<|>|<=|>=|<>"
	IS_ADDITION_OPERATORS  = "[+|-]"
	IS_OPERATOR_OR         = "or"
	IS_MULTIPLIER_OPERATOR = "\\*|/"
	IS_OPERATOR_AND        = "\\w{0}and\\w{0}"
	IS_ASSIGNMENT_OPERATOR = ":="
	//  \\.? ===> encontrar ponto 0 ou 1 vez
	IS_DIGIT = "^(\\d+)$"
	//  \\.{1} ===> encontrar ponto exatamento 1 vez entre inteiros (ou no fim)
	IS_FLOAT  = "^(\\d+\\.{1}\\d*)$"
	IS_NUMBER = "^(\\d+(\\.)?\\d*)$"
	//  [\\w\\W]* ===> palavra e digito, ou simbolo em qualquer ordem
	COMMENT = "\\{{1}[\\w\\W]*\\}{1}"

	COMENTARIO_AULA = "^//[\\w\\W]*"

	IS_OPEN_COMMENT   = "^{$"
	IS_CLOSED_COMMENT = "^}$"

	IS_LETTER_OR_UNDERSCORE = "^[a-zA-Z_]+$"
	IS_WORD_OR_DIGIT        = "^(\\w|\\d)+$"
)

type Lexical struct {
	table []entities.Symbol
}

func NewLexical() Lexical { return Lexical{table: make([]entities.Symbol, 0)} }

func (l Lexical) Analyze(code []string) ([]entities.Symbol, error) {
	log.Println("Lexical analyzes has started")

	for lineNumber, line := range code {
		log.Printf("line %d: %s", lineNumber+1, line)
		for letterIndex := 0; letterIndex < len(line); letterIndex++ {
			letter := string(line[letterIndex])

			if letter == " " {
				continue
			}

			switch {
			case l.isComment(letter, line, &letterIndex, lineNumber):
				continue
			case l.isDelimiter(letter, line, &letterIndex, lineNumber):
				continue
			case l.isRelacionalOrAssignmentOperator(letter, line, &letterIndex, lineNumber):
				continue
			case l.isKeyWordOrIdentifierOrAndOr(letter, line, &letterIndex, lineNumber):
				continue
			case l.isAdditionOperator(letter, line, &letterIndex, lineNumber):
				continue
			case l.isMultiplierOperator(letter, line, &letterIndex, lineNumber):
				continue
			case l.isNumber(letter, line, &letterIndex, lineNumber):
				continue
			}

			return nil, ErrInvalidSymbol
		}
	}

	fmt.Println(l.table)
	return l.table, nil
}

func (l *Lexical) isComment(letter string, line string, letterIndex *int, lineNumber int) bool {
	ok := infra.MatchString(IS_OPEN_COMMENT, letter)
	if !ok {
		return false
	}
	comment := l.buildMultilineComment(letterIndex, line, IS_CLOSED_COMMENT)
	l.addSymbolToTable(comment, "COMMENT", lineNumber)
	return true
}

func (l *Lexical) isAdditionOperator(letter string, line string, letterIndex *int, lineNumber int) bool {
	ok := infra.MatchString(IS_ADDITION_OPERATORS, letter)
	if !ok {
		return false
	}
	l.addSymbolToTable(letter, "OPERADORES ADITIVOS", lineNumber)
	return true
}

func (l *Lexical) isMultiplierOperator(letter string, line string, letterIndex *int, lineNumber int) bool {
	ok := infra.MatchString(IS_MULTIPLIER_OPERATOR, letter)
	if !ok {
		return false
	}
	l.addSymbolToTable(letter, "OPERADORES MULTIPLICATIVOS", lineNumber)
	return true
}

func (l *Lexical) isDelimiter(letter string, line string, letterIndex *int, lineNumber int) bool {
	ok := infra.MatchString(IS_DELIMITER, letter)
	if !ok {
		return false
	}
	if *letterIndex+2 <= len(line) && infra.MatchString(IS_ASSIGNMENT_OPERATOR, line[*letterIndex:*letterIndex+2]) {
		l.addSymbolToTable(line[*letterIndex:*letterIndex+2], "ATRIBUICAO", lineNumber)
		*letterIndex++
	} else {
		l.addSymbolToTable(letter, "DELIMITADOR", lineNumber)
	}
	return true
}

func (l *Lexical) isRelacionalOrAssignmentOperator(letter string, line string, letterIndex *int, lineNumber int) bool {
	ok := infra.MatchString(IS_RELACIONAL_OPERATOR, letter)
	if !ok {
		return false

	}

	if *letterIndex+2 <= len(line) && infra.MatchString(IS_ASSIGNMENT_OPERATOR, line[*letterIndex:*letterIndex+2]) {
		l.addSymbolToTable(line[*letterIndex:*letterIndex+2], "OPERADORES RELACIONAIS", lineNumber)
		*letterIndex++
	} else {
		l.addSymbolToTable(letter, "OPERADORES RELACIONAIS", lineNumber)
	}
	return true
}

func (l *Lexical) isNumber(letter string, line string, letterIndex *int, lineNumber int) bool {
	ok := infra.MatchString(IS_LETTER_OR_UNDERSCORE, letter)
	if !ok {
		return false
	}

	word := l.buildWord(letterIndex, line, IS_WORD_OR_DIGIT)
	if ok, _ := regexp.MatchString(IS_KEY_WORD, word); ok {
		l.addSymbolToTable(word, "PALAVRA CHAVE", lineNumber)
		return true
	} else if ok, _ := regexp.MatchString(IS_IDENTIFIER, word); ok {
		l.addSymbolToTable(word, "IDENTIFICADOR", lineNumber)
		return true
	} else if ok, _ := regexp.MatchString(IS_OPERATOR_AND, word); ok {
		l.addSymbolToTable(word, "OPERATOR AND", lineNumber)
		return true
	} else if ok, _ := regexp.MatchString(IS_OPERATOR_OR, word); ok {
		l.addSymbolToTable(word, "OPERATOR OR", lineNumber)
		return true
	}
	return false
}

func (l *Lexical) isKeyWordOrIdentifierOrAndOr(letter string, line string, letterIndex *int, lineNumber int) bool {
	ok := infra.MatchString(IS_DIGIT, letter)
	if !ok {
		return false
	}

	word := l.buildWord(letterIndex, line, IS_NUMBER)
	if ok, _ := regexp.MatchString(IS_DIGIT, word); ok {
		l.addSymbolToTable(word, "INTEGER", lineNumber)
		return true
	} else if ok, _ := regexp.MatchString(IS_FLOAT, word); ok {
		l.addSymbolToTable(word, "FLOAT", lineNumber)
		return true
	}
	return false
}

func (l *Lexical) buildWord(i *int, line string, pattern string) string {
	word := ""
	initWord := *i
	endWord := *i + 1

	for ; endWord <= len(line); endWord++ {
		//fmt.Println(i, endWord, line[initWord:endWord])
		if ok := infra.MatchString(pattern, line[initWord:endWord]); ok {
			word = line[initWord:endWord]
		} else {
			break
		}
	}
	*i = endWord - 2
	return word
}
func (l *Lexical) buildMultilineComment(i *int, line string, pattern string) string {
	word := ""
	initWord := *i
	endWord := *i + 1

	for ; endWord <= len(line); endWord++ {
		//fmt.Println(i, endWord, line[initWord:endWord])
		if ok := infra.MatchString(pattern, line[endWord-2:endWord]); ok {
			break
		} else {
			word = line[initWord:endWord]
		}
	}
	*i = endWord - 2
	return word
}

func (l *Lexical) addSymbolToTable(word string, classification string, i int) {
	l.table = append(l.table, entities.Symbol{
		Token:          word,
		Classification: classification,
		Line:           i + 1,
	})
}
