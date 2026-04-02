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
	return rules.NewRepeatRule(1, 0, []utils.Course{}, "")
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

// VisitGpaRepeatRule
//
// Rule: 'GPA Repeat Restriction' DASH course
func (v *RequisiteVisitor) VisitGpaRepeatRule(ctx *parser.GpaRepeatRuleContext) any {
	return v.visitGpaRepeatRule(ctx)
}

func (v *RequisiteVisitor) visitGpaRepeatRule(ctx *parser.GpaRepeatRuleContext) *rules.GpaRepeatRule {
	courses := extractCoursesFromCourseList(v.Visit(ctx.Course()).(conditions.Condition))
	return rules.NewGpaRepeatRule(courses[0])
}

// ======================= Degree Satisfaction Rule ======================

// VisitPrefixDegreeSatisfactionRule
//
// Rule: MAY_NOT_BE_USED_TO_SATISFY DEGREE_LEVEL? PREFIX 'degree requirements'
func (v *RequisiteVisitor) VisitPrefixDegreeSatisfactionRule(ctx *parser.PrefixDegreeSatisfactionRuleContext) any {
	return v.visitPrefixDegreeSatisfactionRule(ctx)
}

func (v *RequisiteVisitor) visitPrefixDegreeSatisfactionRule(ctx *parser.PrefixDegreeSatisfactionRuleContext) *rules.DegreeSatisfactionRule {
	prefix := ctx.PREFIX().GetText()
	degreeLevel := mapDegreeLevel(v.getTextOrDefault(ctx.DEGREE_LEVEL(), ""))
	return rules.NewDegreeSatisfactionRuleFromPrefix([]string{prefix}, degreeLevel)
}

// VisitNamedDegreeSatisfactionRule
//
// Rule: MAY_NOT_BE_USED_TO_SATISFY 'degree requirements in' DEGREE_LEVEL? degree
func (v *RequisiteVisitor) VisitNamedDegreeSatisfactionRule(ctx *parser.NamedDegreeSatisfactionRuleContext) any {
	return v.visitNamedDegreeSatisfactionRule(ctx)
}

func (v *RequisiteVisitor) visitNamedDegreeSatisfactionRule(ctx *parser.NamedDegreeSatisfactionRuleContext) *rules.DegreeSatisfactionRule {
	degree := v.Visit(ctx.Degree()).(string)
	degreeLevel := mapDegreeLevel(v.getTextOrDefault(ctx.DEGREE_LEVEL(), ""))
	return rules.NewDegreeSatisfactionRuleFromDegree([]string{degree}, degreeLevel)
}

// VisitMultiPrefixForDegreeSatisfactionRule
//
// Rule: MAY_NOT_BE_USED_TO_SATISFY 'degree requirements for' (OR? THE? DEGREE_LEVEL? PREFIX)+ 'degree plans'
func (v *RequisiteVisitor) VisitMultiPrefixForDegreeSatisfactionRule(ctx *parser.MultiPrefixForDegreeSatisfactionRuleContext) any {
	return v.visitMultiPrefixForDegreeSatisfactionRule(ctx)
}

func (v *RequisiteVisitor) visitMultiPrefixForDegreeSatisfactionRule(ctx *parser.MultiPrefixForDegreeSatisfactionRuleContext) *rules.DegreeSatisfactionRule {
	prefixes := make([]string, 0, len(ctx.AllPREFIX()))
	for _, p := range ctx.AllPREFIX() {
		prefixes = append(prefixes, p.GetText())
	}
	degreeLevels := ctx.AllDEGREE_LEVEL()
	degreeLevel := mapDegreeLevel(v.getTextOrDefault(v.firstOrNil(degreeLevels), ""))
	return rules.NewDegreeSatisfactionRuleFromPrefix(prefixes, degreeLevel)

}

