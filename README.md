# UTD Course Requisite Parser 

ANTLR4 based approach for parsing course requirements. 

Currently, the grammar successfully parses* 97% of the 25f course requirements. At this point adding the remaining 20 ish 
 inputs is not worth focusing on since they are all mostly unique edge cases or typos. 

> * Antlr parses them into an AST, not parsed yet into a useful json object. Currently, only 10% of the way of that but the hardest 
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
./build.sh
```

# Scripts
There are also some scripts in the scripts folder to help development. 
> These are AI slop, but since it's not actually part of the project, I couldn't be bothered to make do better.

completed.bat
 - Prints out how many visitors are created and how many are missing
 - `--missing` flag prints out markdown list for README.md

extract_inputs.bat 
 - Prints out all the inputs that are caputured for a specific rule 
 - Useful for creating the unit tests 

test_grammar.bat 
 - Runs the grammar and makes sure that changes don't make the grammar worse 

---

# Visitor Completion (21 / 151 = 13%)

- [x] VisitAllGradeCourseList
- [x] VisitAtLeastGradeLevelStandingCondition
- [x] VisitCourseAlternativeCondition
- [x] VisitCrossListedCourse
- [x] VisitDegree
- [x] VisitEitherGradeCourseList
- [x] VisitFullCourseList
- [x] VisitGpaGradeCondition
- [x] VisitGradeAtLeastCondition
- [x] VisitGradeCourseListAlternativeCondition
- [x] VisitGradeLevelMajorStandingCondition
- [x] VisitGradeLevelStandingCondition
- [x] VisitGradeListCondition
- [x] VisitMinimumGradeLevelStandingCondition
- [x] VisitParenCourse
- [x] VisitParenGradeCourseList
- [x] VisitPrefixGradeLevelStandingCondition
- [x] VisitShorthandCourseList
- [x] VisitSimpleCourse
- [x] VisitSimpleGradeCondition
- [x] VisitTitle
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
- [ ] VisitAnyCoreSCHCondition
- [ ] VisitAnyMajorCourseExpr
- [ ] VisitAnyPreviousMajorCourseCondition
- [ ] VisitAppendAcademicPlanReq
- [ ] VisitApScoreCondition
- [ ] VisitApScoreExpr
- [ ] VisitBachelorsOrMastersCondition
- [ ] VisitBareRepeatRule
- [ ] VisitBothHonorsCondition
- [ ] VisitCombinedRepeatMaxHoursRule
- [ ] VisitCompleteHoursFromFollowingCondition
- [ ] VisitCompleteNExpr
- [ ] VisitCompleteNFromFollowingCondition
- [ ] VisitCompleteNOfFollowingCondition
- [ ] VisitCompletionOfCoreCondition
- [ ] VisitComputerScholarsExpr
- [ ] VisitComputerScholarsReq
- [ ] VisitConcurrentEnrollmentCondition
- [ ] VisitConcurrentEnrollmentExpr
- [ ] VisitCoreCondition
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
- [ ] VisitDegreeTypePrefixMajorCondition
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
- [ ] VisitGpaInCourseCondition
- [ ] VisitGpaRepeatReq
- [ ] VisitGradeExpr
- [ ] VisitGradeLevelPrefixMajorCondition
- [ ] VisitGradeLevelStandingExpr
- [ ] VisitGraduateLevelStandingCondition
- [ ] VisitGraduateStandingExpr
- [ ] VisitGraduateStandingInCondition
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
- [ ] VisitMinimumGpaCondition
- [ ] VisitMinimumHoursCondition
- [ ] VisitMinimumHoursExpr
- [ ] VisitMultiPrefixDegreeSatisfactionRule
- [ ] VisitMultiPrefixForDegreeSatisfactionRule
- [ ] VisitNamedDegreeSatisfactionRule
- [ ] VisitNamedDegreeTypeMajorCondition
- [ ] VisitNamedLivingLearningRule
- [ ] VisitNamedMajorCondition
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
- [ ] VisitPrefixMajorCondition
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
- [ ] VisitSemesterCreditHoursCondition
- [ ] VisitSemesterCreditHoursExpr
- [ ] VisitSentence
- [ ] VisitSingleHonorsCondition
- [ ] VisitStudentDegreeSatisfactionRule
- [ ] VisitTopicsVaryRepeatRule
- [ ] VisitUndergraduateDegreeCondition
- [ ] VisitUniversityGpaCondition
- [ ] VisitUpperDivisionClassesCondition
- [ ] VisitUpperDivisionClassesExpr
- [ ] VisitUpperDivisionCreditsCondition
- [ ] VisitUpperDivisionHoursExpr
- [ ] VisitUpperDivisionSCHCondition
- [ ] VisitUpperDivisionStandingExpr
- [ ] VisitUteachConsentExpr
- [ ] VisitWorkshopSectionCondition