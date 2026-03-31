package conditions

import "parser/utils"

type PlacementTestScoreCondition struct {
	Name     string
	ScoreMin int
	ScoreMax int
}

func NewPlacementTestScoreCondition(name string, scoreMin, scoreMax int) *PlacementTestScoreCondition {
	return &PlacementTestScoreCondition{
		Name:     name,
		ScoreMin: scoreMin,
		ScoreMax: scoreMax,
	}
}

func (c *PlacementTestScoreCondition) Fulfils(userInfo utils.UserInfo) (bool, error) {
	return false, nil
}

type APScoreCondition struct {
	Score int
}

func NewAPScoreCondition(score int) *APScoreCondition {
	return &APScoreCondition{
		Score: score,
	}
}

func (c *APScoreCondition) Fulfils(userInfo utils.UserInfo) (bool, error) {
	return false, nil
}

type AleksScoreCondition struct {
	Score int
}

func NewAleksScoreCondition(score int) *AleksScoreCondition {
	return &AleksScoreCondition{
		Score: score,
	}
}

func (c *AleksScoreCondition) Fulfils(userInfo utils.UserInfo) (bool, error) {
	return false, nil
}
