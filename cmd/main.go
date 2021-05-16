package main

import (
	"github.com/znobrega/compiler/internal/analyzer"
	"github.com/znobrega/compiler/internal/compiler"
	"github.com/znobrega/compiler/internal/infra"
	"log"
	"time"
)

func main() {
	log.Println("Compiler")

	code, err := infra.ReadFile("code3")
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

	initCompiler := time.Now()
	err = compiler.Compile()
	log.Printf("execution time: %v", time.Since(initCompiler))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("compilation has finished succesfully")
}
