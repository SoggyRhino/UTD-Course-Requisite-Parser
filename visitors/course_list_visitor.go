package visitors

import (
	"parser/conditions"
	"parser/parser"
)

/*
	 course_list
		: course (OR course)+                   # fullCourseList
		| course (OR COURSE_NUMBER)+            # shorthandCourseList
		;
*/

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
	return conditions.OrCondition{Conditions: list}
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
		list[i] = conditions.CourseCondition{
			Course: conditions.NewCourse(prefix, number.GetText()),
		}
	}
	return conditions.OrCondition{Conditions: list}
}
