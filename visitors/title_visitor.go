package visitors

import (
	"parser/parser"

	"github.com/antlr4-go/antlr/v4"
)

// VisitTitle matches what is essentially a name / propper noun
//
// Rule: (CAPITALIZED | CORE)+ ( 'and' (CAPITALIZED | CORE)+ )*
func (v *RequisiteVisitor) VisitTitle(ctx *parser.TitleContext) any {
	return v.visitTitle(ctx)
}

func (v *RequisiteVisitor) visitTitle(ctx *parser.TitleContext) string {
	start := ctx.GetStart().GetStart()
	stop := ctx.GetStop().GetStop()

	return ctx.GetStart().GetInputStream().GetTextFromInterval(antlr.Interval{Start: start, Stop: stop})
}
