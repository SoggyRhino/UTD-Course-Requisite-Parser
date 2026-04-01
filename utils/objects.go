package utils

type Course struct {
	Prefix  string
	Number  string
	Section string
}

type Grade string
type UserInfo struct {
	Taken map[Course]Grade
}

type DegreeLevel string

const (
	Undergraduate DegreeLevel = "Undergraduate"
	Graduate      DegreeLevel = "Graduate"
	PhD           DegreeLevel = "PhD"
	AnyDegree     DegreeLevel = "Any Degree"
)

type StudentGroup string

const (
	ComputerScholarsProgram StudentGroup = "Computer Scholars Program"
	CollegeVHonors                       = "Collegium V Honors"
	LiberalArtsHonors                    = "Liberal Arts Honors"
	SCVG                                 = "SCVG"
	DMHP                                 = "DMHP"
	DLAH                                 = "DLAH"
)

type GradeLevel string

const (
	Senior    GradeLevel = "Senior"
	Junior    GradeLevel = "Junior"
	Sophomore GradeLevel = "Sophomore"
	Freshman  GradeLevel = "Freshman"
	AnyGrade  GradeLevel = "Any Level"
)

type Consent string

const (
	InstructorConsent Consent = "Instructor Consent"
	DepartmentConsent         = "Department Consent"
	UTeachConsent             = "UTeach Consent"
)

type Standing string

const (
	GoodAcademicStanding  Standing = "Good Academic Standing"
	UpperDivisionStanding          = "Upper Division Standing"
)
