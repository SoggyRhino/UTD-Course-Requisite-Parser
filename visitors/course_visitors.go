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
		return conditions.NewCourseCondition(prefixes[0].GetText(), number, "")
	}
	conds := make([]conditions.Condition, len(prefixes))
	for i, p := range prefixes {
		conds[i] = conditions.NewCourseCondition(p.GetText(), number, "")
	}
	return conditions.NewOrCondition(conds...)
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
	return conditions.NewOrCondition(left, right)
}

// VisitFullCourseList expands a list of courses
//
// Rule: course (OR course)+
func (v *RequisiteVisitor) VisitFullCourseList(ctx *parser.FullCourseListContext) any {
	return v.visitFullCourseList(ctx)
}

func (v *RequisiteVisitor) visitFullCourseList(ctx *parser.FullCourseListContext) conditions.Condition {
	list := make([]conditions.Condition, len(ctx.AllCourse()))
	for i := range list {
		list[i] = v.Visit(ctx.Course(i)).(conditions.Condition)
	}
	return conditions.NewOrCondition(list...)
}

// VisitShorthandCourseList expands a list of courses where prefix is implicit
//
// Rule: course (OR COURSE_NUMBER)+
func (v *RequisiteVisitor) VisitShorthandCourseList(ctx *parser.ShorthandCourseListContext) any {
	return v.visitShorthandCourseList(ctx)
}

func (v *RequisiteVisitor) visitShorthandCourseList(ctx *parser.ShorthandCourseListContext) conditions.Condition {
	prefix := ctx.PREFIX().GetText()
	numbers := ctx.AllCOURSE_NUMBER()

	list := make([]conditions.Condition, len(numbers))
	for i, number := range numbers {
		list[i] = conditions.NewCourseCondition(prefix, number.GetText(), "")
	}
	return conditions.NewOrCondition(list...)
}

// VisitEitherGradeCourseList
//
// Rule : 'either'? course (OR 'in'? course)*
func (v *RequisiteVisitor) VisitEitherGradeCourseList(ctx *parser.EitherGradeCourseListContext) any {
	return v.visitEitherGradeCourseList(ctx)
}

func (v *RequisiteVisitor) visitEitherGradeCourseList(ctx *parser.EitherGradeCourseListContext) *conditions.OrCondition {
	list := make([]conditions.Condition, len(ctx.AllCourse()))
	for i := range list {
		list[i] = v.Visit(ctx.Course(i)).(conditions.Condition)
	}
	return conditions.NewOrCondition(list...)
}

// VisitAllGradeCourseList
//
// Rule : course (AND course)*
func (v *RequisiteVisitor) VisitAllGradeCourseList(ctx *parser.AllGradeCourseListContext) any {
	return v.visitAllGradeCourseList(ctx)
}

func (v *RequisiteVisitor) visitAllGradeCourseList(ctx *parser.AllGradeCourseListContext) *conditions.AndCondition {
	list := make([]conditions.Condition, len(ctx.AllCourse()))
	for i := range list {
		list[i] = v.Visit(ctx.Course(i)).(conditions.Condition)
	}
	return conditions.NewAndCondition(list...)
}

// VisitParenGradeCourseList
//
// Rule : '(' course (OR course)* ')'
func (v *RequisiteVisitor) VisitParenGradeCourseList(ctx *parser.ParenGradeCourseListContext) any {
	return v.visitParenGradeCourseList(ctx)
}

func (v *RequisiteVisitor) visitParenGradeCourseList(ctx *parser.ParenGradeCourseListContext) *conditions.OrCondition {
	list := make([]conditions.Condition, len(ctx.AllCourse()))
	for i := range list {
		list[i] = v.Visit(ctx.Course(i)).(conditions.Condition)
	}
	return conditions.NewOrCondition(list...)
}
