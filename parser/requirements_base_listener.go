// Code generated from ../Requirements.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Requirements
import "github.com/antlr4-go/antlr/v4"

// BaseRequirementsListener is a complete listener for a parse tree produced by RequirementsParser.
type BaseRequirementsListener struct{}

var _ RequirementsListener = &BaseRequirementsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRequirementsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRequirementsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRequirementsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRequirementsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseRequirementsListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseRequirementsListener) ExitProg(ctx *ProgContext) {}

// EnterSentence is called when production sentence is entered.
func (s *BaseRequirementsListener) EnterSentence(ctx *SentenceContext) {}

// ExitSentence is called when production sentence is exited.
func (s *BaseRequirementsListener) ExitSentence(ctx *SentenceContext) {}

// EnterSameAsParenReq is called when production sameAsParenReq is entered.
func (s *BaseRequirementsListener) EnterSameAsParenReq(ctx *SameAsParenReqContext) {}

// ExitSameAsParenReq is called when production sameAsParenReq is exited.
func (s *BaseRequirementsListener) ExitSameAsParenReq(ctx *SameAsParenReqContext) {}

// EnterPreOrCoReq is called when production preOrCoReq is entered.
func (s *BaseRequirementsListener) EnterPreOrCoReq(ctx *PreOrCoReqContext) {}

// ExitPreOrCoReq is called when production preOrCoReq is exited.
func (s *BaseRequirementsListener) ExitPreOrCoReq(ctx *PreOrCoReqContext) {}

// EnterPrereqAndCoreqReq is called when production prereqAndCoreqReq is entered.
func (s *BaseRequirementsListener) EnterPrereqAndCoreqReq(ctx *PrereqAndCoreqReqContext) {}

// ExitPrereqAndCoreqReq is called when production prereqAndCoreqReq is exited.
func (s *BaseRequirementsListener) ExitPrereqAndCoreqReq(ctx *PrereqAndCoreqReqContext) {}

// EnterComputerScholarsReq is called when production computerScholarsReq is entered.
func (s *BaseRequirementsListener) EnterComputerScholarsReq(ctx *ComputerScholarsReqContext) {}

// ExitComputerScholarsReq is called when production computerScholarsReq is exited.
func (s *BaseRequirementsListener) ExitComputerScholarsReq(ctx *ComputerScholarsReqContext) {}

// EnterMajorReq is called when production majorReq is entered.
func (s *BaseRequirementsListener) EnterMajorReq(ctx *MajorReqContext) {}

// ExitMajorReq is called when production majorReq is exited.
func (s *BaseRequirementsListener) ExitMajorReq(ctx *MajorReqContext) {}

// EnterDegreeSatisfactionReq is called when production degreeSatisfactionReq is entered.
func (s *BaseRequirementsListener) EnterDegreeSatisfactionReq(ctx *DegreeSatisfactionReqContext) {}

// ExitDegreeSatisfactionReq is called when production degreeSatisfactionReq is exited.
func (s *BaseRequirementsListener) ExitDegreeSatisfactionReq(ctx *DegreeSatisfactionReqContext) {}

// EnterRepeatLimitTimesReq is called when production repeatLimitTimesReq is entered.
func (s *BaseRequirementsListener) EnterRepeatLimitTimesReq(ctx *RepeatLimitTimesReqContext) {}

// ExitRepeatLimitTimesReq is called when production repeatLimitTimesReq is exited.
func (s *BaseRequirementsListener) ExitRepeatLimitTimesReq(ctx *RepeatLimitTimesReqContext) {}

// EnterAcademicPlanReq is called when production academicPlanReq is entered.
func (s *BaseRequirementsListener) EnterAcademicPlanReq(ctx *AcademicPlanReqContext) {}

// ExitAcademicPlanReq is called when production academicPlanReq is exited.
func (s *BaseRequirementsListener) ExitAcademicPlanReq(ctx *AcademicPlanReqContext) {}

// EnterRepeatLimitHoursReq is called when production repeatLimitHoursReq is entered.
func (s *BaseRequirementsListener) EnterRepeatLimitHoursReq(ctx *RepeatLimitHoursReqContext) {}

// ExitRepeatLimitHoursReq is called when production repeatLimitHoursReq is exited.
func (s *BaseRequirementsListener) ExitRepeatLimitHoursReq(ctx *RepeatLimitHoursReqContext) {}

// EnterAppendAcademicPlanReq is called when production appendAcademicPlanReq is entered.
func (s *BaseRequirementsListener) EnterAppendAcademicPlanReq(ctx *AppendAcademicPlanReqContext) {}

// ExitAppendAcademicPlanReq is called when production appendAcademicPlanReq is exited.
func (s *BaseRequirementsListener) ExitAppendAcademicPlanReq(ctx *AppendAcademicPlanReqContext) {}

// EnterRepeatReq is called when production repeatReq is entered.
func (s *BaseRequirementsListener) EnterRepeatReq(ctx *RepeatReqContext) {}

// ExitRepeatReq is called when production repeatReq is exited.
func (s *BaseRequirementsListener) ExitRepeatReq(ctx *RepeatReqContext) {}

// EnterExprReq is called when production exprReq is entered.
func (s *BaseRequirementsListener) EnterExprReq(ctx *ExprReqContext) {}

// ExitExprReq is called when production exprReq is exited.
func (s *BaseRequirementsListener) ExitExprReq(ctx *ExprReqContext) {}

// EnterExactCoreqNoticeReq is called when production exactCoreqNoticeReq is entered.
func (s *BaseRequirementsListener) EnterExactCoreqNoticeReq(ctx *ExactCoreqNoticeReqContext) {}

// ExitExactCoreqNoticeReq is called when production exactCoreqNoticeReq is exited.
func (s *BaseRequirementsListener) ExitExactCoreqNoticeReq(ctx *ExactCoreqNoticeReqContext) {}

// EnterGpaRepeatReq is called when production gpaRepeatReq is entered.
func (s *BaseRequirementsListener) EnterGpaRepeatReq(ctx *GpaRepeatReqContext) {}

// ExitGpaRepeatReq is called when production gpaRepeatReq is exited.
func (s *BaseRequirementsListener) ExitGpaRepeatReq(ctx *GpaRepeatReqContext) {}

// EnterSchoolReq is called when production schoolReq is entered.
func (s *BaseRequirementsListener) EnterSchoolReq(ctx *SchoolReqContext) {}

// ExitSchoolReq is called when production schoolReq is exited.
func (s *BaseRequirementsListener) ExitSchoolReq(ctx *SchoolReqContext) {}

// EnterCoreqReq is called when production coreqReq is entered.
func (s *BaseRequirementsListener) EnterCoreqReq(ctx *CoreqReqContext) {}

// ExitCoreqReq is called when production coreqReq is exited.
func (s *BaseRequirementsListener) ExitCoreqReq(ctx *CoreqReqContext) {}

// EnterPrereqReq is called when production prereqReq is entered.
func (s *BaseRequirementsListener) EnterPrereqReq(ctx *PrereqReqContext) {}

// ExitPrereqReq is called when production prereqReq is exited.
func (s *BaseRequirementsListener) ExitPrereqReq(ctx *PrereqReqContext) {}

// EnterCreditForReq is called when production creditForReq is entered.
func (s *BaseRequirementsListener) EnterCreditForReq(ctx *CreditForReqContext) {}

// ExitCreditForReq is called when production creditForReq is exited.
func (s *BaseRequirementsListener) ExitCreditForReq(ctx *CreditForReqContext) {}

// EnterSameAsReq is called when production sameAsReq is entered.
func (s *BaseRequirementsListener) EnterSameAsReq(ctx *SameAsReqContext) {}

