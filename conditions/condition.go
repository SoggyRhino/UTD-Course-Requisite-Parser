package conditions

type Grade int //todo expand

type Course struct {
	Prefix string
	Number string
	Section string
}

func NewCourse(prefix string, number string) Course {
	return Course{
		Prefix: prefix,
		Number: number,
	}
}
func NewCourseWithSection(prefix string, number string, section string) Course {
	return Course{
		Prefix: prefix,
		Number: number,
		Section: section,
	}
}

type UserInfo struct {
	Taken map[Course]Grade
}

type Condition interface {
	Fulfils(UserInfo) (bool, error)
}
