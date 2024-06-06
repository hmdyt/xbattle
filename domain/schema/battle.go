package schema

import (
	"fmt"
	"github.com/hmdyt/xbattle/entity"
)

type BattleStartReq struct {
	PlayerID string `json:"player_id"`
}

func (r BattleStartReq) Validate() error {
	if r.PlayerID == "" {
		return fmt.Errorf("player_id is required")
	}
	return nil
}

type BattleStartRes struct {
	Turn int `json:"turn"`
}

func NewBattleStartRes(ent entity.Battle) BattleStartRes {
	return BattleStartRes{
		Turn: ent.Turn,
	}
}

type BattleAttackReq struct {
	PlayerID string `json:"player_id"`
}

func (r BattleAttackReq) Validate() error {
	if r.PlayerID == "" {
		return fmt.Errorf("player_id is required")
	}
	return nil
}

type BattleAttackRes struct {
	Turn int `json:"turn"`
}

func NewBattleAttackRes(ent entity.Battle) BattleAttackRes {
	return BattleAttackRes{
		Turn: ent.Turn,
	}
}
