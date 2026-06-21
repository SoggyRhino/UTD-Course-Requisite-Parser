package rules

import (
	"encoding/json"
	"fmt"
	"parser/objects/constants"
	"strings"
)

type CreditForRule struct {
	Courses CourseCollection `json:"courses,omitempty"`
}

func (c *CreditForRule) Fulfils(userInfo constants.UserInfo) *constants.Evaluation {
	if c.Courses.HasCredit(userInfo) {
		return &constants.Evaluation{
			Name:    "Credit For Rule",
			Status:  constants.StatusDefiniteFail,
			Summary: fmt.Sprintf("Student violates credit for rule: %s", c.Courses.String()),
		}
	}
	return &constants.Evaluation{
		Name:    "Credit For Rule",
		Status:  constants.StatusPass,
		Summary: "Student satisfies credit for rule",
	}
}

func (c *CreditForRule) MarshalJSON() ([]byte, error) {
	type Alias CreditForRule
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "credit_for",
		Alias: (*Alias)(c),
	})
}

type rawCreditForRule struct {
	Courses json.RawMessage `json:"courses"`
}

func (c *CreditForRule) UnmarshalJSON(b []byte) error {
	var raw rawCreditForRule
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	col, err := UnmarshalCourseCollection(raw.Courses)
	if err != nil {
		return err
	}
	c.Courses = col
	return nil
}

func NewCreditForRule(courses CourseCollection) *CreditForRule {
	return &CreditForRule{Courses: courses}
}

type CourseCollection interface {
	HasCredit(userInfo constants.UserInfo) bool
	String() string
}

type courseCollectionEnvelope struct {
	Type string `json:"type"`
}

func UnmarshalCourseCollection(b []byte) (CourseCollection, error) {
	var env courseCollectionEnvelope
	if err := json.Unmarshal(b, &env); err != nil {
		return nil, err
	}
	switch env.Type {
	case "and":
		var c AndCourseCollection
		return &c, json.Unmarshal(b, &c)
	case "or":
		var c OrCourseCollection
		return &c, json.Unmarshal(b, &c)
	case "simple":
		var c SimpleCourseCollection
		return &c, json.Unmarshal(b, &c)
	default:
		return nil, fmt.Errorf("unknown course collection type: %s", env.Type)
	}
}

type AndCourseCollection struct {
	Courses []CourseCollection `json:"courses,omitempty"`
}

func (a *AndCourseCollection) MarshalJSON() ([]byte, error) {
	type Alias AndCourseCollection
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "and",
		Alias: (*Alias)(a),
	})
}

type rawAndCourseCollection struct {
	Courses []any `json:"courses"`
}

func (a *AndCourseCollection) UnmarshalJSON(b []byte) error {
	var raw rawAndCourseCollection
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	a.Courses = make([]CourseCollection, len(raw.Courses))
	for i, rawCondAny := range raw.Courses {
		rawCond, _ := json.Marshal(rawCondAny)
		cond, err := UnmarshalCourseCollection(rawCond)
		if err != nil {
			return err
		}
		a.Courses[i] = cond
	}
	return nil
}

func NewAndCourseCollection(col1, col2 CourseCollection) *AndCourseCollection {
	var flattenedCourses []CourseCollection

	if and1, isAnd1 := col1.(*AndCourseCollection); isAnd1 {
		flattenedCourses = append(flattenedCourses, and1.Courses...)
	} else {
		flattenedCourses = append(flattenedCourses, col1)
	}

	if and2, isAnd2 := col2.(*AndCourseCollection); isAnd2 {
		flattenedCourses = append(flattenedCourses, and2.Courses...)
	} else {
		flattenedCourses = append(flattenedCourses, col2)
	}

	return &AndCourseCollection{
		Courses: flattenedCourses,
	}
}

func (a *AndCourseCollection) HasCredit(userInfo constants.UserInfo) bool {
	for _, col := range a.Courses {
		if !col.HasCredit(userInfo) {
			return false
		}
	}
	return true
}

func (a *AndCourseCollection) String() string {
	strs := make([]string, len(a.Courses))
	for i, col := range a.Courses {
		strs[i] = col.String()
	}
	return "(" + strings.Join(strs, " AND ") + ")"
}

type OrCourseCollection struct {
	Courses []CourseCollection `json:"courses,omitempty"`
}

func (o *OrCourseCollection) MarshalJSON() ([]byte, error) {
	type Alias OrCourseCollection
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "or",
		Alias: (*Alias)(o),
	})
}

type rawOrCourseCollection struct {
	Courses []any `json:"courses"`
}

func (o *OrCourseCollection) UnmarshalJSON(b []byte) error {
	var raw rawOrCourseCollection
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	o.Courses = make([]CourseCollection, len(raw.Courses))
	for i, rawCondAny := range raw.Courses {
		rawCond, _ := json.Marshal(rawCondAny)
		cond, err := UnmarshalCourseCollection(rawCond)
		if err != nil {
			return err
		}
		o.Courses[i] = cond
	}
	return nil
}

func NewOrCourseCollection(col1, col2 CourseCollection) *OrCourseCollection {
	var flattenedCourses []CourseCollection

	if or1, isOr1 := col1.(*OrCourseCollection); isOr1 {
		flattenedCourses = append(flattenedCourses, or1.Courses...)
	} else {
		flattenedCourses = append(flattenedCourses, col1)
	}

	if or2, isOr2 := col2.(*OrCourseCollection); isOr2 {
		flattenedCourses = append(flattenedCourses, or2.Courses...)
	} else {
		flattenedCourses = append(flattenedCourses, col2)
	}

	return &OrCourseCollection{
		Courses: flattenedCourses,
	}
}

func (o *OrCourseCollection) HasCredit(userInfo constants.UserInfo) bool {
	for _, col := range o.Courses {
		if col.HasCredit(userInfo) {
			return true
		}
	}
	return false
}

func (o *OrCourseCollection) String() string {
	strs := make([]string, len(o.Courses))
	for i, col := range o.Courses {
		strs[i] = col.String()
	}
	return "(" + strings.Join(strs, " OR ") + ")"
}

type SimpleCourseCollection struct {
	Course []constants.Course `json:"course,omitempty"`
}

func (s *SimpleCourseCollection) MarshalJSON() ([]byte, error) {
	type Alias SimpleCourseCollection
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "simple",
		Alias: (*Alias)(s),
	})
}

func NewSimpleCourseCollection(course []constants.Course) *SimpleCourseCollection {
	return &SimpleCourseCollection{Course: course}
}

func (s *SimpleCourseCollection) HasCredit(userInfo constants.UserInfo) bool {
	for _, c := range s.Course {
		if _, taken := userInfo.Taken[c]; taken {
			return true
		}
		for _, enrolled := range userInfo.CurrentEnrollment {
			if enrolled == c {
				return true
			}
		}
	}
	return false
}

func (s *SimpleCourseCollection) String() string {
	strs := make([]string, len(s.Course))
	for i, c := range s.Course {
		strs[i] = c.String()
	}
	return strings.Join(strs, ", ")
}
