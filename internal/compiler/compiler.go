package compiler

import (
	"github.com/znobrega/compiler/internal/analyzer"
	"log"
)

type Compiler struct {
	Code              []string
	LexicalAnalyzer   analyzer.Lexical
	SyntacticAnalyzer analyzer.Syntactic
}

func New() Compiler {
	return Compiler{}
}

func (c *Compiler) Build(code []string) {
	c.Code = code
}

func (c *Compiler) WithLexicalAnalyzer(lexical analyzer.Lexical) {
	c.LexicalAnalyzer = lexical
}

func (c *Compiler) WithSyntacticAnalyzer(syntactic analyzer.Syntactic) {
	c.SyntacticAnalyzer = syntactic
}

func (c *Compiler) Compile() error {
	table, err := c.LexicalAnalyzer.Analyze(c.Code)
	if err != nil {
		return err
	}
	log.Println("lexical analysis finished")

	err = c.SyntacticAnalyzer.Analyze(table)
	if err != nil {
		return err
	}
	log.Println("syntactic analysis finished")

	return nil
}
