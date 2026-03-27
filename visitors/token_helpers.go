package visitors

import (
	"parser/conditions"
)

// mapGradeLevel
//
// GRADE_LEVEL
//
//	 : [Ff]'reshm'[ae]'n'
//		| [Ss]'ophomore'
//		| [Jj]'unior' | 'JR'
//		| [Ss]'enior' | 'SR'
func mapGradeLevel(text string) conditions.GradeLevel {
	switch text {
	case "Freshman", "freshman", "Freshmen", "freshmen":
		return conditions.Freshman
	case "Sophomore", "sophomore":
		return conditions.Sophomore
	case "Junior", "junior", "Jr":
		return conditions.Junior
	case "Senior", "senior", "Sr":
		return conditions.Senior
	default:
		panic("Invalid Grade Level: " + text)
	}
}
