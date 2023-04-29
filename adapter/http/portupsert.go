package http

import (
	"fmt"
	"io"
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
		fmt.Println("method not allowed")
		w.WriteHeader(405)
		return
	}

	var file io.ReadCloser = nil

	r.ParseMultipartForm(100000)

	if r.MultipartForm != nil {
		println("----", len(r.MultipartForm.File))
		for _, f := range r.MultipartForm.File {
			if len(f) < 1 {
				return
			}

			f, err := f[0].Open()
			if err != nil {
				w.WriteHeader(500)
				break
			}

			file = f
			break
		}
	}

	if file == nil {
		file = r.Body
	}

	defer file.Close()

	err := h.svc.UpsertPortFile(file)
	if err != nil {
		fmt.Println("error upserting file ", err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}
