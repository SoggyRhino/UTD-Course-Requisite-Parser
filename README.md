# UTD Course Requisite Parser 

ANTLR4 based approach for parsing course requirements. 

Currently, the grammar successfully parses* ~97% of the unique, 25f course requirements. At this point adding the remaining 20 ish 
 inputs is not worth focusing on since they are all mostly unique edge cases and/or typos.

> \*Antlr parses them into an AST, not parsed yet into a useful JSON object. Currently, only 24% are done, but the hardest 
> visitors are done.

# Build 

Currently, this is a work in progress. I am building up the visitors from the leaf nodes, so everything will pretty
much not be functional until all visitors are completed. 

Antlr4 is uses grammar to generate go code for a parser, lexer and visitors. This project
ignores the first 2 and only uses visitors. Visitors are functions that are called when a node in the ast is visited.
They essentially act as a visitor for the ast work as reducers where we can take a node, say `course: PREFIX NUMBER`, and 
transform it into a useful object. 

Since antlr4 generates code, you must run the following command
```bash
cd .. scripts 
./build.bat
```

# Scripts
There are also some scripts in the scripts folder to help development. 
> These are AI slop, but since it's not actually part of the project, I couldn't be bothered to make something better.

completed.bat
 - Prints out how many visitors are created and how many are missing
 - `--missing` flag prints out markdown list for README.md

extract_inputs.bat 
 - Prints out all the inputs that are caputured for a specific rule 
 - Useful for creating the unit tests 

test_grammar.bat 
 - Runs the grammar and makes sure that changes don't make the grammar worse 

