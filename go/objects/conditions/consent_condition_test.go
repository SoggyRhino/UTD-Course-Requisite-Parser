package conditions

import (
	"parser/objects/constants"
	"testing"
)

func TestConsentCondition_Fulfils(t *testing.T) {
	cond := NewConsentCondition("Instructor")
	info := constants.UserInfo{}
	eval := cond.Fulfils(info, false)

	if eval == nil {
		t.Fatal("Expected evaluation, got nil")
	}

	if eval.Status != constants.StatusPossibleFail {
		t.Errorf("Expected status %v, got %v", constants.StatusPossibleFail, eval.Status)
	}

	expectedSummary := "Requires manual consent/approval"
	if eval.Summary != expectedSummary {
		t.Errorf("Expected summary %q, got %q", expectedSummary, eval.Summary)
	}
}

func TestConsentCondition_JSON(t *testing.T) {
	cond := NewConsentCondition("Instructor")
	assertJSONRoundTrip[ConsentCondition](t, cond, "consent")
}
