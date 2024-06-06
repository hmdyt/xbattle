package handler

import (
	"github.com/hmdyt/xbattle/domain"
	"github.com/hmdyt/xbattle/domain/schema"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BattleHandler struct {
	battle *domain.Battle
}

func NewBattleHandler(battle *domain.Battle) *BattleHandler {
	return &BattleHandler{
		battle: battle,
	}
}

func (h *BattleHandler) Start(c echo.Context) error {
	var req schema.BattleStartReq
	if err := c.Bind(&req); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := req.Validate(); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	battle := h.battle.Start(req.PlayerID)
	return c.JSON(http.StatusOK, schema.NewBattleStartRes(battle))
}

func (h *BattleHandler) Attack(c echo.Context) error {
	var req schema.BattleAttackReq
	if err := c.Bind(&req); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := req.Validate(); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	battle, err := h.battle.Attack(req.PlayerID)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, schema.NewBattleAttackRes(battle))
}
