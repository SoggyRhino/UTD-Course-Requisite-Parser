package visitors

import (
	"parser/conditions"
	"strconv"
)

// mapGradeLevel maps variants of grade levels to a GradeLevel
//
// GRADE_LEVEL
//
//	: [Ff]'reshm'[ae]'n'
//	| [Ss]'ophomore'
//	| [Jj]'unior' | 'JR'
//	| [Ss]'enior' | 'SR'
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

// mapDegreeLevel maps variants of degree levels to a DegreeLevel
//
// DEGREE_LEVEL : 'MS' | 'BS' | 'PHD' | 'PhD';
func mapDegreeLevel(text string) conditions.DegreeLevel {
	switch text {
	case "MS":
		return conditions.Graduate
	case "BS":
		return conditions.Undergraduate
	case "PHD", "PhD":
		return conditions.PhD
	default:
		panic("Invalid Degree Level: " + text)
	}
}

// mapDivisionType maps variants of division types to a DegreeLevel
//
// DIVISION_TYPE
//
//	 : [Uu]'ndergraduate' | 'Ugrd' | 'ugrd'
//		| [Gg]'raduate'      | 'GRAD' | 'grad'
//		| [Dd]'octoral'
func mapDivisionType(text string) conditions.DegreeLevel {
	switch text {
	case "Undergraduate", "undergraduate", "Ugrd", "ugrd":
		return conditions.Undergraduate
	case "Graduate", "graduate", "GRAD", "grad":
		return conditions.Graduate
	case "Doctoral", "doctoral":
		return conditions.PhD
	default:
		panic("Invalid Division Type: " + text)
	}
}

// mapGPA parse a GPA from a string
func mapGPA(text string) float64 {
	float, err := strconv.ParseFloat(text, 64)
	if err != nil || float < 0 || float > 4 {
		panic("Invalid GPA: " + text)
	}
	return float
}
