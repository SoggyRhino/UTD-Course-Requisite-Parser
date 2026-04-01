package visitors

import (
	"parser/utils"
	"strconv"
	"strings"
)

// mapGradeLevel maps variants of grade levels to a GradeLevel
//
// GRADE_LEVEL
//
//	: [Ff]'reshm'[ae]'n'
//	| [Ss]'ophomore'
//	| [Jj]'unior' | 'JR'
//	| [Ss]'enior' | 'SR'
func mapGradeLevel(text string) utils.GradeLevel {
	switch text {
	case "Freshman", "freshman", "Freshmen", "freshmen":
		return utils.Freshman
	case "Sophomore", "sophomore":
		return utils.Sophomore
	case "Junior", "junior", "Jr":
		return utils.Junior
	case "Senior", "senior", "Sr":
		return utils.Senior
	case "":
		return utils.AnyGrade

	default:
		//todo look into better option than panic
		panic("Invalid Grade Level: " + text)
	}
}

// mapDegreeLevel maps variants of degree levels to a DegreeLevel
//
// DEGREE_LEVEL : 'MS' | 'BS' | 'PHD' | 'PhD';
func mapDegreeLevel(text string) utils.DegreeLevel {
	switch text {
	case "MS":
		return utils.Graduate
	case "BS":
		return utils.Undergraduate
	case "PHD", "PhD":
		return utils.PhD
	case "":
		return utils.AnyDegree
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
func mapDivisionType(text string) utils.DegreeLevel {
	switch text {
	case "Undergraduate", "undergraduate", "Ugrd", "ugrd":
		return utils.Undergraduate
	case "Graduate", "graduate", "GRAD", "grad":
		return utils.Graduate
	case "Doctoral", "doctoral":
		return utils.PhD
	case "":
		return utils.AnyDegree
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

func mapInt(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil || num < 0 {
		//todo look into if we need to separate small in and int
		panic("Invalid Int: " + text)
	}
	return num
}

func mapNumberString(text string) int {
	switch strings.ToLower(text) {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	default:
		panic("Invalid Number String: " + text)
	}
}

func stripChars(text string, chars ...string) string {
	for _, char := range chars {
		text = strings.Replace(text, char, "", -1)
	}
	return text
}

func mapToStudentGroup(text string) utils.StudentGroup {
	switch text {
	case "Computer Scholars Program", "Computing Scholars Program":
		return utils.ComputerScholarsProgram
	case "Collegium V Honors", "CV Honors":
		return utils.CollegeVHonors
	case "Liberal Arts Honors":
		return utils.LiberalArtsHonors
	case "SCVG":
		return utils.SCVG
	case "DMHP":
		return utils.DMHP
	case "DLAH":
		return utils.DLAH
	default:
		panic("Invalid Student Group: " + text)
	}
}
