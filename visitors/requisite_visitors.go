package visitors

import (
	"parser/conditions"
	"parser/constants"
	"parser/parser"
	"parser/rules"
)

// VisitExcludeNoticeReq
//
// Rule: requisite AND EXCLUDE_DMHP_LLC_NOTICE
func (v *RequisiteVisitor) VisitExcludeNoticeReq(ctx *parser.ExcludeNoticeReqContext) any {
	return v.visitExcludeNoticeReq(ctx)
}

func (v *RequisiteVisitor) visitExcludeNoticeReq(ctx *parser.ExcludeNoticeReqContext) any {
	v.appendNotice(constants.ExcludeDMHPLLCNotice)
	return v.Visit(ctx.Requisite())
}

// VisitAppendAcademicPlanReq
//
// Rule: requisite AND academic_plan_condition
func (v *RequisiteVisitor) VisitAppendAcademicPlanReq(ctx *parser.AppendAcademicPlanReqContext) any {
	return v.visitAppendAcademicPlanReq(ctx)
}

func (v *RequisiteVisitor) visitAppendAcademicPlanReq(ctx *parser.AppendAcademicPlanReqContext) any {
	condition := v.Visit(ctx.Academic_plan_condition()).(conditions.Condition)
	v.appendPreReq(condition)
	return v.Visit(ctx.Requisite())
}

// VisitAcademicPlanReq
//
// Rule: academic_plan_condition
func (v *RequisiteVisitor) VisitAcademicPlanReq(ctx *parser.AcademicPlanReqContext) any {
	return v.visitAcademicPlanReq(ctx)
}

func (v *RequisiteVisitor) visitAcademicPlanReq(ctx *parser.AcademicPlanReqContext) *conditions.AcademicYearCondition {
	cond := v.Visit(ctx.Academic_plan_condition()).(*conditions.AcademicYearCondition)
	v.appendPreReq(cond)
	return cond
}

// VisitExactCoreqNoticeReq
//
// Rule: EXACT_COREQ_NOTICE
func (v *RequisiteVisitor) VisitExactCoreqNoticeReq(ctx *parser.ExactCoreqNoticeReqContext) any {
	return v.visitExactCoreqNoticeReq(ctx)
}

func (v *RequisiteVisitor) visitExactCoreqNoticeReq(ctx *parser.ExactCoreqNoticeReqContext) any {
	v.appendNotice(constants.ExactCoReqNotice)
	return nil
}

// VisitComputerScholarsReq
//
// Rule: COMPUTER_SCHOLARS_PROGRAM
func (v *RequisiteVisitor) VisitComputerScholarsReq(ctx *parser.ComputerScholarsReqContext) any {
	return v.visitComputerScholarsReq(ctx)
}

func (v *RequisiteVisitor) visitComputerScholarsReq(ctx *parser.ComputerScholarsReqContext) any {
	v.appendNotice(constants.ComputerScholarsProgramNotice)
	return nil
}

// VisitGpaRepeatReq
//
// Rule: gpa_repeate_rule
func (v *RequisiteVisitor) VisitGpaRepeatReq(ctx *parser.GpaRepeatReqContext) any {
	return v.visitGpaRepeatReq(ctx)
}

func (v *RequisiteVisitor) visitGpaRepeatReq(ctx *parser.GpaRepeatReqContext) *rules.GpaRepeatRule {
	rule := v.Visit(ctx.Gpa_repeate_rule()).(*rules.GpaRepeatRule)
	v.appendRule(rule)
	return rule
}

// VisitRepeatLimitHoursReq
//
// Rule: repeat_limit_hours_rule
func (v *RequisiteVisitor) VisitRepeatLimitHoursReq(ctx *parser.RepeatLimitHoursReqContext) any {
	return v.visitRepeatLimitHoursReq(ctx)
}

func (v *RequisiteVisitor) visitRepeatLimitHoursReq(ctx *parser.RepeatLimitHoursReqContext) *rules.RepeatRule {
	rule := v.Visit(ctx.Repeat_limit_hours_rule()).(*rules.RepeatRule)
	v.appendRule(rule)
	return rule
}

// VisitRepeatLimitTimesReq
//
// Rule: repeat_limit_times_rule
func (v *RequisiteVisitor) VisitRepeatLimitTimesReq(ctx *parser.RepeatLimitTimesReqContext) any {
	return v.visitRepeatLimitTimesReq(ctx)
}

func (v *RequisiteVisitor) visitRepeatLimitTimesReq(ctx *parser.RepeatLimitTimesReqContext) *rules.RepeatRule {
	rule := v.Visit(ctx.Repeat_limit_times_rule()).(*rules.RepeatRule)
	v.appendRule(rule)
	return rule
}

// VisitRepeatReq
//
// Rule: repeat_rule
func (v *RequisiteVisitor) VisitRepeatReq(ctx *parser.RepeatReqContext) any {
	return v.visitRepeatReq(ctx)
}

func (v *RequisiteVisitor) visitRepeatReq(ctx *parser.RepeatReqContext) *rules.RepeatRule {
	rule := v.Visit(ctx.Repeat_rule()).(*rules.RepeatRule)
	v.appendRule(rule)
	return rule
}

// VisitDegreeSatisfactionReq
//
// Rule: degree_satisfaction_rule
func (v *RequisiteVisitor) VisitDegreeSatisfactionReq(ctx *parser.DegreeSatisfactionReqContext) any {
	return v.visitDegreeSatisfactionReq(ctx)
}

