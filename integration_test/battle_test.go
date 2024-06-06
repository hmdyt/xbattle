package integration_test

import (
	"github.com/go-resty/resty/v2"
	"github.com/hmdyt/xbattle/domain/schema"
	"github.com/hmdyt/xbattle/handler"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Testバトルスタートから終了まで(t *testing.T) {
	e := handler.Build()
	testServer := httptest.NewServer(e.Server.Handler)
	t.Cleanup(testServer.Close)
	client := resty.New().SetBaseURL(testServer.URL)

	startRes := start(t, client, "Player01")
	assert.Equal(t, 0, startRes.Turn)

	attackRes := attack(t, client, "Player01")
	assert.Equal(t, 1, attackRes.Turn)

	attackRes = attack(t, client, "Player01")
	assert.Equal(t, 2, attackRes.Turn)
}

func start(t *testing.T, client *resty.Client, playerID string) schema.BattleStartRes {
	var res schema.BattleStartRes
	resp, err := client.R().
		SetBody(map[string]interface{}{"player_id": playerID}).
		SetResult(&res).
		Post("/battle/start")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode())
	return res
}

func attack(t *testing.T, client *resty.Client, playerID string) schema.BattleAttackRes {
	var res schema.BattleAttackRes
	resp, err := client.R().
		SetBody(map[string]interface{}{"player_id": playerID}).
		SetResult(&res).
		Post("/battle/attack")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode())
	return res
}
