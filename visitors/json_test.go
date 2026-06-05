package visitors

import (
	"encoding/json"
	"fmt"
	"parser/conditions"
	"parser/constants"
	"parser/parser"
	"parser/rules"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/google/go-cmp/cmp"
)

func TestJSONSerialization(t *testing.T) {
	testCases := map[string]struct {
		Input    string
		Expected Requirements
	}{
		"Simple Prerequisites": {
			Input: "Prerequisites: ACCT 2301 and ACCT 2302.",
			Expected: Requirements{
				PreReqs: conditions.NewAndCondition(
					conditions.NewCourseCondition("ACCT", "2301", ""),
					conditions.NewCourseCondition("ACCT", "2302", "")),
			},
		},
		"Implicit Prerequisite": {
			Input: "MATH 1325 or MATH 2413.",
			Expected: Requirements{
				PreReqs: conditions.NewOrCondition(
					conditions.NewCourseCondition("MATH", "1325", ""),
					conditions.NewCourseCondition("MATH", "2413", "")),
			},
		},
		"Instructor Consent": {
			Input: "Prerequisite: instructor consent.",
			Expected: Requirements{
				PreReqs: conditions.NewConsentCondition(constants.InstructorConsent),
			},
		},
		"Standing": {
			Input: "Prerequisite: Upper-division standing.",
			Expected: Requirements{
				PreReqs: conditions.NewGenericStandingCondition(constants.UpperDivisionStanding),
			},
		},
		"GPA": {
			Input: "Prerequisite: a university grade point average of at least 2.750.",
			Expected: Requirements{
				PreReqs: conditions.NewGpaCondition(2.75),
			},
		},
		"Major": {
			Input: "JSOM Ugrd Majors Only.",
			Expected: Requirements{
				PreReqs: conditions.NewMajorConditionWithDegreeLevel("JSOM", constants.Undergraduate),
			},
		},
		"Alternative": {
			Input: "Prerequisite: BIOL 2311 or equivalent.",
			Expected: Requirements{
				PreReqs: conditions.NewAlternativeCondition(conditions.NewCourseCondition("BIOL", "2311", "")),
			},
		},
		"Grade Level Standing": {
			Input: "Prerequisite: Junior or Senior standing.",
			Expected: Requirements{
				PreReqs: conditions.NewOrCondition(
					conditions.NewGradeLevelCondition(constants.Junior),
					conditions.NewGradeLevelCondition(constants.Senior)),
			},
		},
		"Graduate Standing": {
			Input: "Prerequisite: Graduate standing in Biology.",
			Expected: Requirements{
				PreReqs: conditions.NewGraduateStandingInConditionWithDegree("Biology"),
			},
		},
		"Core": {
			Input: "Prerequisite: Completion of 050 Creative Arts Core.",
			Expected: Requirements{
				PreReqs: conditions.NewCoreCondition("050", "Creative Arts"),
			},
		},
		"Upper Division Hours": {
			Input: "Prerequisite: 6 SCH of upper-division ARTS courses.",
			Expected: Requirements{
				PreReqs: conditions.NewUpperDivisionCreditHoursCondition(6, "ARTS"),
			},
		},
		"Repeat Rule": {
			Input: "Repeat Restriction.",
			Expected: Requirements{
				Rules: []rules.Rule{rules.NewRepeatRule(1, 0, nil, "")},
			},
		},
		"Repeat Limit Hours": {
			Input: "Repeat Limit - ACCT 7323 may only be repeated for a maximum of 9 semester credit hours.",
			Expected: Requirements{
				Rules: []rules.Rule{rules.NewRepeatRule(0, 9, []constants.Course{{Prefix: "ACCT", Number: "7323"}}, "")},
			},
		},
		"Repeat Limit Times": {
			Input: "Repeat Limit - UNIV 4074 may be repeated up to 3 times.",
			Expected: Requirements{
				Rules: []rules.Rule{rules.NewRepeatRule(3, 0, []constants.Course{{Prefix: "UNIV", Number: "4074"}}, "")},
			},
		},
		"GPA Repeat Rule": {
			Input: "GPA Repeat Restriction - MIS 6309.",
			Expected: Requirements{
				Rules: []rules.Rule{rules.NewGpaRepeatRule(constants.Course{Prefix: "MIS", Number: "6309"})},
			},
		},
		"Degree Satisfaction": {
			Input: "May not be used to satisfy BS INTS degree requirements.",
			Expected: Requirements{
				Rules: []rules.Rule{rules.NewDegreeSatisfactionRuleFromPrefix([]string{"INTS"}, constants.Undergraduate)},
			},
		},
		"Credit For Rule": {
			Input: "Credit cannot be received for both courses, BCOM 1300 and BCOM 3300.",
			Expected: Requirements{
				Rules: []rules.Rule{rules.NewCreditForRule(
					rules.NewAndCourseCollection(
						rules.NewSimpleCourseCollection([]constants.Course{{Prefix: "BCOM", Number: "1300"}}),
						rules.NewSimpleCourseCollection([]constants.Course{{Prefix: "BCOM", Number: "3300"}}),
					)),
				},
			},
		},
		"Living Learning Rule": {
			Input: "ARHM & ATEC Living Learning Community.",
			Expected: Requirements{
				Rules: []rules.Rule{rules.NewLivingLearningRuleFromPrefixes([]string{"ARHM", "ATEC"})},
			},
		},
		"School Rule": {
			Input: "Open to students in the School of Engineering and Computer Science only.",
			Expected: Requirements{
				Rules: []rules.Rule{rules.NewSchoolRule([]string{"Engineering and Computer Science"})},
			},
		},
		"Same As Rule": {
			Input: "(Same as MATH 3335 and STAT 3335)",
			Expected: Requirements{
				Rules: []rules.Rule{rules.NewSameAsRule([]constants.Course{
					{Prefix: "MATH", Number: "3335"},
					{Prefix: "STAT", Number: "3335"},
				})},
			},
		},
		"Academic Plan Req": {
			Input: "Academic Plan Not Equal to BSANMSNF.",
			Expected: Requirements{
				PreReqs: conditions.NewAcademicYearCondition("BSANMSNF", false),
			},
		},
		"Exact Section": {
			Input: "Corequisite: BIOL 2311.001.",
			Expected: Requirements{
				CoReqs: conditions.NewExactSectionCondition(constants.Course{Prefix: "BIOL", Number: "2311", Section: "001"}),
			},
		},
		"Exclude Notice Req": {
			Input: "BLAW 2301 Repeat Restriction and non-DMHP/non-LLC (DMLC, DFLC, DHLC) Student Group Only.",
			Expected: Requirements{
				Rules:   []rules.Rule{rules.NewCourseRepeatRule([]constants.Course{{Prefix: "BLAW", Number: "2301"}})},
				Notices: []constants.Notice{constants.ExcludeDMHPLLCNotice},
			},
		},
		"Append Academic Plan Req": {
			Input: "GPA Repeat Restriction - MIS 6309 and Academic Plan Not Equal to BSANMSNF.",
			Expected: Requirements{
				PreReqs: conditions.NewAcademicYearCondition("BSANMSNF", false),
				Rules:   []rules.Rule{rules.NewGpaRepeatRule(constants.Course{Prefix: "MIS", Number: "6309"})},
			},
		},
		"Exact Coreq Notice Req": {
			Input: "Check class notes to make sure you are selecting the matching corequisite section",
			Expected: Requirements{
				Notices: []constants.Notice{constants.ExactCoReqNotice},
			},
		},
		"Computer Scholars Req": {
			Input: "Computing Scholars Program.",
			Expected: Requirements{
				Notices: []constants.Notice{constants.ComputerScholarsProgramNotice},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			output, err := toFromJson(tc.Input)
			if err != nil {
				t.Fatalf("Failed to parse: %v", err)
			}

			if diff := cmp.Diff(tc.Expected, *output, cmp.AllowUnexported(conditions.OrCondition{}, conditions.AndCondition{})); diff != "" {
				t.Errorf("Unexpected output (-want +got):\n%s", diff)
			}
		})
	}

}

func toFromJson(input string) (*Requirements, error) {
	stream := antlr.NewInputStream(input)
	lexer := parser.NewRequirementsLexer(stream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewRequirementsParser(tokens)

	tree := p.Prog()
	visitor := NewRequisiteVisitor(tokens)
	visitor.Visit(tree)

	output, err := json.Marshal(visitor.Requirements)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %v", err)
	}

	var reqs *Requirements
	if err := json.Unmarshal(output, &reqs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal back: %v", err)
	}
	return reqs, nil
}
