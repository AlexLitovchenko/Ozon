package main

import (
	"test/DataBase"
	"test/Handler"
)

func main() {
	DataBase.ConnectDB()
	var handler Handler.Handle
	handler.Handles()
}
