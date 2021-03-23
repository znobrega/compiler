package main

import (
	"fmt"
	"github.com/znobrega/compiler/internal/analyzer"
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

	lexycalAnalyzer := analyzer.NewLexycal()
	lexycalAnalyzer.Analyze(code)
}
