package conditions

type StudentGroup string

const (
	CollegeVHonors    StudentGroup = "Collegium V Honors"
	LiberalArtsHonors              = "Liberal Arts Honors"
	SCVG                           = "SCVG"
	DMHP                           = "DMHP"
	DLAH                           = "DLAH"
)

type StudentGroupCondition struct {
	Groups StudentGroup
}

func NewStudentGroupCondition(groups StudentGroup) *StudentGroupCondition {
	return &StudentGroupCondition{
		Groups: groups,
	}
}

func (c *StudentGroupCondition) Fulfils(userInfo UserInfo) (bool, error) {
	return false, nil
}
