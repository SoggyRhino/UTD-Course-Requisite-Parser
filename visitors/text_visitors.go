package visitors

import (
	"parser/parser"
)

// VisitTitle matches what is essentially a name / propper noun
//
// Rule: (CAPITALIZED | CORE)+ ( 'and' (CAPITALIZED | CORE)+ )*
func (v *RequisiteVisitor) VisitTitle(ctx *parser.TitleContext) any {
	return v.visitTitle(ctx)
}

func (v *RequisiteVisitor) visitTitle(ctx *parser.TitleContext) string {
	return v.getText(ctx.BaseParserRuleContext)
}

// VisitDegree
//
// Rule: degree : degree_atom+ (AND degree_atom+)* ;
func (v *RequisiteVisitor) VisitDegree(ctx *parser.DegreeContext) any {
	return v.visitDegree(ctx)
}

func (v *RequisiteVisitor) visitDegree(ctx *parser.DegreeContext) string {
	return v.getText(ctx.BaseParserRuleContext)
}

// VisitDegree_list
//
// Rule: degree_list : degree ((COMMA | OR | COMMA AND) degree)* ;
func (v *RequisiteVisitor) VisitDegree_list(ctx *parser.Degree_listContext) any {
	return v.visitDegreeList(ctx)
}

func (v *RequisiteVisitor) visitDegreeList(ctx *parser.Degree_listContext) []string {
	degrees := make([]string, len(ctx.AllDegree()))
	for i, degree := range ctx.AllDegree() {
		degrees[i] = v.visitDegree(degree.(*parser.DegreeContext))
	}

	return degrees
}

// VisitPlacement_test_name
//
// Rule : PREFIX PLACEMENT_KW TEST_KW | title PLACEMENT_KW TEST_KW?
func (v *RequisiteVisitor) VisitPlacement_test_name(ctx *parser.Placement_test_nameContext) any {
	return v.visitPlacementTestName(ctx)
}

func (v *RequisiteVisitor) visitPlacementTestName(ctx *parser.Placement_test_nameContext) string {
	return v.getText(ctx.BaseParserRuleContext)
}
