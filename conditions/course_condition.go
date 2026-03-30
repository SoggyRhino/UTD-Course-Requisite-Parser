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
