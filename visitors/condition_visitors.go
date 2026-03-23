package visitors

import (
	"parser/conditions"
	"parser/parser"

	"github.com/antlr4-go/antlr/v4"
)

// ======================= grade condition =======================

func (v *RequisiteVisitor) applyGrade(node antlr.ParseTree, gradeText string) conditions.Condition {
	gradedCond, _ := v.Visit(node).(conditions.GradedCondition)
	gradedCond.AppendGrade(conditions.Grade(gradeText))
	return gradedCond
}

// VisitSimpleGradeCondition
//
// Rule: course WITH_GRADE GRADE OR_BETTER?
func (v *RequisiteVisitor) VisitSimpleGradeCondition(ctx *parser.SimpleGradeConditionContext) any {
	return v.visitSimpleGradeCondition(ctx)
}

func (v *RequisiteVisitor) visitSimpleGradeCondition(ctx *parser.SimpleGradeConditionContext) conditions.Condition {
	return v.applyGrade(ctx.Course(), ctx.GRADE().GetText())
}

// VisitGpaGradeCondition
//
// Rule: course WITH_GRADE 'greater than or equal to' GRADE ('(' GPA ')')?
func (v *RequisiteVisitor) VisitGpaGradeCondition(ctx *parser.GpaGradeConditionContext) any {
	return v.visitGpaGradeCondition(ctx)
}

func (v *RequisiteVisitor) visitGpaGradeCondition(ctx *parser.GpaGradeConditionContext) conditions.Condition {
	return v.applyGrade(ctx.Course(), ctx.GRADE().GetText())
}

// VisitGradeListCondition
//
// Rule: grade_course_list WITH_GRADE GRADE OR_BETTER?
func (v *RequisiteVisitor) VisitGradeListCondition(ctx *parser.GradeListConditionContext) any {
	return v.visitGradeListCondition(ctx)
}

func (v *RequisiteVisitor) visitGradeListCondition(ctx *parser.GradeListConditionContext) conditions.Condition {
	return v.applyGrade(ctx.Grade_course_list(), ctx.GRADE().GetText())
}

// VisitGradeAtLeastCondition
//
// Rule: A_GRADE_OF_AT_LEAST GRADE OR_BETTER? 'in' grade_course_list
func (v *RequisiteVisitor) VisitGradeAtLeastCondition(ctx *parser.GradeAtLeastConditionContext) any {
	return v.visitGradeAtLeastCondition(ctx)
}

func (v *RequisiteVisitor) visitGradeAtLeastCondition(ctx *parser.GradeAtLeastConditionContext) conditions.Condition {
	return v.applyGrade(ctx.Grade_course_list(), ctx.GRADE().GetText())
}

// ======================= (course) alternative condition =======================

// VisitCourseAlternativeCondition
//
// Rule: course OR EQUIVALENT
func (v *RequisiteVisitor) VisitCourseAlternativeCondition(ctx *parser.CourseAlternativeConditionContext) any {
	return v.visitCourseAlternativeCondition(ctx)
}
func (v *RequisiteVisitor) visitCourseAlternativeCondition(ctx *parser.CourseAlternativeConditionContext) *conditions.AlternativeCondition {
	course := v.Visit(ctx.Course()).(conditions.Condition)
	return conditions.NewAlternativeCondition(course)
}

// VisitGradeCourseListAlternativeCondition
//
// Rule: course_list OR EQUIVALENT
func (v *RequisiteVisitor) VisitGradeCourseListAlternativeCondition(ctx *parser.GradeCourseListAlternativeConditionContext) any {
	return v.visitGradeCourseListAlternativeCondition(ctx)
}

func (v *RequisiteVisitor) visitGradeCourseListAlternativeCondition(ctx *parser.GradeCourseListAlternativeConditionContext) *conditions.AlternativeCondition {
	courseList := v.Visit(ctx.Course_list()).(conditions.Condition)
	return conditions.NewAlternativeCondition(courseList)
}
