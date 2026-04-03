package visitors

import (
	"parser/conditions"
	"parser/parser"
	"parser/rules"
)

// VisitParenExpr matches an expression in the format (expr)
//
// Rule: '(' expr ')'
func (v *RequisiteVisitor) VisitParenExpr(ctx *parser.ParenExprContext) any {
	return v.visitParenExpr(ctx)
}

func (v *RequisiteVisitor) visitParenExpr(ctx *parser.ParenExprContext) conditions.Condition {
	return v.Visit(ctx.Expr()).(conditions.Condition)
}

// VisitOrExpr
//
// Rule: expr COMMA? OR expr
func (v *RequisiteVisitor) VisitOrExpr(ctx *parser.OrExprContext) any {
	return v.visitOrExpr(ctx)
}

func (v *RequisiteVisitor) visitOrExpr(ctx *parser.OrExprContext) conditions.Condition {
	cond1 := v.Visit(ctx.Expr(0)).(conditions.Condition)
	cond2 := v.Visit(ctx.Expr(1)).(conditions.Condition)

	return conditions.NewOrConditionFromExpr(cond1, cond2)
}

// VisitAndExpr
//
// Rule: expr COMMA? AND expr
func (v *RequisiteVisitor) VisitAndExpr(ctx *parser.AndExprContext) any {
	return v.visitAndExpr(ctx)
}

func (v *RequisiteVisitor) visitAndExpr(ctx *parser.AndExprContext) conditions.Condition {
	cond1 := v.Visit(ctx.Expr(0)).(conditions.Condition)
	cond2 := v.Visit(ctx.Expr(1)).(conditions.Condition)

	return conditions.NewAndConditionFromExpr(cond1, cond2)
}

// VisitAmpersandExpr
//
// Rule: expr COMMA? AMPERSAND expr
func (v *RequisiteVisitor) VisitAmpersandExpr(ctx *parser.AmpersandExprContext) any {
	return v.visitAmpersandExpr(ctx)
}

func (v *RequisiteVisitor) visitAmpersandExpr(ctx *parser.AmpersandExprContext) conditions.Condition {
	cond1 := v.Visit(ctx.Expr(0)).(conditions.Condition)
	cond2 := v.Visit(ctx.Expr(1)).(conditions.Condition)

	return conditions.NewAndConditionFromExpr(cond1, cond2)
}

// VisitEquivalentExpr
//
// Rule: expr OR EQUIVALENT
func (v *RequisiteVisitor) VisitEquivalentExpr(ctx *parser.EquivalentExprContext) any {
	return v.visitEquivalentExpr(ctx)
}

func (v *RequisiteVisitor) visitEquivalentExpr(ctx *parser.EquivalentExprContext) conditions.Condition {
	cond := v.Visit(ctx.Expr()).(conditions.Condition)
	return conditions.NewAlternativeCondition(cond)
}

// VisitConsentExpr
//
// Rule: consent_condition
func (v *RequisiteVisitor) VisitConsentExpr(ctx *parser.ConsentExprContext) any {
	return v.visitConsentExpr(ctx)
}

func (v *RequisiteVisitor) visitConsentExpr(ctx *parser.ConsentExprContext) conditions.Condition {
	return v.Visit(ctx.Consent_condition()).(conditions.Condition)
}

// VisitStandingExpr
//
// Rule: standing_condition
func (v *RequisiteVisitor) VisitStandingExpr(ctx *parser.StandingExprContext) any {
	return v.visitStandingExpr(ctx)
}

func (v *RequisiteVisitor) visitStandingExpr(ctx *parser.StandingExprContext) conditions.Condition {
	return v.Visit(ctx.Standing_condition()).(conditions.Condition)
}

// VisitGpaExpr
//
// Rule: gpa_condition
func (v *RequisiteVisitor) VisitGpaExpr(ctx *parser.GpaExprContext) any {
	return v.visitGpaExpr(ctx)
}

func (v *RequisiteVisitor) visitGpaExpr(ctx *parser.GpaExprContext) conditions.Condition {
	return v.Visit(ctx.Gpa_condition()).(conditions.Condition)
}

