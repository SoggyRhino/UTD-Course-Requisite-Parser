package constants

type UserInfo struct {
	Taken             map[Course]Grade `json:"taken,omitempty"`
	CurrentEnrollment []Course         `json:"current_enrollment,omitempty"`
	GPA               float64          `json:"gpa,omitempty"`
	Major             string           `json:"major,omitempty"`
	GradeLevel        GradeLevel       `json:"grade_level,omitempty"`
	Standing          []Standing       `json:"standing,omitempty"`
	Groups            []StudentGroup   `json:"groups,omitempty"`
	DegreeLevel       DegreeLevel      `json:"degree_level,omitempty"`
	AcademicPlan      string           `json:"academic_plan,omitempty"`
	TotalSCH          int              `json:"total_sch,omitempty"`
}
