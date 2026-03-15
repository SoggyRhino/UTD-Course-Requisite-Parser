package visitors

import (
	"parser/conditions"
	"parser/parser"
)

// VisitParenCourse matches a course in the format (course)
//
// Rule:  '(' course ')'
func (v *RequisiteVisitor) VisitParenCourse(ctx *parser.ParenCourseContext) any {
	return v.visitParenCourse(ctx)
}

func (v *RequisiteVisitor) visitParenCourse(ctx *parser.ParenCourseContext) conditions.Condition {
	return v.Visit(ctx.Course()).(conditions.Condition)
}

// VisitSimpleCourse matches a course in the format PREFX COURSE_NUMBER
// it also expands courses listed as PREFIX '/' PREFIX '/' ... '/' COURSE_NUMBER (ie CS/SW 2305)
//
// Rule: PREFIX ('/' PREFIX)* COURSE_NUMBER
func (v *RequisiteVisitor) VisitSimpleCourse(ctx *parser.SimpleCourseContext) any {
	return v.visitSimpleCourse(ctx)
}

func (v *RequisiteVisitor) visitSimpleCourse(ctx *parser.SimpleCourseContext) conditions.Condition {
	number := ctx.COURSE_NUMBER().GetText()
	prefixes := ctx.AllPREFIX()
	if len(prefixes) == 1 {
		return conditions.CourseCondition{Course: conditions.NewCourse(prefixes[0].GetText(), number)}
	}
	conds := make([]conditions.Condition, len(prefixes))
	for i, p := range prefixes {
		conds[i] = conditions.CourseCondition{Course: conditions.NewCourse(p.GetText(), number)}
	}
	return conditions.OrCondition{Conditions: conds}
}

// VisitCrossListedCourse expands courses separated by a slash into an or condition
//
// Rule:  course '/' course
func (v *RequisiteVisitor) VisitCrossListedCourse(ctx *parser.CrossListedCourseContext) interface{} {
	return v.visitCrossListedCourse(ctx)
}

func (v *RequisiteVisitor) visitCrossListedCourse(ctx *parser.CrossListedCourseContext) conditions.Condition {
	left := v.Visit(ctx.Course(0)).(conditions.Condition)
	right := v.Visit(ctx.Course(1)).(conditions.Condition)
	return conditions.OrCondition{Conditions: []conditions.Condition{left, right}}
}
