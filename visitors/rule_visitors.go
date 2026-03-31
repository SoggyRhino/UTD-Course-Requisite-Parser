package visitors

import (
	"parser/conditions"
	"parser/parser"
	rules "parser/rules"
	"parser/utils"
)

// ======================= Repeat Restriction Rule =======================

// VisitCourseRepeatRule
//
// Rule: course REPEAT_RESTRICTION
func (v *RequisiteVisitor) VisitCourseRepeatRule(ctx *parser.CourseRepeatRuleContext) any {
	return v.visitCourseRepeatRule(ctx)
}

func (v *RequisiteVisitor) visitCourseRepeatRule(ctx *parser.CourseRepeatRuleContext) *rules.RepeatRule {
	courses := extractCoursesFromCourseList(v.Visit(ctx.Course()).(conditions.Condition))
	return rules.NewCourseRepeatRule(courses)
}

// VisitInternshipRepeatRule
//
// Rule: PREFIX 'Internship' REPEAT_RESTRICTION
func (v *RequisiteVisitor) VisitInternshipRepeatRule(ctx *parser.InternshipRepeatRuleContext) any {
	return v.visitInternshipRepeatRule(ctx)
}

func (v *RequisiteVisitor) visitInternshipRepeatRule(ctx *parser.InternshipRepeatRuleContext) *rules.RepeatRule {
	return rules.NewInternshipRepeatRule(ctx.PREFIX().GetText())
}

// VisitBareRepeatRule
//
// Rule: REPEAT_RESTRICTION
func (v *RequisiteVisitor) VisitBareRepeatRule(ctx *parser.BareRepeatRuleContext) any {
	return v.visitBareRepeatRule(ctx)
}

func (v *RequisiteVisitor) visitBareRepeatRule(ctx *parser.BareRepeatRuleContext) *rules.RepeatRule {
	//todo better blank value
	return rules.NewRepeatRule(0, 0, []utils.Course{}, "")
}

// VisitRepeatMaxHoursRule
//
// Rule: REPEAT_LIMIT DASH course ONLY_KW? 'may' ONLY_KW? 'be repeated' ONLY_KW? 'for'? 'a maximum of' SMALL_INT SEMESTER_CREDIT_HOURS
func (v *RequisiteVisitor) VisitRepeatMaxHoursRule(ctx *parser.RepeatMaxHoursRuleContext) any {
	return v.visitRepeatMaxHoursRule(ctx)
}

func (v *RequisiteVisitor) visitRepeatMaxHoursRule(ctx *parser.RepeatMaxHoursRuleContext) *rules.RepeatRule {
	hours := mapInt(ctx.SMALL_INT().GetText())
	courses := extractCoursesFromCourseList(v.Visit(ctx.Course()).(conditions.Condition))
	return rules.NewRepeatRule(0, hours, courses, "")
}

// VisitRepeatHoursMaxSuffixRule
//
// Rule: REPEAT_LIMIT DASH course ONLY_KW? 'may' ONLY_KW? 'be repeated' ONLY_KW? 'for'? SMALL_INT SEMESTER_CREDIT_HOURS 'maximum'
func (v *RequisiteVisitor) VisitRepeatHoursMaxSuffixRule(ctx *parser.RepeatHoursMaxSuffixRuleContext) any {
	return v.visitRepeatHoursMaxSuffixRule(ctx)
}

func (v *RequisiteVisitor) visitRepeatHoursMaxSuffixRule(ctx *parser.RepeatHoursMaxSuffixRuleContext) *rules.RepeatRule {
	hours := mapInt(ctx.SMALL_INT().GetText())
	courses := extractCoursesFromCourseList(v.Visit(ctx.Course()).(conditions.Condition))
	return rules.NewRepeatRule(0, hours, courses, "")
}

// VisitCourseRepeatMaxHoursRule
//
// Rule: course REPEAT_LIMIT DASH 'This' COURSE_KW ('may' | 'can') ONLY_KW? 'be repeated' ONLY_KW? 'for'? 'a maximum of' SMALL_INT SEMESTER_CREDIT_HOURS
func (v *RequisiteVisitor) VisitCourseRepeatMaxHoursRule(ctx *parser.CourseRepeatMaxHoursRuleContext) any {
	return v.visitCourseRepeatMaxHoursRule(ctx)
}

func (v *RequisiteVisitor) visitCourseRepeatMaxHoursRule(ctx *parser.CourseRepeatMaxHoursRuleContext) *rules.RepeatRule {
	hours := mapInt(ctx.SMALL_INT().GetText())
	courses := extractCoursesFromCourseList(v.Visit(ctx.Course()).(conditions.Condition))
	return rules.NewRepeatRule(0, hours, courses, "")
}

