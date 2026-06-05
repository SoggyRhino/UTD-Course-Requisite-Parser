package constants

import (
	"strconv"
	"strings"
)

type Course struct {
	Prefix  string `json:"prefix,omitempty"`
	Number  string `json:"number,omitempty"`
	Section string `json:"section,omitempty"`
}

func (c Course) Hours() int {
	if h, err := strconv.Atoi(c.Number[1:2]); err == nil {
		return h
	}
	return 3
}

func (c Course) CourseNumberInt() int {
	n, err := strconv.Atoi(strings.TrimSpace(c.Number))
	if err != nil {
		return -1
	}
	return n
}

func (c Course) IsUpperDivision() bool {
	return c.CourseNumberInt() >= 3000
}

type Grade string
type UserInfo struct {
	Taken             map[Course]Grade `json:"taken,omitempty"`
	CurrentEnrollment []Course         `json:"current_enrollment,omitempty"`
	GPA               float64          `json:"gpa,omitempty"`
	Major             string           `json:"major,omitempty"`
	GradeLevel        GradeLevel       `json:"grade_level,omitempty"`
	Standing          Standing         `json:"standing,omitempty"`
	Groups            []StudentGroup   `json:"groups,omitempty"`
	DegreeLevel       DegreeLevel      `json:"degree_level,omitempty"`
	AcademicPlan      string           `json:"academic_plan,omitempty"`
	TotalSCH          int              `json:"total_sch,omitempty"`
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

type Notice string

const (
	ExcludeDMHPLLCNotice          Notice = "non-DMHP/non-LLC (DMLC, DFLC, DHLC) Student Group Only"
	ExactCoReqNotice                     = "Matching Corequisite Section Required"
	ComputerScholarsProgramNotice        = "Computer Scholars Program Only"
)
