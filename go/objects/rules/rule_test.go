package rules

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUnmarshalRule(t *testing.T) {
	testCases := []struct {
		name    string
		json    string
		wantErr bool
	}{
		{"Invalid JSON", "{invalid}", true},
		{"Unknown Type", `{"type": "unknown"}`, true},
		{"Repeat", `{"type": "repeat", "count": 2}`, false},
		{"GPA Repeat", `{"type": "gpa_repeat", "course": {"prefix": "CS", "number": "1337"}}`, false},
		{"Credit For", `{"type": "credit_for", "courses": {"type": "simple", "course": [{"prefix": "CS", "number": "1337"}]}}`, false},
		{"Degree Satisfaction", `{"type": "degree_satisfaction", "prefixes": ["CS"]}`, false},
		{"Living Learning", `{"type": "living_learning", "prefixes": ["CS"]}`, false},
		{"School", `{"type": "school", "schools": ["ECS"]}`, false},
		{"Same As", `{"type": "same_as", "courses": [{"prefix": "CS", "number": "1337"}]}`, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := UnmarshalRule([]byte(tc.json))
			if (err != nil) != tc.wantErr {
				t.Errorf("UnmarshalRule() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
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