# Visitor Completion (39 / 151 = 26%)
- [x] VisitAllGradeCourseList
- [x] VisitAnyCoreSCHCondition
- [x] VisitAtLeastGradeLevelStandingCondition
- [x] VisitBachelorsOrMastersCondition
- [x] VisitCoreCondition
- [x] VisitCourseAlternativeCondition
- [x] VisitCrossListedCourse
- [x] VisitDegree
- [x] VisitDegreeTypePrefixMajorCondition
- [x] VisitEitherGradeCourseList
- [x] VisitFullCourseList
- [x] VisitGpaGradeCondition
- [x] VisitGpaInCourseCondition
- [x] VisitGradeAtLeastCondition
- [x] VisitGradeCourseListAlternativeCondition
- [x] VisitGradeLevelMajorStandingCondition
- [x] VisitGradeLevelPrefixMajorCondition
- [x] VisitGradeLevelStandingCondition
- [x] VisitGradeListCondition
- [x] VisitGraduateLevelStandingCondition
- [x] VisitGraduateStandingInCondition
- [x] VisitMinimumGpaCondition
- [x] VisitMinimumGradeLevelStandingCondition
- [x] VisitMinimumHoursCondition
- [x] VisitMinimumHoursOfCondition
- [x] VisitNamedDegreeTypeMajorCondition
- [x] VisitNamedMajorCondition
- [x] VisitParenCourse
- [x] VisitParenGradeCourseList
- [x] VisitPrefixGradeLevelStandingCondition
- [x] VisitPrefixMajorCondition
- [x] VisitSemesterCreditHoursCondition
- [x] VisitShorthandCourseList
- [x] VisitSimpleCourse
- [x] VisitSimpleGradeCondition
- [x] VisitTitle
- [x] VisitUndergraduateDegreeCondition
- [x] VisitUniversityGpaCondition
- [x] VisitUpperDivisionSCHCondition
- [ ] VisitAcademic
- [ ] VisitAcademicPlanReq
- [ ] VisitAleksScoreCondition
- [ ] VisitAleksScoreExpr
- [ ] VisitAlternativeExpr
- [ ] VisitAlternativeRepeateExpr
- [ ] VisitAmpersandExpr
- [ ] VisitAndExpr
- [ ] VisitAndRepeateExpr
- [ ] VisitAnyCoreExpr
- [ ] VisitAnyMajorCourseExpr
- [ ] VisitAnyPreviousMajorCourseCondition
- [ ] VisitAppendAcademicPlanReq
- [ ] VisitApScoreCondition
- [ ] VisitApScoreExpr
- [ ] VisitBareRepeatRule
- [ ] VisitBothHonorsCondition
- [ ] VisitCombinedRepeatMaxHoursRule
- [ ] VisitCompleteHoursFromFollowingCondition
- [ ] VisitCompleteNExpr
- [ ] VisitCompleteNFromFollowingCondition
- [ ] VisitCompleteNOfFollowingCondition
- [ ] VisitComputerScholarsExpr
- [ ] VisitComputerScholarsReq
- [ ] VisitConcurrentEnrollmentCondition
- [ ] VisitConcurrentEnrollmentExpr
- [ ] VisitCoreExpr
- [ ] VisitCoreqReq
- [ ] VisitCourseExpr
- [ ] VisitCourseRepeateExpr
- [ ] VisitCourseRepeatLimitRule
- [ ] VisitCourseRepeatMaxHoursRule
- [ ] VisitCourseRepeatRule
- [ ] VisitCredit
- [ ] VisitCreditForReq
- [ ] VisitDegreeExpr
- [ ] VisitDegreeSatisfactionReq
- [ ] VisitDepartmentConsentExpr
- [ ] VisitElectivesDegreeSatisfactionRule
- [ ] VisitEquivalentExpr
- [ ] VisitExactCoreqNoticeReq
- [ ] VisitExactSectionCondition
- [ ] VisitExactSectionExpr
- [ ] VisitExcludeNoticeReq
- [ ] VisitExprReq
- [ ] VisitFourThousandLevelCondition
- [ ] VisitGoodAcademicStandingExpr
- [ ] VisitGpa
- [ ] VisitGpaExpr
- [ ] VisitGpaRepeatReq
- [ ] VisitGradeExpr
- [ ] VisitGradeLevelStandingExpr
- [ ] VisitGraduateStandingExpr
- [ ] VisitGroup
- [ ] VisitGroupCondition
- [ ] VisitGroupExpr
- [ ] VisitHonorsExpr
- [ ] VisitInstructorConsentExpr
- [ ] VisitInternshipRepeatRule
- [ ] VisitLivingLearningExpr
- [ ] VisitLivingLearningReq
- [ ] VisitMajorExpr
- [ ] VisitMajorReq
- [ ] VisitMathDegreeSatisfactionRule
- [ ] VisitMinimumHoursExpr
- [ ] VisitMultiPrefixDegreeSatisfactionRule
- [ ] VisitMultiPrefixForDegreeSatisfactionRule
- [ ] VisitNamedDegreeSatisfactionRule
- [ ] VisitNamedLivingLearningRule
- [ ] VisitOfMultiPrefixDegreeSatisfactionRule
- [ ] VisitOrExpr
- [ ] VisitParenExpr
- [ ] VisitPlacement
- [ ] VisitPlacementScoreComparisonCondition
- [ ] VisitPlacementScoreMinimumCondition
- [ ] VisitPlacementScoreRangeCondition
- [ ] VisitPlacementTestExpr
- [ ] VisitPrefixDegreeSatisfactionRule
- [ ] VisitPrefixLivingLearningRule
- [ ] VisitPreOrCoReq
- [ ] VisitPrereqAndCoreqReq
- [ ] VisitPrereqReq
- [ ] VisitProg
- [ ] VisitRepeatHoursMaxSuffixRule
- [ ] VisitRepeatLimitHoursExpr
- [ ] VisitRepeatLimitHoursReq
- [ ] VisitRepeatLimitTimesReq
- [ ] VisitRepeatMaxHoursRule
- [ ] VisitRepeatMaxTimesRule
- [ ] VisitRepeatReq
- [ ] VisitRepeatRuleExpr
- [ ] VisitRepeatUpToTimesRule
- [ ] VisitResearchCondition
- [ ] VisitResearchExpr
- [ ] VisitSameAsParenReq
- [ ] VisitSameAsReq
- [ ] VisitSchool
- [ ] VisitSchoolDegreeSatisfactionRule
- [ ] VisitSchoolReq
- [ ] VisitSchoolsDegreeSatisfactionRule
- [ ] VisitSemesterCreditHoursExpr
- [ ] VisitSentence
- [ ] VisitSingleHonorsCondition
- [ ] VisitStudentDegreeSatisfactionRule
- [ ] VisitTopicsVaryRepeatRule
- [ ] VisitUpperDivisionClassesCondition
- [ ] VisitUpperDivisionClassesExpr
- [ ] VisitUpperDivisionHoursExpr
- [ ] VisitUpperDivisionStandingExpr
- [ ] VisitUteachConsentExpr
- [ ] VisitWorkshopSectionCondition