package domain

import (
	"github.com/hmdyt/xbattle/entity"
	"github.com/hmdyt/xbattle/misc"
	"github.com/hmdyt/xbattle/repository"
)

type Battle struct {
	battleRepo repository.BattleRepository
	uuidGen    misc.UuidGenerator
}

func NewBattle(battleRepo repository.BattleRepository, uuidGen misc.UuidGenerator) *Battle {
	return &Battle{
		battleRepo: battleRepo,
		uuidGen:    uuidGen,
	}
}

func (b *Battle) Start(playerID string) entity.Battle {
	battle := b.battleRepo.Upsert(entity.Battle{
		PlayerID:      playerID,
		ID:            b.uuidGen.Generate(),
		EnemyPlayerID: "ENEMY",
		Turn:          0,
	})
	return battle
}

func (b *Battle) Attack(playerID string) (entity.Battle, error) {
	battle, err := b.battleRepo.GetByPlayerID(playerID)
	if err != nil {
		return entity.Battle{}, err
	}

	battle.Turn += 1
	return b.battleRepo.Upsert(battle), nil
}
