package http

import (
	"net/http"
	"strings"

	"github.com/mvrilo/go-port-svc/domain"
)

type PortUpsertHandler struct {
	svc domain.PortUpsertService
}

func NewPortUpsertHandler(svc domain.PortUpsertService) *PortUpsertHandler {
	return &PortUpsertHandler{svc}
}

func (h *PortUpsertHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.ToLower(r.Method) != "post" {
		w.WriteHeader(405)
		return
	}

	defer r.Body.Close()
	err := h.svc.UpsertPortFile(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}