// VisitGroupExpr
//
// Rule: group_condition
func (v *RequisiteVisitor) VisitGroupExpr(ctx *parser.GroupExprContext) any {
	return v.visitGroupExpr(ctx)
}

func (v *RequisiteVisitor) visitGroupExpr(ctx *parser.GroupExprContext) conditions.Condition {
	return v.Visit(ctx.Group_condition()).(conditions.Condition)
}

// VisitConcurrentEnrollmentExpr
//
// Rule: concurrent_enrollment_condition
func (v *RequisiteVisitor) VisitConcurrentEnrollmentExpr(ctx *parser.ConcurrentEnrollmentExprContext) any {
	return v.visitConcurrentEnrollmentExpr(ctx)
}

func (v *RequisiteVisitor) visitConcurrentEnrollmentExpr(ctx *parser.ConcurrentEnrollmentExprContext) conditions.Condition {
	return v.Visit(ctx.Concurrent_enrollment_condition()).(conditions.Condition)
}

// VisitGradeExpr
//
// Rule: grade_condition
func (v *RequisiteVisitor) VisitGradeExpr(ctx *parser.GradeExprContext) any {
	return v.visitGradeExpr(ctx)
}

func (v *RequisiteVisitor) visitGradeExpr(ctx *parser.GradeExprContext) conditions.Condition {
	return v.Visit(ctx.Grade_condition()).(conditions.Condition)
}

// VisitAlternativeExpr
//
// Rule: alternative_condition
func (v *RequisiteVisitor) VisitAlternativeExpr(ctx *parser.AlternativeExprContext) any {
	return v.visitAlternativeExpr(ctx)
}

func (v *RequisiteVisitor) visitAlternativeExpr(ctx *parser.AlternativeExprContext) conditions.Condition {
	return v.Visit(ctx.Alternative_condition()).(conditions.Condition)
}

// VisitGradeLevelStandingExpr
//
// Rule: grade_level_standing_condition
func (v *RequisiteVisitor) VisitGradeLevelStandingExpr(ctx *parser.GradeLevelStandingExprContext) any {
	return v.visitGradeLevelStandingExpr(ctx)
}

func (v *RequisiteVisitor) visitGradeLevelStandingExpr(ctx *parser.GradeLevelStandingExprContext) conditions.Condition {
	return v.Visit(ctx.Grade_level_standing_condition()).(conditions.Condition)
}

// VisitGraduateStandingExpr
//
// Rule: graduate_standing_condition
func (v *RequisiteVisitor) VisitGraduateStandingExpr(ctx *parser.GraduateStandingExprContext) any {
	return v.visitGraduateStandingExpr(ctx)
}

func (v *RequisiteVisitor) visitGraduateStandingExpr(ctx *parser.GraduateStandingExprContext) conditions.Condition {
	return v.Visit(ctx.Graduate_standing_condition()).(conditions.Condition)
}

// VisitMajorExpr
//
// Rule: major_condition
func (v *RequisiteVisitor) VisitMajorExpr(ctx *parser.MajorExprContext) any {
	return v.visitMajorExpr(ctx)
}

func (v *RequisiteVisitor) visitMajorExpr(ctx *parser.MajorExprContext) conditions.Condition {
	return v.Visit(ctx.Major_condition()).(conditions.Condition)
}

// VisitDegreeExpr
//
// Rule: degree_condition
func (v *RequisiteVisitor) VisitDegreeExpr(ctx *parser.DegreeExprContext) any {
	return v.visitDegreeExpr(ctx)
}

func (v *RequisiteVisitor) visitDegreeExpr(ctx *parser.DegreeExprContext) conditions.Condition {
	return v.Visit(ctx.Degree_condition()).(conditions.Condition)
}

// VisitCoreExpr
//
// Rule: core_condition
func (v *RequisiteVisitor) VisitCoreExpr(ctx *parser.CoreExprContext) any {
	return v.visitCoreExpr(ctx)
}

