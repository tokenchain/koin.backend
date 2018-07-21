package main

import (
	"github.com/koinkoin-io/koinkoin.backend/pkg/db"
	"fmt"
	"github.com/kataras/iris"
	"github.com/koinkoin-io/koinkoin.backend/pkg/util"
	"time"
)

var startTime  = time.Now()

func Uptime() string {
	return time.Since(startTime).Round(time.Second).String()
}

func main() {
	fmt.Println("Launching koinkoin.io ! Bet, win and repeat.")
	defer db.CloseDb()

	app := iris.Default()
	RouteAll(app)
	app.Run(iris.Addr(":" + util.GetEnvOrDefault("port", "80")))
}
