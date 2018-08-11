package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/koin-bet/koin.backend/pkg/util"
	"time"
	"github.com/koin-bet/koin.backend/pkg/worker"
)

var startTime  = time.Now()

func Uptime() string {
	return time.Since(startTime).Round(time.Second).String()
}

func main() {
	fmt.Println("Launching koin.bet ! Bet, win and repeat.")

	app := iris.Default()
	RouteAll(app)
	dispatcher := worker.NewDispatcher(worker.DefaultJobQueue, *worker.DefaultMaxWorkers)
	dispatcher.Run()

	app.Run(iris.Addr(":" + util.GetEnvOrDefault("port", "80")))
}