func (v *RequisiteVisitor) visitCoreExpr(ctx *parser.CoreExprContext) conditions.Condition {
	return v.Visit(ctx.Core_condition()).(conditions.Condition)
}

// VisitAnyCoreExpr
//
// Rule: any_core_condition
func (v *RequisiteVisitor) VisitAnyCoreExpr(ctx *parser.AnyCoreExprContext) any {
	return v.visitAnyCoreExpr(ctx)
}

func (v *RequisiteVisitor) visitAnyCoreExpr(ctx *parser.AnyCoreExprContext) conditions.Condition {
	return v.Visit(ctx.Any_core_condition()).(conditions.Condition)
}

// VisitCompleteNExpr
//
// Rule: complete_n_condition
func (v *RequisiteVisitor) VisitCompleteNExpr(ctx *parser.CompleteNExprContext) any {
	return v.visitCompleteNExpr(ctx)
}

func (v *RequisiteVisitor) visitCompleteNExpr(ctx *parser.CompleteNExprContext) conditions.Condition {
	return v.Visit(ctx.Complete_n_condition()).(conditions.Condition)
}

// VisitSemesterCreditHoursExpr
//
// Rule: semester_credit_hours_condition
func (v *RequisiteVisitor) VisitSemesterCreditHoursExpr(ctx *parser.SemesterCreditHoursExprContext) any {
	return v.visitSemesterCreditHoursExpr(ctx)
}

func (v *RequisiteVisitor) visitSemesterCreditHoursExpr(ctx *parser.SemesterCreditHoursExprContext) conditions.Condition {
	return v.Visit(ctx.Semester_credit_hours_condition()).(conditions.Condition)
}

// VisitMinimumHoursExpr
//
// Rule: minimum_hours_condition
func (v *RequisiteVisitor) VisitMinimumHoursExpr(ctx *parser.MinimumHoursExprContext) any {
	return v.visitMinimumHoursExpr(ctx)
}

func (v *RequisiteVisitor) visitMinimumHoursExpr(ctx *parser.MinimumHoursExprContext) conditions.Condition {
	return v.Visit(ctx.Minimum_hours_condition()).(conditions.Condition)
}

// VisitUpperDivisionHoursExpr
//
// Rule: upper_division_hours_condition
func (v *RequisiteVisitor) VisitUpperDivisionHoursExpr(ctx *parser.UpperDivisionHoursExprContext) any {
	return v.visitUpperDivisionHoursExpr(ctx)
}

func (v *RequisiteVisitor) visitUpperDivisionHoursExpr(ctx *parser.UpperDivisionHoursExprContext) conditions.Condition {
	return v.Visit(ctx.Upper_division_hours_condition()).(conditions.Condition)
}

// VisitUpperDivisionClassesExpr
//
// Rule: uppper_division_classes_condition
func (v *RequisiteVisitor) VisitUpperDivisionClassesExpr(ctx *parser.UpperDivisionClassesExprContext) any {
	return v.visitUpperDivisionClassesExpr(ctx)
}

func (v *RequisiteVisitor) visitUpperDivisionClassesExpr(ctx *parser.UpperDivisionClassesExprContext) conditions.Condition {
	return v.Visit(ctx.Upper_division_classes_condition()).(conditions.Condition)
}

// VisitResearchExpr
//
// Rule: research_condition
func (v *RequisiteVisitor) VisitResearchExpr(ctx *parser.ResearchExprContext) any {
	return v.visitResearchExpr(ctx)
}

func (v *RequisiteVisitor) visitResearchExpr(ctx *parser.ResearchExprContext) conditions.Condition {
	return v.Visit(ctx.Research_condition()).(conditions.Condition)
}

// VisitPlacementTestExpr
//
// Rule: placement_test_condition
func (v *RequisiteVisitor) VisitPlacementTestExpr(ctx *parser.PlacementTestExprContext) any {
	return v.visitPlacementTestExpr(ctx)
}

