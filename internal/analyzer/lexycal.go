package analyzer

import (
	"fmt"
	"github.com/znobrega/compiler/internal/entities"
	"log"
	"regexp"
)

const (
	//  comparacao comum a qualquer uma do conjunto de Strings
	PALAVRAS_CHAVES = "^(program|var|integer|real|boolean|procedure|begin|end|if|then|else|while|do|not|case|true|false)(;|\\.)?$"
	//  \\w+ ===> qual letra ou digito seguinte. O primeiro caracter ja foi confirmado como letra
	IDENTIDICADOR = "\\_*\\w+[\\_\\w+]*"
	//   | . | : | , | ( | ) ===> compara se e igual a alguma das Strings
	DELIMITADORES = "(;|\\.|:|,|\\(|\\))"
	//  [^=] ===> deixar claro que o '=' de ':=' nao eh operador relacional
	OPERADORES_RELACIONAIS = "=|<|>|<=|>=|<>"
	//  +|- ===> String igual a + ou a -
	OPERADORES_ADITIVOS = "[+|-]"
	//  or ===> comparacao comum de Strings
	OPERADOR_ADITIVO_OR = "or"
	//  *|/ ===> String igual a * ou a /
	OPERADORES_MULTIPLICATIVOS = "\\*|/"
	//  and ===> and, sem letra ou digito antes e depois
	OPERADOR_MULTIPLICATIVO_AND = "\\w{0}and\\w{0}"
	//  := ===> comparacao comum de Strings
	ATRIBUICAO = ":="
	//  \\.? ===> encontrar ponto 0 ou 1 vez
	NUMEROS_INTEIROS = "\\d+"
	//  \\.{1} ===> encontrar ponto exatamento 1 vez entre inteiros (ou no fim)
	NUMEROS_REAIS = "\\d+\\.{1}\\d*"
	//  [\\w\\W]* ===> palavra e digito, ou simbolo em qualquer ordem
	COMENTARIO = "\\{{1}[\\w\\W]*\\}{1}"

	COMENTARIO_AULA = "^//[\\w\\W]*"

	IS_WORD_OR_DIGIT = "^(\\w|\\d)+$"
)

type Lexycal struct {
	table []entities.Symbol
}

func NewLexycal() Lexycal { return Lexycal{table: make([]entities.Symbol, 0)} }

func (l *Lexycal) Analyze(code []string) error {
	log.Println("Lexycal analyzes has started")

	for lineNumber, line := range code {
		log.Printf("line %d: %s", lineNumber, line)
		for i := 0; i < len(line); i++ {
			letter := string(line[i])

			if letter == " " {
				continue
			}

			if ok := l.MatchString(DELIMITADORES, letter); ok {
				if i+2 <= len(line) && l.MatchString(ATRIBUICAO, line[i:i+2]) {
					l.addSymbolToTable(line[i:i+2], "ATRIBUICAO", lineNumber)
					i++
				} else {
					l.addSymbolToTable(letter, "DELIMITADOR", lineNumber)
				}
			} else if ok := l.MatchString(OPERADORES_RELACIONAIS, letter); ok {
				if i+2 <= len(line) && l.MatchString(ATRIBUICAO, line[i:i+2]) {
					l.addSymbolToTable(line[i:i+2], "OPERADORES RELACIONAIS", lineNumber)
					i++
				} else {
					l.addSymbolToTable(letter, "OPERADORES RELACIONAIS", lineNumber)
				}
			} else if ok := l.MatchString(OPERADORES_ADITIVOS, letter); ok {
				l.addSymbolToTable(letter, "OPERADORES ADITIVOS", lineNumber)
			} else if ok := l.MatchString(OPERADORES_MULTIPLICATIVOS, letter); ok {
				l.addSymbolToTable(letter, "OPERADORES MULTIPLICATIVOS", lineNumber)
			} else if ok := l.MatchString(IS_WORD_OR_DIGIT, letter); ok {
				word := l.buildWord(&i, line)
				if ok, _ := regexp.MatchString(PALAVRAS_CHAVES, word); ok {
					l.addSymbolToTable(word, "PALAVRA CHAVE", lineNumber)
				} else if ok, _ := regexp.MatchString(IDENTIDICADOR, word); ok {
					l.addSymbolToTable(word, "IDENTIDICADOR", lineNumber)
				} else if ok, _ := regexp.MatchString(ATRIBUICAO, word); ok {
					l.addSymbolToTable(word, ""+
						"", lineNumber)
				} else if ok, _ := regexp.MatchString(NUMEROS_INTEIROS, word); ok {
					l.addSymbolToTable(word, "NUMEROS_INTEIROS", lineNumber)
				}
			}
		}
	}

	fmt.Println(l.table)
	return nil
}

func (l *Lexycal) buildWord(i *int, line string) string {
	word := ""
	for endWord := *i + 1; endWord < len(line); endWord++ {
		initWord := *i
		fmt.Println(i, endWord, line[initWord:endWord])
		if ok := l.MatchString(IS_WORD_OR_DIGIT, line[initWord:endWord]); ok {
			word = line[initWord:endWord]
		} else {
			*i = endWord - 2
			break
		}
	}
	return word
}

func (l *Lexycal) MatchString(expression, letter string) bool {
	ok, err := regexp.MatchString(expression, letter)
	if err != nil {
		panic(err)
	}
	return ok
}

func (l *Lexycal) addSymbolToTable(word string, classification string, i int) {
	l.table = append(l.table, entities.Symbol{
		Token:          word,
		Classification: classification,
		Line:           i + 1,
	})
}
