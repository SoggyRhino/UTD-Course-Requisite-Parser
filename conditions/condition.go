package conditions

import (
	"encoding/json"
	"fmt"
	"parser/constants"
)

type Condition interface {
	Fulfils(constants.UserInfo) (bool, error)
}

type GradedCondition interface {
	Condition
	AppendGrade(constants.Grade)
}

type conditionEnvelope struct {
	Type string `json:"type"`
}

func UnmarshalCondition(b []byte) (Condition, error) {
	var env conditionEnvelope
	if err := json.Unmarshal(b, &env); err != nil {
		return nil, err
	}
	switch env.Type {
	case "alternative":
		var c AlternativeCondition
		return &c, json.Unmarshal(b, &c)
	case "consent":
		var c ConsentCondition
		return &c, json.Unmarshal(b, &c)
	case "course":
		var c CourseCondition
		return &c, json.Unmarshal(b, &c)
	case "core":
		var c CoreCondition
		return &c, json.Unmarshal(b, &c)
	case "credit_hours":
		var c CreditHoursCondition
		return &c, json.Unmarshal(b, &c)
	case "credit_hours_from":
		var c CreditHoursFromCondition
		return &c, json.Unmarshal(b, &c)
	case "upper_division_courses":
		var c UpperDivisionCoursesCondition
		return &c, json.Unmarshal(b, &c)
	case "research":
		var c ResearchCondition
		return &c, json.Unmarshal(b, &c)
	case "n_courses":
		var c NCoursesCondition
		return &c, json.Unmarshal(b, &c)
	case "major":
		var c MajorCondition
		return &c, json.Unmarshal(b, &c)
	case "degree":
		var c DegreeCondition
		return &c, json.Unmarshal(b, &c)
	case "gpa":
		var c GPACondition
		return &c, json.Unmarshal(b, &c)
	case "or":
		var c OrCondition
		return &c, json.Unmarshal(b, &c)
	case "and":
		var c AndCondition
		return &c, json.Unmarshal(b, &c)
	case "concurrent_enrollment":
		var c ConcurrentEnrollmentCondition
		return &c, json.Unmarshal(b, &c)
	case "exact_section":
		var c ExactSectionCondition
		return &c, json.Unmarshal(b, &c)
	case "any_previous_major_course":
		var c AnyPreviousMajorCourseCondition
		return &c, json.Unmarshal(b, &c)
	case "academic_year":
		var c AcademicYearCondition
		return &c, json.Unmarshal(b, &c)
	case "placement_test_score":
		var c PlacementTestScoreCondition
		return &c, json.Unmarshal(b, &c)
	case "ap_score":
		var c APScoreCondition
		return &c, json.Unmarshal(b, &c)
	case "aleks_score":
		var c AleksScoreCondition
		return &c, json.Unmarshal(b, &c)
	case "grade_level":
		var c GradeLevelCondition
		return &c, json.Unmarshal(b, &c)
	case "graduate_standing_in":
		var c GraduateStandingInCondition
		return &c, json.Unmarshal(b, &c)
	case "generic_standing":
		var c GenericStandingCondition
		return &c, json.Unmarshal(b, &c)
	case "student_group":
		var c StudentGroupCondition
		return &c, json.Unmarshal(b, &c)
	default:
		return nil, fmt.Errorf("unknown condition type: %s", env.Type)
	}
}
