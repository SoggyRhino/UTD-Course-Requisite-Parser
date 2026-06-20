package conditions

import (
	"encoding/json"
	"parser/objects/constants"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUnmarshalCondition(t *testing.T) {
	testCases := []struct {
		name    string
		json    string
		wantErr bool
	}{
		{"Invalid JSON", "{invalid}", true},
		{"Unknown Type", `{"type": "unknown"}`, true},
		{"Alternative", `{"type": "alternative", "condition": {"type": "course", "course": {"prefix": "CS", "number": "1337"}}}`, false},
		{"Consent", `{"type": "consent", "consent": "Instructor"}`, false},
		{"Course", `{"type": "course", "course": {"prefix": "CS", "number": "1337"}}`, false},
		{"Core", `{"type": "core", "core_number": "010"}`, false},
		{"Credit Hours", `{"type": "credit_hours", "hours": 45}`, false},
		{"Credit Hours From", `{"type": "credit_hours_from", "hours": 6, "courses": [{"prefix": "CS", "number": "1337"}]}`, false},
		{"Upper Division", `{"type": "upper_division_courses", "hours": 6}`, false},
		{"Research", `{"type": "research", "hours": 3}`, false},
		{"N Courses", `{"type": "n_courses", "n": 2, "courses": [{"prefix": "CS", "number": "1337"}]}`, false},
		{"Major", `{"type": "major", "degree": "CS"}`, false},
		{"Degree", `{"type": "degree", "degree": "BS"}`, false},
		{"GPA", `{"type": "gpa", "gpa": 3.0}`, false},
		{"Or", `{"type": "or", "conditions": [{"type": "course", "course": {"prefix": "CS", "number": "1337"}}]}`, false},
		{"And", `{"type": "and", "conditions": [{"type": "course", "course": {"prefix": "CS", "number": "1337"}}]}`, false},
		{"Concurrent", `{"type": "concurrent_enrollment", "course": {"prefix": "CS", "number": "1337"}}`, false},
		{"Exact Section", `{"type": "exact_section", "course": {"prefix": "CS", "number": "1337", "section": "001"}}`, false},
		{"Any Previous Major", `{"type": "any_previous_major_course", "prefix": "CS"}`, false},
		{"Academic Year", `{"type": "academic_year", "plan": "2023"}`, false},
		{"Placement", `{"type": "placement_test_score", "name": "Math"}`, false},
		{"AP", `{"type": "ap_score", "score": 4}`, false},
		{"Aleks", `{"type": "aleks_score", "score": 85}`, false},
		{"Grade Level", `{"type": "grade_level", "grade_level": "Senior"}`, false},
		{"Graduate Standing", `{"type": "graduate_standing_in", "degree": "CS"}`, false},
		{"Generic Standing", `{"type": "generic_standing", "standing": "Good"}`, false},
		{"Student Group", `{"type": "student_group", "groups": "Honors"}`, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := UnmarshalCondition([]byte(tc.json))
			if (err != nil) != tc.wantErr {
				t.Errorf("UnmarshalCondition() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func assertEval(t *testing.T, expected constants.Evaluation, got *constants.Evaluation) {
	t.Helper()
	if got == nil {
		t.Fatal("Expected evaluation, got nil")
	}
	if diff := cmp.Diff(expected, *got); diff != "" {
		t.Errorf("Unexpected evaluation (-want +got):\n%s", diff)
	}
}

func assertJSONRoundTrip[T any, PT interface {
	*T
}](t *testing.T, original PT, expectedType string) {
	t.Helper()

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var envelope struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &envelope); err != nil {
		t.Fatalf("Failed to unmarshal type field: %v", err)
	}
	if envelope.Type != expectedType {
		t.Errorf("Expected type %q, got %q", expectedType, envelope.Type)
	}

	decoded := PT(new(T))
	if err := json.Unmarshal(data, decoded); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if diff := cmp.Diff(original, decoded); diff != "" {
		t.Errorf("Round-trip mismatch (-want +got):\n%s", diff)
	}
}
