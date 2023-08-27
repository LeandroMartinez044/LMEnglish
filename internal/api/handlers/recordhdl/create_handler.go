package recordhdl

import (
	"encoding/json"
	"io/ioutil"
	"lmenglish/internal/core/ports"
	"net/http"
)

type Handler struct {
	recordsrv ports.RecordService
}

func New(recordSrv ports.RecordService) *Handler {
	return &Handler{recordsrv: recordSrv}
}

type Record struct {
	Sentence string `json:"sentence"`
	Word     string `json:"word"`
}

func (h *Handler) Record(w http.ResponseWriter, r *http.Request) {
	// 1- It reads body to convert to struct
	reqBody, _ := ioutil.ReadAll(r.Body)
	var record Record
	json.Unmarshal(reqBody, &record)

	h.recordsrv.Create(record.Sentence)

	// 4- response sentence
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode("success")
}
