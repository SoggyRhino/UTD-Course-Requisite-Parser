package conditions

import (
	"encoding/json"
	"fmt"
	"parser/objects/constants"
)

var gradeIndex = map[constants.Grade]int{
	"D-": 0,
	"D":  1,
	"D+": 2,
	"C-": 3,
	"C":  4,
	"C+": 5,
	"B-": 6,
	"B":  7,
	"B+": 8,
	"A-": 9,
	"A":  10,
}

func gradeAtLeast(actual, min constants.Grade) bool {
	ai, aOk := gradeIndex[actual]
	mi, mOk := gradeIndex[min]
	if !aOk || !mOk {
		return false
	}
	return ai >= mi
}

type CourseCondition struct {
	Course   constants.Course `json:"course,omitempty"`
	MinGrade constants.Grade  `json:"min_grade,omitempty"`
}

func (c *CourseCondition) MarshalJSON() ([]byte, error) {
	type Alias CourseCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "course",
		Alias: (*Alias)(c),
	})
}

func NewCourseCondition(prefix, number, grade string) *CourseCondition {
	return &CourseCondition{
		Course: constants.Course{
			Prefix: prefix,
			Number: number,
		},
		MinGrade: constants.Grade(grade),
	}
}

func (c *CourseCondition) AppendGrade(grade constants.Grade) {
	c.MinGrade = grade
}

func (c *CourseCondition) Fulfils(info constants.UserInfo, allowCoReq bool) *constants.Evaluation {
	courseLabel := fmt.Sprintf("%s %s", c.Course.Prefix, c.Course.Number)

	for taken, grade := range info.Taken {
		if taken.Prefix == c.Course.Prefix && taken.Number == c.Course.Number {
			if c.MinGrade == "" || gradeAtLeast(grade, c.MinGrade) {
				return &constants.Evaluation{
					Name:    courseLabel,
					Status:  constants.StatusPass,
					Summary: fmt.Sprintf("Completed %s with grade %s", courseLabel, grade),
				}
			}
			return &constants.Evaluation{
				Name:    courseLabel,
				Status:  constants.StatusDefiniteFail,
				Summary: fmt.Sprintf("Completed %s but grade %s does not meet minimum %s", courseLabel, grade, c.MinGrade),
			}
		}
	}
	for _, enrolled := range info.CurrentEnrollment {
		if enrolled.Prefix == c.Course.Prefix && enrolled.Number == c.Course.Number {
			if allowCoReq {
				return &constants.Evaluation{
					Name:    courseLabel,
					Status:  constants.StatusPass,
					Summary: fmt.Sprintf("Currently enrolled in %s (co-req allowed)", courseLabel),
				}
			}
			return &constants.Evaluation{
				Name:    courseLabel,
				Status:  constants.StatusPossibleFail,
				Summary: fmt.Sprintf("Currently enrolled in %s — awaiting final grade", courseLabel),
			}
		}
	}

	return &constants.Evaluation{
		Name:    courseLabel,
		Status:  constants.StatusDefiniteFail,
		Summary: fmt.Sprintf("Have not taken %s", courseLabel),
	}
}

type CoreCondition struct {
	CoreNumber    string `json:"core_number,omitempty"`
	CoreTitle     string `json:"core_title,omitempty"`
	SemesterHours int    `json:"semester_hours,omitempty"`
}

func (c *CoreCondition) MarshalJSON() ([]byte, error) {
	type Alias CoreCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "core",
		Alias: (*Alias)(c),
	})
}

func NewCoreCondition(coreNumber, coreTitle string) *CoreCondition {
	return &CoreCondition{
		CoreNumber: coreNumber,
		CoreTitle:  coreTitle,
	}
}

func NewCoreConditionWithSemesterHours(coreNumber, coreTitle string, semesterHours int) *CoreCondition {
	return &CoreCondition{
		CoreNumber:    coreNumber,
		CoreTitle:     coreTitle,
		SemesterHours: semesterHours,
	}
}

func (c *CoreCondition) Fulfils(_ constants.UserInfo, _ bool) *constants.Evaluation {
	label := c.CoreTitle
	if label == "" {
		label = c.CoreNumber
	}

	//todo build list of core classes
	summary := fmt.Sprintf("Cannot verify core requirement %q automatically", label)
	if c.SemesterHours > 0 {
		summary = fmt.Sprintf("Cannot verify %d SCH of core requirement %q automatically", c.SemesterHours, label)
	}
	return &constants.Evaluation{
		Name:    fmt.Sprintf("Core %s", label),
		Status:  constants.StatusUnknown,
		Summary: summary,
	}
}

type CreditHoursCondition struct {
	Hours int `json:"hours"`
}

func (c *CreditHoursCondition) MarshalJSON() ([]byte, error) {
	type Alias CreditHoursCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "credit_hours",
		Alias: (*Alias)(c),
	})
}

