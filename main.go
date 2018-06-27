package main

import "./pkg/db"

func main() {
	defer db.CloseDb()
	//todo : make router to use theses services
}
