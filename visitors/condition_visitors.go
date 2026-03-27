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

// ======================= grade level standing condition condition =======================

// VisitGradeLevelStandingCondition
//
// Rule: GRADE_LEVEL (OR GRADE_LEVEL)* LEVEL? STANDING
func (v *RequisiteVisitor) VisitGradeLevelStandingCondition(ctx *parser.GradeLevelStandingConditionContext) any {
	return v.visitGradeLevelStandingCondition(ctx)
}

func (v *RequisiteVisitor) visitGradeLevelStandingCondition(ctx *parser.GradeLevelStandingConditionContext) *conditions.OrCondition {
	conds := make([]conditions.Condition, len(ctx.AllGRADE_LEVEL()))
	for i, gradeLevel := range ctx.AllGRADE_LEVEL() {
		conds[i] = conditions.NewGradeLevelCondition(mapGradeLevel(gradeLevel.GetText()))
	}

	return conditions.NewOrCondition(conds...)
}

// VisitGradeLevelMajorStandingCondition
//
// Rule: GRADE_LEVEL degree MAJOR_KW STANDING
func (v *RequisiteVisitor) VisitGradeLevelMajorStandingCondition(ctx *parser.GradeLevelMajorStandingConditionContext) any {
	return v.visitGradeLevelMajorStandingCondition(ctx)
}

func (v *RequisiteVisitor) visitGradeLevelMajorStandingCondition(ctx *parser.GradeLevelMajorStandingConditionContext) *conditions.GradeLevelCondition {
	return conditions.NewGradeLevelConditionWithDegree(mapGradeLevel(ctx.GRADE_LEVEL().GetText()), ctx.Degree().GetText())
}

// VisitMinimumGradeLevelStandingCondition
//
// Rule: MINIMUM_OF GRADE_LEVEL STANDING
func (v *RequisiteVisitor) VisitMinimumGradeLevelStandingCondition(ctx *parser.MinimumGradeLevelStandingConditionContext) any {
	return v.visitMinimumGradeLevelStandingCondition(ctx)
}

func (v *RequisiteVisitor) visitMinimumGradeLevelStandingCondition(ctx *parser.MinimumGradeLevelStandingConditionContext) *conditions.GradeLevelCondition {
	return conditions.NewGradeLevelCondition(mapGradeLevel(ctx.GRADE_LEVEL().GetText()))
}

// VisitAtLeastGradeLevelStandingCondition
//
// Rule: AT_LEAST GRADE_LEVEL (DASH LEVEL | LEVEL)? STANDING
func (v *RequisiteVisitor) VisitAtLeastGradeLevelStandingCondition(ctx *parser.AtLeastGradeLevelStandingConditionContext) any {
	return v.visitAtLeastGradeLevelStandingCondition(ctx)
}

func (v *RequisiteVisitor) visitAtLeastGradeLevelStandingCondition(ctx *parser.AtLeastGradeLevelStandingConditionContext) *conditions.GradeLevelCondition {
	return conditions.NewGradeLevelCondition(mapGradeLevel(ctx.GRADE_LEVEL().GetText()))
}

// VisitPrefixGradeLevelStandingCondition
//
// Rule: PREFIX MAJOR_KW? ONLY_KW 'with' GRADE_LEVEL (OR GRADE_LEVEL)* LEVEL? STANDING
func (v *RequisiteVisitor) VisitPrefixGradeLevelStandingCondition(ctx *parser.PrefixGradeLevelStandingConditionContext) any {
	return v.visitPrefixGradeLevelStandingCondition(ctx)
}

//todo post-process conditions and collapse and/ors with single conditions

func (v *RequisiteVisitor) visitPrefixGradeLevelStandingCondition(ctx *parser.PrefixGradeLevelStandingConditionContext) *conditions.OrCondition {
	//todo map/standardize major/degree/school etc
	degree := ctx.PREFIX().GetText()

	conds := make([]conditions.Condition, len(ctx.AllGRADE_LEVEL()))
	for i, gradeLevel := range ctx.AllGRADE_LEVEL() {
		conds[i] = conditions.NewGradeLevelConditionWithDegree(mapGradeLevel(gradeLevel.GetText()), degree)
	}

	return conditions.NewOrCondition(conds...)
}

// ======================= graduate standing condition =======================

// VisitGraduateStandingInCondition
//
// Rule: 'Graduate standing in' degree
func (v *RequisiteVisitor) VisitGraduateStandingInCondition(ctx *parser.GraduateStandingInConditionContext) any {
	return v.visitGraduateStandingInCondition(ctx)
}

func (v *RequisiteVisitor) visitGraduateStandingInCondition(ctx *parser.GraduateStandingInConditionContext) *conditions.GraduateStandingInCondition {
	return conditions.NewGraduateStandingInConditionWithDegree(ctx.Degree().GetText())
}

// VisitGraduateLevelStandingCondition
//
// Rule: 'Graduate Level Standing'
func (v *RequisiteVisitor) VisitGraduateLevelStandingCondition(ctx *parser.GraduateLevelStandingConditionContext) any {
	return v.visitGraduateLevelStandingCondition(ctx)
}

func (v *RequisiteVisitor) visitGraduateLevelStandingCondition(ctx *parser.GraduateLevelStandingConditionContext) *conditions.GraduateStandingInCondition {
	return conditions.NewGraduateStandingInCondition()
}

// ======================= GPA condition =======================

// VisitUniversityGpaCondition
//
// Rule: UNIVERSITY_GPA_KW GPA
func (v *RequisiteVisitor) VisitUniversityGpaCondition(ctx *parser.UniversityGpaConditionContext) any {
	return v.visitUniversityGpaCondition(ctx)
}

func (v *RequisiteVisitor) visitUniversityGpaCondition(ctx *parser.UniversityGpaConditionContext) *conditions.GPACondition {
	return conditions.NewGpaCondition(mapGPA(ctx.GPA().GetText()))
}

// VisitMinimumGpaCondition
//
// Rule: 'Minimum GPA requirement' GPA
func (v *RequisiteVisitor) VisitMinimumGpaCondition(ctx *parser.MinimumGpaConditionContext) any {
	return v.visitMinimumGpaCondition(ctx)
}

func (v *RequisiteVisitor) visitMinimumGpaCondition(ctx *parser.MinimumGpaConditionContext) *conditions.GPACondition {
	return conditions.NewGpaCondition(mapGPA(ctx.GPA().GetText()))
}

// VisitGpaInCourseCondition
//
// Rule: 'a GPA of' GPA OR_BETTER 'in' degree COURSE_KW?
func (v *RequisiteVisitor) VisitGpaInCourseCondition(ctx *parser.GpaInCourseConditionContext) any {
	return v.visitGpaInCourseCondition(ctx)
}

func (v *RequisiteVisitor) visitGpaInCourseCondition(ctx *parser.GpaInCourseConditionContext) *conditions.GPACondition {
	degree := v.Visit(ctx.Degree()).(string)
	return conditions.NewGpaConditionWithDegree(mapGPA(ctx.GPA().GetText()), degree)
}
