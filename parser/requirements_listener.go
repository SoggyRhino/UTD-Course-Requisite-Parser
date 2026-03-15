// Code generated from ../Requirements.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Requirements
import "github.com/antlr4-go/antlr/v4"

// RequirementsListener is a complete listener for a parse tree produced by RequirementsParser.
type RequirementsListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterSentence is called when entering the sentence production.
	EnterSentence(c *SentenceContext)

	// EnterSameAsParenReq is called when entering the sameAsParenReq production.
	EnterSameAsParenReq(c *SameAsParenReqContext)

	// EnterPreOrCoReq is called when entering the preOrCoReq production.
	EnterPreOrCoReq(c *PreOrCoReqContext)

	// EnterPrereqAndCoreqReq is called when entering the prereqAndCoreqReq production.
	EnterPrereqAndCoreqReq(c *PrereqAndCoreqReqContext)

	// EnterComputerScholarsReq is called when entering the computerScholarsReq production.
	EnterComputerScholarsReq(c *ComputerScholarsReqContext)

	// EnterMajorReq is called when entering the majorReq production.
	EnterMajorReq(c *MajorReqContext)

	// EnterDegreeSatisfactionReq is called when entering the degreeSatisfactionReq production.
	EnterDegreeSatisfactionReq(c *DegreeSatisfactionReqContext)

	// EnterRepeatLimitTimesReq is called when entering the repeatLimitTimesReq production.
	EnterRepeatLimitTimesReq(c *RepeatLimitTimesReqContext)

	// EnterAcademicPlanReq is called when entering the academicPlanReq production.
	EnterAcademicPlanReq(c *AcademicPlanReqContext)

	// EnterRepeatLimitHoursReq is called when entering the repeatLimitHoursReq production.
	EnterRepeatLimitHoursReq(c *RepeatLimitHoursReqContext)

	// EnterAppendAcademicPlanReq is called when entering the appendAcademicPlanReq production.
	EnterAppendAcademicPlanReq(c *AppendAcademicPlanReqContext)

	// EnterRepeatReq is called when entering the repeatReq production.
	EnterRepeatReq(c *RepeatReqContext)

	// EnterExprReq is called when entering the exprReq production.
	EnterExprReq(c *ExprReqContext)

	// EnterExactCoreqNoticeReq is called when entering the exactCoreqNoticeReq production.
	EnterExactCoreqNoticeReq(c *ExactCoreqNoticeReqContext)

	// EnterGpaRepeatReq is called when entering the gpaRepeatReq production.
	EnterGpaRepeatReq(c *GpaRepeatReqContext)

	// EnterSchoolReq is called when entering the schoolReq production.
	EnterSchoolReq(c *SchoolReqContext)

	// EnterCoreqReq is called when entering the coreqReq production.
	EnterCoreqReq(c *CoreqReqContext)

	// EnterPrereqReq is called when entering the prereqReq production.
	EnterPrereqReq(c *PrereqReqContext)

	// EnterCreditForReq is called when entering the creditForReq production.
	EnterCreditForReq(c *CreditForReqContext)

	// EnterSameAsReq is called when entering the sameAsReq production.
	EnterSameAsReq(c *SameAsReqContext)

	// EnterExcludeNoticeReq is called when entering the excludeNoticeReq production.
	EnterExcludeNoticeReq(c *ExcludeNoticeReqContext)

	// EnterLivingLearningReq is called when entering the livingLearningReq production.
	EnterLivingLearningReq(c *LivingLearningReqContext)

	// EnterGroupExpr is called when entering the groupExpr production.
	EnterGroupExpr(c *GroupExprContext)

	// EnterCourseExpr is called when entering the courseExpr production.
	EnterCourseExpr(c *CourseExprContext)

	// EnterAlternativeExpr is called when entering the alternativeExpr production.
	EnterAlternativeExpr(c *AlternativeExprContext)

	// EnterUpperDivisionStandingExpr is called when entering the upperDivisionStandingExpr production.
	EnterUpperDivisionStandingExpr(c *UpperDivisionStandingExprContext)

	// EnterSemesterCreditHoursExpr is called when entering the semesterCreditHoursExpr production.
	EnterSemesterCreditHoursExpr(c *SemesterCreditHoursExprContext)

	// EnterResearchExpr is called when entering the researchExpr production.
	EnterResearchExpr(c *ResearchExprContext)

	// EnterComputerScholarsExpr is called when entering the computerScholarsExpr production.
	EnterComputerScholarsExpr(c *ComputerScholarsExprContext)

	// EnterCompleteNExpr is called when entering the completeNExpr production.
	EnterCompleteNExpr(c *CompleteNExprContext)

	// EnterAnyCoreExpr is called when entering the anyCoreExpr production.
	EnterAnyCoreExpr(c *AnyCoreExprContext)

	// EnterHonorsExpr is called when entering the honorsExpr production.
	EnterHonorsExpr(c *HonorsExprContext)

	// EnterGradeExpr is called when entering the gradeExpr production.
	EnterGradeExpr(c *GradeExprContext)

	// EnterDepartmentConsentExpr is called when entering the departmentConsentExpr production.
	EnterDepartmentConsentExpr(c *DepartmentConsentExprContext)

	// EnterParenExpr is called when entering the parenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterAleksScoreExpr is called when entering the aleksScoreExpr production.
	EnterAleksScoreExpr(c *AleksScoreExprContext)

	// EnterAmpersandExpr is called when entering the ampersandExpr production.
	EnterAmpersandExpr(c *AmpersandExprContext)

	// EnterLivingLearningExpr is called when entering the livingLearningExpr production.
	EnterLivingLearningExpr(c *LivingLearningExprContext)

	// EnterUteachConsentExpr is called when entering the uteachConsentExpr production.
	EnterUteachConsentExpr(c *UteachConsentExprContext)

	// EnterGoodAcademicStandingExpr is called when entering the goodAcademicStandingExpr production.
	EnterGoodAcademicStandingExpr(c *GoodAcademicStandingExprContext)

	// EnterCoreExpr is called when entering the coreExpr production.
	EnterCoreExpr(c *CoreExprContext)

	// EnterGpaExpr is called when entering the gpaExpr production.
	EnterGpaExpr(c *GpaExprContext)

	// EnterMajorExpr is called when entering the majorExpr production.
	EnterMajorExpr(c *MajorExprContext)

	// EnterEquivalentExpr is called when entering the equivalentExpr production.
	EnterEquivalentExpr(c *EquivalentExprContext)

	// EnterInstructorConsentExpr is called when entering the instructorConsentExpr production.
	EnterInstructorConsentExpr(c *InstructorConsentExprContext)

	// EnterMinimumHoursExpr is called when entering the minimumHoursExpr production.
	EnterMinimumHoursExpr(c *MinimumHoursExprContext)

	// EnterRepeatLimitHoursExpr is called when entering the repeatLimitHoursExpr production.
	EnterRepeatLimitHoursExpr(c *RepeatLimitHoursExprContext)

	// EnterRepeatRuleExpr is called when entering the repeatRuleExpr production.
	EnterRepeatRuleExpr(c *RepeatRuleExprContext)

	// EnterApScoreExpr is called when entering the apScoreExpr production.
	EnterApScoreExpr(c *ApScoreExprContext)

	// EnterAnyMajorCourseExpr is called when entering the anyMajorCourseExpr production.
	EnterAnyMajorCourseExpr(c *AnyMajorCourseExprContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterUpperDivisionHoursExpr is called when entering the upperDivisionHoursExpr production.
	EnterUpperDivisionHoursExpr(c *UpperDivisionHoursExprContext)

	// EnterUpperDivisionClassesExpr is called when entering the upperDivisionClassesExpr production.
	EnterUpperDivisionClassesExpr(c *UpperDivisionClassesExprContext)

	// EnterPlacementTestExpr is called when entering the placementTestExpr production.
	EnterPlacementTestExpr(c *PlacementTestExprContext)

	// EnterConcurrentEnrollmentExpr is called when entering the concurrentEnrollmentExpr production.
	EnterConcurrentEnrollmentExpr(c *ConcurrentEnrollmentExprContext)

	// EnterGradeLevelStandingExpr is called when entering the gradeLevelStandingExpr production.
	EnterGradeLevelStandingExpr(c *GradeLevelStandingExprContext)

	// EnterDegreeExpr is called when entering the degreeExpr production.
	EnterDegreeExpr(c *DegreeExprContext)

	// EnterGraduateStandingExpr is called when entering the graduateStandingExpr production.
	EnterGraduateStandingExpr(c *GraduateStandingExprContext)

	// EnterExactSectionExpr is called when entering the exactSectionExpr production.
	EnterExactSectionExpr(c *ExactSectionExprContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterSimpleCourse is called when entering the simpleCourse production.
	EnterSimpleCourse(c *SimpleCourseContext)

	// EnterParenCourse is called when entering the parenCourse production.
	EnterParenCourse(c *ParenCourseContext)

	// EnterCrossListedCourse is called when entering the crossListedCourse production.
	EnterCrossListedCourse(c *CrossListedCourseContext)

	// EnterFullCourseList is called when entering the fullCourseList production.
	EnterFullCourseList(c *FullCourseListContext)

	// EnterShorthandCourseList is called when entering the shorthandCourseList production.
	EnterShorthandCourseList(c *ShorthandCourseListContext)

	// EnterTitle is called when entering the title production.
	EnterTitle(c *TitleContext)

	// EnterSimpleGradeCondition is called when entering the simpleGradeCondition production.
	EnterSimpleGradeCondition(c *SimpleGradeConditionContext)

	// EnterGpaGradeCondition is called when entering the gpaGradeCondition production.
	EnterGpaGradeCondition(c *GpaGradeConditionContext)

	// EnterParenGradeCondition is called when entering the parenGradeCondition production.
	EnterParenGradeCondition(c *ParenGradeConditionContext)

	// EnterGradeAtLeastCondition is called when entering the gradeAtLeastCondition production.
	EnterGradeAtLeastCondition(c *GradeAtLeastConditionContext)

	// EnterGrade_course_list is called when entering the grade_course_list production.
	EnterGrade_course_list(c *Grade_course_listContext)

	// EnterCourseAlternativeCondition is called when entering the courseAlternativeCondition production.
	EnterCourseAlternativeCondition(c *CourseAlternativeConditionContext)

	// EnterParenAlternativeCondition is called when entering the parenAlternativeCondition production.
	EnterParenAlternativeCondition(c *ParenAlternativeConditionContext)

	// EnterGradeAlternativeCondition is called when entering the gradeAlternativeCondition production.
	EnterGradeAlternativeCondition(c *GradeAlternativeConditionContext)

	// EnterGradeLevelStandingCondition is called when entering the gradeLevelStandingCondition production.
	EnterGradeLevelStandingCondition(c *GradeLevelStandingConditionContext)

	// EnterGradeLevelMajorStandingCondition is called when entering the gradeLevelMajorStandingCondition production.
	EnterGradeLevelMajorStandingCondition(c *GradeLevelMajorStandingConditionContext)

	// EnterMinimumGradeLevelStandingCondition is called when entering the minimumGradeLevelStandingCondition production.
	EnterMinimumGradeLevelStandingCondition(c *MinimumGradeLevelStandingConditionContext)

	// EnterAtLeastGradeLevelStandingCondition is called when entering the atLeastGradeLevelStandingCondition production.
	EnterAtLeastGradeLevelStandingCondition(c *AtLeastGradeLevelStandingConditionContext)

	// EnterPrefixGradeLevelStandingCondition is called when entering the prefixGradeLevelStandingCondition production.
	EnterPrefixGradeLevelStandingCondition(c *PrefixGradeLevelStandingConditionContext)

	// EnterGraduateStandingInCondition is called when entering the graduateStandingInCondition production.
	EnterGraduateStandingInCondition(c *GraduateStandingInConditionContext)

	// EnterGraduateLevelStandingCondition is called when entering the graduateLevelStandingCondition production.
	EnterGraduateLevelStandingCondition(c *GraduateLevelStandingConditionContext)

	// EnterUniversityGpaCondition is called when entering the universityGpaCondition production.
	EnterUniversityGpaCondition(c *UniversityGpaConditionContext)

	// EnterMinimumGpaCondition is called when entering the minimumGpaCondition production.
	EnterMinimumGpaCondition(c *MinimumGpaConditionContext)

	// EnterGpaInCourseCondition is called when entering the gpaInCourseCondition production.
	EnterGpaInCourseCondition(c *GpaInCourseConditionContext)

	// EnterPrefixMajorCondition is called when entering the prefixMajorCondition production.
	EnterPrefixMajorCondition(c *PrefixMajorConditionContext)

	// EnterGradeLevelPrefixMajorCondition is called when entering the gradeLevelPrefixMajorCondition production.
	EnterGradeLevelPrefixMajorCondition(c *GradeLevelPrefixMajorConditionContext)

	// EnterDegreeTypePrefixMajorCondition is called when entering the degreeTypePrefixMajorCondition production.
	EnterDegreeTypePrefixMajorCondition(c *DegreeTypePrefixMajorConditionContext)

	// EnterNamedMajorCondition is called when entering the namedMajorCondition production.
	EnterNamedMajorCondition(c *NamedMajorConditionContext)

	// EnterNamedDegreeTypeMajorCondition is called when entering the namedDegreeTypeMajorCondition production.
	EnterNamedDegreeTypeMajorCondition(c *NamedDegreeTypeMajorConditionContext)

	// EnterUndergraduateDegreeCondition is called when entering the undergraduateDegreeCondition production.
	EnterUndergraduateDegreeCondition(c *UndergraduateDegreeConditionContext)

	// EnterBachelorsOrMastersCondition is called when entering the bachelorsOrMastersCondition production.
	EnterBachelorsOrMastersCondition(c *BachelorsOrMastersConditionContext)

	// EnterDegree_list is called when entering the degree_list production.
	EnterDegree_list(c *Degree_listContext)

	// EnterDegree is called when entering the degree production.
	EnterDegree(c *DegreeContext)

	// EnterCoreCondition is called when entering the coreCondition production.
	EnterCoreCondition(c *CoreConditionContext)

	// EnterAnyCoreSCHCondition is called when entering the anyCoreSCHCondition production.
	EnterAnyCoreSCHCondition(c *AnyCoreSCHConditionContext)

	// EnterCompletionOfCoreCondition is called when entering the completionOfCoreCondition production.
	EnterCompletionOfCoreCondition(c *CompletionOfCoreConditionContext)

	// EnterSemesterCreditHoursCondition is called when entering the semesterCreditHoursCondition production.
	EnterSemesterCreditHoursCondition(c *SemesterCreditHoursConditionContext)

	// EnterMinimumHoursCondition is called when entering the minimumHoursCondition production.
	EnterMinimumHoursCondition(c *MinimumHoursConditionContext)

	// EnterUpperDivisionSCHCondition is called when entering the upperDivisionSCHCondition production.
	EnterUpperDivisionSCHCondition(c *UpperDivisionSCHConditionContext)

	// EnterUpperDivisionCreditsCondition is called when entering the upperDivisionCreditsCondition production.
	EnterUpperDivisionCreditsCondition(c *UpperDivisionCreditsConditionContext)

	// EnterUpperDivisionClassesCondition is called when entering the upperDivisionClassesCondition production.
	EnterUpperDivisionClassesCondition(c *UpperDivisionClassesConditionContext)

	// EnterFourThousandLevelCondition is called when entering the fourThousandLevelCondition production.
	EnterFourThousandLevelCondition(c *FourThousandLevelConditionContext)

	// EnterResearchCondition is called when entering the researchCondition production.
	EnterResearchCondition(c *ResearchConditionContext)

	// EnterCompleteNOfFollowingCondition is called when entering the completeNOfFollowingCondition production.
	EnterCompleteNOfFollowingCondition(c *CompleteNOfFollowingConditionContext)

	// EnterCompleteNFromFollowingCondition is called when entering the completeNFromFollowingCondition production.
	EnterCompleteNFromFollowingCondition(c *CompleteNFromFollowingConditionContext)

	// EnterCompleteHoursFromFollowingCondition is called when entering the completeHoursFromFollowingCondition production.
	EnterCompleteHoursFromFollowingCondition(c *CompleteHoursFromFollowingConditionContext)

	// EnterPlacementScoreComparisonCondition is called when entering the placementScoreComparisonCondition production.
	EnterPlacementScoreComparisonCondition(c *PlacementScoreComparisonConditionContext)

	// EnterPlacementScoreRangeCondition is called when entering the placementScoreRangeCondition production.
	EnterPlacementScoreRangeCondition(c *PlacementScoreRangeConditionContext)

	// EnterPlacementScoreMinimumCondition is called when entering the placementScoreMinimumCondition production.
	EnterPlacementScoreMinimumCondition(c *PlacementScoreMinimumConditionContext)

	// EnterPlacement_test_name is called when entering the placement_test_name production.
	EnterPlacement_test_name(c *Placement_test_nameContext)

	// EnterApScoreCondition is called when entering the apScoreCondition production.
	EnterApScoreCondition(c *ApScoreConditionContext)

	// EnterAleksScoreCondition is called when entering the aleksScoreCondition production.
	EnterAleksScoreCondition(c *AleksScoreConditionContext)

	// EnterBothHonorsCondition is called when entering the bothHonorsCondition production.
	EnterBothHonorsCondition(c *BothHonorsConditionContext)

	// EnterSingleHonorsCondition is called when entering the singleHonorsCondition production.
	EnterSingleHonorsCondition(c *SingleHonorsConditionContext)

	// EnterGroupCondition is called when entering the groupCondition production.
	EnterGroupCondition(c *GroupConditionContext)

	// EnterGroup is called when entering the group production.
	EnterGroup(c *GroupContext)

	// EnterConcurrentEnrollmentCondition is called when entering the concurrentEnrollmentCondition production.
	EnterConcurrentEnrollmentCondition(c *ConcurrentEnrollmentConditionContext)

	// EnterExactSectionCondition is called when entering the exactSectionCondition production.
	EnterExactSectionCondition(c *ExactSectionConditionContext)

	// EnterWorkshopSectionCondition is called when entering the workshopSectionCondition production.
	EnterWorkshopSectionCondition(c *WorkshopSectionConditionContext)

	// EnterAnyPreviousMajorCourseCondition is called when entering the anyPreviousMajorCourseCondition production.
	EnterAnyPreviousMajorCourseCondition(c *AnyPreviousMajorCourseConditionContext)

	// EnterAcademic_plan_condition is called when entering the academic_plan_condition production.
	EnterAcademic_plan_condition(c *Academic_plan_conditionContext)

	// EnterCourseRepeatRule is called when entering the courseRepeatRule production.
	EnterCourseRepeatRule(c *CourseRepeatRuleContext)

	// EnterInternshipRepeatRule is called when entering the internshipRepeatRule production.
	EnterInternshipRepeatRule(c *InternshipRepeatRuleContext)

	// EnterBareRepeatRule is called when entering the bareRepeatRule production.
	EnterBareRepeatRule(c *BareRepeatRuleContext)

	// EnterRepeatMaxHoursRule is called when entering the repeatMaxHoursRule production.
	EnterRepeatMaxHoursRule(c *RepeatMaxHoursRuleContext)

	// EnterRepeatHoursMaxSuffixRule is called when entering the repeatHoursMaxSuffixRule production.
	EnterRepeatHoursMaxSuffixRule(c *RepeatHoursMaxSuffixRuleContext)

	// EnterCourseRepeatMaxHoursRule is called when entering the courseRepeatMaxHoursRule production.
	EnterCourseRepeatMaxHoursRule(c *CourseRepeatMaxHoursRuleContext)

	// EnterCombinedRepeatMaxHoursRule is called when entering the combinedRepeatMaxHoursRule production.
	EnterCombinedRepeatMaxHoursRule(c *CombinedRepeatMaxHoursRuleContext)

	// EnterTopicsVaryRepeatRule is called when entering the topicsVaryRepeatRule production.
	EnterTopicsVaryRepeatRule(c *TopicsVaryRepeatRuleContext)

	// EnterCourseRepeatLimitRule is called when entering the courseRepeatLimitRule production.
	EnterCourseRepeatLimitRule(c *CourseRepeatLimitRuleContext)

	// EnterRepeatUpToTimesRule is called when entering the repeatUpToTimesRule production.
	EnterRepeatUpToTimesRule(c *RepeatUpToTimesRuleContext)

	// EnterRepeatMaxTimesRule is called when entering the repeatMaxTimesRule production.
	EnterRepeatMaxTimesRule(c *RepeatMaxTimesRuleContext)

	// EnterAndRepeateExpr is called when entering the andRepeateExpr production.
	EnterAndRepeateExpr(c *AndRepeateExprContext)

	// EnterAlternativeRepeateExpr is called when entering the alternativeRepeateExpr production.
	EnterAlternativeRepeateExpr(c *AlternativeRepeateExprContext)

	// EnterCourseRepeateExpr is called when entering the courseRepeateExpr production.
	EnterCourseRepeateExpr(c *CourseRepeateExprContext)

	// EnterGpa_repeate_rule is called when entering the gpa_repeate_rule production.
	EnterGpa_repeate_rule(c *Gpa_repeate_ruleContext)

	// EnterSchoolDegreeSatisfactionRule is called when entering the schoolDegreeSatisfactionRule production.
	EnterSchoolDegreeSatisfactionRule(c *SchoolDegreeSatisfactionRuleContext)

	// EnterPrefixDegreeSatisfactionRule is called when entering the prefixDegreeSatisfactionRule production.
	EnterPrefixDegreeSatisfactionRule(c *PrefixDegreeSatisfactionRuleContext)

	// EnterOfMultiPrefixDegreeSatisfactionRule is called when entering the ofMultiPrefixDegreeSatisfactionRule production.
	EnterOfMultiPrefixDegreeSatisfactionRule(c *OfMultiPrefixDegreeSatisfactionRuleContext)

	// EnterElectivesDegreeSatisfactionRule is called when entering the electivesDegreeSatisfactionRule production.
	EnterElectivesDegreeSatisfactionRule(c *ElectivesDegreeSatisfactionRuleContext)

	// EnterSchoolsDegreeSatisfactionRule is called when entering the schoolsDegreeSatisfactionRule production.
	EnterSchoolsDegreeSatisfactionRule(c *SchoolsDegreeSatisfactionRuleContext)

	// EnterMultiPrefixForDegreeSatisfactionRule is called when entering the multiPrefixForDegreeSatisfactionRule production.
	EnterMultiPrefixForDegreeSatisfactionRule(c *MultiPrefixForDegreeSatisfactionRuleContext)

	// EnterStudentDegreeSatisfactionRule is called when entering the studentDegreeSatisfactionRule production.
	EnterStudentDegreeSatisfactionRule(c *StudentDegreeSatisfactionRuleContext)

	// EnterMathDegreeSatisfactionRule is called when entering the mathDegreeSatisfactionRule production.
	EnterMathDegreeSatisfactionRule(c *MathDegreeSatisfactionRuleContext)

	// EnterNamedDegreeSatisfactionRule is called when entering the namedDegreeSatisfactionRule production.
	EnterNamedDegreeSatisfactionRule(c *NamedDegreeSatisfactionRuleContext)

	// EnterMultiPrefixDegreeSatisfactionRule is called when entering the multiPrefixDegreeSatisfactionRule production.
	EnterMultiPrefixDegreeSatisfactionRule(c *MultiPrefixDegreeSatisfactionRuleContext)

	// EnterCredit_for_rule is called when entering the credit_for_rule production.
	EnterCredit_for_rule(c *Credit_for_ruleContext)

	// EnterPrefixLivingLearningRule is called when entering the prefixLivingLearningRule production.
	EnterPrefixLivingLearningRule(c *PrefixLivingLearningRuleContext)

	// EnterNamedLivingLearningRule is called when entering the namedLivingLearningRule production.
	EnterNamedLivingLearningRule(c *NamedLivingLearningRuleContext)

	// EnterSchool_rule is called when entering the school_rule production.
	EnterSchool_rule(c *School_ruleContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitSentence is called when exiting the sentence production.
	ExitSentence(c *SentenceContext)

	// ExitSameAsParenReq is called when exiting the sameAsParenReq production.
	ExitSameAsParenReq(c *SameAsParenReqContext)

	// ExitPreOrCoReq is called when exiting the preOrCoReq production.
	ExitPreOrCoReq(c *PreOrCoReqContext)

	// ExitPrereqAndCoreqReq is called when exiting the prereqAndCoreqReq production.
	ExitPrereqAndCoreqReq(c *PrereqAndCoreqReqContext)

	// ExitComputerScholarsReq is called when exiting the computerScholarsReq production.
	ExitComputerScholarsReq(c *ComputerScholarsReqContext)

	// ExitMajorReq is called when exiting the majorReq production.
	ExitMajorReq(c *MajorReqContext)

	// ExitDegreeSatisfactionReq is called when exiting the degreeSatisfactionReq production.
	ExitDegreeSatisfactionReq(c *DegreeSatisfactionReqContext)

	// ExitRepeatLimitTimesReq is called when exiting the repeatLimitTimesReq production.
	ExitRepeatLimitTimesReq(c *RepeatLimitTimesReqContext)

	// ExitAcademicPlanReq is called when exiting the academicPlanReq production.
	ExitAcademicPlanReq(c *AcademicPlanReqContext)

	// ExitRepeatLimitHoursReq is called when exiting the repeatLimitHoursReq production.
	ExitRepeatLimitHoursReq(c *RepeatLimitHoursReqContext)

	// ExitAppendAcademicPlanReq is called when exiting the appendAcademicPlanReq production.
	ExitAppendAcademicPlanReq(c *AppendAcademicPlanReqContext)

	// ExitRepeatReq is called when exiting the repeatReq production.
	ExitRepeatReq(c *RepeatReqContext)

	// ExitExprReq is called when exiting the exprReq production.
	ExitExprReq(c *ExprReqContext)

	// ExitExactCoreqNoticeReq is called when exiting the exactCoreqNoticeReq production.
	ExitExactCoreqNoticeReq(c *ExactCoreqNoticeReqContext)

	// ExitGpaRepeatReq is called when exiting the gpaRepeatReq production.
	ExitGpaRepeatReq(c *GpaRepeatReqContext)

	// ExitSchoolReq is called when exiting the schoolReq production.
	ExitSchoolReq(c *SchoolReqContext)

	// ExitCoreqReq is called when exiting the coreqReq production.
	ExitCoreqReq(c *CoreqReqContext)

	// ExitPrereqReq is called when exiting the prereqReq production.
	ExitPrereqReq(c *PrereqReqContext)

	// ExitCreditForReq is called when exiting the creditForReq production.
	ExitCreditForReq(c *CreditForReqContext)

	// ExitSameAsReq is called when exiting the sameAsReq production.
	ExitSameAsReq(c *SameAsReqContext)

	// ExitExcludeNoticeReq is called when exiting the excludeNoticeReq production.
	ExitExcludeNoticeReq(c *ExcludeNoticeReqContext)

	// ExitLivingLearningReq is called when exiting the livingLearningReq production.
	ExitLivingLearningReq(c *LivingLearningReqContext)

	// ExitGroupExpr is called when exiting the groupExpr production.
	ExitGroupExpr(c *GroupExprContext)

	// ExitCourseExpr is called when exiting the courseExpr production.
	ExitCourseExpr(c *CourseExprContext)

	// ExitAlternativeExpr is called when exiting the alternativeExpr production.
	ExitAlternativeExpr(c *AlternativeExprContext)

	// ExitUpperDivisionStandingExpr is called when exiting the upperDivisionStandingExpr production.
	ExitUpperDivisionStandingExpr(c *UpperDivisionStandingExprContext)

	// ExitSemesterCreditHoursExpr is called when exiting the semesterCreditHoursExpr production.
	ExitSemesterCreditHoursExpr(c *SemesterCreditHoursExprContext)

	// ExitResearchExpr is called when exiting the researchExpr production.
	ExitResearchExpr(c *ResearchExprContext)

	// ExitComputerScholarsExpr is called when exiting the computerScholarsExpr production.
	ExitComputerScholarsExpr(c *ComputerScholarsExprContext)

	// ExitCompleteNExpr is called when exiting the completeNExpr production.
	ExitCompleteNExpr(c *CompleteNExprContext)

	// ExitAnyCoreExpr is called when exiting the anyCoreExpr production.
	ExitAnyCoreExpr(c *AnyCoreExprContext)

	// ExitHonorsExpr is called when exiting the honorsExpr production.
	ExitHonorsExpr(c *HonorsExprContext)

	// ExitGradeExpr is called when exiting the gradeExpr production.
	ExitGradeExpr(c *GradeExprContext)

	// ExitDepartmentConsentExpr is called when exiting the departmentConsentExpr production.
	ExitDepartmentConsentExpr(c *DepartmentConsentExprContext)

	// ExitParenExpr is called when exiting the parenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitAleksScoreExpr is called when exiting the aleksScoreExpr production.
	ExitAleksScoreExpr(c *AleksScoreExprContext)

	// ExitAmpersandExpr is called when exiting the ampersandExpr production.
	ExitAmpersandExpr(c *AmpersandExprContext)

	// ExitLivingLearningExpr is called when exiting the livingLearningExpr production.
	ExitLivingLearningExpr(c *LivingLearningExprContext)

	// ExitUteachConsentExpr is called when exiting the uteachConsentExpr production.
	ExitUteachConsentExpr(c *UteachConsentExprContext)

	// ExitGoodAcademicStandingExpr is called when exiting the goodAcademicStandingExpr production.
	ExitGoodAcademicStandingExpr(c *GoodAcademicStandingExprContext)

	// ExitCoreExpr is called when exiting the coreExpr production.
	ExitCoreExpr(c *CoreExprContext)

	// ExitGpaExpr is called when exiting the gpaExpr production.
	ExitGpaExpr(c *GpaExprContext)

	// ExitMajorExpr is called when exiting the majorExpr production.
	ExitMajorExpr(c *MajorExprContext)

	// ExitEquivalentExpr is called when exiting the equivalentExpr production.
	ExitEquivalentExpr(c *EquivalentExprContext)

	// ExitInstructorConsentExpr is called when exiting the instructorConsentExpr production.
	ExitInstructorConsentExpr(c *InstructorConsentExprContext)

	// ExitMinimumHoursExpr is called when exiting the minimumHoursExpr production.
	ExitMinimumHoursExpr(c *MinimumHoursExprContext)

	// ExitRepeatLimitHoursExpr is called when exiting the repeatLimitHoursExpr production.
	ExitRepeatLimitHoursExpr(c *RepeatLimitHoursExprContext)

	// ExitRepeatRuleExpr is called when exiting the repeatRuleExpr production.
	ExitRepeatRuleExpr(c *RepeatRuleExprContext)

	// ExitApScoreExpr is called when exiting the apScoreExpr production.
	ExitApScoreExpr(c *ApScoreExprContext)

	// ExitAnyMajorCourseExpr is called when exiting the anyMajorCourseExpr production.
	ExitAnyMajorCourseExpr(c *AnyMajorCourseExprContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitUpperDivisionHoursExpr is called when exiting the upperDivisionHoursExpr production.
	ExitUpperDivisionHoursExpr(c *UpperDivisionHoursExprContext)

	// ExitUpperDivisionClassesExpr is called when exiting the upperDivisionClassesExpr production.
	ExitUpperDivisionClassesExpr(c *UpperDivisionClassesExprContext)

	// ExitPlacementTestExpr is called when exiting the placementTestExpr production.
	ExitPlacementTestExpr(c *PlacementTestExprContext)

	// ExitConcurrentEnrollmentExpr is called when exiting the concurrentEnrollmentExpr production.
	ExitConcurrentEnrollmentExpr(c *ConcurrentEnrollmentExprContext)

	// ExitGradeLevelStandingExpr is called when exiting the gradeLevelStandingExpr production.
	ExitGradeLevelStandingExpr(c *GradeLevelStandingExprContext)

	// ExitDegreeExpr is called when exiting the degreeExpr production.
	ExitDegreeExpr(c *DegreeExprContext)

	// ExitGraduateStandingExpr is called when exiting the graduateStandingExpr production.
	ExitGraduateStandingExpr(c *GraduateStandingExprContext)

	// ExitExactSectionExpr is called when exiting the exactSectionExpr production.
	ExitExactSectionExpr(c *ExactSectionExprContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitSimpleCourse is called when exiting the simpleCourse production.
	ExitSimpleCourse(c *SimpleCourseContext)

	// ExitParenCourse is called when exiting the parenCourse production.
	ExitParenCourse(c *ParenCourseContext)

	// ExitCrossListedCourse is called when exiting the crossListedCourse production.
	ExitCrossListedCourse(c *CrossListedCourseContext)

	// ExitFullCourseList is called when exiting the fullCourseList production.
	ExitFullCourseList(c *FullCourseListContext)

	// ExitShorthandCourseList is called when exiting the shorthandCourseList production.
	ExitShorthandCourseList(c *ShorthandCourseListContext)

	// ExitTitle is called when exiting the title production.
	ExitTitle(c *TitleContext)

	// ExitSimpleGradeCondition is called when exiting the simpleGradeCondition production.
	ExitSimpleGradeCondition(c *SimpleGradeConditionContext)

	// ExitGpaGradeCondition is called when exiting the gpaGradeCondition production.
	ExitGpaGradeCondition(c *GpaGradeConditionContext)

	// ExitParenGradeCondition is called when exiting the parenGradeCondition production.
	ExitParenGradeCondition(c *ParenGradeConditionContext)

	// ExitGradeAtLeastCondition is called when exiting the gradeAtLeastCondition production.
	ExitGradeAtLeastCondition(c *GradeAtLeastConditionContext)

	// ExitGrade_course_list is called when exiting the grade_course_list production.
	ExitGrade_course_list(c *Grade_course_listContext)

	// ExitCourseAlternativeCondition is called when exiting the courseAlternativeCondition production.
	ExitCourseAlternativeCondition(c *CourseAlternativeConditionContext)

	// ExitParenAlternativeCondition is called when exiting the parenAlternativeCondition production.
	ExitParenAlternativeCondition(c *ParenAlternativeConditionContext)

	// ExitGradeAlternativeCondition is called when exiting the gradeAlternativeCondition production.
	ExitGradeAlternativeCondition(c *GradeAlternativeConditionContext)

	// ExitGradeLevelStandingCondition is called when exiting the gradeLevelStandingCondition production.
	ExitGradeLevelStandingCondition(c *GradeLevelStandingConditionContext)

	// ExitGradeLevelMajorStandingCondition is called when exiting the gradeLevelMajorStandingCondition production.
	ExitGradeLevelMajorStandingCondition(c *GradeLevelMajorStandingConditionContext)

	// ExitMinimumGradeLevelStandingCondition is called when exiting the minimumGradeLevelStandingCondition production.
	ExitMinimumGradeLevelStandingCondition(c *MinimumGradeLevelStandingConditionContext)

	// ExitAtLeastGradeLevelStandingCondition is called when exiting the atLeastGradeLevelStandingCondition production.
	ExitAtLeastGradeLevelStandingCondition(c *AtLeastGradeLevelStandingConditionContext)

	// ExitPrefixGradeLevelStandingCondition is called when exiting the prefixGradeLevelStandingCondition production.
	ExitPrefixGradeLevelStandingCondition(c *PrefixGradeLevelStandingConditionContext)

	// ExitGraduateStandingInCondition is called when exiting the graduateStandingInCondition production.
	ExitGraduateStandingInCondition(c *GraduateStandingInConditionContext)

	// ExitGraduateLevelStandingCondition is called when exiting the graduateLevelStandingCondition production.
	ExitGraduateLevelStandingCondition(c *GraduateLevelStandingConditionContext)

	// ExitUniversityGpaCondition is called when exiting the universityGpaCondition production.
	ExitUniversityGpaCondition(c *UniversityGpaConditionContext)

	// ExitMinimumGpaCondition is called when exiting the minimumGpaCondition production.
	ExitMinimumGpaCondition(c *MinimumGpaConditionContext)

	// ExitGpaInCourseCondition is called when exiting the gpaInCourseCondition production.
	ExitGpaInCourseCondition(c *GpaInCourseConditionContext)

	// ExitPrefixMajorCondition is called when exiting the prefixMajorCondition production.
	ExitPrefixMajorCondition(c *PrefixMajorConditionContext)

	// ExitGradeLevelPrefixMajorCondition is called when exiting the gradeLevelPrefixMajorCondition production.
	ExitGradeLevelPrefixMajorCondition(c *GradeLevelPrefixMajorConditionContext)

	// ExitDegreeTypePrefixMajorCondition is called when exiting the degreeTypePrefixMajorCondition production.
	ExitDegreeTypePrefixMajorCondition(c *DegreeTypePrefixMajorConditionContext)

	// ExitNamedMajorCondition is called when exiting the namedMajorCondition production.
	ExitNamedMajorCondition(c *NamedMajorConditionContext)

	// ExitNamedDegreeTypeMajorCondition is called when exiting the namedDegreeTypeMajorCondition production.
	ExitNamedDegreeTypeMajorCondition(c *NamedDegreeTypeMajorConditionContext)

	// ExitUndergraduateDegreeCondition is called when exiting the undergraduateDegreeCondition production.
	ExitUndergraduateDegreeCondition(c *UndergraduateDegreeConditionContext)

	// ExitBachelorsOrMastersCondition is called when exiting the bachelorsOrMastersCondition production.
	ExitBachelorsOrMastersCondition(c *BachelorsOrMastersConditionContext)

	// ExitDegree_list is called when exiting the degree_list production.
	ExitDegree_list(c *Degree_listContext)

	// ExitDegree is called when exiting the degree production.
	ExitDegree(c *DegreeContext)

	// ExitCoreCondition is called when exiting the coreCondition production.
	ExitCoreCondition(c *CoreConditionContext)

	// ExitAnyCoreSCHCondition is called when exiting the anyCoreSCHCondition production.
	ExitAnyCoreSCHCondition(c *AnyCoreSCHConditionContext)

	// ExitCompletionOfCoreCondition is called when exiting the completionOfCoreCondition production.
	ExitCompletionOfCoreCondition(c *CompletionOfCoreConditionContext)

	// ExitSemesterCreditHoursCondition is called when exiting the semesterCreditHoursCondition production.
	ExitSemesterCreditHoursCondition(c *SemesterCreditHoursConditionContext)

	// ExitMinimumHoursCondition is called when exiting the minimumHoursCondition production.
	ExitMinimumHoursCondition(c *MinimumHoursConditionContext)

	// ExitUpperDivisionSCHCondition is called when exiting the upperDivisionSCHCondition production.
	ExitUpperDivisionSCHCondition(c *UpperDivisionSCHConditionContext)

	// ExitUpperDivisionCreditsCondition is called when exiting the upperDivisionCreditsCondition production.
	ExitUpperDivisionCreditsCondition(c *UpperDivisionCreditsConditionContext)

	// ExitUpperDivisionClassesCondition is called when exiting the upperDivisionClassesCondition production.
	ExitUpperDivisionClassesCondition(c *UpperDivisionClassesConditionContext)

	// ExitFourThousandLevelCondition is called when exiting the fourThousandLevelCondition production.
	ExitFourThousandLevelCondition(c *FourThousandLevelConditionContext)

	// ExitResearchCondition is called when exiting the researchCondition production.
	ExitResearchCondition(c *ResearchConditionContext)

	// ExitCompleteNOfFollowingCondition is called when exiting the completeNOfFollowingCondition production.
	ExitCompleteNOfFollowingCondition(c *CompleteNOfFollowingConditionContext)

	// ExitCompleteNFromFollowingCondition is called when exiting the completeNFromFollowingCondition production.
	ExitCompleteNFromFollowingCondition(c *CompleteNFromFollowingConditionContext)

	// ExitCompleteHoursFromFollowingCondition is called when exiting the completeHoursFromFollowingCondition production.
	ExitCompleteHoursFromFollowingCondition(c *CompleteHoursFromFollowingConditionContext)

	// ExitPlacementScoreComparisonCondition is called when exiting the placementScoreComparisonCondition production.
	ExitPlacementScoreComparisonCondition(c *PlacementScoreComparisonConditionContext)

	// ExitPlacementScoreRangeCondition is called when exiting the placementScoreRangeCondition production.
	ExitPlacementScoreRangeCondition(c *PlacementScoreRangeConditionContext)

	// ExitPlacementScoreMinimumCondition is called when exiting the placementScoreMinimumCondition production.
	ExitPlacementScoreMinimumCondition(c *PlacementScoreMinimumConditionContext)

	// ExitPlacement_test_name is called when exiting the placement_test_name production.
	ExitPlacement_test_name(c *Placement_test_nameContext)

	// ExitApScoreCondition is called when exiting the apScoreCondition production.
	ExitApScoreCondition(c *ApScoreConditionContext)

	// ExitAleksScoreCondition is called when exiting the aleksScoreCondition production.
	ExitAleksScoreCondition(c *AleksScoreConditionContext)

	// ExitBothHonorsCondition is called when exiting the bothHonorsCondition production.
	ExitBothHonorsCondition(c *BothHonorsConditionContext)

	// ExitSingleHonorsCondition is called when exiting the singleHonorsCondition production.
	ExitSingleHonorsCondition(c *SingleHonorsConditionContext)

	// ExitGroupCondition is called when exiting the groupCondition production.
	ExitGroupCondition(c *GroupConditionContext)

	// ExitGroup is called when exiting the group production.
	ExitGroup(c *GroupContext)

	// ExitConcurrentEnrollmentCondition is called when exiting the concurrentEnrollmentCondition production.
	ExitConcurrentEnrollmentCondition(c *ConcurrentEnrollmentConditionContext)

	// ExitExactSectionCondition is called when exiting the exactSectionCondition production.
	ExitExactSectionCondition(c *ExactSectionConditionContext)

	// ExitWorkshopSectionCondition is called when exiting the workshopSectionCondition production.
	ExitWorkshopSectionCondition(c *WorkshopSectionConditionContext)

	// ExitAnyPreviousMajorCourseCondition is called when exiting the anyPreviousMajorCourseCondition production.
	ExitAnyPreviousMajorCourseCondition(c *AnyPreviousMajorCourseConditionContext)

	// ExitAcademic_plan_condition is called when exiting the academic_plan_condition production.
	ExitAcademic_plan_condition(c *Academic_plan_conditionContext)

	// ExitCourseRepeatRule is called when exiting the courseRepeatRule production.
	ExitCourseRepeatRule(c *CourseRepeatRuleContext)

	// ExitInternshipRepeatRule is called when exiting the internshipRepeatRule production.
	ExitInternshipRepeatRule(c *InternshipRepeatRuleContext)

	// ExitBareRepeatRule is called when exiting the bareRepeatRule production.
	ExitBareRepeatRule(c *BareRepeatRuleContext)

	// ExitRepeatMaxHoursRule is called when exiting the repeatMaxHoursRule production.
	ExitRepeatMaxHoursRule(c *RepeatMaxHoursRuleContext)

	// ExitRepeatHoursMaxSuffixRule is called when exiting the repeatHoursMaxSuffixRule production.
	ExitRepeatHoursMaxSuffixRule(c *RepeatHoursMaxSuffixRuleContext)

	// ExitCourseRepeatMaxHoursRule is called when exiting the courseRepeatMaxHoursRule production.
	ExitCourseRepeatMaxHoursRule(c *CourseRepeatMaxHoursRuleContext)

	// ExitCombinedRepeatMaxHoursRule is called when exiting the combinedRepeatMaxHoursRule production.
	ExitCombinedRepeatMaxHoursRule(c *CombinedRepeatMaxHoursRuleContext)

	// ExitTopicsVaryRepeatRule is called when exiting the topicsVaryRepeatRule production.
	ExitTopicsVaryRepeatRule(c *TopicsVaryRepeatRuleContext)

	// ExitCourseRepeatLimitRule is called when exiting the courseRepeatLimitRule production.
	ExitCourseRepeatLimitRule(c *CourseRepeatLimitRuleContext)

	// ExitRepeatUpToTimesRule is called when exiting the repeatUpToTimesRule production.
	ExitRepeatUpToTimesRule(c *RepeatUpToTimesRuleContext)

	// ExitRepeatMaxTimesRule is called when exiting the repeatMaxTimesRule production.
	ExitRepeatMaxTimesRule(c *RepeatMaxTimesRuleContext)

	// ExitAndRepeateExpr is called when exiting the andRepeateExpr production.
	ExitAndRepeateExpr(c *AndRepeateExprContext)

	// ExitAlternativeRepeateExpr is called when exiting the alternativeRepeateExpr production.
	ExitAlternativeRepeateExpr(c *AlternativeRepeateExprContext)

	// ExitCourseRepeateExpr is called when exiting the courseRepeateExpr production.
	ExitCourseRepeateExpr(c *CourseRepeateExprContext)

	// ExitGpa_repeate_rule is called when exiting the gpa_repeate_rule production.
	ExitGpa_repeate_rule(c *Gpa_repeate_ruleContext)

	// ExitSchoolDegreeSatisfactionRule is called when exiting the schoolDegreeSatisfactionRule production.
	ExitSchoolDegreeSatisfactionRule(c *SchoolDegreeSatisfactionRuleContext)

	// ExitPrefixDegreeSatisfactionRule is called when exiting the prefixDegreeSatisfactionRule production.
	ExitPrefixDegreeSatisfactionRule(c *PrefixDegreeSatisfactionRuleContext)

	// ExitOfMultiPrefixDegreeSatisfactionRule is called when exiting the ofMultiPrefixDegreeSatisfactionRule production.
	ExitOfMultiPrefixDegreeSatisfactionRule(c *OfMultiPrefixDegreeSatisfactionRuleContext)

	// ExitElectivesDegreeSatisfactionRule is called when exiting the electivesDegreeSatisfactionRule production.
	ExitElectivesDegreeSatisfactionRule(c *ElectivesDegreeSatisfactionRuleContext)

	// ExitSchoolsDegreeSatisfactionRule is called when exiting the schoolsDegreeSatisfactionRule production.
	ExitSchoolsDegreeSatisfactionRule(c *SchoolsDegreeSatisfactionRuleContext)

	// ExitMultiPrefixForDegreeSatisfactionRule is called when exiting the multiPrefixForDegreeSatisfactionRule production.
	ExitMultiPrefixForDegreeSatisfactionRule(c *MultiPrefixForDegreeSatisfactionRuleContext)

	// ExitStudentDegreeSatisfactionRule is called when exiting the studentDegreeSatisfactionRule production.
	ExitStudentDegreeSatisfactionRule(c *StudentDegreeSatisfactionRuleContext)

	// ExitMathDegreeSatisfactionRule is called when exiting the mathDegreeSatisfactionRule production.
	ExitMathDegreeSatisfactionRule(c *MathDegreeSatisfactionRuleContext)

	// ExitNamedDegreeSatisfactionRule is called when exiting the namedDegreeSatisfactionRule production.
	ExitNamedDegreeSatisfactionRule(c *NamedDegreeSatisfactionRuleContext)

	// ExitMultiPrefixDegreeSatisfactionRule is called when exiting the multiPrefixDegreeSatisfactionRule production.
	ExitMultiPrefixDegreeSatisfactionRule(c *MultiPrefixDegreeSatisfactionRuleContext)

	// ExitCredit_for_rule is called when exiting the credit_for_rule production.
	ExitCredit_for_rule(c *Credit_for_ruleContext)

	// ExitPrefixLivingLearningRule is called when exiting the prefixLivingLearningRule production.
	ExitPrefixLivingLearningRule(c *PrefixLivingLearningRuleContext)

	// ExitNamedLivingLearningRule is called when exiting the namedLivingLearningRule production.
	ExitNamedLivingLearningRule(c *NamedLivingLearningRuleContext)

	// ExitSchool_rule is called when exiting the school_rule production.
	ExitSchool_rule(c *School_ruleContext)
}
