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
	CollegeVHonors    StudentGroup = "Collegium V Honors"
	LiberalArtsHonors              = "Liberal Arts Honors"
	SCVG                           = "SCVG"
	DMHP                           = "DMHP"
	DLAH                           = "DLAH"
)

type GradeLevel string

const (
	Senior    GradeLevel = "Senior"
	Junior    GradeLevel = "Junior"
	Sophomore GradeLevel = "Sophomore"
	Freshman  GradeLevel = "Freshman"
	AnyGrade  GradeLevel = "Any Level"
)
