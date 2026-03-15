package visitors

import (
	"parser/parser"
	"testing"
)

/*

title : (CAPITALIZED | CORE)+ ( 'and' (CAPITALIZED | CORE)+ )* ;

*/

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
