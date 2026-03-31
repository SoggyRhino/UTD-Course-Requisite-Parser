package visitors

import (
	"parser/parser"

	"github.com/antlr4-go/antlr/v4"
)

var _ parser.RequirementsVisitor = (*RequisiteVisitor)(nil)

type RequisiteVisitor struct {
	parser.BaseRequirementsVisitor
	Tokens *antlr.CommonTokenStream
	Errors []error
}

func NewRequisiteVisitor(tokens *antlr.CommonTokenStream) *RequisiteVisitor {
	return &RequisiteVisitor{
		Tokens: tokens,
		Errors: make([]error, 0),
	}
}

func (v *RequisiteVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *RequisiteVisitor) ReportError(ctx antlr.BaseParserRuleContext, err error) {
	//todo use ctx for better error handling
	v.Errors = append(v.Errors, err)
}

func (v *RequisiteVisitor) getText(ctx antlr.BaseParserRuleContext) string {
	start := ctx.GetStart().GetStart()
	stop := ctx.GetStop().GetStop()

	return ctx.GetStart().GetInputStream().GetTextFromInterval(antlr.Interval{Start: start, Stop: stop})
}

func (v *RequisiteVisitor) getTextOrDefault(node antlr.TerminalNode, str string) string {
	if node != nil {
		return node.GetText()
	}
	return str
}

func (v *RequisiteVisitor) firstOrNil(nodes []antlr.TerminalNode) antlr.TerminalNode {
	if len(nodes) == 0 {
		return nil
	}
	return nodes[0]
}
