package main

import (
	"encoding/json"
	"log"
	"os"
	"parser/parser"
	"parser/visitors"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input := "scripts/inputs/input.txt"

	file, err := os.ReadFile(input)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	for _, line := range strings.Split(string(file), "\n") {
		func() {
			defer func() {
				if r := recover(); r != nil {
					log.Fatalf("Panic occurred for line: %s\n%v\n\n", line, r)
				}

			}()
			stream := antlr.NewInputStream(line)
			lexer := parser.NewRequirementsLexer(stream)
			tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
			p := parser.NewRequirementsParser(tokens)

			tree := p.Prog()
			visitor := visitors.NewRequisiteVisitor(tokens)
			visitor.Visit(tree)

			_, err := json.Marshal(visitor.Requirements)
			if err != nil {
				log.Fatalf("failed to marshal: %v", err)
			}
		}()

	}

}
