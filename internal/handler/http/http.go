package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tong_simulator/internal/controller"
)

type Data struct {
	Encoder int64 `json:"encoder"`
	Loadcel int64 `json:"loadcell"`
}

type Handler struct {
	ctrl *controller.Controller
}

func New(s *controller.Controller) *Handler {
	return &Handler{
		ctrl: s,
	}
}

func (h *Handler) GetSensors(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	data := &Data{
		h.ctrl.GetEncoder(),
		h.ctrl.GetLoadCell(),
	}
	json.NewEncoder(writer).Encode(data)
}

func (h *Handler) Dump(writer http.ResponseWriter, request *http.Request) {
	h.ctrl.Dump()
	writer.WriteHeader(http.StatusOK)
}

func (h *Handler) SetDumpLevel(writer http.ResponseWriter, request *http.Request) {
	dumpLevel, err := strconv.ParseInt(request.URL.Query().Get(`dump_level`), 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	h.ctrl.SetDumpLevel(dumpLevel)
	writer.WriteHeader(http.StatusOK)
}
