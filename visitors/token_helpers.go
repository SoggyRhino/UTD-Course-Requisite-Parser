package visitors

import (
	"parser/conditions"
	"strconv"
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
		//todo look into better option than panic
		panic("Invalid Grade Level: " + text)
	}
}

func mapGPA(text string) float64 {
	float, err := strconv.ParseFloat(text, 64)
	if err != nil {
		panic("Invalid GPA: " + text)
	}
	return float
}
