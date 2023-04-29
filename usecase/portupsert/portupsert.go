package portupsert

import (
	"encoding/json"
	"io"

	"github.com/mvrilo/go-port-svc/domain"
)

type PortUpsertService struct {
	db domain.PortUpsertStorage
}

func New(db domain.PortUpsertStorage) *PortUpsertService {
	return &PortUpsertService{db}
}

func (p *PortUpsertService) UpsertPort(shortname string, port *domain.Port) error {
	_, err := p.db.GetPort(shortname)
	if err == domain.ErrPortNotFound {
		return p.db.InsertPort(shortname, port)
	}
	if err != nil {
		return err
	}
	return p.db.UpdatePort(shortname, port)
}

func (p *PortUpsertService) UpsertPortFile(portfile io.Reader) error {
	dec := json.NewDecoder(portfile)

	var portmap domain.PortMap
	if err := dec.Decode(&portmap); err != nil {
		return err
	}

	for shortname, port := range portmap {
		if err := p.UpsertPort(shortname, port); err != nil {
			return err
		}
	}

	return nil
}

var _ domain.PortUpsertService = (*PortUpsertService)(nil)