// ExitSameAsReq is called when production sameAsReq is exited.
func (s *BaseRequirementsListener) ExitSameAsReq(ctx *SameAsReqContext) {}

// EnterExcludeNoticeReq is called when production excludeNoticeReq is entered.
func (s *BaseRequirementsListener) EnterExcludeNoticeReq(ctx *ExcludeNoticeReqContext) {}

// ExitExcludeNoticeReq is called when production excludeNoticeReq is exited.
func (s *BaseRequirementsListener) ExitExcludeNoticeReq(ctx *ExcludeNoticeReqContext) {}

// EnterLivingLearningReq is called when production livingLearningReq is entered.
func (s *BaseRequirementsListener) EnterLivingLearningReq(ctx *LivingLearningReqContext) {}

// ExitLivingLearningReq is called when production livingLearningReq is exited.
func (s *BaseRequirementsListener) ExitLivingLearningReq(ctx *LivingLearningReqContext) {}

// EnterGroupExpr is called when production groupExpr is entered.
func (s *BaseRequirementsListener) EnterGroupExpr(ctx *GroupExprContext) {}

// ExitGroupExpr is called when production groupExpr is exited.
func (s *BaseRequirementsListener) ExitGroupExpr(ctx *GroupExprContext) {}

// EnterCourseExpr is called when production courseExpr is entered.
func (s *BaseRequirementsListener) EnterCourseExpr(ctx *CourseExprContext) {}

// ExitCourseExpr is called when production courseExpr is exited.
func (s *BaseRequirementsListener) ExitCourseExpr(ctx *CourseExprContext) {}

// EnterAlternativeExpr is called when production alternativeExpr is entered.
func (s *BaseRequirementsListener) EnterAlternativeExpr(ctx *AlternativeExprContext) {}

// ExitAlternativeExpr is called when production alternativeExpr is exited.
func (s *BaseRequirementsListener) ExitAlternativeExpr(ctx *AlternativeExprContext) {}

// EnterUpperDivisionStandingExpr is called when production upperDivisionStandingExpr is entered.
func (s *BaseRequirementsListener) EnterUpperDivisionStandingExpr(ctx *UpperDivisionStandingExprContext) {
}

// ExitUpperDivisionStandingExpr is called when production upperDivisionStandingExpr is exited.
func (s *BaseRequirementsListener) ExitUpperDivisionStandingExpr(ctx *UpperDivisionStandingExprContext) {
}

// EnterSemesterCreditHoursExpr is called when production semesterCreditHoursExpr is entered.
func (s *BaseRequirementsListener) EnterSemesterCreditHoursExpr(ctx *SemesterCreditHoursExprContext) {
}

// ExitSemesterCreditHoursExpr is called when production semesterCreditHoursExpr is exited.
func (s *BaseRequirementsListener) ExitSemesterCreditHoursExpr(ctx *SemesterCreditHoursExprContext) {}

// EnterResearchExpr is called when production researchExpr is entered.
func (s *BaseRequirementsListener) EnterResearchExpr(ctx *ResearchExprContext) {}

// ExitResearchExpr is called when production researchExpr is exited.
func (s *BaseRequirementsListener) ExitResearchExpr(ctx *ResearchExprContext) {}

// EnterComputerScholarsExpr is called when production computerScholarsExpr is entered.
func (s *BaseRequirementsListener) EnterComputerScholarsExpr(ctx *ComputerScholarsExprContext) {}

// ExitComputerScholarsExpr is called when production computerScholarsExpr is exited.
func (s *BaseRequirementsListener) ExitComputerScholarsExpr(ctx *ComputerScholarsExprContext) {}

// EnterCompleteNExpr is called when production completeNExpr is entered.
func (s *BaseRequirementsListener) EnterCompleteNExpr(ctx *CompleteNExprContext) {}

// ExitCompleteNExpr is called when production completeNExpr is exited.
func (s *BaseRequirementsListener) ExitCompleteNExpr(ctx *CompleteNExprContext) {}

// EnterAnyCoreExpr is called when production anyCoreExpr is entered.
func (s *BaseRequirementsListener) EnterAnyCoreExpr(ctx *AnyCoreExprContext) {}

// ExitAnyCoreExpr is called when production anyCoreExpr is exited.
func (s *BaseRequirementsListener) ExitAnyCoreExpr(ctx *AnyCoreExprContext) {}

// EnterHonorsExpr is called when production honorsExpr is entered.
func (s *BaseRequirementsListener) EnterHonorsExpr(ctx *HonorsExprContext) {}

// ExitHonorsExpr is called when production honorsExpr is exited.
func (s *BaseRequirementsListener) ExitHonorsExpr(ctx *HonorsExprContext) {}

// EnterGradeExpr is called when production gradeExpr is entered.
func (s *BaseRequirementsListener) EnterGradeExpr(ctx *GradeExprContext) {}

// ExitGradeExpr is called when production gradeExpr is exited.
func (s *BaseRequirementsListener) ExitGradeExpr(ctx *GradeExprContext) {}

// EnterDepartmentConsentExpr is called when production departmentConsentExpr is entered.
func (s *BaseRequirementsListener) EnterDepartmentConsentExpr(ctx *DepartmentConsentExprContext) {}

// ExitDepartmentConsentExpr is called when production departmentConsentExpr is exited.
func (s *BaseRequirementsListener) ExitDepartmentConsentExpr(ctx *DepartmentConsentExprContext) {}

