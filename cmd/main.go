package main

import (
	"github.com/znobrega/compiler/internal/analyzer"
	"github.com/znobrega/compiler/internal/compiler"
	"github.com/znobrega/compiler/internal/infra"
	"log"
)

func main() {
	log.Println("Compiler")

	code, err := infra.ReadFile("code")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("code readed, number of lines:", len(code))

	lexicalAnalyzer := analyzer.NewLexical()
	syntacticAnalyzer := analyzer.NewSyntactic()

	compiler := compiler.New()
	compiler.Build(code)
	compiler.WithLexicalAnalyzer(lexicalAnalyzer)
	compiler.WithSyntacticAnalyzer(syntacticAnalyzer)
	err = compiler.Compile()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("compilation has finished")
}
