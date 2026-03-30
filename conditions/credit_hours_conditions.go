package conditions

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
