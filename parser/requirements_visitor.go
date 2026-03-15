// Code generated from ../Requirements.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Requirements
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by RequirementsParser.
type RequirementsVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RequirementsParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by RequirementsParser#sentence.
	VisitSentence(ctx *SentenceContext) interface{}

	// Visit a parse tree produced by RequirementsParser#sameAsParenReq.
	VisitSameAsParenReq(ctx *SameAsParenReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#preOrCoReq.
	VisitPreOrCoReq(ctx *PreOrCoReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#prereqAndCoreqReq.
	VisitPrereqAndCoreqReq(ctx *PrereqAndCoreqReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#computerScholarsReq.
	VisitComputerScholarsReq(ctx *ComputerScholarsReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#majorReq.
	VisitMajorReq(ctx *MajorReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#degreeSatisfactionReq.
	VisitDegreeSatisfactionReq(ctx *DegreeSatisfactionReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#repeatLimitTimesReq.
	VisitRepeatLimitTimesReq(ctx *RepeatLimitTimesReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#academicPlanReq.
	VisitAcademicPlanReq(ctx *AcademicPlanReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#repeatLimitHoursReq.
	VisitRepeatLimitHoursReq(ctx *RepeatLimitHoursReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#appendAcademicPlanReq.
	VisitAppendAcademicPlanReq(ctx *AppendAcademicPlanReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#repeatReq.
	VisitRepeatReq(ctx *RepeatReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#exprReq.
	VisitExprReq(ctx *ExprReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#exactCoreqNoticeReq.
	VisitExactCoreqNoticeReq(ctx *ExactCoreqNoticeReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#gpaRepeatReq.
	VisitGpaRepeatReq(ctx *GpaRepeatReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#schoolReq.
	VisitSchoolReq(ctx *SchoolReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#coreqReq.
	VisitCoreqReq(ctx *CoreqReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#prereqReq.
	VisitPrereqReq(ctx *PrereqReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#creditForReq.
	VisitCreditForReq(ctx *CreditForReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#sameAsReq.
	VisitSameAsReq(ctx *SameAsReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#excludeNoticeReq.
	VisitExcludeNoticeReq(ctx *ExcludeNoticeReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#livingLearningReq.
	VisitLivingLearningReq(ctx *LivingLearningReqContext) interface{}

	// Visit a parse tree produced by RequirementsParser#groupExpr.
	VisitGroupExpr(ctx *GroupExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#courseExpr.
	VisitCourseExpr(ctx *CourseExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#alternativeExpr.
	VisitAlternativeExpr(ctx *AlternativeExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#upperDivisionStandingExpr.
	VisitUpperDivisionStandingExpr(ctx *UpperDivisionStandingExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#semesterCreditHoursExpr.
	VisitSemesterCreditHoursExpr(ctx *SemesterCreditHoursExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#researchExpr.
	VisitResearchExpr(ctx *ResearchExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#computerScholarsExpr.
	VisitComputerScholarsExpr(ctx *ComputerScholarsExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#completeNExpr.
	VisitCompleteNExpr(ctx *CompleteNExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#anyCoreExpr.
	VisitAnyCoreExpr(ctx *AnyCoreExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#honorsExpr.
	VisitHonorsExpr(ctx *HonorsExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#gradeExpr.
	VisitGradeExpr(ctx *GradeExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#departmentConsentExpr.
	VisitDepartmentConsentExpr(ctx *DepartmentConsentExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#parenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#aleksScoreExpr.
	VisitAleksScoreExpr(ctx *AleksScoreExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#ampersandExpr.
	VisitAmpersandExpr(ctx *AmpersandExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#livingLearningExpr.
	VisitLivingLearningExpr(ctx *LivingLearningExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#uteachConsentExpr.
	VisitUteachConsentExpr(ctx *UteachConsentExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#goodAcademicStandingExpr.
	VisitGoodAcademicStandingExpr(ctx *GoodAcademicStandingExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#coreExpr.
	VisitCoreExpr(ctx *CoreExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#gpaExpr.
	VisitGpaExpr(ctx *GpaExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#majorExpr.
	VisitMajorExpr(ctx *MajorExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#equivalentExpr.
	VisitEquivalentExpr(ctx *EquivalentExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#instructorConsentExpr.
	VisitInstructorConsentExpr(ctx *InstructorConsentExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#minimumHoursExpr.
	VisitMinimumHoursExpr(ctx *MinimumHoursExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#repeatLimitHoursExpr.
	VisitRepeatLimitHoursExpr(ctx *RepeatLimitHoursExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#repeatRuleExpr.
	VisitRepeatRuleExpr(ctx *RepeatRuleExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#apScoreExpr.
	VisitApScoreExpr(ctx *ApScoreExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#anyMajorCourseExpr.
	VisitAnyMajorCourseExpr(ctx *AnyMajorCourseExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#upperDivisionHoursExpr.
	VisitUpperDivisionHoursExpr(ctx *UpperDivisionHoursExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#upperDivisionClassesExpr.
	VisitUpperDivisionClassesExpr(ctx *UpperDivisionClassesExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#placementTestExpr.
	VisitPlacementTestExpr(ctx *PlacementTestExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#concurrentEnrollmentExpr.
	VisitConcurrentEnrollmentExpr(ctx *ConcurrentEnrollmentExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#gradeLevelStandingExpr.
	VisitGradeLevelStandingExpr(ctx *GradeLevelStandingExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#degreeExpr.
	VisitDegreeExpr(ctx *DegreeExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#graduateStandingExpr.
	VisitGraduateStandingExpr(ctx *GraduateStandingExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#exactSectionExpr.
	VisitExactSectionExpr(ctx *ExactSectionExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#simpleCourse.
	VisitSimpleCourse(ctx *SimpleCourseContext) interface{}

	// Visit a parse tree produced by RequirementsParser#parenCourse.
	VisitParenCourse(ctx *ParenCourseContext) interface{}

	// Visit a parse tree produced by RequirementsParser#crossListedCourse.
	VisitCrossListedCourse(ctx *CrossListedCourseContext) interface{}

	// Visit a parse tree produced by RequirementsParser#fullCourseList.
	VisitFullCourseList(ctx *FullCourseListContext) interface{}

	// Visit a parse tree produced by RequirementsParser#shorthandCourseList.
	VisitShorthandCourseList(ctx *ShorthandCourseListContext) interface{}

	// Visit a parse tree produced by RequirementsParser#title.
	VisitTitle(ctx *TitleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#simpleGradeCondition.
	VisitSimpleGradeCondition(ctx *SimpleGradeConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#gpaGradeCondition.
	VisitGpaGradeCondition(ctx *GpaGradeConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#parenGradeCondition.
	VisitParenGradeCondition(ctx *ParenGradeConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#gradeAtLeastCondition.
	VisitGradeAtLeastCondition(ctx *GradeAtLeastConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#grade_course_list.
	VisitGrade_course_list(ctx *Grade_course_listContext) interface{}

	// Visit a parse tree produced by RequirementsParser#courseAlternativeCondition.
	VisitCourseAlternativeCondition(ctx *CourseAlternativeConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#parenAlternativeCondition.
	VisitParenAlternativeCondition(ctx *ParenAlternativeConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#gradeAlternativeCondition.
	VisitGradeAlternativeCondition(ctx *GradeAlternativeConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#gradeLevelStandingCondition.
	VisitGradeLevelStandingCondition(ctx *GradeLevelStandingConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#gradeLevelMajorStandingCondition.
	VisitGradeLevelMajorStandingCondition(ctx *GradeLevelMajorStandingConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#minimumGradeLevelStandingCondition.
	VisitMinimumGradeLevelStandingCondition(ctx *MinimumGradeLevelStandingConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#atLeastGradeLevelStandingCondition.
	VisitAtLeastGradeLevelStandingCondition(ctx *AtLeastGradeLevelStandingConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#prefixGradeLevelStandingCondition.
	VisitPrefixGradeLevelStandingCondition(ctx *PrefixGradeLevelStandingConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#graduateStandingInCondition.
	VisitGraduateStandingInCondition(ctx *GraduateStandingInConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#graduateLevelStandingCondition.
	VisitGraduateLevelStandingCondition(ctx *GraduateLevelStandingConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#universityGpaCondition.
	VisitUniversityGpaCondition(ctx *UniversityGpaConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#minimumGpaCondition.
	VisitMinimumGpaCondition(ctx *MinimumGpaConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#gpaInCourseCondition.
	VisitGpaInCourseCondition(ctx *GpaInCourseConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#prefixMajorCondition.
	VisitPrefixMajorCondition(ctx *PrefixMajorConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#gradeLevelPrefixMajorCondition.
	VisitGradeLevelPrefixMajorCondition(ctx *GradeLevelPrefixMajorConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#degreeTypePrefixMajorCondition.
	VisitDegreeTypePrefixMajorCondition(ctx *DegreeTypePrefixMajorConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#namedMajorCondition.
	VisitNamedMajorCondition(ctx *NamedMajorConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#namedDegreeTypeMajorCondition.
	VisitNamedDegreeTypeMajorCondition(ctx *NamedDegreeTypeMajorConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#undergraduateDegreeCondition.
	VisitUndergraduateDegreeCondition(ctx *UndergraduateDegreeConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#bachelorsOrMastersCondition.
	VisitBachelorsOrMastersCondition(ctx *BachelorsOrMastersConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#degree_list.
	VisitDegree_list(ctx *Degree_listContext) interface{}

	// Visit a parse tree produced by RequirementsParser#degree.
	VisitDegree(ctx *DegreeContext) interface{}

	// Visit a parse tree produced by RequirementsParser#coreCondition.
	VisitCoreCondition(ctx *CoreConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#anyCoreSCHCondition.
	VisitAnyCoreSCHCondition(ctx *AnyCoreSCHConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#completionOfCoreCondition.
	VisitCompletionOfCoreCondition(ctx *CompletionOfCoreConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#semesterCreditHoursCondition.
	VisitSemesterCreditHoursCondition(ctx *SemesterCreditHoursConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#minimumHoursCondition.
	VisitMinimumHoursCondition(ctx *MinimumHoursConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#upperDivisionSCHCondition.
	VisitUpperDivisionSCHCondition(ctx *UpperDivisionSCHConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#upperDivisionCreditsCondition.
	VisitUpperDivisionCreditsCondition(ctx *UpperDivisionCreditsConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#upperDivisionClassesCondition.
	VisitUpperDivisionClassesCondition(ctx *UpperDivisionClassesConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#fourThousandLevelCondition.
	VisitFourThousandLevelCondition(ctx *FourThousandLevelConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#researchCondition.
	VisitResearchCondition(ctx *ResearchConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#completeNOfFollowingCondition.
	VisitCompleteNOfFollowingCondition(ctx *CompleteNOfFollowingConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#completeNFromFollowingCondition.
	VisitCompleteNFromFollowingCondition(ctx *CompleteNFromFollowingConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#completeHoursFromFollowingCondition.
	VisitCompleteHoursFromFollowingCondition(ctx *CompleteHoursFromFollowingConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#placementScoreComparisonCondition.
	VisitPlacementScoreComparisonCondition(ctx *PlacementScoreComparisonConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#placementScoreRangeCondition.
	VisitPlacementScoreRangeCondition(ctx *PlacementScoreRangeConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#placementScoreMinimumCondition.
	VisitPlacementScoreMinimumCondition(ctx *PlacementScoreMinimumConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#placement_test_name.
	VisitPlacement_test_name(ctx *Placement_test_nameContext) interface{}

	// Visit a parse tree produced by RequirementsParser#apScoreCondition.
	VisitApScoreCondition(ctx *ApScoreConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#aleksScoreCondition.
	VisitAleksScoreCondition(ctx *AleksScoreConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#bothHonorsCondition.
	VisitBothHonorsCondition(ctx *BothHonorsConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#singleHonorsCondition.
	VisitSingleHonorsCondition(ctx *SingleHonorsConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#groupCondition.
	VisitGroupCondition(ctx *GroupConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#group.
	VisitGroup(ctx *GroupContext) interface{}

	// Visit a parse tree produced by RequirementsParser#concurrentEnrollmentCondition.
	VisitConcurrentEnrollmentCondition(ctx *ConcurrentEnrollmentConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#exactSectionCondition.
	VisitExactSectionCondition(ctx *ExactSectionConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#workshopSectionCondition.
	VisitWorkshopSectionCondition(ctx *WorkshopSectionConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#anyPreviousMajorCourseCondition.
	VisitAnyPreviousMajorCourseCondition(ctx *AnyPreviousMajorCourseConditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#academic_plan_condition.
	VisitAcademic_plan_condition(ctx *Academic_plan_conditionContext) interface{}

	// Visit a parse tree produced by RequirementsParser#courseRepeatRule.
	VisitCourseRepeatRule(ctx *CourseRepeatRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#internshipRepeatRule.
	VisitInternshipRepeatRule(ctx *InternshipRepeatRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#bareRepeatRule.
	VisitBareRepeatRule(ctx *BareRepeatRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#repeatMaxHoursRule.
	VisitRepeatMaxHoursRule(ctx *RepeatMaxHoursRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#repeatHoursMaxSuffixRule.
	VisitRepeatHoursMaxSuffixRule(ctx *RepeatHoursMaxSuffixRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#courseRepeatMaxHoursRule.
	VisitCourseRepeatMaxHoursRule(ctx *CourseRepeatMaxHoursRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#combinedRepeatMaxHoursRule.
	VisitCombinedRepeatMaxHoursRule(ctx *CombinedRepeatMaxHoursRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#topicsVaryRepeatRule.
	VisitTopicsVaryRepeatRule(ctx *TopicsVaryRepeatRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#courseRepeatLimitRule.
	VisitCourseRepeatLimitRule(ctx *CourseRepeatLimitRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#repeatUpToTimesRule.
	VisitRepeatUpToTimesRule(ctx *RepeatUpToTimesRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#repeatMaxTimesRule.
	VisitRepeatMaxTimesRule(ctx *RepeatMaxTimesRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#andRepeateExpr.
	VisitAndRepeateExpr(ctx *AndRepeateExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#alternativeRepeateExpr.
	VisitAlternativeRepeateExpr(ctx *AlternativeRepeateExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#courseRepeateExpr.
	VisitCourseRepeateExpr(ctx *CourseRepeateExprContext) interface{}

	// Visit a parse tree produced by RequirementsParser#gpa_repeate_rule.
	VisitGpa_repeate_rule(ctx *Gpa_repeate_ruleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#schoolDegreeSatisfactionRule.
	VisitSchoolDegreeSatisfactionRule(ctx *SchoolDegreeSatisfactionRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#prefixDegreeSatisfactionRule.
	VisitPrefixDegreeSatisfactionRule(ctx *PrefixDegreeSatisfactionRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#ofMultiPrefixDegreeSatisfactionRule.
	VisitOfMultiPrefixDegreeSatisfactionRule(ctx *OfMultiPrefixDegreeSatisfactionRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#electivesDegreeSatisfactionRule.
	VisitElectivesDegreeSatisfactionRule(ctx *ElectivesDegreeSatisfactionRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#schoolsDegreeSatisfactionRule.
	VisitSchoolsDegreeSatisfactionRule(ctx *SchoolsDegreeSatisfactionRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#multiPrefixForDegreeSatisfactionRule.
	VisitMultiPrefixForDegreeSatisfactionRule(ctx *MultiPrefixForDegreeSatisfactionRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#studentDegreeSatisfactionRule.
	VisitStudentDegreeSatisfactionRule(ctx *StudentDegreeSatisfactionRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#mathDegreeSatisfactionRule.
	VisitMathDegreeSatisfactionRule(ctx *MathDegreeSatisfactionRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#namedDegreeSatisfactionRule.
	VisitNamedDegreeSatisfactionRule(ctx *NamedDegreeSatisfactionRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#multiPrefixDegreeSatisfactionRule.
	VisitMultiPrefixDegreeSatisfactionRule(ctx *MultiPrefixDegreeSatisfactionRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#credit_for_rule.
	VisitCredit_for_rule(ctx *Credit_for_ruleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#prefixLivingLearningRule.
	VisitPrefixLivingLearningRule(ctx *PrefixLivingLearningRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#namedLivingLearningRule.
	VisitNamedLivingLearningRule(ctx *NamedLivingLearningRuleContext) interface{}

	// Visit a parse tree produced by RequirementsParser#school_rule.
	VisitSchool_rule(ctx *School_ruleContext) interface{}
}
