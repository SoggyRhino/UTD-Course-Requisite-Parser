package conditions

import (
	"encoding/json"
	"parser/constants"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
