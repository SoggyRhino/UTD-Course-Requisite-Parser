package conditions

type ConcurrentEnrollmentCondition struct {
	Course Course
}

func NewConcurrentEnrollmentCondition(course Course) *ConcurrentEnrollmentCondition {
	return &ConcurrentEnrollmentCondition{
		Course: course,
	}
}

func (c *ConcurrentEnrollmentCondition) Fulfils(userInfo UserInfo) (bool, error) {
	return false, nil
}

type ExactSectionCondition struct {
	Course Course
}

func NewExactSectionCondition(course Course) *ExactSectionCondition {
	return &ExactSectionCondition{
		Course: course,
	}
}

func (c *ExactSectionCondition) Fulfils(userInfo UserInfo) (bool, error) {
	return false, nil
}

type AnyPreviousMajorCourseCondition struct {
	Prefix string
}

func NewAnyPreviousMajorCourseCondition(prefix string) *AnyPreviousMajorCourseCondition {
	return &AnyPreviousMajorCourseCondition{
		Prefix: prefix,
	}
}

func (c *AnyPreviousMajorCourseCondition) Fulfils(userInfo UserInfo) (bool, error) {
	return false, nil
}

type AcademicYearCondition struct {
	Plan  string
	Equal bool // must be in plan or must not be in plan
}

func NewAcademicYearCondition(plan string, equal bool) *AcademicYearCondition {
	return &AcademicYearCondition{
		Plan:  plan,
		Equal: equal,
	}
}

func (c *AcademicYearCondition) Fulfils(userInfo UserInfo) (bool, error) {
	return false, nil
}
