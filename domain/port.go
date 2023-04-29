package domain

import (
	"errors"
	"io"
)

var ErrPortNotFound = errors.New("Port not found")

type Port struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []any     `json:"alias"`
	Regions     []any     `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

type PortMap map[string]*Port

type PortUpsertService interface {
	UpsertPort(shortname string, port *Port) error
	UpsertPortFile(portfile io.Reader) error
}

type PortUpsertStorage interface {
	GetPort(shortname string) (*Port, error)
	InsertPort(shortname string, port *Port) error
	UpdatePort(shortname string, port *Port) error
}
