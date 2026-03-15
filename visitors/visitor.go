package visitors

import (
	"parser/parser"

	"github.com/antlr4-go/antlr/v4"
)

type RequisiteVisitor struct {
	parser.BaseRequirementsVisitor
	Tokens *antlr.CommonTokenStream
}

func NewRequisiteVisitor(tokens *antlr.CommonTokenStream) *RequisiteVisitor {
	return &RequisiteVisitor{Tokens: tokens}
}

func (v *RequisiteVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}
