package visitors

import (
	"parser/conditions"
	"parser/parser"
	rules "parser/rules"
	"parser/utils"
)

// ======================= Repeat Restriction Rule =======================

// VisitCourseRepeatRule
//
// Rule: course REPEAT_RESTRICTION
func (v *RequisiteVisitor) VisitCourseRepeatRule(ctx *parser.CourseRepeatRuleContext) any {
	return v.visitCourseRepeatRule(ctx)
}

func (v *RequisiteVisitor) visitCourseRepeatRule(ctx *parser.CourseRepeatRuleContext) []*rules.RepeatRule {
	course := extractCoursesFromCourseList(v.Visit(ctx.Course()).(conditions.Condition))
	list := make([]*rules.RepeatRule, len(course))
	for i, c := range course {
		list[i] = rules.NewCourseRepeatRule(c)
	}

	return list
}

// VisitInternshipRepeatRule
//
// Rule: PREFIX 'Internship' REPEAT_RESTRICTION
func (v *RequisiteVisitor) VisitInternshipRepeatRule(ctx *parser.InternshipRepeatRuleContext) any {
	return v.visitInternshipRepeatRule(ctx)
}

func (v *RequisiteVisitor) visitInternshipRepeatRule(ctx *parser.InternshipRepeatRuleContext) *rules.RepeatRule {
	return rules.NewInternshipRepeatRule(ctx.PREFIX().GetText())
}

// VisitBareRepeatRule
//
// Rule: REPEAT_RESTRICTION
func (v *RequisiteVisitor) VisitBareRepeatRule(ctx *parser.BareRepeatRuleContext) any {
	return v.visitBareRepeatRule(ctx)
}

func (v *RequisiteVisitor) visitBareRepeatRule(ctx *parser.BareRepeatRuleContext) *rules.RepeatRule {
	//todo better blank value
	return rules.NewRepeatRule(0, 0, utils.Course{}, "")
}
