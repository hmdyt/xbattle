package repository

import (
	"fmt"
	"github.com/hmdyt/xbattle/entity"
)

var battleRepoMemory = map[string]interface{}{}

type BattleRepository interface {
	GetByPlayerID(playerID string) (entity.Battle, error)
	Upsert(entity.Battle) entity.Battle
}

type OnMemoryBattleRepository struct {
}

func NewOnMemoryBattleRepository() BattleRepository {
	return &OnMemoryBattleRepository{}
}

func (o *OnMemoryBattleRepository) Upsert(battle entity.Battle) entity.Battle {
	battleRepoMemory[battle.ID] = battle
	return battle
}

func (o *OnMemoryBattleRepository) GetByPlayerID(playerID string) (entity.Battle, error) {
	for _, v := range battleRepoMemory {
		battle := v.(entity.Battle)
		if battle.PlayerID == playerID {
			return battle, nil
		}
	}

	return entity.Battle{}, fmt.Errorf("battle not found playerID: %s", playerID)
}
