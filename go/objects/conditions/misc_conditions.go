package conditions

import (
	"encoding/json"
	"fmt"
	"parser/objects/constants"
)

type ConcurrentEnrollmentCondition struct {
	Course constants.Course `json:"course,omitempty"`
}

func (c *ConcurrentEnrollmentCondition) MarshalJSON() ([]byte, error) {
	type Alias ConcurrentEnrollmentCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "concurrent_enrollment",
		Alias: (*Alias)(c),
	})
}

func NewConcurrentEnrollmentCondition(course constants.Course) *ConcurrentEnrollmentCondition {
	return &ConcurrentEnrollmentCondition{
		Course: course,
	}
}

func (c *ConcurrentEnrollmentCondition) Fulfils(userInfo constants.UserInfo) *constants.Evaluation {
	label := fmt.Sprintf("%s %s", c.Course.Prefix, c.Course.Number)
	for taken := range userInfo.Taken {
		if taken.Prefix == c.Course.Prefix && taken.Number == c.Course.Number {
			return &constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: fmt.Sprintf("Already completed %s", label),
			}
		}
	}
	for _, enrolled := range userInfo.CurrentEnrollment {
		if enrolled.Prefix == c.Course.Prefix && enrolled.Number == c.Course.Number {
			return &constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: fmt.Sprintf("Currently enrolled in %s", label),
			}
		}
	}
	return &constants.Evaluation{
		Status:  constants.StatusPossibleFail,
		Summary: fmt.Sprintf("Requires concurrent enrollment in %s", label),
	}
}

type ExactSectionCondition struct {
	Course constants.Course `json:"course,omitempty"`
}

func (c *ExactSectionCondition) MarshalJSON() ([]byte, error) {
	type Alias ExactSectionCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "exact_section",
		Alias: (*Alias)(c),
	})
}

func NewExactSectionCondition(course constants.Course) *ExactSectionCondition {
	return &ExactSectionCondition{
		Course: course,
	}
}

func (c *ExactSectionCondition) Fulfils(userInfo constants.UserInfo) *constants.Evaluation {
	label := fmt.Sprintf("%s %s section %s", c.Course.Prefix, c.Course.Number, c.Course.Section)
	for taken := range userInfo.Taken {
		if taken.Prefix == c.Course.Prefix && taken.Number == c.Course.Number && (c.Course.Section == "" || taken.Section == c.Course.Section) {
			return &constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: fmt.Sprintf("Completed %s", label),
			}
		}
	}
	for _, enrolled := range userInfo.CurrentEnrollment {
		if enrolled.Prefix == c.Course.Prefix && enrolled.Number == c.Course.Number && (c.Course.Section == "" || enrolled.Section == c.Course.Section) {
			return &constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: fmt.Sprintf("Currently enrolled in %s", label),
			}
		}
	}
	return &constants.Evaluation{
		Status:  constants.StatusDefiniteFail,
		Summary: fmt.Sprintf("Must be in %s", label),
	}
}

type AnyPreviousMajorCourseCondition struct {
	Prefix string `json:"prefix,omitempty"`
}

func (c *AnyPreviousMajorCourseCondition) MarshalJSON() ([]byte, error) {
	type Alias AnyPreviousMajorCourseCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "any_previous_major_course",
		Alias: (*Alias)(c),
	})
}

func NewAnyPreviousMajorCourseCondition(prefix string) *AnyPreviousMajorCourseCondition {
	return &AnyPreviousMajorCourseCondition{
		Prefix: prefix,
	}
}

func (c *AnyPreviousMajorCourseCondition) Fulfils(userInfo constants.UserInfo) *constants.Evaluation {
	for taken := range userInfo.Taken {
		if taken.Prefix == c.Prefix {
			return &constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: fmt.Sprintf("Completed at least one course with prefix %s (%s %s)", c.Prefix, taken.Prefix, taken.Number),
			}
		}
	}
	return &constants.Evaluation{
		Status:  constants.StatusDefiniteFail,
		Summary: fmt.Sprintf("No previous courses with prefix %s found", c.Prefix),
	}
}

type AcademicYearCondition struct {
	Plan  string `json:"plan,omitempty"`
	Equal bool   `json:"equal,omitempty"`
}

func (c *AcademicYearCondition) MarshalJSON() ([]byte, error) {
	type Alias AcademicYearCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "academic_year",
		Alias: (*Alias)(c),
	})
}

func NewAcademicYearCondition(plan string, equal bool) *AcademicYearCondition {
	return &AcademicYearCondition{
		Plan:  plan,
		Equal: equal,
	}
}

func (c *AcademicYearCondition) Fulfils(userInfo constants.UserInfo) *constants.Evaluation {
	match := userInfo.AcademicPlan == c.Plan
	if c.Equal {
		if match {
			return &constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: fmt.Sprintf("Academic plan %s matches required %s", userInfo.AcademicPlan, c.Plan),
			}
		}
		return &constants.Evaluation{
			Status:  constants.StatusDefiniteFail,
			Summary: fmt.Sprintf("Academic plan %s does not match required %s", userInfo.AcademicPlan, c.Plan),
		}
	}

	// Not equal (usually "before" or "after" but here just checking mismatch)
	if !match {
		return &constants.Evaluation{
			Status:  constants.StatusPass,
			Summary: fmt.Sprintf("Academic plan %s is not %s", userInfo.AcademicPlan, c.Plan),
		}
	}
	return &constants.Evaluation{
		Status:  constants.StatusDefiniteFail,
		Summary: fmt.Sprintf("Academic plan %s matches prohibited %s", userInfo.AcademicPlan, c.Plan),
	}
}
