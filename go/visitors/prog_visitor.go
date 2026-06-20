package visitors

import (
	"parser/parser"
)

// VisitProg
//
// Rule: sentence+ EOF ;
func (v *RequisiteVisitor) VisitProg(ctx *parser.ProgContext) any {
	for _, s := range ctx.AllSentence() {
		v.Visit(s)
	}

	return v.Requirements
}

// VisitSentence
//
// sentence : requisite PERIOD? ;
func (v *RequisiteVisitor) VisitSentence(ctx *parser.SentenceContext) any {
	return v.Visit(ctx.Requisite())
}
