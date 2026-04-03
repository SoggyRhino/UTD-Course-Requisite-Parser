package visitors

import (
	"parser/conditions"
	"parser/parser"
	"parser/rules"
	"parser/utils"
)

// VisitExcludeNoticeReq
//
// Rule: requisite AND EXCLUDE_DMHP_LLC_NOTICE
func (v *RequisiteVisitor) VisitExcludeNoticeReq(ctx *parser.ExcludeNoticeReqContext) any {
	return v.visitExcludeNoticeReq(ctx)
}

func (v *RequisiteVisitor) visitExcludeNoticeReq(ctx *parser.ExcludeNoticeReqContext) any {
	v.Requirements.Notices = append(v.Requirements.Notices, utils.ExcludeDMHPLLCNotice)
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

func (v *RequisiteVisitor) visitAcademicPlanReq(ctx *parser.AcademicPlanReqContext) any {
	condition := v.Visit(ctx.Academic_plan_condition()).(conditions.Condition)
	v.appendPreReq(condition)
	return nil
}

// VisitExactCoreqNoticeReq
//
// Rule: EXACT_COREQ_NOTICE
func (v *RequisiteVisitor) VisitExactCoreqNoticeReq(ctx *parser.ExactCoreqNoticeReqContext) any {
	return v.visitExactCoreqNoticeReq(ctx)
}

func (v *RequisiteVisitor) visitExactCoreqNoticeReq(ctx *parser.ExactCoreqNoticeReqContext) any {
	v.Requirements.Notices = append(v.Requirements.Notices, utils.ExactCoReqNotice)
	return nil
}

// VisitComputerScholarsReq
//
// Rule: COMPUTER_SCHOLARS_PROGRAM
func (v *RequisiteVisitor) VisitComputerScholarsReq(ctx *parser.ComputerScholarsReqContext) any {
	return v.visitComputerScholarsReq(ctx)
}

func (v *RequisiteVisitor) visitComputerScholarsReq(ctx *parser.ComputerScholarsReqContext) any {
	v.Requirements.Notices = append(v.Requirements.Notices, utils.ExcludeDMHPLLCNotice)
	return nil
}

// VisitGpaRepeatReq
//
// Rule: gpa_repeate_rule
func (v *RequisiteVisitor) VisitGpaRepeatReq(ctx *parser.GpaRepeatReqContext) any {
	return v.visitGpaRepeatReq(ctx)
}

func (v *RequisiteVisitor) visitGpaRepeatReq(ctx *parser.GpaRepeatReqContext) *rules.GpaRepeatRule {
	v.Requirements.Rules = append(v.Requirements.Rules, v.Visit(ctx.Gpa_repeate_rule()).(rules.Rule))
	return nil
}

// VisitRepeatLimitHoursReq
//
// Rule: repeat_limit_hours_rule
func (v *RequisiteVisitor) VisitRepeatLimitHoursReq(ctx *parser.RepeatLimitHoursReqContext) any {
	return v.visitRepeatLimitHoursReq(ctx)
}

func (v *RequisiteVisitor) visitRepeatLimitHoursReq(ctx *parser.RepeatLimitHoursReqContext) *rules.RepeatRule {
	v.Requirements.Rules = append(v.Requirements.Rules, v.Visit(ctx.Repeat_limit_hours_rule()).(rules.Rule))
	return nil
}

// VisitRepeatLimitTimesReq
//
// Rule: repeat_limit_times_rule
func (v *RequisiteVisitor) VisitRepeatLimitTimesReq(ctx *parser.RepeatLimitTimesReqContext) any {
	return v.visitRepeatLimitTimesReq(ctx)
}

func (v *RequisiteVisitor) visitRepeatLimitTimesReq(ctx *parser.RepeatLimitTimesReqContext) *rules.RepeatRule {
	v.Requirements.Rules = append(v.Requirements.Rules, v.Visit(ctx.Repeat_limit_times_rule()).(rules.Rule))
	return nil
}

// VisitRepeatReq
//
// Rule: repeat_rule
func (v *RequisiteVisitor) VisitRepeatReq(ctx *parser.RepeatReqContext) any {
	return v.visitRepeatReq(ctx)
}

func (v *RequisiteVisitor) visitRepeatReq(ctx *parser.RepeatReqContext) *rules.RepeatRule {
	v.Requirements.Rules = append(v.Requirements.Rules, v.Visit(ctx.Repeat_rule()).(rules.Rule))
	return nil
}

// VisitDegreeSatisfactionReq
//
// Rule: degree_satisfaction_rule
func (v *RequisiteVisitor) VisitDegreeSatisfactionReq(ctx *parser.DegreeSatisfactionReqContext) any {
	return v.visitDegreeSatisfactionReq(ctx)
}

func (v *RequisiteVisitor) visitDegreeSatisfactionReq(ctx *parser.DegreeSatisfactionReqContext) *rules.DegreeSatisfactionRule {
	v.Requirements.Rules = append(v.Requirements.Rules, v.Visit(ctx.Degree_satisfaction_rule()).(rules.Rule))
	return nil
}

// VisitCreditForReq
//
// Rule: credit_for_rule
func (v *RequisiteVisitor) VisitCreditForReq(ctx *parser.CreditForReqContext) any {
	return v.visitCreditForReq(ctx)
}

func (v *RequisiteVisitor) visitCreditForReq(ctx *parser.CreditForReqContext) *rules.CreditForRule {
	v.Requirements.Rules = append(v.Requirements.Rules, v.Visit(ctx.Credit_for_rule()).(rules.Rule))
	return nil
}

// VisitLivingLearningReq
//
// Rule: living_learning_rule
func (v *RequisiteVisitor) VisitLivingLearningReq(ctx *parser.LivingLearningReqContext) any {
	return v.visitLivingLearningReq(ctx)
}

func (v *RequisiteVisitor) visitLivingLearningReq(ctx *parser.LivingLearningReqContext) *rules.LivingLearningRule {
	v.Requirements.Rules = append(v.Requirements.Rules, v.Visit(ctx.Living_learning_rule()).(rules.Rule))
	return nil
}

// VisitSchoolReq
//
// Rule: school_rule
func (v *RequisiteVisitor) VisitSchoolReq(ctx *parser.SchoolReqContext) any {
	return v.visitSchoolReq(ctx)
}

func (v *RequisiteVisitor) visitSchoolReq(ctx *parser.SchoolReqContext) *rules.SchoolRule {
	v.Requirements.Rules = append(v.Requirements.Rules, v.Visit(ctx.School_rule()).(rules.Rule))
	return nil
}

// VisitMajorReq
//
// Rule: major_condition
func (v *RequisiteVisitor) VisitMajorReq(ctx *parser.MajorReqContext) any {
	return v.visitMajorReq(ctx)
}

func (v *RequisiteVisitor) visitMajorReq(ctx *parser.MajorReqContext) conditions.Condition {
	v.appendPreReq(v.Visit(ctx.Major_condition()).(conditions.Condition))
	return nil
}

// VisitPrereqReq
//
// Rule: PREREQ_KW COLON? expr
func (v *RequisiteVisitor) VisitPrereqReq(ctx *parser.PrereqReqContext) any {
	return v.visitPrereqReq(ctx)
}

func (v *RequisiteVisitor) visitPrereqReq(ctx *parser.PrereqReqContext) any {
	condition := v.Visit(ctx.Expr())
	switch res := condition.(type) {
	case conditions.Condition:
		v.appendPreReq(res)
	case rules.Rule:
		v.Requirements.Rules = append(v.Requirements.Rules, res)
	}
	return nil
}

// VisitCoreqReq
//
// Rule: COREQ_KW  COLON? expr
func (v *RequisiteVisitor) VisitCoreqReq(ctx *parser.CoreqReqContext) any {
	return v.visitCoreqReq(ctx)
}

func (v *RequisiteVisitor) visitCoreqReq(ctx *parser.CoreqReqContext) any {
	condition := v.Visit(ctx.Expr())
	switch res := condition.(type) {
	case conditions.Condition:
		v.Requirements.CoReqs = res
	case rules.Rule:
		v.Requirements.Rules = append(v.Requirements.Rules, res)
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
	PreCond := v.Visit(ctx.Expr(0))
	CoCond := v.Visit(ctx.Expr(1))

	switch res := PreCond.(type) {
	case conditions.Condition:
		v.appendPreReq(res)
	case rules.Rule:
		v.Requirements.Rules = append(v.Requirements.Rules, res)
	}

	switch res := CoCond.(type) {
	case conditions.Condition:
		v.Requirements.CoReqs = res
	case rules.Rule:
		v.Requirements.Rules = append(v.Requirements.Rules, res)
	}

	return nil
}

// VisitPreOrCoReq
//
// Rule: PRE_OR_CO_KW COLON? expr
func (v *RequisiteVisitor) VisitPreOrCoReq(ctx *parser.PreOrCoReqContext) any {
	return v.visitPreOrCoReq(ctx)
}

func (v *RequisiteVisitor) visitPreOrCoReq(ctx *parser.PreOrCoReqContext) any {
	condition := v.Visit(ctx.Expr())

	switch res := condition.(type) {
	case conditions.Condition:
		v.Requirements.PreOrCoReqs = res
	case rules.Rule:
		v.Requirements.Rules = append(v.Requirements.Rules, res)
	}
	return nil
}

// VisitSameAsReq
//
// Rule: same_as_rule
func (v *RequisiteVisitor) VisitSameAsReq(ctx *parser.SameAsReqContext) any {
	return v.visitSameAsReq(ctx)
}

func (v *RequisiteVisitor) visitSameAsReq(ctx *parser.SameAsReqContext) any {
	v.Requirements.Rules = append(v.Requirements.Rules, v.Visit(ctx.Same_as_rule()).(rules.Rule))
	return nil
}

// VisitExprReq
//
// Rule: expr
func (v *RequisiteVisitor) VisitExprReq(ctx *parser.ExprReqContext) any {
	return v.visitExprReq(ctx)
}

func (v *RequisiteVisitor) visitExprReq(ctx *parser.ExprReqContext) any {
	condition := v.Visit(ctx.Expr())

	switch res := condition.(type) {
	case conditions.Condition:
		v.appendPreReq(res)
	case rules.Rule:
		v.Requirements.Rules = append(v.Requirements.Rules, res)
	}
	return nil
}
