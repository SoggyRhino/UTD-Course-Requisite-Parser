package visitors

import (
	"parser/parser"
	"testing"
)

func TestVisitTitle(t *testing.T) {

	testCases := map[string]struct {
		Input  string
		Result string
	}{
		"Test 1": {
			Input:  "Information Technology and Management",
			Result: "Information Technology and Management",
		},
		"Test 2": {
			Input:  "Korean Language",
			Result: "Korean Language",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[string](t, tc.Input, rule((*parser.RequirementsParser).Title), tc.Result)
		})

	}
}

func TestVisitDegree(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result string
	}{
		"Lowercase 2 words": {
			Input:  "computer science",
			Result: "computer science",
		},
		"Capitalized with and": {
			Input:  "Visual and Performing Arts",
			Result: "Visual and Performing Arts",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[string](t, tc.Input, rule((*parser.RequirementsParser).Degree), tc.Result)
		})
	}
}

func TestVisitDegreeList(t *testing.T) {

	testCases := map[string]struct {
		Input  string
		Result []string
	}{
		"Or separated": {
			Input:  "psychology or computer science or neuroscience",
			Result: []string{"psychology", "computer science", "neuroscience"},
		},
		"And with comma": {
			Input: "Engineering and Computer Science, Actuarial Science, Data Science, and Cognitive Science",
			Result: []string{
				"Engineering and Computer Science",
				"Actuarial Science",
				"Data Science",
				"Cognitive Science",
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[[]string](t, tc.Input, rule((*parser.RequirementsParser).Degree_list), tc.Result)
		})
	}

}
