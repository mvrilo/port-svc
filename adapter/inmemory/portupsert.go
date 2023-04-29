package inmemory

import (
	"fmt"

	"github.com/mvrilo/go-port-svc/domain"
)

type PortUpsertInMemoryStorage struct {
	db domain.PortMap
}

func NewPortUpsertInMemoryStorage(db domain.PortMap) *PortUpsertInMemoryStorage {
	return &PortUpsertInMemoryStorage{db: db}
}

func (p *PortUpsertInMemoryStorage) debug() {
	fmt.Printf("current number of records: %d\n", len(p.db))
}

func (p *PortUpsertInMemoryStorage) GetPort(shortname string) (*domain.Port, error) {
	port, ok := p.db[shortname]
	if !ok {
		return nil, domain.ErrPortNotFound
	}
	return port, nil
}

func (p *PortUpsertInMemoryStorage) InsertPort(shortname string, port *domain.Port) error {
	p.db[shortname] = port
	p.debug()
	return nil
}

func (p *PortUpsertInMemoryStorage) UpdatePort(shortname string, port *domain.Port) error {
	p.db[shortname] = port
	p.debug()
	return nil
}
