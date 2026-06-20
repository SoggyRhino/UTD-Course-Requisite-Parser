package rules

import (
	"encoding/json"
	"fmt"
	"parser/objects/constants"
)

type CreditForRule struct {
	Courses CourseCollection `json:"courses,omitempty"`
}

func (c *CreditForRule) Fulfils(userInfo constants.UserInfo) *constants.Evaluation {
	return &constants.Evaluation{
		Name:    "Credit For Rule",
		Status:  constants.StatusDefiniteFail,
		Summary: "Not implemented",
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

func (c *CreditForRule) UnmarshalJSON(b []byte) error {
	var raw struct {
		Courses json.RawMessage `json:"courses"`
	}
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
	isCourseCollection() bool //todo change this actual work when evaluating
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

func (a *AndCourseCollection) UnmarshalJSON(b []byte) error {
	var raw struct {
		Courses []json.RawMessage `json:"courses"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	a.Courses = make([]CourseCollection, len(raw.Courses))
	for i, rawCol := range raw.Courses {
		col, err := UnmarshalCourseCollection(rawCol)
		if err != nil {
			return err
		}
		a.Courses[i] = col
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

func (a *AndCourseCollection) isCourseCollection() bool {
	return true
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

func (o *OrCourseCollection) UnmarshalJSON(b []byte) error {
	var raw struct {
		Courses []json.RawMessage `json:"courses"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	o.Courses = make([]CourseCollection, len(raw.Courses))
	for i, rawCol := range raw.Courses {
		col, err := UnmarshalCourseCollection(rawCol)
		if err != nil {
			return err
		}
		o.Courses[i] = col
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

func (o *OrCourseCollection) isCourseCollection() bool {
	return true
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

func (s *SimpleCourseCollection) isCourseCollection() bool {
	return true
}