func NewCreditHoursCondition(hours int) *CreditHoursCondition {
	return &CreditHoursCondition{Hours: hours}
}

func (c *CreditHoursCondition) Fulfils(info constants.UserInfo, _ bool) *constants.Evaluation {
	if info.TotalSCH >= c.Hours {
		return &constants.Evaluation{
			Name:    "Credit Hours",
			Status:  constants.StatusPass,
			Summary: fmt.Sprintf("Has %d SCH (requires %d)", info.TotalSCH, c.Hours),
		}
	}
	return &constants.Evaluation{
		Name:    "Credit Hours",
		Status:  constants.StatusDefiniteFail,
		Summary: fmt.Sprintf("Has %d SCH but requires %d", info.TotalSCH, c.Hours),
	}
}

type CreditHoursFromCondition struct {
	Hours   int                `json:"hours,omitempty"`
	Courses []constants.Course `json:"courses,omitempty"`
}

func (c *CreditHoursFromCondition) MarshalJSON() ([]byte, error) {
	type Alias CreditHoursFromCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "credit_hours_from",
		Alias: (*Alias)(c),
	})
}

func NewCreditHoursFromCondition(hours int, courses []constants.Course) *CreditHoursFromCondition {
	return &CreditHoursFromCondition{Hours: hours, Courses: courses}
}

func (c *CreditHoursFromCondition) Fulfils(info constants.UserInfo, _ bool) *constants.Evaluation {
	var completed, inProgress []string
	var earnedHours, projectedHours int

	for _, required := range c.Courses {
		label := fmt.Sprintf("%s %s", required.Prefix, required.Number)
		found := false
		for taken := range info.Taken {
			if taken.Prefix == required.Prefix && taken.Number == required.Number {
				completed = append(completed, label)
				earnedHours += taken.Hours()
				found = true
				break
			}
		}
		if found {
			continue
		}
		for _, enrolled := range info.CurrentEnrollment {
			if enrolled.Prefix == required.Prefix && enrolled.Number == required.Number {
				inProgress = append(inProgress, label)
				projectedHours += enrolled.Hours()
				break
			}
		}
	}

	if earnedHours >= c.Hours {
		return &constants.Evaluation{
			Name:    "Credit Hours From",
			Status:  constants.StatusPass,
			Summary: fmt.Sprintf("Earned %d SCH from specified courses (requires %d)", earnedHours, c.Hours),
		}
	}
	if projectedHours >= c.Hours {
		return &constants.Evaluation{
			Name:    "Credit Hours From",
			Status:  constants.StatusPossibleFail,
			Summary: fmt.Sprintf("Earned %d SCH; %d more in progress — may reach required %d SCH", earnedHours, projectedHours-earnedHours, c.Hours),
		}
	}
	return &constants.Evaluation{
		Name:    "Credit Hours From",
		Status:  constants.StatusDefiniteFail,
		Summary: fmt.Sprintf("Only %d SCH earned from specified courses (requires %d)", earnedHours, c.Hours),
	}
}

type UpperDivisionCoursesCondition struct {
	Hours  int    `json:"hours,omitempty"`
	Count  int    `json:"count,omitempty"`
	Prefix string `json:"prefix,omitempty"`
}

func (c *UpperDivisionCoursesCondition) MarshalJSON() ([]byte, error) {
	type Alias UpperDivisionCoursesCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "upper_division_courses",
		Alias: (*Alias)(c),
	})
}

func NewUpperDivisionCreditHoursCondition(hours int, prefix string) *UpperDivisionCoursesCondition {
	return &UpperDivisionCoursesCondition{Hours: hours, Prefix: prefix}
}

func NewUpperDivisionCountCondition(count int, prefix string) *UpperDivisionCoursesCondition {
	return &UpperDivisionCoursesCondition{Count: count, Prefix: prefix}
}

