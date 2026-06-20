package visitors

import (
	conditions2 "parser/objects/conditions"
	"parser/objects/constants"
	"parser/parser"
)

// VisitParenCourse matches a course in the format (course)
//
// Rule:  '(' course ')'
func (v *RequisiteVisitor) VisitParenCourse(ctx *parser.ParenCourseContext) any {
	return v.visitParenCourse(ctx)
}

func (v *RequisiteVisitor) visitParenCourse(ctx *parser.ParenCourseContext) conditions2.Condition {
	return v.Visit(ctx.Course()).(conditions2.Condition)
}

// VisitSimpleCourse matches a course in the format PREFX COURSE_NUMBER
// it also expands courses listed as PREFIX '/' PREFIX '/' ... '/' COURSE_NUMBER (ie CS/SW 2305)
//
// Rule: PREFIX ('/' PREFIX)* COURSE_NUMBER
func (v *RequisiteVisitor) VisitSimpleCourse(ctx *parser.SimpleCourseContext) any {
	return v.visitSimpleCourse(ctx)
}

func (v *RequisiteVisitor) visitSimpleCourse(ctx *parser.SimpleCourseContext) conditions2.Condition {
	number := ctx.COURSE_NUMBER().GetText()
	prefixes := ctx.AllPREFIX()
	if len(prefixes) == 1 {
		return conditions2.NewCourseCondition(prefixes[0].GetText(), number, "")
	}
	conds := make([]conditions2.Condition, len(prefixes))
	for i, p := range prefixes {
		conds[i] = conditions2.NewCourseCondition(p.GetText(), number, "")
	}
	return conditions2.NewOrCondition(conds...)
}

// VisitCrossListedCourse expands courses separated by a slash into an or condition
//
// Rule:  course '/' course
func (v *RequisiteVisitor) VisitCrossListedCourse(ctx *parser.CrossListedCourseContext) interface{} {
	return v.visitCrossListedCourse(ctx)
}

func (v *RequisiteVisitor) visitCrossListedCourse(ctx *parser.CrossListedCourseContext) conditions2.Condition {
	left := v.Visit(ctx.Course(0)).(conditions2.Condition)
	right := v.Visit(ctx.Course(1)).(conditions2.Condition)
	return conditions2.NewOrCondition(left, right)
}

// VisitFullCourseList expands a list of courses
//
// Rule: course (OR course)+
func (v *RequisiteVisitor) VisitFullCourseList(ctx *parser.FullCourseListContext) any {
	return v.visitFullCourseList(ctx)
}

func (v *RequisiteVisitor) visitFullCourseList(ctx *parser.FullCourseListContext) conditions2.Condition {
	list := make([]conditions2.Condition, len(ctx.AllCourse()))
	for i := range list {
		list[i] = v.Visit(ctx.Course(i)).(conditions2.Condition)
	}
	return conditions2.NewOrCondition(list...)
}

// VisitShorthandCourseList expands a list of courses where prefix is implicit
//
// Rule: course (OR COURSE_NUMBER)+
func (v *RequisiteVisitor) VisitShorthandCourseList(ctx *parser.ShorthandCourseListContext) any {
	return v.visitShorthandCourseList(ctx)
}

func (v *RequisiteVisitor) visitShorthandCourseList(ctx *parser.ShorthandCourseListContext) conditions2.Condition {
	prefix := ctx.PREFIX().GetText()
	numbers := ctx.AllCOURSE_NUMBER()

	list := make([]conditions2.Condition, len(numbers))
	for i, number := range numbers {
		list[i] = conditions2.NewCourseCondition(prefix, number.GetText(), "")
	}
	return conditions2.NewOrCondition(list...)
}

// VisitEitherGradeCourseList
//
// Rule: 'either'? course (OR 'in'? course)*
func (v *RequisiteVisitor) VisitEitherGradeCourseList(ctx *parser.EitherGradeCourseListContext) any {
	return v.visitEitherGradeCourseList(ctx)
}

func (v *RequisiteVisitor) visitEitherGradeCourseList(ctx *parser.EitherGradeCourseListContext) conditions2.Condition {
	list := make([]conditions2.Condition, len(ctx.AllCourse()))
	for i := range list {
		list[i] = v.Visit(ctx.Course(i)).(conditions2.Condition)
	}
	return conditions2.NewOrCondition(list...)
}

// VisitAllGradeCourseList
//
// Rule: course (AND course)*
func (v *RequisiteVisitor) VisitAllGradeCourseList(ctx *parser.AllGradeCourseListContext) any {
	return v.visitAllGradeCourseList(ctx)
}

func (v *RequisiteVisitor) visitAllGradeCourseList(ctx *parser.AllGradeCourseListContext) conditions2.Condition {
	list := make([]conditions2.Condition, len(ctx.AllCourse()))
	for i := range list {
		list[i] = v.Visit(ctx.Course(i)).(conditions2.Condition)
	}
	return conditions2.NewAndCondition(list...)
}

// VisitParenGradeCourseList
//
// Rule: '(' course (OR course)* ')'
func (v *RequisiteVisitor) VisitParenGradeCourseList(ctx *parser.ParenGradeCourseListContext) any {
	return v.visitParenGradeCourseList(ctx)
}

func (v *RequisiteVisitor) visitParenGradeCourseList(ctx *parser.ParenGradeCourseListContext) conditions2.Condition {
	list := make([]conditions2.Condition, len(ctx.AllCourse()))
	for i := range list {
		list[i] = v.Visit(ctx.Course(i)).(conditions2.Condition)
	}
	return conditions2.NewOrCondition(list...)
}

func extractCoursesFromCourseList(condition conditions2.Condition) []constants.Course {
	courses := make([]constants.Course, 0)

	switch cond := condition.(type) {
	case *conditions2.CourseCondition:
		courses = append(courses, cond.Course)
	case *conditions2.OrCondition:
		for _, c := range cond.Conditions {
			courses = append(courses, extractCoursesFromCourseList(c)...)
		}
	case *conditions2.AndCondition:
		for _, c := range cond.Conditions {
			courses = append(courses, extractCoursesFromCourseList(c)...)
		}
	}

	return courses
}
