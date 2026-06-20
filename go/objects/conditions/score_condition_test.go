package conditions

import (
	"parser/objects/constants"
	"testing"
)

func TestScoreConditionsFulfils(t *testing.T) {
	info := constants.UserInfo{}

	t.Run("PlacementTestScore", func(t *testing.T) {
		cond := NewPlacementTestScoreCondition("Math", 80, 100)
		expected := constants.Evaluation{
			Name:    "Placement Test Score Math",
			Status:  constants.StatusUnknown,
			Summary: "Cannot verify Math placement test score (requires 80-100) automatically",
		}
		assertEval(t, expected, cond.Fulfils(info, false))
	})

	t.Run("APScore", func(t *testing.T) {
		cond := NewAPScoreCondition(4)
		expected := constants.Evaluation{
			Name:    "AP Score",
			Status:  constants.StatusUnknown,
			Summary: "Cannot verify AP score of 4 automatically",
		}
		assertEval(t, expected, cond.Fulfils(info, false))
	})

	t.Run("AleksScore", func(t *testing.T) {
		cond := NewAleksScoreCondition(85)
		expected := constants.Evaluation{
			Name:    "Aleks Score",
			Status:  constants.StatusUnknown,
			Summary: "Cannot verify ALEKS score of 85 automatically",
		}
		assertEval(t, expected, cond.Fulfils(info, false))
	})
}

func TestPlacementTestScoreCondition_JSON(t *testing.T) {
	cond := NewPlacementTestScoreCondition("Math", 80, 100)
	assertJSONRoundTrip[PlacementTestScoreCondition](t, cond, "placement_test_score")
}

func TestAPScoreCondition_JSON(t *testing.T) {
	cond := NewAPScoreCondition(4)
	assertJSONRoundTrip[APScoreCondition](t, cond, "ap_score")
}

func TestAleksScoreCondition_JSON(t *testing.T) {
	cond := NewAleksScoreCondition(85)
	assertJSONRoundTrip[AleksScoreCondition](t, cond, "aleks_score")
}