// VisitOfMultiPrefixDegreeSatisfactionRule
//
// Rule: MAY_NOT_BE_USED_TO_SATISFY 'the degree requirements of' (OR? THE? DEGREE_LEVEL? PREFIX)+ 'degree plans'
func (v *RequisiteVisitor) VisitOfMultiPrefixDegreeSatisfactionRule(ctx *parser.OfMultiPrefixDegreeSatisfactionRuleContext) any {
	return v.visitOfMultiPrefixDegreeSatisfactionRule(ctx)
}

func (v *RequisiteVisitor) visitOfMultiPrefixDegreeSatisfactionRule(ctx *parser.OfMultiPrefixDegreeSatisfactionRuleContext) *rules.DegreeSatisfactionRule {
	prefixes := make([]string, 0, len(ctx.AllPREFIX()))
	for _, p := range ctx.AllPREFIX() {
		prefixes = append(prefixes, p.GetText())
	}
	degreeLevels := ctx.AllDEGREE_LEVEL()
	degreeLevel := mapDegreeLevel(v.getTextOrDefault(v.firstOrNil(degreeLevels), ""))
	return rules.NewDegreeSatisfactionRuleFromPrefix(prefixes, degreeLevel)
}

// VisitSchoolDegreeSatisfactionRule
//
// Rule: MAY_NOT_BE_USED_TO_SATISFY 'degree requirements for' DEGREE_LEVEL? 'majors in'? 'the School of' degree+
func (v *RequisiteVisitor) VisitSchoolDegreeSatisfactionRule(ctx *parser.SchoolDegreeSatisfactionRuleContext) any {
	return v.visitSchoolDegreeSatisfactionRule(ctx)
}

func (v *RequisiteVisitor) visitSchoolDegreeSatisfactionRule(ctx *parser.SchoolDegreeSatisfactionRuleContext) *rules.DegreeSatisfactionRule {
	schools := make([]string, 0, len(ctx.AllDegree()))
	for _, d := range ctx.AllDegree() {
		schools = append(schools, v.Visit(d).(string))
	}
	degreeLevel := mapDegreeLevel(v.getTextOrDefault(ctx.DEGREE_LEVEL(), ""))
	return rules.NewDegreeSatisfactionRuleFromSchool(schools, degreeLevel)
}

// VisitSchoolsDegreeSatisfactionRule
//
// Rule: MAY_NOT_BE_USED_TO_SATISFY 'degree requirements for' DEGREE_LEVEL? 'majors in'? 'Schools of'? degree_list+
func (v *RequisiteVisitor) VisitSchoolsDegreeSatisfactionRule(ctx *parser.SchoolsDegreeSatisfactionRuleContext) any {
	return v.visitSchoolsDegreeSatisfactionRule(ctx)
}

func (v *RequisiteVisitor) visitSchoolsDegreeSatisfactionRule(ctx *parser.SchoolsDegreeSatisfactionRuleContext) *rules.DegreeSatisfactionRule {
	schools := make([]string, 0)
	for _, degrees := range ctx.AllDegree_list() {
		courses := v.Visit(degrees).([]string)
		schools = append(schools, courses...)
	}
	degreeLevel := mapDegreeLevel(v.getTextOrDefault(ctx.DEGREE_LEVEL(), ""))
	return rules.NewDegreeSatisfactionRuleFromSchool(schools, degreeLevel)
}

// VisitStudentDegreeSatisfactionRule
//
// Rule: MAY_NOT_BE_USED_TO_SATISFY 'degree requirements by students in' degree
func (v *RequisiteVisitor) VisitStudentDegreeSatisfactionRule(ctx *parser.StudentDegreeSatisfactionRuleContext) any {
	return v.visitStudentDegreeSatisfactionRule(ctx)
}

func (v *RequisiteVisitor) visitStudentDegreeSatisfactionRule(ctx *parser.StudentDegreeSatisfactionRuleContext) *rules.DegreeSatisfactionRule {
	degree := v.Visit(ctx.Degree()).(string)
	return rules.NewDegreeSatisfactionRuleFromDegree([]string{degree}, utils.AnyDegree)

}

