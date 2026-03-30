package conditions

type CourseCondition struct {
	Course   Course
	MinGrade Grade
}

func NewCourseCondition(prefix, number, grade string) *CourseCondition {
	return &CourseCondition{
		Course: Course{
			Prefix: prefix,
			Number: number,
		},
		MinGrade: Grade(grade),
	}
}

func (c *CourseCondition) Fulfils(userInfo UserInfo) (bool, error) {
	return false, nil
}

func (c *CourseCondition) AppendGrade(grade Grade) {
	c.MinGrade = grade
}

type CoreCondition struct {
	CoreNumber    string
	CoreTitle     string
	SemesterHours int
}

func (c *CoreCondition) Fulfils(userInfo UserInfo) (bool, error) {
	return false, nil
}

func NewCoreCondition(courseNumber, coreTitle string) *CoreCondition {
	return &CoreCondition{
		CoreNumber: courseNumber,
		CoreTitle:  coreTitle,
	}
}

func NewCoreConditionWithSemesterHours(courseNumber, coreTitle string, semesterHours int) *CoreCondition {
	return &CoreCondition{
		CoreNumber:    courseNumber,
		SemesterHours: semesterHours,
	}
}

type CreditHoursCondition struct {
	Hours int
}

func NewCreditHoursCondition(hours int) *CreditHoursCondition {
	return &CreditHoursCondition{
		Hours: hours,
	}
}

func (c *CreditHoursCondition) Fulfils(userInfo UserInfo) (bool, error) {
	return false, nil
}

type CreditHoursFromCondition struct {
	Hours   int
	Courses []Course
}

func NewCreditHoursFromCondition(hours int, courses []Course) *CreditHoursFromCondition {
	return &CreditHoursFromCondition{
		Hours:   hours,
		Courses: courses,
	}
}

func (c *CreditHoursFromCondition) Fulfils(userInfo UserInfo) (bool, error) {
	return false, nil
}

type UpperDivisionCoursesCondition struct {
	Hours  int
	Count  int
	Prefix string
}

func NewUpperDivisionCreditHoursCondition(hours int, prefix string) *UpperDivisionCoursesCondition {
	return &UpperDivisionCoursesCondition{
		Hours:  hours,
		Prefix: prefix,
	}
}

func NewUpperDivisionCountCondition(count int, prefix string) *UpperDivisionCoursesCondition {
	return &UpperDivisionCoursesCondition{
		Count:  count,
		Prefix: prefix,
	}
}

func (c *UpperDivisionCoursesCondition) Fulfils(userInfo UserInfo) (bool, error) {
	return false, nil
}

type ResearchCondition struct {
	Hours       int
	DegreeLevel DegreeLevel
}

func NewResearchCondition(hours int, degreeLevel DegreeLevel) *ResearchCondition {
	return &ResearchCondition{
		Hours:       hours,
		DegreeLevel: degreeLevel,
	}
}

func (c *ResearchCondition) Fulfils(userInfo UserInfo) (bool, error) {
	return false, nil
}

type NCoursesCondition struct {
	N       int
	Courses []Course
}

func NewNCoursesCondition(n int, courses []Course) *NCoursesCondition {
	return &NCoursesCondition{
		N:       n,
		Courses: courses,
	}
}

func (c *NCoursesCondition) Fulfils(userInfo UserInfo) (bool, error) {
	return false, nil
}
