package main

import (
	"./pkg/db"
	"fmt"
)

func main() {
	fmt.Println("Launching koinkoin.io !")
	defer db.CloseDb()
	//todo : make router to use theses services
}
