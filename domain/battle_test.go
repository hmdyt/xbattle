package domain_test

import (
	"github.com/hmdyt/xbattle/domain"
	"github.com/hmdyt/xbattle/misc"
	"github.com/hmdyt/xbattle/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBattle_Integration(t *testing.T) {
	uuidGen := misc.NewUuidGeneratorV7()
	battleRepo := repository.NewOnMemoryBattleRepository()
	battleDomain := domain.NewBattle(battleRepo, uuidGen)

	playerID := "Player01"

	battle := battleDomain.Start(playerID)
	assert.Equal(t, battle.Turn, 0)

	battle, err := battleDomain.Attack(playerID)
	assert.Nil(t, err)
	assert.Equal(t, battle.Turn, 1)

	battle, err = battleDomain.Attack(playerID)
	assert.Nil(t, err)
	assert.Equal(t, battle.Turn, 2)
}