func (v *RequisiteVisitor) visitDegreeSatisfactionReq(ctx *parser.DegreeSatisfactionReqContext) *rules.DegreeSatisfactionRule {
	rule := v.Visit(ctx.Degree_satisfaction_rule()).(*rules.DegreeSatisfactionRule)
	v.appendRule(rule)
	return rule
}

// VisitCreditForReq
//
// Rule: credit_for_rule
func (v *RequisiteVisitor) VisitCreditForReq(ctx *parser.CreditForReqContext) any {
	return v.visitCreditForReq(ctx)
}

func (v *RequisiteVisitor) visitCreditForReq(ctx *parser.CreditForReqContext) *rules.CreditForRule {
	rule := v.Visit(ctx.Credit_for_rule()).(*rules.CreditForRule)
	v.appendRule(rule)
	return rule
}

// VisitLivingLearningReq
//
// Rule: living_learning_rule
func (v *RequisiteVisitor) VisitLivingLearningReq(ctx *parser.LivingLearningReqContext) any {
	return v.visitLivingLearningReq(ctx)
}

func (v *RequisiteVisitor) visitLivingLearningReq(ctx *parser.LivingLearningReqContext) *rules.LivingLearningRule {
	rule := v.Visit(ctx.Living_learning_rule()).(*rules.LivingLearningRule)
	v.appendRule(rule)
	return rule
}

// VisitSchoolReq
//
// Rule: school_rule
func (v *RequisiteVisitor) VisitSchoolReq(ctx *parser.SchoolReqContext) any {
	return v.visitSchoolReq(ctx)
}

func (v *RequisiteVisitor) visitSchoolReq(ctx *parser.SchoolReqContext) *rules.SchoolRule {
	rule := v.Visit(ctx.School_rule()).(*rules.SchoolRule)
	v.appendRule(rule)
	return rule
}

// VisitMajorReq
//
// Rule: major_condition
func (v *RequisiteVisitor) VisitMajorReq(ctx *parser.MajorReqContext) any {
	return v.visitMajorReq(ctx)
}

func (v *RequisiteVisitor) visitMajorReq(ctx *parser.MajorReqContext) conditions.Condition {
	cond := v.Visit(ctx.Major_condition()).(conditions.Condition)
	v.appendPreReq(cond)
	return cond
}

// VisitPrereqReq
//
// Rule: PREREQ_KW COLON? expr
func (v *RequisiteVisitor) VisitPrereqReq(ctx *parser.PrereqReqContext) any {
	return v.visitPrereqReq(ctx)
}

func (v *RequisiteVisitor) visitPrereqReq(ctx *parser.PrereqReqContext) conditions.Condition {
	if cond, ok := v.Visit(ctx.Expr()).(conditions.Condition); ok {
		v.appendPreReq(cond)
		return cond
	}
	return nil
}

// VisitCoreqReq
//
// Rule: COREQ_KW  COLON? expr
func (v *RequisiteVisitor) VisitCoreqReq(ctx *parser.CoreqReqContext) any {
	return v.visitCoreqReq(ctx)
}

func (v *RequisiteVisitor) visitCoreqReq(ctx *parser.CoreqReqContext) conditions.Condition {
	if cond, ok := v.Visit(ctx.Expr()).(conditions.Condition); ok {
		v.appendCoReq(cond)
		return cond
	}
	return nil
}

// VisitPrereqAndCoreqReq
//
// Rule: PREREQ_KW COLON? expr AND COREQ_KW COLON expr
func (v *RequisiteVisitor) VisitPrereqAndCoreqReq(ctx *parser.PrereqAndCoreqReqContext) any {
	return v.visitPrereqAndCoreqReq(ctx)
}

func (v *RequisiteVisitor) visitPrereqAndCoreqReq(ctx *parser.PrereqAndCoreqReqContext) any {
	if cond, ok := v.Visit(ctx.Expr(0)).(conditions.Condition); ok {
		v.appendPreReq(cond)
	}
	if cond, ok := v.Visit(ctx.Expr(1)).(conditions.Condition); ok {
		v.appendCoReq(cond)
	}
	return nil
}

// VisitPreOrCoReq
//
// Rule: PRE_OR_CO_KW COLON? expr
func (v *RequisiteVisitor) VisitPreOrCoReq(ctx *parser.PreOrCoReqContext) any {
	return v.visitPreOrCoReq(ctx)
}

func (v *RequisiteVisitor) visitPreOrCoReq(ctx *parser.PreOrCoReqContext) conditions.Condition {
	cond := v.Visit(ctx.Expr()).(conditions.Condition)
	v.appendPreOrCoReq(cond)
	return cond
}

// VisitSameAsReq
//
// Rule: same_as_rule
func (v *RequisiteVisitor) VisitSameAsReq(ctx *parser.SameAsReqContext) any {
	return v.visitSameAsReq(ctx)
}

func (v *RequisiteVisitor) visitSameAsReq(ctx *parser.SameAsReqContext) *rules.SameAsRule {
	rule := v.Visit(ctx.Same_as_rule()).(*rules.SameAsRule)
	v.appendRule(rule)
	return rule
}

// VisitExprReq
//
// Rule: expr
func (v *RequisiteVisitor) VisitExprReq(ctx *parser.ExprReqContext) any {
	return v.visitExprReq(ctx)
}

func (v *RequisiteVisitor) visitExprReq(ctx *parser.ExprReqContext) conditions.Condition {
	cond := v.Visit(ctx.Expr()).(conditions.Condition)
	v.appendPreReq(cond)
	return cond
}