// VisitMathDegreeSatisfactionRule
//
// Rule: MAY_NOT_BE_USED_TO_SATISFY 'mathematics requirements by students in Mathematics'
func (v *RequisiteVisitor) VisitMathDegreeSatisfactionRule(ctx *parser.MathDegreeSatisfactionRuleContext) any {
	return v.visitMathDegreeSatisfactionRule(ctx)
}

func (v *RequisiteVisitor) visitMathDegreeSatisfactionRule(ctx *parser.MathDegreeSatisfactionRuleContext) *rules.DegreeSatisfactionRule {
	return rules.NewMathDegreeSatisfactionRule()
}

// VisitElectivesDegreeSatisfactionRule
//
// Rule: degree_satisfaction_rule AND 'may not be used to satisfy electives'
func (v *RequisiteVisitor) VisitElectivesDegreeSatisfactionRule(ctx *parser.ElectivesDegreeSatisfactionRuleContext) any {
	return v.visitElectivesDegreeSatisfactionRule(ctx)
}

func (v *RequisiteVisitor) visitElectivesDegreeSatisfactionRule(ctx *parser.ElectivesDegreeSatisfactionRuleContext) *rules.DegreeSatisfactionRule {
	inner := v.Visit(ctx.Degree_satisfaction_rule()).(*rules.DegreeSatisfactionRule)
	return rules.NewDegreeSatisfactionRuleFromElectives(inner)
}

// ======================= Credit Rules ======================

// VisitCreditForRule
//
// Rule: CREDIT_PREFIX (COLON|COMMA)? credit_expr
func (v *RequisiteVisitor) VisitCreditForRule(ctx *parser.CreditForRuleContext) any {
	return v.visitCreditForRule(ctx)
}

func (v *RequisiteVisitor) visitCreditForRule(ctx *parser.CreditForRuleContext) *rules.CreditForRule {
	courseCollection := v.Visit(ctx.Credit_expr()).(rules.CourseCollection)
	return rules.NewCreditForRule(courseCollection)
}

// VisitParenCreditExpr
//
// Rule: '(' credit_expr ')'
func (v *RequisiteVisitor) VisitParenCreditExpr(ctx *parser.ParenCreditExprContext) any {
	return v.visitParenCreditExpr(ctx)
}

func (v *RequisiteVisitor) visitParenCreditExpr(ctx *parser.ParenCreditExprContext) rules.CourseCollection {
	return v.Visit(ctx.Credit_expr()).(rules.CourseCollection)
}

// VisitOrCreditExpr
//
// Rule: credit_expr COMMA? OR credit_expr
func (v *RequisiteVisitor) VisitOrCreditExpr(ctx *parser.OrCreditExprContext) any {
	return v.visitOrCreditExpr(ctx)
}

func (v *RequisiteVisitor) visitOrCreditExpr(ctx *parser.OrCreditExprContext) rules.CourseCollection {
	col1 := v.Visit(ctx.Credit_expr(0)).(rules.CourseCollection)
	col2 := v.Visit(ctx.Credit_expr(1)).(rules.CourseCollection)

	return rules.NewOrCourseCollection(col1, col2)
}

// VisitAndCreditExpr
//
// Rule: credit_expr COMMA? AND credit_expr
func (v *RequisiteVisitor) VisitAndCreditExpr(ctx *parser.AndCreditExprContext) any {
	return v.visitAndCreditExpr(ctx)
}

func (v *RequisiteVisitor) visitAndCreditExpr(ctx *parser.AndCreditExprContext) rules.CourseCollection {
	col1 := v.Visit(ctx.Credit_expr(0)).(rules.CourseCollection)
	col2 := v.Visit(ctx.Credit_expr(1)).(rules.CourseCollection)

	return rules.NewAndCourseCollection(col1, col2)
}