// VisitCombinedRepeatMaxHoursRule
//
// Rule: REPEAT_LIMIT DASH course AND course 'combined may only be repeated for a maximum of' SMALL_INT SEMESTER_CREDIT_HOURS
func (v *RequisiteVisitor) VisitCombinedRepeatMaxHoursRule(ctx *parser.CombinedRepeatMaxHoursRuleContext) any {
	return v.visitCombinedRepeatMaxHoursRule(ctx)
}

func (v *RequisiteVisitor) visitCombinedRepeatMaxHoursRule(ctx *parser.CombinedRepeatMaxHoursRuleContext) *rules.RepeatRule {
	hours := mapInt(ctx.SMALL_INT().GetText())

	var courses []utils.Course
	courses = append(courses, extractCoursesFromCourseList(v.Visit(ctx.Course(0)).(conditions.Condition))...)
	courses = append(courses, extractCoursesFromCourseList(v.Visit(ctx.Course(1)).(conditions.Condition))...)

	return rules.NewRepeatRule(0, hours, courses, "")
}

// VisitTopicsVaryRepeatRule
//
// Rule: course REPEAT_LIMIT DASH 'May be repeated for credit as topics vary' SMALL_INT SEMESTER_CREDIT_HOURS 'maximum'
func (v *RequisiteVisitor) VisitTopicsVaryRepeatRule(ctx *parser.TopicsVaryRepeatRuleContext) any {
	return v.visitTopicsVaryRepeatRule(ctx)
}

func (v *RequisiteVisitor) visitTopicsVaryRepeatRule(ctx *parser.TopicsVaryRepeatRuleContext) *rules.RepeatRule {
	hours := mapInt(ctx.SMALL_INT().GetText())
	courses := extractCoursesFromCourseList(v.Visit(ctx.Course()).(conditions.Condition))
	return rules.NewRepeatRule(0, hours, courses, "")
}

// VisitCourseRepeatLimitRule
//
// Rule: course REPEAT_LIMIT
func (v *RequisiteVisitor) VisitCourseRepeatLimitRule(ctx *parser.CourseRepeatLimitRuleContext) any {
	return v.visitCourseRepeatLimitRule(ctx)
}

func (v *RequisiteVisitor) visitCourseRepeatLimitRule(ctx *parser.CourseRepeatLimitRuleContext) *rules.RepeatRule {
	courses := extractCoursesFromCourseList(v.Visit(ctx.Course()).(conditions.Condition))
	return rules.NewCourseRepeatRule(courses)
}

// VisitRepeatUpToTimesRule
//
// Rule: REPEAT_LIMIT DASH course 'may' ONLY_KW? 'be repeated up to' SMALL_INT 'times'
func (v *RequisiteVisitor) VisitRepeatUpToTimesRule(ctx *parser.RepeatUpToTimesRuleContext) any {
	return v.visitRepeatUpToTimesRule(ctx)
}

func (v *RequisiteVisitor) visitRepeatUpToTimesRule(ctx *parser.RepeatUpToTimesRuleContext) *rules.RepeatRule {
	count := mapInt(ctx.SMALL_INT().GetText())
	courses := extractCoursesFromCourseList(v.Visit(ctx.Course()).(conditions.Condition))
	return rules.NewRepeatRule(count, 0, courses, "")
}

// VisitRepeatMaxTimesRule
//
// Rule: REPEAT_LIMIT DASH course 'may' ONLY_KW? 'be repeated' 'a maximum of' SMALL_INT 'times'
func (v *RequisiteVisitor) VisitRepeatMaxTimesRule(ctx *parser.RepeatMaxTimesRuleContext) any {
	return v.visitRepeatMaxTimesRule(ctx)
}

func (v *RequisiteVisitor) visitRepeatMaxTimesRule(ctx *parser.RepeatMaxTimesRuleContext) *rules.RepeatRule {
	count := mapInt(ctx.SMALL_INT().GetText())
	courses := extractCoursesFromCourseList(v.Visit(ctx.Course()).(conditions.Condition))
	return rules.NewRepeatRule(count, 0, courses, "")
}

func (v *RequisiteVisitor) VisitGpaRepeatRule(ctx *parser.GpaRepeatRuleContext) any {
	return v.visitGpaRepeatRule(ctx)
}

func (v *RequisiteVisitor) visitGpaRepeatRule(ctx *parser.GpaRepeatRuleContext) *rules.GpaRepeatRule {
	courses := extractCoursesFromCourseList(v.Visit(ctx.Course()).(conditions.Condition))
	return rules.NewGpaRepeatRule(courses[0])
}
