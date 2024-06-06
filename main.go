package integration_test_

import "github.com/hmdyt/xbattle/handler"

func main() {
	e := handler.Build()
	e.Logger.Fatal(e.Start(":1323"))
}
