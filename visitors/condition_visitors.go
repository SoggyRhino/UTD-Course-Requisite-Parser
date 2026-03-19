package visitors

import (
	"parser/conditions"
	"parser/parser"
)

// VisitSimpleGradeCondition
//
// Rule: course WITH_GRADE GRADE OR_BETTER?
func (v *RequisiteVisitor) VisitSimpleGradeCondition(ctx *parser.SimpleGradeConditionContext) any {
	return v.visitSimpleGradeCondition(ctx)
}

func (v *RequisiteVisitor) visitSimpleGradeCondition(ctx *parser.SimpleGradeConditionContext) conditions.Condition {
	gradedCond, _ := v.Visit(ctx.Course()).(conditions.GradedCondition)
	gradedCond.AppendGrade(conditions.Grade(ctx.GRADE().GetText()))
	return gradedCond
}

// VisitGpaGradeCondition
//
// Rule: course WITH_GRADE 'greater than or equal to' GRADE ('(' GPA ')')?
func (v *RequisiteVisitor) VisitGpaGradeCondition(ctx *parser.GpaGradeConditionContext) any {
	return v.visitGpaGradeCondition(ctx)
}

func (v *RequisiteVisitor) visitGpaGradeCondition(ctx *parser.GpaGradeConditionContext) conditions.Condition {
	gradedCond, _ := v.Visit(ctx.Course()).(conditions.GradedCondition)
	gradedCond.AppendGrade(conditions.Grade(ctx.GRADE().GetText()))
	return gradedCond
}

// VisitGradeListCondition
//
// Rule: grade_course_list WITH_GRADE GRADE OR_BETTER?
func (v *RequisiteVisitor) VisitGradeListCondition(ctx *parser.GradeListConditionContext) any {
	return v.visitGradeListCondition(ctx)
}

func (v *RequisiteVisitor) visitGradeListCondition(ctx *parser.GradeListConditionContext) conditions.Condition {
	gradedCond, _ := v.Visit(ctx.Grade_course_list()).(conditions.GradedCondition)
	gradedCond.AppendGrade(conditions.Grade(ctx.GRADE().GetText()))
	return gradedCond
}

// VisitGradeAtLeastCondition
//
// Rule: A_GRADE_OF_AT_LEAST GRADE OR_BETTER? 'in' grade_course_list
func (v *RequisiteVisitor) VisitGradeAtLeastCondition(ctx *parser.GradeAtLeastConditionContext) any {
	return v.visitGradeAtLeastCondition(ctx)
}

func (v *RequisiteVisitor) visitGradeAtLeastCondition(ctx *parser.GradeAtLeastConditionContext) conditions.Condition {
	gradedCond, _ := v.Visit(ctx.Grade_course_list()).(conditions.GradedCondition)
	gradedCond.AppendGrade(conditions.Grade(ctx.GRADE().GetText()))
	return gradedCond
}