func (v *RequisiteVisitor) visitPlacementTestExpr(ctx *parser.PlacementTestExprContext) conditions.Condition {
	return v.Visit(ctx.Placement_test_condition()).(conditions.Condition)
}

// VisitApScoreExpr
//
// Rule: ap_score_condition
func (v *RequisiteVisitor) VisitApScoreExpr(ctx *parser.ApScoreExprContext) any {
	return v.visitApScoreExpr(ctx)
}

func (v *RequisiteVisitor) visitApScoreExpr(ctx *parser.ApScoreExprContext) conditions.Condition {
	return v.Visit(ctx.Ap_score_condition()).(conditions.Condition)
}

// VisitAleksScoreExpr
//
// Rule: aleks_score_condition
func (v *RequisiteVisitor) VisitAleksScoreExpr(ctx *parser.AleksScoreExprContext) any {
	return v.visitAleksScoreExpr(ctx)
}

func (v *RequisiteVisitor) visitAleksScoreExpr(ctx *parser.AleksScoreExprContext) conditions.Condition {
	return v.Visit(ctx.Aleks_score_condition()).(conditions.Condition)
}

// VisitExactSectionExpr
//
// Rule: exact_section_condition
func (v *RequisiteVisitor) VisitExactSectionExpr(ctx *parser.ExactSectionExprContext) any {
	return v.visitExactSectionExpr(ctx)
}

func (v *RequisiteVisitor) visitExactSectionExpr(ctx *parser.ExactSectionExprContext) conditions.Condition {
	return v.Visit(ctx.Exact_section_condition()).(conditions.Condition)
}

// VisitAnyMajorCourseExpr
//
// Rule: any_major_course_condition
func (v *RequisiteVisitor) VisitAnyMajorCourseExpr(ctx *parser.AnyMajorCourseExprContext) any {
	return v.visitAnyMajorCourseExpr(ctx)
}

func (v *RequisiteVisitor) visitAnyMajorCourseExpr(ctx *parser.AnyMajorCourseExprContext) conditions.Condition {
	return v.Visit(ctx.Any_major_course_condition()).(conditions.Condition)
}

// VisitLivingLearningExpr
//
// Rule: living_learning_rule
func (v *RequisiteVisitor) VisitLivingLearningExpr(ctx *parser.LivingLearningExprContext) any {
	return v.visitLivingLearningExpr(ctx)
}

func (v *RequisiteVisitor) visitLivingLearningExpr(ctx *parser.LivingLearningExprContext) *rules.LivingLearningRule {
	return v.Visit(ctx.Living_learning_rule()).(*rules.LivingLearningRule)
}

// VisitRepeatRuleExpr
//
// Rule: repeat_rule
func (v *RequisiteVisitor) VisitRepeatRuleExpr(ctx *parser.RepeatRuleExprContext) any {
	return v.visitRepeatRuleExpr(ctx)
}

func (v *RequisiteVisitor) visitRepeatRuleExpr(ctx *parser.RepeatRuleExprContext) *rules.RepeatRule {
	return v.Visit(ctx.Repeat_rule()).(*rules.RepeatRule)
}

// VisitRepeatLimitHoursExpr
//
// Rule: repeat_limit_hours_rule
func (v *RequisiteVisitor) VisitRepeatLimitHoursExpr(ctx *parser.RepeatLimitHoursExprContext) any {
	return v.visitRepeatLimitHoursExpr(ctx)
}

func (v *RequisiteVisitor) visitRepeatLimitHoursExpr(ctx *parser.RepeatLimitHoursExprContext) *rules.RepeatRule {
	return v.Visit(ctx.Repeat_limit_hours_rule()).(*rules.RepeatRule)
}

// VisitCourseExpr
//
// Rule: course
func (v *RequisiteVisitor) VisitCourseExpr(ctx *parser.CourseExprContext) any {
	return v.visitCourseExpr(ctx)
}

func (v *RequisiteVisitor) visitCourseExpr(ctx *parser.CourseExprContext) conditions.Condition {
	return v.Visit(ctx.Course()).(conditions.Condition)
}
