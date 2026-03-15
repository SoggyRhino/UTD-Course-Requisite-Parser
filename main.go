package main

import (
	"fmt"
	"parser/conditions"
	"parser/parser"
	"parser/visitors"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input := "Prerequisite: ACN 6340 or HCS 6340."
	stream := antlr.NewInputStream(input)
	lexer := parser.NewRequirementsLexer(stream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewRequirementsParser(tokens)

	tree := p.Prog()
	visitor := &visitors.RequisiteVisitor{
		BaseRequirementsVisitor: parser.BaseRequirementsVisitor{
			BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{},
		},
	}

	result := visitor.Visit(tree)
	if courses, ok := result.([]conditions.Course); ok {
		fmt.Printf("Parsed Courses: %v\n", courses)
	} else {
		fmt.Println("No courses found or unexpected result type")
	}
}
