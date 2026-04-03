package visitors

import (
	"parser/conditions"
	"parser/parser"
	"parser/utils"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// ======================= consent conditions =======================

// VisitInstructorConsentCondition
//
// Rule: INSTRUCTOR_CONSENT
func (v *RequisiteVisitor) VisitInstructorConsentCondition(ctx *parser.InstructorConsentConditionContext) any {
	return v.visitInstructorConsentCondition(ctx)
}

func (v *RequisiteVisitor) visitInstructorConsentCondition(ctx *parser.InstructorConsentConditionContext) *conditions.ConsentCondition {
	return conditions.NewConsentCondition(utils.InstructorConsent)
}

// VisitDepartmentConsentCondition
//
// Rule: DEPARTMENT_CONSENT
func (v *RequisiteVisitor) VisitDepartmentConsentCondition(ctx *parser.DepartmentConsentConditionContext) any {
	return v.visitDepartmentConsentCondition(ctx)
}

func (v *RequisiteVisitor) visitDepartmentConsentCondition(ctx *parser.DepartmentConsentConditionContext) *conditions.ConsentCondition {
	return conditions.NewConsentCondition(utils.DepartmentConsent)
}

// VisitUteachConsentCondition
//
// Rule: UTEACH_CONSENT
func (v *RequisiteVisitor) VisitUteachConsentCondition(ctx *parser.UteachConsentConditionContext) any {
	return v.visitUteachConsentCondition(ctx)
}

func (v *RequisiteVisitor) visitUteachConsentCondition(ctx *parser.UteachConsentConditionContext) *conditions.ConsentCondition {
	return conditions.NewConsentCondition(utils.UTeachConsent)
}

// ======================= standing conditions =======================

// VisitUpperDivisionStandingCondition
//
// Rule: UPPER_DVISION_STANDING
func (v *RequisiteVisitor) VisitUpperDivisionStandingCondition(ctx *parser.UpperDivisionStandingConditionContext) any {
	return v.visitUpperDivisionStandingCondition(ctx)
}

func (v *RequisiteVisitor) visitUpperDivisionStandingCondition(ctx *parser.UpperDivisionStandingConditionContext) *conditions.GenericStandingCondition {
	return conditions.NewGenericStandingCondition(utils.UpperDivisionStanding)
}

// VisitGoodAcademicStandingCondition
//
// Rule: GOOD_ACADEMIC_STANDING
func (v *RequisiteVisitor) VisitGoodAcademicStandingCondition(ctx *parser.GoodAcademicStandingConditionContext) any {
	return v.visitGoodAcademicStandingCondition(ctx)
}

func (v *RequisiteVisitor) visitGoodAcademicStandingCondition(ctx *parser.GoodAcademicStandingConditionContext) *conditions.GenericStandingCondition {
	return conditions.NewGenericStandingCondition(utils.GoodAcademicStanding)
}

// ======================= grade condition =======================

func (v *RequisiteVisitor) applyGrade(node antlr.ParseTree, gradeText string) conditions.Condition {
	gradedCond, _ := v.Visit(node).(conditions.GradedCondition)
	gradedCond.AppendGrade(utils.Grade(gradeText))
	return gradedCond
}

// VisitSimpleGradeCondition
//
// Rule: course WITH_GRADE GRADE OR_BETTER?
func (v *RequisiteVisitor) VisitSimpleGradeCondition(ctx *parser.SimpleGradeConditionContext) any {
	return v.visitSimpleGradeCondition(ctx)
}

func (v *RequisiteVisitor) visitSimpleGradeCondition(ctx *parser.SimpleGradeConditionContext) conditions.Condition {
	return v.applyGrade(ctx.Course(), ctx.GRADE().GetText())
}

// VisitGpaGradeCondition
//
// Rule: course WITH_GRADE 'greater than or equal to' GRADE ('(' GPA ')')?
func (v *RequisiteVisitor) VisitGpaGradeCondition(ctx *parser.GpaGradeConditionContext) any {
	return v.visitGpaGradeCondition(ctx)
}

func (v *RequisiteVisitor) visitGpaGradeCondition(ctx *parser.GpaGradeConditionContext) conditions.Condition {
	return v.applyGrade(ctx.Course(), ctx.GRADE().GetText())
}

// VisitGradeListCondition
//
// Rule: grade_course_list WITH_GRADE GRADE OR_BETTER?
func (v *RequisiteVisitor) VisitGradeListCondition(ctx *parser.GradeListConditionContext) any {
	return v.visitGradeListCondition(ctx)
}

func (v *RequisiteVisitor) visitGradeListCondition(ctx *parser.GradeListConditionContext) conditions.Condition {
	return v.applyGrade(ctx.Grade_course_list(), ctx.GRADE().GetText())
}

// VisitGradeAtLeastCondition
//
// Rule: A_GRADE_OF_AT_LEAST GRADE OR_BETTER? 'in' grade_course_list
func (v *RequisiteVisitor) VisitGradeAtLeastCondition(ctx *parser.GradeAtLeastConditionContext) any {
	return v.visitGradeAtLeastCondition(ctx)
}

func (v *RequisiteVisitor) visitGradeAtLeastCondition(ctx *parser.GradeAtLeastConditionContext) conditions.Condition {
	return v.applyGrade(ctx.Grade_course_list(), ctx.GRADE().GetText())
}

// ======================= (course) alternative condition =======================

// VisitCourseAlternativeCondition
//
// Rule: course OR EQUIVALENT
func (v *RequisiteVisitor) VisitCourseAlternativeCondition(ctx *parser.CourseAlternativeConditionContext) any {
	return v.visitCourseAlternativeCondition(ctx)
}
func (v *RequisiteVisitor) visitCourseAlternativeCondition(ctx *parser.CourseAlternativeConditionContext) *conditions.AlternativeCondition {
	course := v.Visit(ctx.Course()).(conditions.Condition)
	return conditions.NewAlternativeCondition(course)
}

// VisitGradeCourseListAlternativeCondition
//
// Rule: course_list OR EQUIVALENT
func (v *RequisiteVisitor) VisitGradeCourseListAlternativeCondition(ctx *parser.GradeCourseListAlternativeConditionContext) any {
	return v.visitGradeCourseListAlternativeCondition(ctx)
}

func (v *RequisiteVisitor) visitGradeCourseListAlternativeCondition(ctx *parser.GradeCourseListAlternativeConditionContext) *conditions.AlternativeCondition {
	courseList := v.Visit(ctx.Course_list()).(conditions.Condition)
	return conditions.NewAlternativeCondition(courseList)
}

// ======================= grade level standing condition =======================

// VisitGradeLevelStandingCondition
//
// Rule: GRADE_LEVEL (OR GRADE_LEVEL)* LEVEL? STANDING
func (v *RequisiteVisitor) VisitGradeLevelStandingCondition(ctx *parser.GradeLevelStandingConditionContext) any {
	return v.visitGradeLevelStandingCondition(ctx)
}

func (v *RequisiteVisitor) visitGradeLevelStandingCondition(ctx *parser.GradeLevelStandingConditionContext) conditions.Condition {
	conds := make([]conditions.Condition, len(ctx.AllGRADE_LEVEL()))
	for i, gradeLevel := range ctx.AllGRADE_LEVEL() {
		conds[i] = conditions.NewGradeLevelCondition(mapGradeLevel(gradeLevel.GetText()))
	}

	return conditions.NewOrCondition(conds...)
}

// VisitGradeLevelMajorStandingCondition
//
// Rule: GRADE_LEVEL degree MAJOR_KW STANDING
func (v *RequisiteVisitor) VisitGradeLevelMajorStandingCondition(ctx *parser.GradeLevelMajorStandingConditionContext) any {
	return v.visitGradeLevelMajorStandingCondition(ctx)
}

func (v *RequisiteVisitor) visitGradeLevelMajorStandingCondition(ctx *parser.GradeLevelMajorStandingConditionContext) *conditions.GradeLevelCondition {
	return conditions.NewGradeLevelConditionWithDegree(mapGradeLevel(ctx.GRADE_LEVEL().GetText()), ctx.Degree().GetText())
}

// VisitMinimumGradeLevelStandingCondition
//
// Rule: MINIMUM_OF GRADE_LEVEL STANDING
func (v *RequisiteVisitor) VisitMinimumGradeLevelStandingCondition(ctx *parser.MinimumGradeLevelStandingConditionContext) any {
	return v.visitMinimumGradeLevelStandingCondition(ctx)
}

func (v *RequisiteVisitor) visitMinimumGradeLevelStandingCondition(ctx *parser.MinimumGradeLevelStandingConditionContext) *conditions.GradeLevelCondition {
	return conditions.NewGradeLevelCondition(mapGradeLevel(ctx.GRADE_LEVEL().GetText()))
}

// VisitAtLeastGradeLevelStandingCondition
//
// Rule: AT_LEAST GRADE_LEVEL (DASH LEVEL | LEVEL)? STANDING
func (v *RequisiteVisitor) VisitAtLeastGradeLevelStandingCondition(ctx *parser.AtLeastGradeLevelStandingConditionContext) any {
	return v.visitAtLeastGradeLevelStandingCondition(ctx)
}

func (v *RequisiteVisitor) visitAtLeastGradeLevelStandingCondition(ctx *parser.AtLeastGradeLevelStandingConditionContext) *conditions.GradeLevelCondition {
	return conditions.NewGradeLevelCondition(mapGradeLevel(ctx.GRADE_LEVEL().GetText()))
}

// VisitPrefixGradeLevelStandingCondition
//
// Rule: PREFIX MAJOR_KW? ONLY_KW 'with' GRADE_LEVEL (OR GRADE_LEVEL)* LEVEL? STANDING
func (v *RequisiteVisitor) VisitPrefixGradeLevelStandingCondition(ctx *parser.PrefixGradeLevelStandingConditionContext) any {
	return v.visitPrefixGradeLevelStandingCondition(ctx)
}

func (v *RequisiteVisitor) visitPrefixGradeLevelStandingCondition(ctx *parser.PrefixGradeLevelStandingConditionContext) conditions.Condition {
	//todo map/standardize major/degree/school etc
	degree := ctx.PREFIX().GetText()

	conds := make([]conditions.Condition, len(ctx.AllGRADE_LEVEL()))
	for i, gradeLevel := range ctx.AllGRADE_LEVEL() {
		conds[i] = conditions.NewGradeLevelConditionWithDegree(mapGradeLevel(gradeLevel.GetText()), degree)
	}

	return conditions.NewOrCondition(conds...)
}

// ======================= graduate standing condition =======================

// VisitGraduateStandingInCondition
//
// Rule: 'Graduate standing in' degree
func (v *RequisiteVisitor) VisitGraduateStandingInCondition(ctx *parser.GraduateStandingInConditionContext) any {
	return v.visitGraduateStandingInCondition(ctx)
}

func (v *RequisiteVisitor) visitGraduateStandingInCondition(ctx *parser.GraduateStandingInConditionContext) *conditions.GraduateStandingInCondition {
	return conditions.NewGraduateStandingInConditionWithDegree(ctx.Degree().GetText())
}

// VisitGraduateLevelStandingCondition
//
// Rule: 'Graduate Level Standing'
func (v *RequisiteVisitor) VisitGraduateLevelStandingCondition(ctx *parser.GraduateLevelStandingConditionContext) any {
	return v.visitGraduateLevelStandingCondition(ctx)
}

func (v *RequisiteVisitor) visitGraduateLevelStandingCondition(ctx *parser.GraduateLevelStandingConditionContext) *conditions.GraduateStandingInCondition {
	return conditions.NewGraduateStandingInCondition()
}

// ======================= GPA condition =======================

// VisitUniversityGpaCondition
//
// Rule: UNIVERSITY_GPA_KW GPA
func (v *RequisiteVisitor) VisitUniversityGpaCondition(ctx *parser.UniversityGpaConditionContext) any {
	return v.visitUniversityGpaCondition(ctx)
}

func (v *RequisiteVisitor) visitUniversityGpaCondition(ctx *parser.UniversityGpaConditionContext) *conditions.GPACondition {
	return conditions.NewGpaCondition(mapGPA(ctx.GPA().GetText()))
}

// VisitMinimumGpaCondition
//
// Rule: 'Minimum GPA requirement' GPA
func (v *RequisiteVisitor) VisitMinimumGpaCondition(ctx *parser.MinimumGpaConditionContext) any {
	return v.visitMinimumGpaCondition(ctx)
}

func (v *RequisiteVisitor) visitMinimumGpaCondition(ctx *parser.MinimumGpaConditionContext) *conditions.GPACondition {
	return conditions.NewGpaCondition(mapGPA(ctx.GPA().GetText()))
}

// VisitGpaInCourseCondition
//
// Rule: 'a GPA of' GPA OR_BETTER 'in' degree COURSE_KW?
func (v *RequisiteVisitor) VisitGpaInCourseCondition(ctx *parser.GpaInCourseConditionContext) any {
	return v.visitGpaInCourseCondition(ctx)
}

func (v *RequisiteVisitor) visitGpaInCourseCondition(ctx *parser.GpaInCourseConditionContext) *conditions.GPACondition {
	degree := v.Visit(ctx.Degree()).(string)
	return conditions.NewGpaConditionWithDegree(mapGPA(ctx.GPA().GetText()), degree)
}

// ======================= Major condition =======================

// VisitPrefixMajorCondition
//
// Rule: PREFIX (OR PREFIX)* (DIVISION_TYPE | DEGREE_LEVEL)? GRADE_LEVEL? MAJOR_KW? ONLY_KW?
func (v *RequisiteVisitor) VisitPrefixMajorCondition(ctx *parser.PrefixMajorConditionContext) any {
	return v.visitPrefixMajorCondition(ctx)
}

func (v *RequisiteVisitor) visitPrefixMajorCondition(ctx *parser.PrefixMajorConditionContext) conditions.Condition {

	var degreeLevel utils.DegreeLevel
	if ctx.DIVISION_TYPE() != nil {
		degreeLevel = mapDivisionType(ctx.DIVISION_TYPE().GetText())
	}

	if ctx.DEGREE_LEVEL() != nil {
		degreeLevel = mapDegreeLevel(ctx.DEGREE_LEVEL().GetText())
	}

	var gradeLevel utils.GradeLevel
	if ctx.GRADE_LEVEL() != nil {
		gradeLevel = mapGradeLevel(ctx.GRADE_LEVEL().GetText())
	}

	conds := make([]conditions.Condition, len(ctx.AllPREFIX()))
	for i, prefix := range ctx.AllPREFIX() {
		degree := prefix.GetText()

		if degreeLevel != "" && gradeLevel != "" {
			conds[i] = conditions.NewMajorConditionWithDegreeAndGradeLevel(degree, degreeLevel, gradeLevel)
		} else if degreeLevel != "" {
			conds[i] = conditions.NewMajorConditionWithDegreeLevel(degree, degreeLevel)
		} else if gradeLevel != "" {
			conds[i] = conditions.NewMajorConditionWithGradeLevel(degree, gradeLevel)
		} else {
			conds[i] = conditions.NewMajorCondition(degree)
		}
	}
	return conditions.NewOrCondition(conds...)
}

// VisitGradeLevelPrefixMajorCondition
//
// Rule: GRADE_LEVEL PREFIX MAJOR_KW? ONLY_KW?
func (v *RequisiteVisitor) VisitGradeLevelPrefixMajorCondition(ctx *parser.GradeLevelPrefixMajorConditionContext) any {
	return v.visitGradeLevelPrefixMajorCondition(ctx)
}

func (v *RequisiteVisitor) visitGradeLevelPrefixMajorCondition(ctx *parser.GradeLevelPrefixMajorConditionContext) *conditions.MajorCondition {
	//todo map prefix to degree
	degree := ctx.PREFIX().GetText()
	gradeLevel := mapGradeLevel(ctx.GRADE_LEVEL().GetText())
	return conditions.NewMajorConditionWithGradeLevel(degree, gradeLevel)
}

// VisitDegreeTypePrefixMajorCondition
//
// Rule: DEGREE_LEVEL PREFIX MAJOR_KW? ONLY_KW?
func (v *RequisiteVisitor) VisitDegreeTypePrefixMajorCondition(ctx *parser.DegreeTypePrefixMajorConditionContext) any {
	return v.visitDegreeTypePrefixMajorCondition(ctx)
}

func (v *RequisiteVisitor) visitDegreeTypePrefixMajorCondition(ctx *parser.DegreeTypePrefixMajorConditionContext) *conditions.MajorCondition {
	degree := ctx.PREFIX().GetText()
	gradeLevel := mapDegreeLevel(ctx.DEGREE_LEVEL().GetText())
	return conditions.NewMajorConditionWithDegreeLevel(degree, gradeLevel)
}

// VisitNamedMajorCondition
//
// Rule: degree MAJOR_KW ONLY_KW?
func (v *RequisiteVisitor) VisitNamedMajorCondition(ctx *parser.NamedMajorConditionContext) any {
	return v.visitNamedMajorCondition(ctx)
}

func (v *RequisiteVisitor) visitNamedMajorCondition(ctx *parser.NamedMajorConditionContext) *conditions.MajorCondition {
	degree := v.Visit(ctx.Degree()).(string)
	return conditions.NewMajorCondition(degree)
}

// VisitNamedDegreeTypeMajorCondition
//
// Rule: degree DEGREE_LEVEL MAJOR_KW? ONLY_KW?
func (v *RequisiteVisitor) VisitNamedDegreeTypeMajorCondition(ctx *parser.NamedDegreeTypeMajorConditionContext) any {
	return v.visitNamedDegreeTypeMajorCondition(ctx)
}

func (v *RequisiteVisitor) visitNamedDegreeTypeMajorCondition(ctx *parser.NamedDegreeTypeMajorConditionContext) *conditions.MajorCondition {
	degree := v.Visit(ctx.Degree()).(string)
	degreeLevel := mapDegreeLevel(ctx.DEGREE_LEVEL().GetText())
	return conditions.NewMajorConditionWithDegreeLevel(degree, degreeLevel)
}

// ======================= degree condition =======================

// VisitUndergraduateDegreeCondition
//
// Rule: 'an undergraduate degree in' title 'and adequate foundation/academic performance in a corresponding area'
func (v *RequisiteVisitor) VisitUndergraduateDegreeCondition(ctx *parser.UndergraduateDegreeConditionContext) any {
	return v.visitUndergraduateDegreeCondition(ctx)
}

func (v *RequisiteVisitor) visitUndergraduateDegreeCondition(ctx *parser.UndergraduateDegreeConditionContext) *conditions.DegreeCondition {
	return conditions.NewDegreeCondition(v.Visit(ctx.Degree()).(string))
}

// VisitBachelorsOrMastersCondition
//
// Rule: 'Bachelor\'s or Master\'s degree in' degree_list
func (v *RequisiteVisitor) VisitBachelorsOrMastersCondition(ctx *parser.BachelorsOrMastersConditionContext) any {
	return v.visitBachelorsOrMastersCondition(ctx)
}

func (v *RequisiteVisitor) visitBachelorsOrMastersCondition(ctx *parser.BachelorsOrMastersConditionContext) conditions.Condition {
	degreeList := v.Visit(ctx.Degree_list()).([]string)
	conds := make([]conditions.Condition, len(degreeList))
	for i, degree := range degreeList {
		conds[i] = conditions.NewDegreeCondition(degree)
	}

	return conditions.NewOrCondition(conds...)
}

// ======================= core condition =======================

// VisitCoreCondition
//
// Rule: COMPLETION_OF ('a' | 'an')? CORE_NUMBER (CORE | '(' CORE ')' )? CORE_KW COURSE_KW?
func (v *RequisiteVisitor) VisitCoreCondition(ctx *parser.CoreConditionContext) any {
	return v.visitCoreCondition(ctx)
}

func (v *RequisiteVisitor) visitCoreCondition(ctx *parser.CoreConditionContext) *conditions.CoreCondition {
	coreNumber := ctx.CORE_NUMBER().GetText()
	coreTitle := stripChars(v.getTextOrDefault(ctx.CORE(), ""), "(", ")")
	if strings.ToLower(coreTitle) == "core" {
		coreTitle = ""
	}

	return conditions.NewCoreCondition(coreNumber, coreTitle)
}

// VisitAnyCoreSCHCondition
//
// Rule: 'any' SMALL_INT SEMESTER_CREDIT_HOURS CORE_NUMBER CORE_KW COURSE_KW
func (v *RequisiteVisitor) VisitAnyCoreSCHCondition(ctx *parser.AnyCoreSCHConditionContext) any {
	return v.visitAnyCoreSCHCondition(ctx)
}

func (v *RequisiteVisitor) visitAnyCoreSCHCondition(ctx *parser.AnyCoreSCHConditionContext) *conditions.CoreCondition {
	hours := mapInt(ctx.SMALL_INT().GetText())
	coreNumber := ctx.CORE_NUMBER().GetText()

	//todo map core number to title / vice versa
	return conditions.NewCoreConditionWithSemesterHours(coreNumber, "", hours)
}

// ======================= SCH condition =======================

// VisitSemesterCreditHoursCondition
//
// Rule: INT SEMESTER_CREDIT_HOURS
func (v *RequisiteVisitor) VisitSemesterCreditHoursCondition(ctx *parser.SemesterCreditHoursConditionContext) any {
	return v.visitSemesterCreditHoursCondition(ctx)
}

func (v *RequisiteVisitor) visitSemesterCreditHoursCondition(ctx *parser.SemesterCreditHoursConditionContext) *conditions.CreditHoursCondition {
	hours := mapInt(ctx.INT().GetText())
	return conditions.NewCreditHoursCondition(hours)
}

// ======================= SCH in Courses condition =======================

// VisitMinimumHoursCondition
//
// Rule: AT_LEAST SMALL_INT 'semester credits of' course_list
func (v *RequisiteVisitor) VisitMinimumHoursCondition(ctx *parser.MinimumHoursConditionContext) any {
	return v.visitMinimumHoursCondition(ctx)
}

func (v *RequisiteVisitor) visitMinimumHoursCondition(ctx *parser.MinimumHoursConditionContext) *conditions.CreditHoursFromCondition {
	hours := mapInt(ctx.SMALL_INT().GetText())
	courses := extractCoursesFromCourseList(v.Visit(ctx.Course_list()).(conditions.Condition))
	return conditions.NewCreditHoursFromCondition(hours, courses)
}

// VisitMinimumHoursOfCondition
//
// Rule: MINIMUM_OF SMALL_INT SEMESTER_CREDIT_HOURS 'in any combination of' course_list
func (v *RequisiteVisitor) VisitMinimumHoursOfCondition(ctx *parser.MinimumHoursOfConditionContext) any {
	return v.visitMinimumHoursOfCondition(ctx)
}

func (v *RequisiteVisitor) visitMinimumHoursOfCondition(ctx *parser.MinimumHoursOfConditionContext) *conditions.CreditHoursFromCondition {
	hours := mapInt(ctx.SMALL_INT().GetText())
	courses := extractCoursesFromCourseList(v.Visit(ctx.Course_list()).(conditions.Condition))
	return conditions.NewCreditHoursFromCondition(hours, courses)
}

// VisitMinimumHoursFromCondition
//
// Rule: MINIMUM_OF SMALL_INT SEMESTER_CREDIT_HOURS 'in any combination of' course_list
func (v *RequisiteVisitor) VisitMinimumHoursFromCondition(ctx *parser.MinimumHoursFromConditionContext) any {
	return v.visitMinimumHoursFromCondition(ctx)
}

func (v *RequisiteVisitor) visitMinimumHoursFromCondition(ctx *parser.MinimumHoursFromConditionContext) *conditions.CreditHoursFromCondition {
	hours := mapInt(ctx.SMALL_INT().GetText())
	courses := extractCoursesFromCourseList(v.Visit(ctx.Course_list()).(conditions.Condition))
	return conditions.NewCreditHoursFromCondition(hours, courses)
}

// ======================= Upper Division condition =======================

// VisitUpperDivisionSCHCondition
//
// Rule: SMALL_INT 'SCH of upper-division' PREFIX COURSE_KW
func (v *RequisiteVisitor) VisitUpperDivisionSCHCondition(ctx *parser.UpperDivisionSCHConditionContext) any {
	return v.visitUpperDivisionSCHCondition(ctx)
}

func (v *RequisiteVisitor) visitUpperDivisionSCHCondition(ctx *parser.UpperDivisionSCHConditionContext) *conditions.UpperDivisionCoursesCondition {
	hours := mapInt(ctx.SMALL_INT().GetText())
	prefix := ctx.PREFIX().GetText()

	return conditions.NewUpperDivisionCreditHoursCondition(hours, prefix)
}

// VisitUpperDivisionCountCondition
//
// Rule: AT_LEAST NUMBER_STRING PREFIX UPPER_DIVISION_COURSE_NUMBER COURSE_KW
func (v *RequisiteVisitor) VisitUpperDivisionCountCondition(ctx *parser.UpperDivisionCountConditionContext) any {
	return v.visitUpperDivisionCountCondition(ctx)
}

func (v *RequisiteVisitor) visitUpperDivisionCountCondition(ctx *parser.UpperDivisionCountConditionContext) *conditions.UpperDivisionCoursesCondition {
	count := mapNumberString(ctx.NUMBER_STRING().GetText())
	prefix := ctx.PREFIX().GetText()
	return conditions.NewUpperDivisionCountCondition(count, prefix)
}

// VisitUpperDivisionSingleCondition
//
// Rule: 'a 4000-level' PREFIX COURSE_KW
func (v *RequisiteVisitor) VisitUpperDivisionSingleCondition(ctx *parser.UpperDivisionSingleConditionContext) any {
	return v.visitUpperDivisionSingleCondition(ctx)
}

func (v *RequisiteVisitor) visitUpperDivisionSingleCondition(ctx *parser.UpperDivisionSingleConditionContext) *conditions.UpperDivisionCoursesCondition {
	prefix := ctx.PREFIX().GetText()
	return conditions.NewUpperDivisionCountCondition(1, prefix)
}

// VisitResearchCondition
//
// Rule: AT_LEAST SMALL_INT SEMESTER_CREDIT_HOURS 'of' DIVISION_TYPE 'research'
func (v *RequisiteVisitor) VisitResearchCondition(ctx *parser.ResearchConditionContext) any {
	return v.visitResearchCondition(ctx)
}

func (v *RequisiteVisitor) visitResearchCondition(ctx *parser.ResearchConditionContext) *conditions.ResearchCondition {
	count := mapInt(ctx.SMALL_INT().GetText())
	degreeLevel := mapDivisionType(ctx.DIVISION_TYPE().GetText())

	return conditions.NewResearchCondition(count, degreeLevel)
}

// ======================= N courses condition =======================

// VisitCompleteNOfFollowingCondition
//
// Rule: COMPLETION_OF NUMBER_STRING 'of the following' COLON course_list
func (v *RequisiteVisitor) VisitCompleteNOfFollowingCondition(ctx *parser.CompleteNOfFollowingConditionContext) any {
	return v.visitCompleteNOfFollowingCondition(ctx)
}

func (v *RequisiteVisitor) visitCompleteNOfFollowingCondition(ctx *parser.CompleteNOfFollowingConditionContext) *conditions.NCoursesCondition {
	count := mapNumberString(ctx.NUMBER_STRING().GetText())
	courses := extractCoursesFromCourseList(v.Visit(ctx.Course_list()).(conditions.Condition))
	return conditions.NewNCoursesCondition(count, courses)
}

// VisitCompleteNFromFollowingCondition
//
// Rule: NUMBER_STRING COURSE_KW 'from the following' DASH course_list
func (v *RequisiteVisitor) VisitCompleteNFromFollowingCondition(ctx *parser.CompleteNFromFollowingConditionContext) any {
	return v.visitCompleteNFromFollowingCondition(ctx)
}

func (v *RequisiteVisitor) visitCompleteNFromFollowingCondition(ctx *parser.CompleteNFromFollowingConditionContext) *conditions.NCoursesCondition {
	count := mapNumberString(ctx.NUMBER_STRING().GetText())
	courses := extractCoursesFromCourseList(v.Visit(ctx.Course_list()).(conditions.Condition))
	return conditions.NewNCoursesCondition(count, courses)
}

// ======================= Placement Test condition =======================

// VisitPlacementScoreComparisonCondition
//
// Rule: 'a'? placement_test_name SCORE_KW (LESS_THAN | GREATER_THAN) INT
func (v *RequisiteVisitor) VisitPlacementScoreComparisonCondition(ctx *parser.PlacementScoreComparisonConditionContext) any {
	return v.visitPlacementScoreComparisonCondition(ctx)
}

func (v *RequisiteVisitor) visitPlacementScoreComparisonCondition(ctx *parser.PlacementScoreComparisonConditionContext) *conditions.PlacementTestScoreCondition {
	name := v.Visit(ctx.Placement_test_name()).(string)
	score := mapInt(ctx.INT().GetText())

	if ctx.LESS_THAN() != nil {
		return conditions.NewPlacementTestScoreCondition(name, 0, score)
	}
	return conditions.NewPlacementTestScoreCondition(name, score, 100)
}

// VisitPlacementScoreRangeCondition
//
// Rule: 'a'? placement_test_name SCORE_KW 'of' INT DASH INT
func (v *RequisiteVisitor) VisitPlacementScoreRangeCondition(ctx *parser.PlacementScoreRangeConditionContext) any {
	return v.visitPlacementScoreRangeCondition(ctx)
}

func (v *RequisiteVisitor) visitPlacementScoreRangeCondition(ctx *parser.PlacementScoreRangeConditionContext) *conditions.PlacementTestScoreCondition {
	name := v.Visit(ctx.Placement_test_name()).(string)
	minScore := mapInt(ctx.INT(0).GetText())
	maxScore := mapInt(ctx.INT(1).GetText())

	return conditions.NewPlacementTestScoreCondition(name, minScore, maxScore)
}

// VisitPlacementScoreMinimumCondition
//
// Rule: 'a'? placement_test_name 'of'? INT OR_BETTER
func (v *RequisiteVisitor) VisitPlacementScoreMinimumCondition(ctx *parser.PlacementScoreMinimumConditionContext) any {
	return v.visitPlacementScoreMinimumCondition(ctx)
}

func (v *RequisiteVisitor) visitPlacementScoreMinimumCondition(ctx *parser.PlacementScoreMinimumConditionContext) *conditions.PlacementTestScoreCondition {
	name := v.Visit(ctx.Placement_test_name()).(string)
	score := mapInt(ctx.INT().GetText())

	return conditions.NewPlacementTestScoreCondition(name, score, 100)
}

// VisitApScoreCondition
//
// Rule: 'AP score of' AT_LEAST SMALL_INT
func (v *RequisiteVisitor) VisitApScoreCondition(ctx *parser.ApScoreConditionContext) any {
	return v.visitApScoreCondition(ctx)
}

func (v *RequisiteVisitor) visitApScoreCondition(ctx *parser.ApScoreConditionContext) *conditions.APScoreCondition {
	score := mapInt(ctx.SMALL_INT().GetText())
	return conditions.NewAPScoreCondition(score)
}

// VisitAleksScoreCondition
//
// Rule: A_SCORE_OF INT '%' 'on ALEKS math placement exam'
func (v *RequisiteVisitor) VisitAleksScoreCondition(ctx *parser.AleksScoreConditionContext) any {
	return v.visitAleksScoreCondition(ctx)
}

func (v *RequisiteVisitor) visitAleksScoreCondition(ctx *parser.AleksScoreConditionContext) *conditions.AleksScoreCondition {
	score := mapInt(ctx.INT().GetText())
	return conditions.NewAleksScoreCondition(score)
}

// ======================= Group condition =======================

// VisitBothGroupCondition
//
// Rule: 'Students in both' group '/' group STUDENT_GROUP ONLY_KW
func (v *RequisiteVisitor) VisitBothGroupCondition(ctx *parser.BothGroupConditionContext) any {
	return v.visitBothGroupCondition(ctx)
}

func (v *RequisiteVisitor) visitBothGroupCondition(ctx *parser.BothGroupConditionContext) conditions.Condition {
	group1 := v.Visit(ctx.Group(0)).(utils.StudentGroup)
	group2 := v.Visit(ctx.Group(1)).(utils.StudentGroup)

	return conditions.NewAndCondition(
		conditions.NewStudentGroupCondition(group1),
		conditions.NewStudentGroupCondition(group2),
	)
}

// VisitGroupListCondition
//
// Rule: group (AND group)* (STUDENT_GROUP | STUDENTS) ONLY_KW?
func (v *RequisiteVisitor) VisitGroupListCondition(ctx *parser.GroupListConditionContext) any {
	return v.visitGroupListCondition(ctx)
}

func (v *RequisiteVisitor) visitGroupListCondition(ctx *parser.GroupListConditionContext) conditions.Condition {
	groupsConditions := make([]conditions.Condition, len(ctx.AllGroup()))
	for i, group := range ctx.AllGroup() {
		groupsConditions[i] = conditions.NewStudentGroupCondition(v.Visit(group).(utils.StudentGroup))
	}

	return conditions.NewOrCondition(groupsConditions...)
}

// VisitSingleGroupCondition
//
// Rule: group
func (v *RequisiteVisitor) VisitSingleGroupCondition(ctx *parser.SingleGroupConditionContext) any {
	return v.visitSingleGroupCondition(ctx)
}

func (v *RequisiteVisitor) visitSingleGroupCondition(ctx *parser.SingleGroupConditionContext) conditions.Condition {
	group := v.Visit(ctx.Group()).(utils.StudentGroup)
	return conditions.NewStudentGroupCondition(group)
}

// ======================= Misc conditions =======================

// VisitConcurrentEnrollmentCondition
//
// Rule: CONCURRENT_ENROLLMENT_IN course
func (v *RequisiteVisitor) VisitConcurrentEnrollmentCondition(ctx *parser.ConcurrentEnrollmentConditionContext) any {
	return v.visitConcurrentEnrollmentCondition(ctx)
}

func (v *RequisiteVisitor) visitConcurrentEnrollmentCondition(ctx *parser.ConcurrentEnrollmentConditionContext) conditions.Condition {
	courses := extractCoursesFromCourseList(v.Visit(ctx.Course()).(conditions.Condition))
	conds := make([]conditions.Condition, len(courses))
	for i, course := range courses {
		conds[i] = conditions.NewConcurrentEnrollmentCondition(course)
	}
	return conditions.NewAndCondition(conds...)
}

// VisitExactSectionCondition
//
// Rule: course PERIOD SECTION_NUMBER
func (v *RequisiteVisitor) VisitExactSectionCondition(ctx *parser.ExactSectionConditionContext) any {
	return v.visitExactSectionCondition(ctx)
}

func (v *RequisiteVisitor) visitExactSectionCondition(ctx *parser.ExactSectionConditionContext) conditions.Condition {
	courses := extractCoursesFromCourseList(v.Visit(ctx.Course()).(conditions.Condition))
	conds := make([]conditions.Condition, len(courses))
	for i, course := range courses {
		conds[i] = conditions.NewConcurrentEnrollmentCondition(
			utils.Course{
				Prefix:  course.Prefix,
				Number:  course.Number,
				Section: ctx.SECTION_NUMBER().GetText(),
			},
		)
	}
	return conditions.NewAndCondition(conds...)
}

// VisitWorkshopSectionCondition
//
// Rule: course 'workshop' SECTION_NUMBER
func (v *RequisiteVisitor) VisitWorkshopSectionCondition(ctx *parser.WorkshopSectionConditionContext) any {
	return v.visitWorkshopSectionCondition(ctx)
}

func (v *RequisiteVisitor) visitWorkshopSectionCondition(ctx *parser.WorkshopSectionConditionContext) conditions.Condition {
	courses := extractCoursesFromCourseList(v.Visit(ctx.Course()).(conditions.Condition))
	conds := make([]conditions.Condition, len(courses))
	for i, course := range courses {
		conds[i] = conditions.NewConcurrentEnrollmentCondition(
			utils.Course{
				Prefix:  course.Prefix,
				Number:  course.Number,
				Section: ctx.SECTION_NUMBER().GetText(),
			},
		)
	}
	return conditions.NewAndCondition(conds...)
}

// VisitAnyPreviousMajorCourseCondition
//
// Rule: ANY_PREVIOUS PREFIX COURSE_KW
func (v *RequisiteVisitor) VisitAnyPreviousMajorCourseCondition(ctx *parser.AnyPreviousMajorCourseConditionContext) any {
	return v.visitAnyPreviousMajorCourseCondition(ctx)
}

func (v *RequisiteVisitor) visitAnyPreviousMajorCourseCondition(ctx *parser.AnyPreviousMajorCourseConditionContext) *conditions.AnyPreviousMajorCourseCondition {
	prefix := ctx.PREFIX().GetText()

	return conditions.NewAnyPreviousMajorCourseCondition(prefix)
}

// VisitAcademicPlanCondition
//
// Rule: 'Academic Plan' (NOT_EQUAL | EQUAL) 'to' ACADEMIC_PLAN
func (v *RequisiteVisitor) VisitAcademicPlanCondition(ctx *parser.AcademicPlanConditionContext) any {
	return v.visitAcademicPlanCondition(ctx)
}

func (v *RequisiteVisitor) visitAcademicPlanCondition(ctx *parser.AcademicPlanConditionContext) *conditions.AcademicYearCondition {
	plan := ctx.ACADEMIC_PLAN().GetText()
	equal := ctx.EQUAL() != nil
	return conditions.NewAcademicYearCondition(plan, equal)
}
