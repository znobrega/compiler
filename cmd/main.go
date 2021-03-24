package main

import (
	"fmt"
	"github.com/znobrega/compiler/internal/analyzer"
	"github.com/znobrega/compiler/internal/compiler"
	"github.com/znobrega/compiler/internal/infra"
	"log"
)

func main() {
	fmt.Println("Compiler")

	code, err := infra.ReadFile("code")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("code readed, number of lines:", len(code))

	lexicalAnalyzer := analyzer.NewLexical()

	compiler := compiler.New()
	compiler.Build(code, lexicalAnalyzer)
	err = compiler.Compile()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("compilation has finish")
}
