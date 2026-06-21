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

	// EnterSemesterCreditHoursExpr is called when entering the semesterCreditHoursExpr production.
	EnterSemesterCreditHoursExpr(c *SemesterCreditHoursExprContext)

	// EnterResearchExpr is called when entering the researchExpr production.
	EnterResearchExpr(c *ResearchExprContext)

	// EnterCompleteNExpr is called when entering the completeNExpr production.
	EnterCompleteNExpr(c *CompleteNExprContext)

	// EnterConsentExpr is called when entering the consentExpr production.
	EnterConsentExpr(c *ConsentExprContext)

	// EnterAnyCoreExpr is called when entering the anyCoreExpr production.
	EnterAnyCoreExpr(c *AnyCoreExprContext)

	// EnterGradeExpr is called when entering the gradeExpr production.
	EnterGradeExpr(c *GradeExprContext)

	// EnterParenExpr is called when entering the parenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterAleksScoreExpr is called when entering the aleksScoreExpr production.
	EnterAleksScoreExpr(c *AleksScoreExprContext)

	// EnterAmpersandExpr is called when entering the ampersandExpr production.
	EnterAmpersandExpr(c *AmpersandExprContext)

	// EnterLivingLearningExpr is called when entering the livingLearningExpr production.
	EnterLivingLearningExpr(c *LivingLearningExprContext)

	// EnterCoreExpr is called when entering the coreExpr production.
	EnterCoreExpr(c *CoreExprContext)

	// EnterGpaExpr is called when entering the gpaExpr production.
	EnterGpaExpr(c *GpaExprContext)

	// EnterMajorExpr is called when entering the majorExpr production.
	EnterMajorExpr(c *MajorExprContext)

	// EnterEquivalentExpr is called when entering the equivalentExpr production.
	EnterEquivalentExpr(c *EquivalentExprContext)

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

	// EnterStandingExpr is called when entering the standingExpr production.
	EnterStandingExpr(c *StandingExprContext)

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

	// EnterDegree_atom is called when entering the degree_atom production.
	EnterDegree_atom(c *Degree_atomContext)

	// EnterDegree is called when entering the degree production.
	EnterDegree(c *DegreeContext)

	// EnterDegree_list is called when entering the degree_list production.
	EnterDegree_list(c *Degree_listContext)

	// EnterInstructorConsentCondition is called when entering the instructorConsentCondition production.
	EnterInstructorConsentCondition(c *InstructorConsentConditionContext)

	// EnterDepartmentConsentCondition is called when entering the departmentConsentCondition production.
	EnterDepartmentConsentCondition(c *DepartmentConsentConditionContext)

	// EnterUteachConsentCondition is called when entering the uteachConsentCondition production.
	EnterUteachConsentCondition(c *UteachConsentConditionContext)

	// EnterUpperDivisionStandingCondition is called when entering the upperDivisionStandingCondition production.
	EnterUpperDivisionStandingCondition(c *UpperDivisionStandingConditionContext)

	// EnterGoodAcademicStandingCondition is called when entering the goodAcademicStandingCondition production.
	EnterGoodAcademicStandingCondition(c *GoodAcademicStandingConditionContext)

	// EnterSimpleGradeCondition is called when entering the simpleGradeCondition production.
	EnterSimpleGradeCondition(c *SimpleGradeConditionContext)

	// EnterGpaGradeCondition is called when entering the gpaGradeCondition production.
	EnterGpaGradeCondition(c *GpaGradeConditionContext)

	// EnterGradeListCondition is called when entering the gradeListCondition production.
	EnterGradeListCondition(c *GradeListConditionContext)

	// EnterGradeAtLeastCondition is called when entering the gradeAtLeastCondition production.
	EnterGradeAtLeastCondition(c *GradeAtLeastConditionContext)

	// EnterEitherGradeCourseList is called when entering the eitherGradeCourseList production.
	EnterEitherGradeCourseList(c *EitherGradeCourseListContext)

	// EnterAllGradeCourseList is called when entering the allGradeCourseList production.
	EnterAllGradeCourseList(c *AllGradeCourseListContext)

	// EnterParenGradeCourseList is called when entering the parenGradeCourseList production.
	EnterParenGradeCourseList(c *ParenGradeCourseListContext)

	// EnterCourseAlternativeCondition is called when entering the courseAlternativeCondition production.
	EnterCourseAlternativeCondition(c *CourseAlternativeConditionContext)

	// EnterGradeCourseListAlternativeCondition is called when entering the gradeCourseListAlternativeCondition production.
	EnterGradeCourseListAlternativeCondition(c *GradeCourseListAlternativeConditionContext)

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

	// EnterCoreCondition is called when entering the coreCondition production.
	EnterCoreCondition(c *CoreConditionContext)

	// EnterAnyCoreSCHCondition is called when entering the anyCoreSCHCondition production.
	EnterAnyCoreSCHCondition(c *AnyCoreSCHConditionContext)

	// EnterSemesterCreditHoursCondition is called when entering the semesterCreditHoursCondition production.
	EnterSemesterCreditHoursCondition(c *SemesterCreditHoursConditionContext)

	// EnterMinimumHoursCondition is called when entering the minimumHoursCondition production.
	EnterMinimumHoursCondition(c *MinimumHoursConditionContext)

	// EnterMinimumHoursOfCondition is called when entering the minimumHoursOfCondition production.
	EnterMinimumHoursOfCondition(c *MinimumHoursOfConditionContext)

	// EnterMinimumHoursFromCondition is called when entering the minimumHoursFromCondition production.
	EnterMinimumHoursFromCondition(c *MinimumHoursFromConditionContext)

	// EnterUpperDivisionSCHCondition is called when entering the upperDivisionSCHCondition production.
	EnterUpperDivisionSCHCondition(c *UpperDivisionSCHConditionContext)

	// EnterUpperDivisionCountCondition is called when entering the upperDivisionCountCondition production.
	EnterUpperDivisionCountCondition(c *UpperDivisionCountConditionContext)

	// EnterUpperDivisionSingleCondition is called when entering the upperDivisionSingleCondition production.
	EnterUpperDivisionSingleCondition(c *UpperDivisionSingleConditionContext)

	// EnterResearchCondition is called when entering the researchCondition production.
	EnterResearchCondition(c *ResearchConditionContext)

	// EnterCompleteNOfFollowingCondition is called when entering the completeNOfFollowingCondition production.
	EnterCompleteNOfFollowingCondition(c *CompleteNOfFollowingConditionContext)

	// EnterCompleteNFromFollowingCondition is called when entering the completeNFromFollowingCondition production.
	EnterCompleteNFromFollowingCondition(c *CompleteNFromFollowingConditionContext)

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

	// EnterBothGroupCondition is called when entering the bothGroupCondition production.
	EnterBothGroupCondition(c *BothGroupConditionContext)

	// EnterGroupListCondition is called when entering the groupListCondition production.
	EnterGroupListCondition(c *GroupListConditionContext)

	// EnterSingleGroupCondition is called when entering the singleGroupCondition production.
	EnterSingleGroupCondition(c *SingleGroupConditionContext)

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

	// EnterAcademicPlanCondition is called when entering the academicPlanCondition production.
	EnterAcademicPlanCondition(c *AcademicPlanConditionContext)

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

	// EnterCourseRepeatLimitRule is called when entering the courseRepeatLimitRule production.
	EnterCourseRepeatLimitRule(c *CourseRepeatLimitRuleContext)

	// EnterRepeatUpToTimesRule is called when entering the repeatUpToTimesRule production.
	EnterRepeatUpToTimesRule(c *RepeatUpToTimesRuleContext)

	// EnterRepeatMaxTimesRule is called when entering the repeatMaxTimesRule production.
	EnterRepeatMaxTimesRule(c *RepeatMaxTimesRuleContext)

	// EnterGpaRepeatRule is called when entering the gpaRepeatRule production.
	EnterGpaRepeatRule(c *GpaRepeatRuleContext)

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

	// EnterCreditForRule is called when entering the creditForRule production.
	EnterCreditForRule(c *CreditForRuleContext)

	// EnterOrCreditExpr is called when entering the orCreditExpr production.
	EnterOrCreditExpr(c *OrCreditExprContext)

	// EnterAmpersandCreditExpr is called when entering the ampersandCreditExpr production.
	EnterAmpersandCreditExpr(c *AmpersandCreditExprContext)

	// EnterCourseCreditExpr is called when entering the courseCreditExpr production.
	EnterCourseCreditExpr(c *CourseCreditExprContext)

	// EnterParenCreditExpr is called when entering the parenCreditExpr production.
	EnterParenCreditExpr(c *ParenCreditExprContext)

	// EnterAndCreditExpr is called when entering the andCreditExpr production.
	EnterAndCreditExpr(c *AndCreditExprContext)

	// EnterPrefixLivingLearningRule is called when entering the prefixLivingLearningRule production.
	EnterPrefixLivingLearningRule(c *PrefixLivingLearningRuleContext)

	// EnterDegreeLivingLearningRule is called when entering the degreeLivingLearningRule production.
	EnterDegreeLivingLearningRule(c *DegreeLivingLearningRuleContext)

	// EnterSchoolRule is called when entering the schoolRule production.
	EnterSchoolRule(c *SchoolRuleContext)

	// EnterSameAsRule is called when entering the sameAsRule production.
	EnterSameAsRule(c *SameAsRuleContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitSentence is called when exiting the sentence production.
	ExitSentence(c *SentenceContext)

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

	// ExitSemesterCreditHoursExpr is called when exiting the semesterCreditHoursExpr production.
	ExitSemesterCreditHoursExpr(c *SemesterCreditHoursExprContext)

	// ExitResearchExpr is called when exiting the researchExpr production.
	ExitResearchExpr(c *ResearchExprContext)

	// ExitCompleteNExpr is called when exiting the completeNExpr production.
	ExitCompleteNExpr(c *CompleteNExprContext)

	// ExitConsentExpr is called when exiting the consentExpr production.
	ExitConsentExpr(c *ConsentExprContext)

	// ExitAnyCoreExpr is called when exiting the anyCoreExpr production.
	ExitAnyCoreExpr(c *AnyCoreExprContext)

	// ExitGradeExpr is called when exiting the gradeExpr production.
	ExitGradeExpr(c *GradeExprContext)

	// ExitParenExpr is called when exiting the parenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitAleksScoreExpr is called when exiting the aleksScoreExpr production.
	ExitAleksScoreExpr(c *AleksScoreExprContext)

	// ExitAmpersandExpr is called when exiting the ampersandExpr production.
	ExitAmpersandExpr(c *AmpersandExprContext)

	// ExitLivingLearningExpr is called when exiting the livingLearningExpr production.
	ExitLivingLearningExpr(c *LivingLearningExprContext)

	// ExitCoreExpr is called when exiting the coreExpr production.
	ExitCoreExpr(c *CoreExprContext)

	// ExitGpaExpr is called when exiting the gpaExpr production.
	ExitGpaExpr(c *GpaExprContext)

	// ExitMajorExpr is called when exiting the majorExpr production.
	ExitMajorExpr(c *MajorExprContext)

	// ExitEquivalentExpr is called when exiting the equivalentExpr production.
	ExitEquivalentExpr(c *EquivalentExprContext)

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

	// ExitStandingExpr is called when exiting the standingExpr production.
	ExitStandingExpr(c *StandingExprContext)

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

	// ExitDegree_atom is called when exiting the degree_atom production.
	ExitDegree_atom(c *Degree_atomContext)

	// ExitDegree is called when exiting the degree production.
	ExitDegree(c *DegreeContext)

	// ExitDegree_list is called when exiting the degree_list production.
	ExitDegree_list(c *Degree_listContext)

	// ExitInstructorConsentCondition is called when exiting the instructorConsentCondition production.
	ExitInstructorConsentCondition(c *InstructorConsentConditionContext)

	// ExitDepartmentConsentCondition is called when exiting the departmentConsentCondition production.
	ExitDepartmentConsentCondition(c *DepartmentConsentConditionContext)

	// ExitUteachConsentCondition is called when exiting the uteachConsentCondition production.
	ExitUteachConsentCondition(c *UteachConsentConditionContext)

	// ExitUpperDivisionStandingCondition is called when exiting the upperDivisionStandingCondition production.
	ExitUpperDivisionStandingCondition(c *UpperDivisionStandingConditionContext)

	// ExitGoodAcademicStandingCondition is called when exiting the goodAcademicStandingCondition production.
	ExitGoodAcademicStandingCondition(c *GoodAcademicStandingConditionContext)

	// ExitSimpleGradeCondition is called when exiting the simpleGradeCondition production.
	ExitSimpleGradeCondition(c *SimpleGradeConditionContext)

	// ExitGpaGradeCondition is called when exiting the gpaGradeCondition production.
	ExitGpaGradeCondition(c *GpaGradeConditionContext)

	// ExitGradeListCondition is called when exiting the gradeListCondition production.
	ExitGradeListCondition(c *GradeListConditionContext)

	// ExitGradeAtLeastCondition is called when exiting the gradeAtLeastCondition production.
	ExitGradeAtLeastCondition(c *GradeAtLeastConditionContext)

	// ExitEitherGradeCourseList is called when exiting the eitherGradeCourseList production.
	ExitEitherGradeCourseList(c *EitherGradeCourseListContext)

	// ExitAllGradeCourseList is called when exiting the allGradeCourseList production.
	ExitAllGradeCourseList(c *AllGradeCourseListContext)

	// ExitParenGradeCourseList is called when exiting the parenGradeCourseList production.
	ExitParenGradeCourseList(c *ParenGradeCourseListContext)

	// ExitCourseAlternativeCondition is called when exiting the courseAlternativeCondition production.
	ExitCourseAlternativeCondition(c *CourseAlternativeConditionContext)

	// ExitGradeCourseListAlternativeCondition is called when exiting the gradeCourseListAlternativeCondition production.
	ExitGradeCourseListAlternativeCondition(c *GradeCourseListAlternativeConditionContext)

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

	// ExitCoreCondition is called when exiting the coreCondition production.
	ExitCoreCondition(c *CoreConditionContext)

	// ExitAnyCoreSCHCondition is called when exiting the anyCoreSCHCondition production.
	ExitAnyCoreSCHCondition(c *AnyCoreSCHConditionContext)

	// ExitSemesterCreditHoursCondition is called when exiting the semesterCreditHoursCondition production.
	ExitSemesterCreditHoursCondition(c *SemesterCreditHoursConditionContext)

	// ExitMinimumHoursCondition is called when exiting the minimumHoursCondition production.
	ExitMinimumHoursCondition(c *MinimumHoursConditionContext)

	// ExitMinimumHoursOfCondition is called when exiting the minimumHoursOfCondition production.
	ExitMinimumHoursOfCondition(c *MinimumHoursOfConditionContext)

	// ExitMinimumHoursFromCondition is called when exiting the minimumHoursFromCondition production.
	ExitMinimumHoursFromCondition(c *MinimumHoursFromConditionContext)

	// ExitUpperDivisionSCHCondition is called when exiting the upperDivisionSCHCondition production.
	ExitUpperDivisionSCHCondition(c *UpperDivisionSCHConditionContext)

	// ExitUpperDivisionCountCondition is called when exiting the upperDivisionCountCondition production.
	ExitUpperDivisionCountCondition(c *UpperDivisionCountConditionContext)

	// ExitUpperDivisionSingleCondition is called when exiting the upperDivisionSingleCondition production.
	ExitUpperDivisionSingleCondition(c *UpperDivisionSingleConditionContext)

	// ExitResearchCondition is called when exiting the researchCondition production.
	ExitResearchCondition(c *ResearchConditionContext)

	// ExitCompleteNOfFollowingCondition is called when exiting the completeNOfFollowingCondition production.
	ExitCompleteNOfFollowingCondition(c *CompleteNOfFollowingConditionContext)

	// ExitCompleteNFromFollowingCondition is called when exiting the completeNFromFollowingCondition production.
	ExitCompleteNFromFollowingCondition(c *CompleteNFromFollowingConditionContext)

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

	// ExitBothGroupCondition is called when exiting the bothGroupCondition production.
	ExitBothGroupCondition(c *BothGroupConditionContext)

	// ExitGroupListCondition is called when exiting the groupListCondition production.
	ExitGroupListCondition(c *GroupListConditionContext)

	// ExitSingleGroupCondition is called when exiting the singleGroupCondition production.
	ExitSingleGroupCondition(c *SingleGroupConditionContext)

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

	// ExitAcademicPlanCondition is called when exiting the academicPlanCondition production.
	ExitAcademicPlanCondition(c *AcademicPlanConditionContext)

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

	// ExitCourseRepeatLimitRule is called when exiting the courseRepeatLimitRule production.
	ExitCourseRepeatLimitRule(c *CourseRepeatLimitRuleContext)

	// ExitRepeatUpToTimesRule is called when exiting the repeatUpToTimesRule production.
	ExitRepeatUpToTimesRule(c *RepeatUpToTimesRuleContext)

	// ExitRepeatMaxTimesRule is called when exiting the repeatMaxTimesRule production.
	ExitRepeatMaxTimesRule(c *RepeatMaxTimesRuleContext)

	// ExitGpaRepeatRule is called when exiting the gpaRepeatRule production.
	ExitGpaRepeatRule(c *GpaRepeatRuleContext)

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

	// ExitCreditForRule is called when exiting the creditForRule production.
	ExitCreditForRule(c *CreditForRuleContext)

	// ExitOrCreditExpr is called when exiting the orCreditExpr production.
	ExitOrCreditExpr(c *OrCreditExprContext)

	// ExitAmpersandCreditExpr is called when exiting the ampersandCreditExpr production.
	ExitAmpersandCreditExpr(c *AmpersandCreditExprContext)

	// ExitCourseCreditExpr is called when exiting the courseCreditExpr production.
	ExitCourseCreditExpr(c *CourseCreditExprContext)

	// ExitParenCreditExpr is called when exiting the parenCreditExpr production.
	ExitParenCreditExpr(c *ParenCreditExprContext)

	// ExitAndCreditExpr is called when exiting the andCreditExpr production.
	ExitAndCreditExpr(c *AndCreditExprContext)

	// ExitPrefixLivingLearningRule is called when exiting the prefixLivingLearningRule production.
	ExitPrefixLivingLearningRule(c *PrefixLivingLearningRuleContext)

	// ExitDegreeLivingLearningRule is called when exiting the degreeLivingLearningRule production.
	ExitDegreeLivingLearningRule(c *DegreeLivingLearningRuleContext)

	// ExitSchoolRule is called when exiting the schoolRule production.
	ExitSchoolRule(c *SchoolRuleContext)

	// ExitSameAsRule is called when exiting the sameAsRule production.
	ExitSameAsRule(c *SameAsRuleContext)
}
