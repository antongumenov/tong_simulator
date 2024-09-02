package main

import (
	"net/http"
	"tong_simulator/internal/controller"
	handler "tong_simulator/internal/handler/http"
)

func runServer(ctrl *controller.Controller) {
	mux := http.NewServeMux()
	h := handler.New(ctrl)
	mux.HandleFunc("/get", h.GetSensors)
	mux.HandleFunc("/dump", h.Dump)
	mux.HandleFunc("/level", h.SetDumpLevel)
	http.ListenAndServe(":5000", mux)
}