// EnterParenExpr is called when production parenExpr is entered.
func (s *BaseRequirementsListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production parenExpr is exited.
func (s *BaseRequirementsListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterAleksScoreExpr is called when production aleksScoreExpr is entered.
func (s *BaseRequirementsListener) EnterAleksScoreExpr(ctx *AleksScoreExprContext) {}

// ExitAleksScoreExpr is called when production aleksScoreExpr is exited.
func (s *BaseRequirementsListener) ExitAleksScoreExpr(ctx *AleksScoreExprContext) {}

// EnterAmpersandExpr is called when production ampersandExpr is entered.
func (s *BaseRequirementsListener) EnterAmpersandExpr(ctx *AmpersandExprContext) {}

// ExitAmpersandExpr is called when production ampersandExpr is exited.
func (s *BaseRequirementsListener) ExitAmpersandExpr(ctx *AmpersandExprContext) {}

// EnterLivingLearningExpr is called when production livingLearningExpr is entered.
func (s *BaseRequirementsListener) EnterLivingLearningExpr(ctx *LivingLearningExprContext) {}

// ExitLivingLearningExpr is called when production livingLearningExpr is exited.
func (s *BaseRequirementsListener) ExitLivingLearningExpr(ctx *LivingLearningExprContext) {}

// EnterUteachConsentExpr is called when production uteachConsentExpr is entered.
func (s *BaseRequirementsListener) EnterUteachConsentExpr(ctx *UteachConsentExprContext) {}

// ExitUteachConsentExpr is called when production uteachConsentExpr is exited.
func (s *BaseRequirementsListener) ExitUteachConsentExpr(ctx *UteachConsentExprContext) {}

// EnterGoodAcademicStandingExpr is called when production goodAcademicStandingExpr is entered.
func (s *BaseRequirementsListener) EnterGoodAcademicStandingExpr(ctx *GoodAcademicStandingExprContext) {
}

// ExitGoodAcademicStandingExpr is called when production goodAcademicStandingExpr is exited.
func (s *BaseRequirementsListener) ExitGoodAcademicStandingExpr(ctx *GoodAcademicStandingExprContext) {
}

// EnterCoreExpr is called when production coreExpr is entered.
func (s *BaseRequirementsListener) EnterCoreExpr(ctx *CoreExprContext) {}

// ExitCoreExpr is called when production coreExpr is exited.
func (s *BaseRequirementsListener) ExitCoreExpr(ctx *CoreExprContext) {}

// EnterGpaExpr is called when production gpaExpr is entered.
func (s *BaseRequirementsListener) EnterGpaExpr(ctx *GpaExprContext) {}

// ExitGpaExpr is called when production gpaExpr is exited.
func (s *BaseRequirementsListener) ExitGpaExpr(ctx *GpaExprContext) {}

// EnterMajorExpr is called when production majorExpr is entered.
func (s *BaseRequirementsListener) EnterMajorExpr(ctx *MajorExprContext) {}

// ExitMajorExpr is called when production majorExpr is exited.
func (s *BaseRequirementsListener) ExitMajorExpr(ctx *MajorExprContext) {}

// EnterEquivalentExpr is called when production equivalentExpr is entered.
func (s *BaseRequirementsListener) EnterEquivalentExpr(ctx *EquivalentExprContext) {}

// ExitEquivalentExpr is called when production equivalentExpr is exited.
func (s *BaseRequirementsListener) ExitEquivalentExpr(ctx *EquivalentExprContext) {}

// EnterInstructorConsentExpr is called when production instructorConsentExpr is entered.
func (s *BaseRequirementsListener) EnterInstructorConsentExpr(ctx *InstructorConsentExprContext) {}

// ExitInstructorConsentExpr is called when production instructorConsentExpr is exited.
func (s *BaseRequirementsListener) ExitInstructorConsentExpr(ctx *InstructorConsentExprContext) {}

// EnterMinimumHoursExpr is called when production minimumHoursExpr is entered.
func (s *BaseRequirementsListener) EnterMinimumHoursExpr(ctx *MinimumHoursExprContext) {}

// ExitMinimumHoursExpr is called when production minimumHoursExpr is exited.
func (s *BaseRequirementsListener) ExitMinimumHoursExpr(ctx *MinimumHoursExprContext) {}

// EnterRepeatLimitHoursExpr is called when production repeatLimitHoursExpr is entered.
func (s *BaseRequirementsListener) EnterRepeatLimitHoursExpr(ctx *RepeatLimitHoursExprContext) {}

// ExitRepeatLimitHoursExpr is called when production repeatLimitHoursExpr is exited.
func (s *BaseRequirementsListener) ExitRepeatLimitHoursExpr(ctx *RepeatLimitHoursExprContext) {}

// EnterRepeatRuleExpr is called when production repeatRuleExpr is entered.
func (s *BaseRequirementsListener) EnterRepeatRuleExpr(ctx *RepeatRuleExprContext) {}

// ExitRepeatRuleExpr is called when production repeatRuleExpr is exited.
func (s *BaseRequirementsListener) ExitRepeatRuleExpr(ctx *RepeatRuleExprContext) {}

// EnterApScoreExpr is called when production apScoreExpr is entered.
func (s *BaseRequirementsListener) EnterApScoreExpr(ctx *ApScoreExprContext) {}

// ExitApScoreExpr is called when production apScoreExpr is exited.
func (s *BaseRequirementsListener) ExitApScoreExpr(ctx *ApScoreExprContext) {}

// EnterAnyMajorCourseExpr is called when production anyMajorCourseExpr is entered.
func (s *BaseRequirementsListener) EnterAnyMajorCourseExpr(ctx *AnyMajorCourseExprContext) {}

// ExitAnyMajorCourseExpr is called when production anyMajorCourseExpr is exited.
func (s *BaseRequirementsListener) ExitAnyMajorCourseExpr(ctx *AnyMajorCourseExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseRequirementsListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseRequirementsListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterUpperDivisionHoursExpr is called when production upperDivisionHoursExpr is entered.
func (s *BaseRequirementsListener) EnterUpperDivisionHoursExpr(ctx *UpperDivisionHoursExprContext) {}

// ExitUpperDivisionHoursExpr is called when production upperDivisionHoursExpr is exited.
func (s *BaseRequirementsListener) ExitUpperDivisionHoursExpr(ctx *UpperDivisionHoursExprContext) {}

// EnterUpperDivisionClassesExpr is called when production upperDivisionClassesExpr is entered.
func (s *BaseRequirementsListener) EnterUpperDivisionClassesExpr(ctx *UpperDivisionClassesExprContext) {
}

// ExitUpperDivisionClassesExpr is called when production upperDivisionClassesExpr is exited.
func (s *BaseRequirementsListener) ExitUpperDivisionClassesExpr(ctx *UpperDivisionClassesExprContext) {
}

// EnterPlacementTestExpr is called when production placementTestExpr is entered.
func (s *BaseRequirementsListener) EnterPlacementTestExpr(ctx *PlacementTestExprContext) {}

// ExitPlacementTestExpr is called when production placementTestExpr is exited.
func (s *BaseRequirementsListener) ExitPlacementTestExpr(ctx *PlacementTestExprContext) {}

// EnterConcurrentEnrollmentExpr is called when production concurrentEnrollmentExpr is entered.
func (s *BaseRequirementsListener) EnterConcurrentEnrollmentExpr(ctx *ConcurrentEnrollmentExprContext) {
}

// ExitConcurrentEnrollmentExpr is called when production concurrentEnrollmentExpr is exited.
func (s *BaseRequirementsListener) ExitConcurrentEnrollmentExpr(ctx *ConcurrentEnrollmentExprContext) {
}

// EnterGradeLevelStandingExpr is called when production gradeLevelStandingExpr is entered.
func (s *BaseRequirementsListener) EnterGradeLevelStandingExpr(ctx *GradeLevelStandingExprContext) {}

// ExitGradeLevelStandingExpr is called when production gradeLevelStandingExpr is exited.
func (s *BaseRequirementsListener) ExitGradeLevelStandingExpr(ctx *GradeLevelStandingExprContext) {}

// EnterDegreeExpr is called when production degreeExpr is entered.
func (s *BaseRequirementsListener) EnterDegreeExpr(ctx *DegreeExprContext) {}

// ExitDegreeExpr is called when production degreeExpr is exited.
func (s *BaseRequirementsListener) ExitDegreeExpr(ctx *DegreeExprContext) {}

// EnterGraduateStandingExpr is called when production graduateStandingExpr is entered.
func (s *BaseRequirementsListener) EnterGraduateStandingExpr(ctx *GraduateStandingExprContext) {}

// ExitGraduateStandingExpr is called when production graduateStandingExpr is exited.
func (s *BaseRequirementsListener) ExitGraduateStandingExpr(ctx *GraduateStandingExprContext) {}

// EnterExactSectionExpr is called when production exactSectionExpr is entered.
func (s *BaseRequirementsListener) EnterExactSectionExpr(ctx *ExactSectionExprContext) {}

// ExitExactSectionExpr is called when production exactSectionExpr is exited.
func (s *BaseRequirementsListener) ExitExactSectionExpr(ctx *ExactSectionExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseRequirementsListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseRequirementsListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterSimpleCourse is called when production simpleCourse is entered.
func (s *BaseRequirementsListener) EnterSimpleCourse(ctx *SimpleCourseContext) {}

// ExitSimpleCourse is called when production simpleCourse is exited.
func (s *BaseRequirementsListener) ExitSimpleCourse(ctx *SimpleCourseContext) {}

// EnterParenCourse is called when production parenCourse is entered.
func (s *BaseRequirementsListener) EnterParenCourse(ctx *ParenCourseContext) {}

// ExitParenCourse is called when production parenCourse is exited.
func (s *BaseRequirementsListener) ExitParenCourse(ctx *ParenCourseContext) {}

// EnterCrossListedCourse is called when production crossListedCourse is entered.
func (s *BaseRequirementsListener) EnterCrossListedCourse(ctx *CrossListedCourseContext) {}

// ExitCrossListedCourse is called when production crossListedCourse is exited.
func (s *BaseRequirementsListener) ExitCrossListedCourse(ctx *CrossListedCourseContext) {}

// EnterFullCourseList is called when production fullCourseList is entered.
func (s *BaseRequirementsListener) EnterFullCourseList(ctx *FullCourseListContext) {}

// ExitFullCourseList is called when production fullCourseList is exited.
func (s *BaseRequirementsListener) ExitFullCourseList(ctx *FullCourseListContext) {}

// EnterShorthandCourseList is called when production shorthandCourseList is entered.
func (s *BaseRequirementsListener) EnterShorthandCourseList(ctx *ShorthandCourseListContext) {}

// ExitShorthandCourseList is called when production shorthandCourseList is exited.
func (s *BaseRequirementsListener) ExitShorthandCourseList(ctx *ShorthandCourseListContext) {}

// EnterTitle is called when production title is entered.
func (s *BaseRequirementsListener) EnterTitle(ctx *TitleContext) {}

// ExitTitle is called when production title is exited.
func (s *BaseRequirementsListener) ExitTitle(ctx *TitleContext) {}

// EnterSimpleGradeCondition is called when production simpleGradeCondition is entered.
func (s *BaseRequirementsListener) EnterSimpleGradeCondition(ctx *SimpleGradeConditionContext) {}

// ExitSimpleGradeCondition is called when production simpleGradeCondition is exited.
func (s *BaseRequirementsListener) ExitSimpleGradeCondition(ctx *SimpleGradeConditionContext) {}

// EnterGpaGradeCondition is called when production gpaGradeCondition is entered.
func (s *BaseRequirementsListener) EnterGpaGradeCondition(ctx *GpaGradeConditionContext) {}

// ExitGpaGradeCondition is called when production gpaGradeCondition is exited.
func (s *BaseRequirementsListener) ExitGpaGradeCondition(ctx *GpaGradeConditionContext) {}

// EnterParenGradeCondition is called when production parenGradeCondition is entered.
func (s *BaseRequirementsListener) EnterParenGradeCondition(ctx *ParenGradeConditionContext) {}

// ExitParenGradeCondition is called when production parenGradeCondition is exited.
func (s *BaseRequirementsListener) ExitParenGradeCondition(ctx *ParenGradeConditionContext) {}

// EnterGradeAtLeastCondition is called when production gradeAtLeastCondition is entered.
func (s *BaseRequirementsListener) EnterGradeAtLeastCondition(ctx *GradeAtLeastConditionContext) {}

// ExitGradeAtLeastCondition is called when production gradeAtLeastCondition is exited.
func (s *BaseRequirementsListener) ExitGradeAtLeastCondition(ctx *GradeAtLeastConditionContext) {}

// EnterGrade_course_list is called when production grade_course_list is entered.
func (s *BaseRequirementsListener) EnterGrade_course_list(ctx *Grade_course_listContext) {}

// ExitGrade_course_list is called when production grade_course_list is exited.
func (s *BaseRequirementsListener) ExitGrade_course_list(ctx *Grade_course_listContext) {}

// EnterCourseAlternativeCondition is called when production courseAlternativeCondition is entered.
func (s *BaseRequirementsListener) EnterCourseAlternativeCondition(ctx *CourseAlternativeConditionContext) {
}

// ExitCourseAlternativeCondition is called when production courseAlternativeCondition is exited.
func (s *BaseRequirementsListener) ExitCourseAlternativeCondition(ctx *CourseAlternativeConditionContext) {
}

// EnterParenAlternativeCondition is called when production parenAlternativeCondition is entered.
func (s *BaseRequirementsListener) EnterParenAlternativeCondition(ctx *ParenAlternativeConditionContext) {
}

// ExitParenAlternativeCondition is called when production parenAlternativeCondition is exited.
func (s *BaseRequirementsListener) ExitParenAlternativeCondition(ctx *ParenAlternativeConditionContext) {
}

// EnterGradeAlternativeCondition is called when production gradeAlternativeCondition is entered.
func (s *BaseRequirementsListener) EnterGradeAlternativeCondition(ctx *GradeAlternativeConditionContext) {
}

// ExitGradeAlternativeCondition is called when production gradeAlternativeCondition is exited.
func (s *BaseRequirementsListener) ExitGradeAlternativeCondition(ctx *GradeAlternativeConditionContext) {
}

// EnterGradeLevelStandingCondition is called when production gradeLevelStandingCondition is entered.
func (s *BaseRequirementsListener) EnterGradeLevelStandingCondition(ctx *GradeLevelStandingConditionContext) {
}

// ExitGradeLevelStandingCondition is called when production gradeLevelStandingCondition is exited.
func (s *BaseRequirementsListener) ExitGradeLevelStandingCondition(ctx *GradeLevelStandingConditionContext) {
}

// EnterGradeLevelMajorStandingCondition is called when production gradeLevelMajorStandingCondition is entered.
func (s *BaseRequirementsListener) EnterGradeLevelMajorStandingCondition(ctx *GradeLevelMajorStandingConditionContext) {
}

// ExitGradeLevelMajorStandingCondition is called when production gradeLevelMajorStandingCondition is exited.
func (s *BaseRequirementsListener) ExitGradeLevelMajorStandingCondition(ctx *GradeLevelMajorStandingConditionContext) {
}

// EnterMinimumGradeLevelStandingCondition is called when production minimumGradeLevelStandingCondition is entered.
func (s *BaseRequirementsListener) EnterMinimumGradeLevelStandingCondition(ctx *MinimumGradeLevelStandingConditionContext) {
}

// ExitMinimumGradeLevelStandingCondition is called when production minimumGradeLevelStandingCondition is exited.
func (s *BaseRequirementsListener) ExitMinimumGradeLevelStandingCondition(ctx *MinimumGradeLevelStandingConditionContext) {
}

// EnterAtLeastGradeLevelStandingCondition is called when production atLeastGradeLevelStandingCondition is entered.
func (s *BaseRequirementsListener) EnterAtLeastGradeLevelStandingCondition(ctx *AtLeastGradeLevelStandingConditionContext) {
}

// ExitAtLeastGradeLevelStandingCondition is called when production atLeastGradeLevelStandingCondition is exited.
func (s *BaseRequirementsListener) ExitAtLeastGradeLevelStandingCondition(ctx *AtLeastGradeLevelStandingConditionContext) {
}

// EnterPrefixGradeLevelStandingCondition is called when production prefixGradeLevelStandingCondition is entered.
func (s *BaseRequirementsListener) EnterPrefixGradeLevelStandingCondition(ctx *PrefixGradeLevelStandingConditionContext) {
}

// ExitPrefixGradeLevelStandingCondition is called when production prefixGradeLevelStandingCondition is exited.
func (s *BaseRequirementsListener) ExitPrefixGradeLevelStandingCondition(ctx *PrefixGradeLevelStandingConditionContext) {
}

// EnterGraduateStandingInCondition is called when production graduateStandingInCondition is entered.
func (s *BaseRequirementsListener) EnterGraduateStandingInCondition(ctx *GraduateStandingInConditionContext) {
}

// ExitGraduateStandingInCondition is called when production graduateStandingInCondition is exited.
func (s *BaseRequirementsListener) ExitGraduateStandingInCondition(ctx *GraduateStandingInConditionContext) {
}

// EnterGraduateLevelStandingCondition is called when production graduateLevelStandingCondition is entered.
func (s *BaseRequirementsListener) EnterGraduateLevelStandingCondition(ctx *GraduateLevelStandingConditionContext) {
}

// ExitGraduateLevelStandingCondition is called when production graduateLevelStandingCondition is exited.
func (s *BaseRequirementsListener) ExitGraduateLevelStandingCondition(ctx *GraduateLevelStandingConditionContext) {
}

// EnterUniversityGpaCondition is called when production universityGpaCondition is entered.
func (s *BaseRequirementsListener) EnterUniversityGpaCondition(ctx *UniversityGpaConditionContext) {}

// ExitUniversityGpaCondition is called when production universityGpaCondition is exited.
func (s *BaseRequirementsListener) ExitUniversityGpaCondition(ctx *UniversityGpaConditionContext) {}

// EnterMinimumGpaCondition is called when production minimumGpaCondition is entered.
func (s *BaseRequirementsListener) EnterMinimumGpaCondition(ctx *MinimumGpaConditionContext) {}

// ExitMinimumGpaCondition is called when production minimumGpaCondition is exited.
func (s *BaseRequirementsListener) ExitMinimumGpaCondition(ctx *MinimumGpaConditionContext) {}

// EnterGpaInCourseCondition is called when production gpaInCourseCondition is entered.
func (s *BaseRequirementsListener) EnterGpaInCourseCondition(ctx *GpaInCourseConditionContext) {}

// ExitGpaInCourseCondition is called when production gpaInCourseCondition is exited.
func (s *BaseRequirementsListener) ExitGpaInCourseCondition(ctx *GpaInCourseConditionContext) {}

// EnterPrefixMajorCondition is called when production prefixMajorCondition is entered.
func (s *BaseRequirementsListener) EnterPrefixMajorCondition(ctx *PrefixMajorConditionContext) {}

// ExitPrefixMajorCondition is called when production prefixMajorCondition is exited.
func (s *BaseRequirementsListener) ExitPrefixMajorCondition(ctx *PrefixMajorConditionContext) {}

// EnterGradeLevelPrefixMajorCondition is called when production gradeLevelPrefixMajorCondition is entered.
func (s *BaseRequirementsListener) EnterGradeLevelPrefixMajorCondition(ctx *GradeLevelPrefixMajorConditionContext) {
}

// ExitGradeLevelPrefixMajorCondition is called when production gradeLevelPrefixMajorCondition is exited.
func (s *BaseRequirementsListener) ExitGradeLevelPrefixMajorCondition(ctx *GradeLevelPrefixMajorConditionContext) {
}

// EnterDegreeTypePrefixMajorCondition is called when production degreeTypePrefixMajorCondition is entered.
func (s *BaseRequirementsListener) EnterDegreeTypePrefixMajorCondition(ctx *DegreeTypePrefixMajorConditionContext) {
}

// ExitDegreeTypePrefixMajorCondition is called when production degreeTypePrefixMajorCondition is exited.
func (s *BaseRequirementsListener) ExitDegreeTypePrefixMajorCondition(ctx *DegreeTypePrefixMajorConditionContext) {
}

// EnterNamedMajorCondition is called when production namedMajorCondition is entered.
func (s *BaseRequirementsListener) EnterNamedMajorCondition(ctx *NamedMajorConditionContext) {}

// ExitNamedMajorCondition is called when production namedMajorCondition is exited.
func (s *BaseRequirementsListener) ExitNamedMajorCondition(ctx *NamedMajorConditionContext) {}

// EnterNamedDegreeTypeMajorCondition is called when production namedDegreeTypeMajorCondition is entered.
func (s *BaseRequirementsListener) EnterNamedDegreeTypeMajorCondition(ctx *NamedDegreeTypeMajorConditionContext) {
}

// ExitNamedDegreeTypeMajorCondition is called when production namedDegreeTypeMajorCondition is exited.
func (s *BaseRequirementsListener) ExitNamedDegreeTypeMajorCondition(ctx *NamedDegreeTypeMajorConditionContext) {
}

// EnterUndergraduateDegreeCondition is called when production undergraduateDegreeCondition is entered.
func (s *BaseRequirementsListener) EnterUndergraduateDegreeCondition(ctx *UndergraduateDegreeConditionContext) {
}

// ExitUndergraduateDegreeCondition is called when production undergraduateDegreeCondition is exited.
func (s *BaseRequirementsListener) ExitUndergraduateDegreeCondition(ctx *UndergraduateDegreeConditionContext) {
}

// EnterBachelorsOrMastersCondition is called when production bachelorsOrMastersCondition is entered.
func (s *BaseRequirementsListener) EnterBachelorsOrMastersCondition(ctx *BachelorsOrMastersConditionContext) {
}

// ExitBachelorsOrMastersCondition is called when production bachelorsOrMastersCondition is exited.
func (s *BaseRequirementsListener) ExitBachelorsOrMastersCondition(ctx *BachelorsOrMastersConditionContext) {
}

// EnterDegree_list is called when production degree_list is entered.
func (s *BaseRequirementsListener) EnterDegree_list(ctx *Degree_listContext) {}

// ExitDegree_list is called when production degree_list is exited.
func (s *BaseRequirementsListener) ExitDegree_list(ctx *Degree_listContext) {}

// EnterDegree is called when production degree is entered.
func (s *BaseRequirementsListener) EnterDegree(ctx *DegreeContext) {}

// ExitDegree is called when production degree is exited.
func (s *BaseRequirementsListener) ExitDegree(ctx *DegreeContext) {}

// EnterCoreCondition is called when production coreCondition is entered.
func (s *BaseRequirementsListener) EnterCoreCondition(ctx *CoreConditionContext) {}

// ExitCoreCondition is called when production coreCondition is exited.
func (s *BaseRequirementsListener) ExitCoreCondition(ctx *CoreConditionContext) {}

// EnterAnyCoreSCHCondition is called when production anyCoreSCHCondition is entered.
func (s *BaseRequirementsListener) EnterAnyCoreSCHCondition(ctx *AnyCoreSCHConditionContext) {}

// ExitAnyCoreSCHCondition is called when production anyCoreSCHCondition is exited.
func (s *BaseRequirementsListener) ExitAnyCoreSCHCondition(ctx *AnyCoreSCHConditionContext) {}

// EnterCompletionOfCoreCondition is called when production completionOfCoreCondition is entered.
func (s *BaseRequirementsListener) EnterCompletionOfCoreCondition(ctx *CompletionOfCoreConditionContext) {
}

// ExitCompletionOfCoreCondition is called when production completionOfCoreCondition is exited.
func (s *BaseRequirementsListener) ExitCompletionOfCoreCondition(ctx *CompletionOfCoreConditionContext) {
}

// EnterSemesterCreditHoursCondition is called when production semesterCreditHoursCondition is entered.
func (s *BaseRequirementsListener) EnterSemesterCreditHoursCondition(ctx *SemesterCreditHoursConditionContext) {
}

// ExitSemesterCreditHoursCondition is called when production semesterCreditHoursCondition is exited.
func (s *BaseRequirementsListener) ExitSemesterCreditHoursCondition(ctx *SemesterCreditHoursConditionContext) {
}

// EnterMinimumHoursCondition is called when production minimumHoursCondition is entered.
func (s *BaseRequirementsListener) EnterMinimumHoursCondition(ctx *MinimumHoursConditionContext) {}

// ExitMinimumHoursCondition is called when production minimumHoursCondition is exited.
func (s *BaseRequirementsListener) ExitMinimumHoursCondition(ctx *MinimumHoursConditionContext) {}

// EnterUpperDivisionSCHCondition is called when production upperDivisionSCHCondition is entered.
func (s *BaseRequirementsListener) EnterUpperDivisionSCHCondition(ctx *UpperDivisionSCHConditionContext) {
}

// ExitUpperDivisionSCHCondition is called when production upperDivisionSCHCondition is exited.
func (s *BaseRequirementsListener) ExitUpperDivisionSCHCondition(ctx *UpperDivisionSCHConditionContext) {
}

// EnterUpperDivisionCreditsCondition is called when production upperDivisionCreditsCondition is entered.
func (s *BaseRequirementsListener) EnterUpperDivisionCreditsCondition(ctx *UpperDivisionCreditsConditionContext) {
}

// ExitUpperDivisionCreditsCondition is called when production upperDivisionCreditsCondition is exited.
func (s *BaseRequirementsListener) ExitUpperDivisionCreditsCondition(ctx *UpperDivisionCreditsConditionContext) {
}

// EnterUpperDivisionClassesCondition is called when production upperDivisionClassesCondition is entered.
func (s *BaseRequirementsListener) EnterUpperDivisionClassesCondition(ctx *UpperDivisionClassesConditionContext) {
}

// ExitUpperDivisionClassesCondition is called when production upperDivisionClassesCondition is exited.
func (s *BaseRequirementsListener) ExitUpperDivisionClassesCondition(ctx *UpperDivisionClassesConditionContext) {
}

// EnterFourThousandLevelCondition is called when production fourThousandLevelCondition is entered.
func (s *BaseRequirementsListener) EnterFourThousandLevelCondition(ctx *FourThousandLevelConditionContext) {
}

// ExitFourThousandLevelCondition is called when production fourThousandLevelCondition is exited.
func (s *BaseRequirementsListener) ExitFourThousandLevelCondition(ctx *FourThousandLevelConditionContext) {
}

// EnterResearchCondition is called when production researchCondition is entered.
func (s *BaseRequirementsListener) EnterResearchCondition(ctx *ResearchConditionContext) {}

// ExitResearchCondition is called when production researchCondition is exited.
func (s *BaseRequirementsListener) ExitResearchCondition(ctx *ResearchConditionContext) {}

// EnterCompleteNOfFollowingCondition is called when production completeNOfFollowingCondition is entered.
func (s *BaseRequirementsListener) EnterCompleteNOfFollowingCondition(ctx *CompleteNOfFollowingConditionContext) {
}

// ExitCompleteNOfFollowingCondition is called when production completeNOfFollowingCondition is exited.
func (s *BaseRequirementsListener) ExitCompleteNOfFollowingCondition(ctx *CompleteNOfFollowingConditionContext) {
}

// EnterCompleteNFromFollowingCondition is called when production completeNFromFollowingCondition is entered.
func (s *BaseRequirementsListener) EnterCompleteNFromFollowingCondition(ctx *CompleteNFromFollowingConditionContext) {
}

// ExitCompleteNFromFollowingCondition is called when production completeNFromFollowingCondition is exited.
func (s *BaseRequirementsListener) ExitCompleteNFromFollowingCondition(ctx *CompleteNFromFollowingConditionContext) {
}

// EnterCompleteHoursFromFollowingCondition is called when production completeHoursFromFollowingCondition is entered.
func (s *BaseRequirementsListener) EnterCompleteHoursFromFollowingCondition(ctx *CompleteHoursFromFollowingConditionContext) {
}

// ExitCompleteHoursFromFollowingCondition is called when production completeHoursFromFollowingCondition is exited.
func (s *BaseRequirementsListener) ExitCompleteHoursFromFollowingCondition(ctx *CompleteHoursFromFollowingConditionContext) {
}

// EnterPlacementScoreComparisonCondition is called when production placementScoreComparisonCondition is entered.
func (s *BaseRequirementsListener) EnterPlacementScoreComparisonCondition(ctx *PlacementScoreComparisonConditionContext) {
}

// ExitPlacementScoreComparisonCondition is called when production placementScoreComparisonCondition is exited.
func (s *BaseRequirementsListener) ExitPlacementScoreComparisonCondition(ctx *PlacementScoreComparisonConditionContext) {
}

// EnterPlacementScoreRangeCondition is called when production placementScoreRangeCondition is entered.
func (s *BaseRequirementsListener) EnterPlacementScoreRangeCondition(ctx *PlacementScoreRangeConditionContext) {
}

// ExitPlacementScoreRangeCondition is called when production placementScoreRangeCondition is exited.
func (s *BaseRequirementsListener) ExitPlacementScoreRangeCondition(ctx *PlacementScoreRangeConditionContext) {
}

// EnterPlacementScoreMinimumCondition is called when production placementScoreMinimumCondition is entered.
func (s *BaseRequirementsListener) EnterPlacementScoreMinimumCondition(ctx *PlacementScoreMinimumConditionContext) {
}

// ExitPlacementScoreMinimumCondition is called when production placementScoreMinimumCondition is exited.
func (s *BaseRequirementsListener) ExitPlacementScoreMinimumCondition(ctx *PlacementScoreMinimumConditionContext) {
}

// EnterPlacement_test_name is called when production placement_test_name is entered.
func (s *BaseRequirementsListener) EnterPlacement_test_name(ctx *Placement_test_nameContext) {}

// ExitPlacement_test_name is called when production placement_test_name is exited.
func (s *BaseRequirementsListener) ExitPlacement_test_name(ctx *Placement_test_nameContext) {}

// EnterApScoreCondition is called when production apScoreCondition is entered.
func (s *BaseRequirementsListener) EnterApScoreCondition(ctx *ApScoreConditionContext) {}

// ExitApScoreCondition is called when production apScoreCondition is exited.
func (s *BaseRequirementsListener) ExitApScoreCondition(ctx *ApScoreConditionContext) {}

// EnterAleksScoreCondition is called when production aleksScoreCondition is entered.
func (s *BaseRequirementsListener) EnterAleksScoreCondition(ctx *AleksScoreConditionContext) {}

// ExitAleksScoreCondition is called when production aleksScoreCondition is exited.
func (s *BaseRequirementsListener) ExitAleksScoreCondition(ctx *AleksScoreConditionContext) {}

// EnterBothHonorsCondition is called when production bothHonorsCondition is entered.
func (s *BaseRequirementsListener) EnterBothHonorsCondition(ctx *BothHonorsConditionContext) {}

// ExitBothHonorsCondition is called when production bothHonorsCondition is exited.
func (s *BaseRequirementsListener) ExitBothHonorsCondition(ctx *BothHonorsConditionContext) {}

// EnterSingleHonorsCondition is called when production singleHonorsCondition is entered.
func (s *BaseRequirementsListener) EnterSingleHonorsCondition(ctx *SingleHonorsConditionContext) {}

// ExitSingleHonorsCondition is called when production singleHonorsCondition is exited.
func (s *BaseRequirementsListener) ExitSingleHonorsCondition(ctx *SingleHonorsConditionContext) {}

// EnterGroupCondition is called when production groupCondition is entered.
func (s *BaseRequirementsListener) EnterGroupCondition(ctx *GroupConditionContext) {}

// ExitGroupCondition is called when production groupCondition is exited.
func (s *BaseRequirementsListener) ExitGroupCondition(ctx *GroupConditionContext) {}

// EnterGroup is called when production group is entered.
func (s *BaseRequirementsListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BaseRequirementsListener) ExitGroup(ctx *GroupContext) {}

// EnterConcurrentEnrollmentCondition is called when production concurrentEnrollmentCondition is entered.
func (s *BaseRequirementsListener) EnterConcurrentEnrollmentCondition(ctx *ConcurrentEnrollmentConditionContext) {
}

// ExitConcurrentEnrollmentCondition is called when production concurrentEnrollmentCondition is exited.
func (s *BaseRequirementsListener) ExitConcurrentEnrollmentCondition(ctx *ConcurrentEnrollmentConditionContext) {
}

// EnterExactSectionCondition is called when production exactSectionCondition is entered.
func (s *BaseRequirementsListener) EnterExactSectionCondition(ctx *ExactSectionConditionContext) {}

// ExitExactSectionCondition is called when production exactSectionCondition is exited.
func (s *BaseRequirementsListener) ExitExactSectionCondition(ctx *ExactSectionConditionContext) {}

// EnterWorkshopSectionCondition is called when production workshopSectionCondition is entered.
func (s *BaseRequirementsListener) EnterWorkshopSectionCondition(ctx *WorkshopSectionConditionContext) {
}

// ExitWorkshopSectionCondition is called when production workshopSectionCondition is exited.
func (s *BaseRequirementsListener) ExitWorkshopSectionCondition(ctx *WorkshopSectionConditionContext) {
}

// EnterAnyPreviousMajorCourseCondition is called when production anyPreviousMajorCourseCondition is entered.
func (s *BaseRequirementsListener) EnterAnyPreviousMajorCourseCondition(ctx *AnyPreviousMajorCourseConditionContext) {
}

// ExitAnyPreviousMajorCourseCondition is called when production anyPreviousMajorCourseCondition is exited.
func (s *BaseRequirementsListener) ExitAnyPreviousMajorCourseCondition(ctx *AnyPreviousMajorCourseConditionContext) {
}

// EnterAcademic_plan_condition is called when production academic_plan_condition is entered.
func (s *BaseRequirementsListener) EnterAcademic_plan_condition(ctx *Academic_plan_conditionContext) {
}

// ExitAcademic_plan_condition is called when production academic_plan_condition is exited.
func (s *BaseRequirementsListener) ExitAcademic_plan_condition(ctx *Academic_plan_conditionContext) {}

// EnterCourseRepeatRule is called when production courseRepeatRule is entered.
func (s *BaseRequirementsListener) EnterCourseRepeatRule(ctx *CourseRepeatRuleContext) {}

// ExitCourseRepeatRule is called when production courseRepeatRule is exited.
func (s *BaseRequirementsListener) ExitCourseRepeatRule(ctx *CourseRepeatRuleContext) {}

// EnterInternshipRepeatRule is called when production internshipRepeatRule is entered.
func (s *BaseRequirementsListener) EnterInternshipRepeatRule(ctx *InternshipRepeatRuleContext) {}

// ExitInternshipRepeatRule is called when production internshipRepeatRule is exited.
func (s *BaseRequirementsListener) ExitInternshipRepeatRule(ctx *InternshipRepeatRuleContext) {}

// EnterBareRepeatRule is called when production bareRepeatRule is entered.
func (s *BaseRequirementsListener) EnterBareRepeatRule(ctx *BareRepeatRuleContext) {}

// ExitBareRepeatRule is called when production bareRepeatRule is exited.
func (s *BaseRequirementsListener) ExitBareRepeatRule(ctx *BareRepeatRuleContext) {}

// EnterRepeatMaxHoursRule is called when production repeatMaxHoursRule is entered.
func (s *BaseRequirementsListener) EnterRepeatMaxHoursRule(ctx *RepeatMaxHoursRuleContext) {}

// ExitRepeatMaxHoursRule is called when production repeatMaxHoursRule is exited.
func (s *BaseRequirementsListener) ExitRepeatMaxHoursRule(ctx *RepeatMaxHoursRuleContext) {}

// EnterRepeatHoursMaxSuffixRule is called when production repeatHoursMaxSuffixRule is entered.
func (s *BaseRequirementsListener) EnterRepeatHoursMaxSuffixRule(ctx *RepeatHoursMaxSuffixRuleContext) {
}

// ExitRepeatHoursMaxSuffixRule is called when production repeatHoursMaxSuffixRule is exited.
func (s *BaseRequirementsListener) ExitRepeatHoursMaxSuffixRule(ctx *RepeatHoursMaxSuffixRuleContext) {
}

// EnterCourseRepeatMaxHoursRule is called when production courseRepeatMaxHoursRule is entered.
func (s *BaseRequirementsListener) EnterCourseRepeatMaxHoursRule(ctx *CourseRepeatMaxHoursRuleContext) {
}

// ExitCourseRepeatMaxHoursRule is called when production courseRepeatMaxHoursRule is exited.
func (s *BaseRequirementsListener) ExitCourseRepeatMaxHoursRule(ctx *CourseRepeatMaxHoursRuleContext) {
}

// EnterCombinedRepeatMaxHoursRule is called when production combinedRepeatMaxHoursRule is entered.
func (s *BaseRequirementsListener) EnterCombinedRepeatMaxHoursRule(ctx *CombinedRepeatMaxHoursRuleContext) {
}

// ExitCombinedRepeatMaxHoursRule is called when production combinedRepeatMaxHoursRule is exited.
func (s *BaseRequirementsListener) ExitCombinedRepeatMaxHoursRule(ctx *CombinedRepeatMaxHoursRuleContext) {
}

// EnterTopicsVaryRepeatRule is called when production topicsVaryRepeatRule is entered.
func (s *BaseRequirementsListener) EnterTopicsVaryRepeatRule(ctx *TopicsVaryRepeatRuleContext) {}

// ExitTopicsVaryRepeatRule is called when production topicsVaryRepeatRule is exited.
func (s *BaseRequirementsListener) ExitTopicsVaryRepeatRule(ctx *TopicsVaryRepeatRuleContext) {}

// EnterCourseRepeatLimitRule is called when production courseRepeatLimitRule is entered.
func (s *BaseRequirementsListener) EnterCourseRepeatLimitRule(ctx *CourseRepeatLimitRuleContext) {}

// ExitCourseRepeatLimitRule is called when production courseRepeatLimitRule is exited.
func (s *BaseRequirementsListener) ExitCourseRepeatLimitRule(ctx *CourseRepeatLimitRuleContext) {}

// EnterRepeatUpToTimesRule is called when production repeatUpToTimesRule is entered.
func (s *BaseRequirementsListener) EnterRepeatUpToTimesRule(ctx *RepeatUpToTimesRuleContext) {}

// ExitRepeatUpToTimesRule is called when production repeatUpToTimesRule is exited.
func (s *BaseRequirementsListener) ExitRepeatUpToTimesRule(ctx *RepeatUpToTimesRuleContext) {}

// EnterRepeatMaxTimesRule is called when production repeatMaxTimesRule is entered.
func (s *BaseRequirementsListener) EnterRepeatMaxTimesRule(ctx *RepeatMaxTimesRuleContext) {}

// ExitRepeatMaxTimesRule is called when production repeatMaxTimesRule is exited.
func (s *BaseRequirementsListener) ExitRepeatMaxTimesRule(ctx *RepeatMaxTimesRuleContext) {}

// EnterAndRepeateExpr is called when production andRepeateExpr is entered.
func (s *BaseRequirementsListener) EnterAndRepeateExpr(ctx *AndRepeateExprContext) {}

// ExitAndRepeateExpr is called when production andRepeateExpr is exited.
func (s *BaseRequirementsListener) ExitAndRepeateExpr(ctx *AndRepeateExprContext) {}

// EnterAlternativeRepeateExpr is called when production alternativeRepeateExpr is entered.
func (s *BaseRequirementsListener) EnterAlternativeRepeateExpr(ctx *AlternativeRepeateExprContext) {}

// ExitAlternativeRepeateExpr is called when production alternativeRepeateExpr is exited.
func (s *BaseRequirementsListener) ExitAlternativeRepeateExpr(ctx *AlternativeRepeateExprContext) {}

// EnterCourseRepeateExpr is called when production courseRepeateExpr is entered.
func (s *BaseRequirementsListener) EnterCourseRepeateExpr(ctx *CourseRepeateExprContext) {}

// ExitCourseRepeateExpr is called when production courseRepeateExpr is exited.
func (s *BaseRequirementsListener) ExitCourseRepeateExpr(ctx *CourseRepeateExprContext) {}

// EnterGpa_repeate_rule is called when production gpa_repeate_rule is entered.
func (s *BaseRequirementsListener) EnterGpa_repeate_rule(ctx *Gpa_repeate_ruleContext) {}

// ExitGpa_repeate_rule is called when production gpa_repeate_rule is exited.
func (s *BaseRequirementsListener) ExitGpa_repeate_rule(ctx *Gpa_repeate_ruleContext) {}

// EnterSchoolDegreeSatisfactionRule is called when production schoolDegreeSatisfactionRule is entered.
func (s *BaseRequirementsListener) EnterSchoolDegreeSatisfactionRule(ctx *SchoolDegreeSatisfactionRuleContext) {
}

// ExitSchoolDegreeSatisfactionRule is called when production schoolDegreeSatisfactionRule is exited.
func (s *BaseRequirementsListener) ExitSchoolDegreeSatisfactionRule(ctx *SchoolDegreeSatisfactionRuleContext) {
}

// EnterPrefixDegreeSatisfactionRule is called when production prefixDegreeSatisfactionRule is entered.
func (s *BaseRequirementsListener) EnterPrefixDegreeSatisfactionRule(ctx *PrefixDegreeSatisfactionRuleContext) {
}

// ExitPrefixDegreeSatisfactionRule is called when production prefixDegreeSatisfactionRule is exited.
func (s *BaseRequirementsListener) ExitPrefixDegreeSatisfactionRule(ctx *PrefixDegreeSatisfactionRuleContext) {
}

// EnterOfMultiPrefixDegreeSatisfactionRule is called when production ofMultiPrefixDegreeSatisfactionRule is entered.
func (s *BaseRequirementsListener) EnterOfMultiPrefixDegreeSatisfactionRule(ctx *OfMultiPrefixDegreeSatisfactionRuleContext) {
}

// ExitOfMultiPrefixDegreeSatisfactionRule is called when production ofMultiPrefixDegreeSatisfactionRule is exited.
func (s *BaseRequirementsListener) ExitOfMultiPrefixDegreeSatisfactionRule(ctx *OfMultiPrefixDegreeSatisfactionRuleContext) {
}

// EnterElectivesDegreeSatisfactionRule is called when production electivesDegreeSatisfactionRule is entered.
func (s *BaseRequirementsListener) EnterElectivesDegreeSatisfactionRule(ctx *ElectivesDegreeSatisfactionRuleContext) {
}

// ExitElectivesDegreeSatisfactionRule is called when production electivesDegreeSatisfactionRule is exited.
func (s *BaseRequirementsListener) ExitElectivesDegreeSatisfactionRule(ctx *ElectivesDegreeSatisfactionRuleContext) {
}

// EnterSchoolsDegreeSatisfactionRule is called when production schoolsDegreeSatisfactionRule is entered.
func (s *BaseRequirementsListener) EnterSchoolsDegreeSatisfactionRule(ctx *SchoolsDegreeSatisfactionRuleContext) {
}

// ExitSchoolsDegreeSatisfactionRule is called when production schoolsDegreeSatisfactionRule is exited.
func (s *BaseRequirementsListener) ExitSchoolsDegreeSatisfactionRule(ctx *SchoolsDegreeSatisfactionRuleContext) {
}

// EnterMultiPrefixForDegreeSatisfactionRule is called when production multiPrefixForDegreeSatisfactionRule is entered.
func (s *BaseRequirementsListener) EnterMultiPrefixForDegreeSatisfactionRule(ctx *MultiPrefixForDegreeSatisfactionRuleContext) {
}

// ExitMultiPrefixForDegreeSatisfactionRule is called when production multiPrefixForDegreeSatisfactionRule is exited.
func (s *BaseRequirementsListener) ExitMultiPrefixForDegreeSatisfactionRule(ctx *MultiPrefixForDegreeSatisfactionRuleContext) {
}

// EnterStudentDegreeSatisfactionRule is called when production studentDegreeSatisfactionRule is entered.
func (s *BaseRequirementsListener) EnterStudentDegreeSatisfactionRule(ctx *StudentDegreeSatisfactionRuleContext) {
}

// ExitStudentDegreeSatisfactionRule is called when production studentDegreeSatisfactionRule is exited.
func (s *BaseRequirementsListener) ExitStudentDegreeSatisfactionRule(ctx *StudentDegreeSatisfactionRuleContext) {
}

// EnterMathDegreeSatisfactionRule is called when production mathDegreeSatisfactionRule is entered.
func (s *BaseRequirementsListener) EnterMathDegreeSatisfactionRule(ctx *MathDegreeSatisfactionRuleContext) {
}

// ExitMathDegreeSatisfactionRule is called when production mathDegreeSatisfactionRule is exited.
func (s *BaseRequirementsListener) ExitMathDegreeSatisfactionRule(ctx *MathDegreeSatisfactionRuleContext) {
}

// EnterNamedDegreeSatisfactionRule is called when production namedDegreeSatisfactionRule is entered.
func (s *BaseRequirementsListener) EnterNamedDegreeSatisfactionRule(ctx *NamedDegreeSatisfactionRuleContext) {
}

// ExitNamedDegreeSatisfactionRule is called when production namedDegreeSatisfactionRule is exited.
func (s *BaseRequirementsListener) ExitNamedDegreeSatisfactionRule(ctx *NamedDegreeSatisfactionRuleContext) {
}

// EnterMultiPrefixDegreeSatisfactionRule is called when production multiPrefixDegreeSatisfactionRule is entered.
func (s *BaseRequirementsListener) EnterMultiPrefixDegreeSatisfactionRule(ctx *MultiPrefixDegreeSatisfactionRuleContext) {
}

// ExitMultiPrefixDegreeSatisfactionRule is called when production multiPrefixDegreeSatisfactionRule is exited.
func (s *BaseRequirementsListener) ExitMultiPrefixDegreeSatisfactionRule(ctx *MultiPrefixDegreeSatisfactionRuleContext) {
}

// EnterCredit_for_rule is called when production credit_for_rule is entered.
func (s *BaseRequirementsListener) EnterCredit_for_rule(ctx *Credit_for_ruleContext) {}

// ExitCredit_for_rule is called when production credit_for_rule is exited.
func (s *BaseRequirementsListener) ExitCredit_for_rule(ctx *Credit_for_ruleContext) {}

// EnterPrefixLivingLearningRule is called when production prefixLivingLearningRule is entered.
func (s *BaseRequirementsListener) EnterPrefixLivingLearningRule(ctx *PrefixLivingLearningRuleContext) {
}

// ExitPrefixLivingLearningRule is called when production prefixLivingLearningRule is exited.
func (s *BaseRequirementsListener) ExitPrefixLivingLearningRule(ctx *PrefixLivingLearningRuleContext) {
}

// EnterNamedLivingLearningRule is called when production namedLivingLearningRule is entered.
func (s *BaseRequirementsListener) EnterNamedLivingLearningRule(ctx *NamedLivingLearningRuleContext) {
}

// ExitNamedLivingLearningRule is called when production namedLivingLearningRule is exited.
func (s *BaseRequirementsListener) ExitNamedLivingLearningRule(ctx *NamedLivingLearningRuleContext) {}

// EnterSchool_rule is called when production school_rule is entered.
func (s *BaseRequirementsListener) EnterSchool_rule(ctx *School_ruleContext) {}

// ExitSchool_rule is called when production school_rule is exited.
func (s *BaseRequirementsListener) ExitSchool_rule(ctx *School_ruleContext) {}
