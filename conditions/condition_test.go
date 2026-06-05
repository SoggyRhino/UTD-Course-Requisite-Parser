package conditions

import (
	"testing"
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
