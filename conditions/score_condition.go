package conditions

import (
	"encoding/json"
	"parser/constants"
)

type PlacementTestScoreCondition struct {
	Name     string `json:"name,omitempty"`
	ScoreMin int    `json:"score_min,omitempty"`
	ScoreMax int    `json:"score_max,omitempty"`
}

func (c *PlacementTestScoreCondition) MarshalJSON() ([]byte, error) {
	type Alias PlacementTestScoreCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "placement_test_score",
		Alias: (*Alias)(c),
	})
}

func NewPlacementTestScoreCondition(name string, scoreMin, scoreMax int) *PlacementTestScoreCondition {
	return &PlacementTestScoreCondition{
		Name:     name,
		ScoreMin: scoreMin,
		ScoreMax: scoreMax,
	}
}

func (c *PlacementTestScoreCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}

type APScoreCondition struct {
	Score int `json:"score,omitempty"`
}

func (c *APScoreCondition) MarshalJSON() ([]byte, error) {
	type Alias APScoreCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "ap_score",
		Alias: (*Alias)(c),
	})
}

func NewAPScoreCondition(score int) *APScoreCondition {
	return &APScoreCondition{
		Score: score,
	}
}

func (c *APScoreCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}

type AleksScoreCondition struct {
	Score int `json:"score,omitempty"`
}

func (c *AleksScoreCondition) MarshalJSON() ([]byte, error) {
	type Alias AleksScoreCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "aleks_score",
		Alias: (*Alias)(c),
	})
}

func NewAleksScoreCondition(score int) *AleksScoreCondition {
	return &AleksScoreCondition{
		Score: score,
	}
}

func (c *AleksScoreCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}