// VisitAmpersandCreditExpr
//
// Rule: credit_expr COMMA? AMPERSAND credit_expr
func (v *RequisiteVisitor) VisitAmpersandCreditExpr(ctx *parser.AmpersandCreditExprContext) any {
	return v.visitAmpersandCreditExpr(ctx)
}

func (v *RequisiteVisitor) visitAmpersandCreditExpr(ctx *parser.AmpersandCreditExprContext) rules.CourseCollection {
	col1 := v.Visit(ctx.Credit_expr(0)).(rules.CourseCollection)
	col2 := v.Visit(ctx.Credit_expr(1)).(rules.CourseCollection)

	return rules.NewAndCourseCollection(col1, col2)
}

// VisitCourseCreditExpr
//
// Rule: course
func (v *RequisiteVisitor) VisitCourseCreditExpr(ctx *parser.CourseCreditExprContext) any {
	return v.visitCourseCreditExpr(ctx)
}

func (v *RequisiteVisitor) visitCourseCreditExpr(ctx *parser.CourseCreditExprContext) rules.CourseCollection {
	courses := extractCoursesFromCourseList(v.Visit(ctx.Course()).(conditions.Condition))
	return rules.NewSimpleCourseCollection(courses)
}

// ======================= Living Learning Rule ======================

// VisitPrefixLivingLearningRule
//
// Rule: PREFIX ('&' PREFIX)* LIVING_LEARNING_COMMUNITY
func (v *RequisiteVisitor) VisitPrefixLivingLearningRule(ctx *parser.PrefixLivingLearningRuleContext) any {
	return v.visitPrefixLivingLearningRule(ctx)
}

func (v *RequisiteVisitor) visitPrefixLivingLearningRule(ctx *parser.PrefixLivingLearningRuleContext) *rules.LivingLearningRule {
	prefixes := make([]string, len(ctx.AllPREFIX()))
	for i, p := range ctx.AllPREFIX() {
		prefixes[i] = p.GetText()
	}
	return rules.NewLivingLearningRuleFromPrefixes(prefixes)
}

// VisitDegreeLivingLearningRule
//
// Rule: degree_list LIVING_LEARNING_COMMUNITY
func (v *RequisiteVisitor) VisitDegreeLivingLearningRule(ctx *parser.DegreeLivingLearningRuleContext) any {
	return v.visitDegreeLivingLearningRule(ctx)
}

func (v *RequisiteVisitor) visitDegreeLivingLearningRule(ctx *parser.DegreeLivingLearningRuleContext) *rules.LivingLearningRule {
	return rules.NewLivingLearningRuleFromDegrees(v.Visit(ctx.Degree_list()).([]string))
}

// ======================= School Rule ======================

// VisitSchoolRule
//
// Rule: 'Open to students in the School of' degree_list ONLY_KW? #schoolRule
func (v *RequisiteVisitor) VisitSchoolRule(ctx *parser.SchoolRuleContext) any {
	return v.visitSchoolRule(ctx)
}

func (v *RequisiteVisitor) visitSchoolRule(ctx *parser.SchoolRuleContext) *rules.SchoolRule {
	schools := v.Visit(ctx.Degree_list()).([]string)
	return rules.NewSchoolRule(schools)
}

// ======================= Same As Rule ======================

// VisitSameAsRule
//
// Rule: '(' SAME_AS course (AND course)* ')'
func (v *RequisiteVisitor) VisitSameAsRule(ctx *parser.SameAsRuleContext) any {
	return v.visitSameAsRule(ctx)
}

func (v *RequisiteVisitor) visitSameAsRule(ctx *parser.SameAsRuleContext) *rules.SameAsRule {
	courses := make([]utils.Course, 0, len(ctx.AllCourse()))
	for _, course := range ctx.AllCourse() {
		courses = append(courses, extractCoursesFromCourseList(v.Visit(course).(conditions.Condition))...)
	}

	return rules.NewSameAsRule(courses)
}
