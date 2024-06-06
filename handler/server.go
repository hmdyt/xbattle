package handler

import (
	"github.com/hmdyt/xbattle/domain"
	"github.com/hmdyt/xbattle/misc"
	"github.com/hmdyt/xbattle/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Build() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	Route(e)
	return e
}

func Route(e *echo.Echo) {
	uuidGen := misc.NewUuidGeneratorV7()
	battleRepo := repository.NewOnMemoryBattleRepository()
	battleDomain := domain.NewBattle(battleRepo, uuidGen)

	battleHandler := NewBattleHandler(battleDomain)

	battle := e.Group("/battle")
	battle.POST("/start", battleHandler.Start)
	battle.POST("/attack", battleHandler.Attack)
}
