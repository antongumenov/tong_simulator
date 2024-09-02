package main

import (
	"tong_simulator/internal/controller"
	"tong_simulator/internal/repository/memory"
)

func main() {
	repo := memory.New()
	ctrl := controller.New(repo)

	go runServer(ctrl)

	runGui(ctrl)

}
