package main

import (
	"github.com/koinkoin-io/koinkoin.backend/pkg/db"
	"fmt"
	"github.com/kataras/iris"

)

func main() {
	fmt.Println("Launching koinkoin.io !")
	defer db.CloseDb()

	app := iris.Default()
	RouteAll(app)
	app.Run(iris.Addr(":8080"))
}
