package main

import (
	"github.com/hmdyt/xbattle/domain"
	"github.com/hmdyt/xbattle/handler"
	"github.com/hmdyt/xbattle/misc"
	"github.com/hmdyt/xbattle/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	Route(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func Route(e *echo.Echo) {
	uuidGen := misc.NewUuidGeneratorV7()
	battleRepo := repository.NewOnMemoryBattleRepository()
	battleDomain := domain.NewBattle(battleRepo, uuidGen)

	battleHandler := handler.NewBattleHandler(battleDomain)

	battle := e.Group("/battle")
	battle.POST("/start", battleHandler.Start)
	battle.POST("/attack", battleHandler.Attack)
}
