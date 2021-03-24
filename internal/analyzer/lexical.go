package analyzer

import (
	"fmt"
	"github.com/znobrega/compiler/internal/entities"
	"log"
	"regexp"
)

const (
	//  comparacao comum a qualquer uma do conjunto de Strings
	IS_KEY_WORD = "^(program|var|integer|real|boolean|procedure|begin|end|if|then|else|while|do|not|case|true|false)(;|\\.)?$"
	//  \\w+ ===> qual letra ou digito seguinte. O primeiro caracter ja foi confirmado como letra
	IS_IDENTIFIER = "\\_*\\w+[\\_\\w+]*"
	//   | . | : | , | ( | ) ===> compara se e igual a alguma das Strings
	IS_DELIMITER = "(;|\\.|:|,|\\(|\\))"
	//  [^=] ===> deixar claro que o '=' de ':=' nao eh operador relacional
	IS_RELACIONAL_OPERATOR = "=|<|>|<=|>=|<>"
	//  +|- ===> String igual a + ou a -
	IS_ADDITION_OPERATORS = "[+|-]"
	//  or ===> comparacao comum de Strings
	IS_OPERATOR_OR = "or"
	//  *|/ ===> String igual a * ou a /
	IS_MULTIPLIER_OPERATOR = "\\*|/"
	//  and ===> and, sem letra ou digito antes e depois
	IS_OPERATOR_AND = "\\w{0}and\\w{0}"
	//  := ===> comparacao comum de Strings
	IS_ASSIGNMENT_OPERATOR = ":="
	//  \\.? ===> encontrar ponto 0 ou 1 vez
	IS_DIGIT = "^(\\d+)$"
	//  \\.{1} ===> encontrar ponto exatamento 1 vez entre inteiros (ou no fim)
	IS_FLOAT  = "^(\\d+\\.{1}\\d*)$"
	IS_NUMBER = "^(\\d+(\\.)?\\d*)$"
	//  [\\w\\W]* ===> palavra e digito, ou simbolo em qualquer ordem
	COMMENT = "\\{{1}[\\w\\W]*\\}{1}"

	COMENTARIO_AULA = "^//[\\w\\W]*"

	IS_LETTER_OR_UNDERSCORE = "^[a-zA-Z_]+$"
	IS_WORD_OR_DIGIT        = "^(\\w|\\d)+$"
)

type Lexical struct {
	table []entities.Symbol
}

func NewLexical() Lexical { return Lexical{table: make([]entities.Symbol, 0)} }

func (l Lexical) Analyze(code []string) error {
	log.Println("Lexical analyzes has started")

	for lineNumber, line := range code {
		log.Printf("line %d: %s", lineNumber+1, line)
		for i := 0; i < len(line); i++ {
			letter := string(line[i])

			if letter == " " {
				continue
			}

			if ok := l.MatchString(IS_DELIMITER, letter); ok {
				if i+2 <= len(line) && l.MatchString(IS_ASSIGNMENT_OPERATOR, line[i:i+2]) {
					l.addSymbolToTable(line[i:i+2], "ATRIBUICAO", lineNumber)
					i++
				} else {
					l.addSymbolToTable(letter, "DELIMITADOR", lineNumber)
				}
			} else if ok := l.MatchString(IS_RELACIONAL_OPERATOR, letter); ok {
				if i+2 <= len(line) && l.MatchString(IS_ASSIGNMENT_OPERATOR, line[i:i+2]) {
					l.addSymbolToTable(line[i:i+2], "OPERADORES RELACIONAIS", lineNumber)
					i++
				} else {
					l.addSymbolToTable(letter, "OPERADORES RELACIONAIS", lineNumber)
				}
			} else if ok := l.MatchString(IS_ADDITION_OPERATORS, letter); ok {
				l.addSymbolToTable(letter, "OPERADORES ADITIVOS", lineNumber)
			} else if ok := l.MatchString(IS_MULTIPLIER_OPERATOR, letter); ok {
				l.addSymbolToTable(letter, "OPERADORES MULTIPLICATIVOS", lineNumber)
			} else if ok := l.MatchString(IS_LETTER_OR_UNDERSCORE, letter); ok {
				word := l.buildWord(&i, line)
				if ok, _ := regexp.MatchString(IS_KEY_WORD, word); ok {
					l.addSymbolToTable(word, "PALAVRA CHAVE", lineNumber)
				} else if ok, _ := regexp.MatchString(IS_IDENTIFIER, word); ok {
					l.addSymbolToTable(word, "IDENTIFICADOR", lineNumber)
				} else if ok, _ := regexp.MatchString(IS_OPERATOR_AND, word); ok {
					l.addSymbolToTable(word, "OPERATOR AND", lineNumber)
				} else if ok, _ := regexp.MatchString(IS_OPERATOR_OR, word); ok {
					l.addSymbolToTable(word, "OPERATOR OR", lineNumber)
				}
			} else if ok := l.MatchString(IS_DIGIT, letter); ok {
				word := l.buildNumber(&i, line)
				if ok, _ := regexp.MatchString(IS_DIGIT, word); ok {
					l.addSymbolToTable(word, "INTEGER", lineNumber)
				} else if ok, _ := regexp.MatchString(IS_FLOAT, word); ok {
					l.addSymbolToTable(word, "FLOAT", lineNumber)
				}
			} else {
				log.Println("invalid symbol:", letter)
				//return ErrInvalidSymbol
			}
		}
	}

	fmt.Println(l.table)
	return nil
}

func (l *Lexical) buildWord(i *int, line string) string {
	word := ""
	initWord := *i
	endWord := *i + 1
	for ; endWord <= len(line); endWord++ {
		//fmt.Println(i, endWord, line[initWord:endWord])
		if ok := l.MatchString(IS_WORD_OR_DIGIT, line[initWord:endWord]); ok {
			word = line[initWord:endWord]
		} else {
			break
		}
	}
	*i = endWord - 2
	return word
}

func (l *Lexical) buildNumber(i *int, line string) string {
	word := ""
	initWord := *i
	endWord := *i + 1
	for ; endWord <= len(line); endWord++ {
		fmt.Println(i, endWord, line[initWord:endWord])
		if ok := l.MatchString(IS_NUMBER, line[initWord:endWord]); ok {
			word = line[initWord:endWord]
		} else {
			break
		}
	}
	*i = endWord - 2
	return word
}

func (l *Lexical) MatchString(expression, letter string) bool {
	ok, err := regexp.MatchString(expression, letter)
	if err != nil {
		// TODO REFACTOR ERROR TREATMENT
		panic(err)
	}
	return ok
}

func (l *Lexical) addSymbolToTable(word string, classification string, i int) {
	l.table = append(l.table, entities.Symbol{
		Token:          word,
		Classification: classification,
		Line:           i + 1,
	})
}
