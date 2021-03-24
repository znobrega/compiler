package compiler

import "github.com/znobrega/compiler/internal/analyzer"

type Compiler struct {
	Code            []string
	LexicalAnalyzer analyzer.Lexical
}

func New() Compiler {
	return Compiler{}
}

func (c *Compiler) Build(code []string, lexical analyzer.Lexical) {
	c.Code = code
	c.LexicalAnalyzer = lexical
}

func (c *Compiler) Compile() error {
	err := c.LexicalAnalyzer.Analyze(c.Code)
	if err != nil {
		return err
	}

	return nil
}