func (c *UpperDivisionCoursesCondition) Fulfils(info constants.UserInfo, _ bool) *constants.Evaluation {
	prefixLabel := c.Prefix
	if prefixLabel == "" {
		prefixLabel = "any prefix"
	}

	var earned, udCount int
	for taken := range info.Taken {
		if c.Prefix != "" && taken.Prefix != c.Prefix {
			continue
		}
		if taken.IsUpperDivision() {
			earned += taken.Hours()
			udCount++
		}
	}

	if c.Hours > 0 {
		if earned >= c.Hours {
			return &constants.Evaluation{
				Name:    "Upper Division Courses",
				Status:  constants.StatusPass,
				Summary: fmt.Sprintf("Has %d upper-division SCH in %s (requires %d)", earned, prefixLabel, c.Hours),
			}
		}
		return &constants.Evaluation{
			Name:    "Upper Division Courses",
			Status:  constants.StatusDefiniteFail,
			Summary: fmt.Sprintf("Has %d upper-division SCH in %s (requires %d)", earned, prefixLabel, c.Hours),
		}
	}

	if c.Count > 0 {
		if udCount >= c.Count {
			return &constants.Evaluation{
				Name:    "Upper Division Courses",
				Status:  constants.StatusPass,
				Summary: fmt.Sprintf("Has completed %d upper-division courses in %s (requires %d)", udCount, prefixLabel, c.Count),
			}
		}
		return &constants.Evaluation{
			Name:    "Upper Division Courses",
			Status:  constants.StatusDefiniteFail,
			Summary: fmt.Sprintf("Has completed %d upper-division courses in %s (requires %d)", udCount, prefixLabel, c.Count),
		}
	}

	return &constants.Evaluation{
		Name:    "Upper Division Courses",
		Status:  constants.StatusInvalidRule,
		Summary: "UpperDivisionCoursesCondition has neither Hours nor Count set",
	}
}

type ResearchCondition struct {
	Hours       int                   `json:"hours,omitempty"`
	DegreeLevel constants.DegreeLevel `json:"degree_level,omitempty"`
}

func (c *ResearchCondition) MarshalJSON() ([]byte, error) {
	type Alias ResearchCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "research",
		Alias: (*Alias)(c),
	})
}

func NewResearchCondition(hours int, degreeLevel constants.DegreeLevel) *ResearchCondition {
	return &ResearchCondition{Hours: hours, DegreeLevel: degreeLevel}
}

func (c *ResearchCondition) Fulfils(info constants.UserInfo, _ bool) *constants.Evaluation {
	if c.DegreeLevel != constants.AnyDegree && c.DegreeLevel != "" {
		if info.DegreeLevel != c.DegreeLevel {
			return &constants.Evaluation{
				Name:    "Research",
				Status:  constants.StatusDefiniteFail,
				Summary: fmt.Sprintf("Requires %s standing; student is %s", c.DegreeLevel, info.DegreeLevel),
			}
		}
	}
	return &constants.Evaluation{
		Name:    "Research",
		Status:  constants.StatusUnknown,
		Summary: fmt.Sprintf("Cannot automatically verify %d research hours for %s students", c.Hours, c.DegreeLevel),
	}
}

type NCoursesCondition struct {
	N       int                `json:"n,omitempty"`
	Courses []constants.Course `json:"courses,omitempty"`
}

func (c *NCoursesCondition) MarshalJSON() ([]byte, error) {
	type Alias NCoursesCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "n_courses",
		Alias: (*Alias)(c),
	})
}

func NewNCoursesCondition(n int, courses []constants.Course) *NCoursesCondition {
	return &NCoursesCondition{N: n, Courses: courses}
}

func (c *NCoursesCondition) Fulfils(info constants.UserInfo, _ bool) *constants.Evaluation {
	var children []constants.Evaluation
	passCount := 0
	possibleCount := 0

	for _, required := range c.Courses {
		label := fmt.Sprintf("%s %s", required.Prefix, required.Number)
		takenEval := func() *constants.Evaluation {
			for taken := range info.Taken {
				if taken.Prefix == required.Prefix && taken.Number == required.Number {
					return &constants.Evaluation{Name: label, Status: constants.StatusPass, Summary: fmt.Sprintf("Completed %s", label)}
				}
			}
			for _, enrolled := range info.CurrentEnrollment {
				if enrolled.Prefix == required.Prefix && enrolled.Number == required.Number {
					return &constants.Evaluation{Name: label, Status: constants.StatusPossibleFail, Summary: fmt.Sprintf("Currently enrolled in %s", label)}
				}
			}
			return &constants.Evaluation{Name: label, Status: constants.StatusDefiniteFail, Summary: fmt.Sprintf("Have not taken %s", label)}
		}()

		children = append(children, *takenEval)
		switch takenEval.Status {
		case constants.StatusPass:
			passCount++
		case constants.StatusPossibleFail:
			possibleCount++
		}
	}

	if passCount >= c.N {
		return &constants.Evaluation{
			Name:     "N Courses",
			Status:   constants.StatusPass,
			Summary:  fmt.Sprintf("Completed %d of required %d courses", passCount, c.N),
			Children: children,
		}
	}
	if passCount+possibleCount >= c.N {
		return &constants.Evaluation{
			Name:     "N Courses",
			Status:   constants.StatusPossibleFail,
			Summary:  fmt.Sprintf("Completed %d courses; %d in progress — may satisfy requirement of %d", passCount, possibleCount, c.N),
			Children: children,
		}
	}
	return &constants.Evaluation{
		Name:     "N Courses",
		Status:   constants.StatusDefiniteFail,
		Summary:  fmt.Sprintf("Completed only %d of required %d courses from the specified list", passCount, c.N),
		Children: children,
	}
}
