package portupsert_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/mvrilo/go-port-svc/adapter/inmemory"
	"github.com/mvrilo/go-port-svc/domain"
	"github.com/mvrilo/go-port-svc/usecase/portupsert"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UpsertPortSuite struct {
	suite.Suite
}

func (s *UpsertPortSuite) TestUpsertPort() {
	type testcase struct {
		name string

		portreq   *domain.Port
		shortname string

		storage domain.PortMap
		after   func(testcase)
	}

	samplePort1 := &domain.Port{
		Name:    "Ajman",
		City:    "Ajman",
		Country: "United Arab Emirates",
		Alias:   []any{},
		Regions: []any{},
		Coordinates: []float64{
			55.5136433,
			25.4052165,
		},
		Province: "Ajman",
		Timezone: "Asia/Dubai",
		Unlocs:   []string{"AEAJM"},
		Code:     "52000",
	}

	samplePort2 := &*samplePort1
	samplePort2.Code = "100"

	emptyStorage := domain.PortMap{}
	nonEmptyStorage := domain.PortMap{"AEAJM": samplePort1}

	tests := []testcase{
		{
			name:      "insert port 1, empty storage",
			shortname: "AEAJM",
			portreq:   samplePort1,
			storage:   emptyStorage,
			after: func(tt testcase) {
				assert.Equal(s.T(), tt.storage[tt.shortname], tt.portreq)
			},
		},
		{
			name:      "update port 1",
			shortname: "AEAJM",
			portreq:   samplePort2,
			storage:   nonEmptyStorage,
			after: func(tt testcase) {
				port := tt.storage[tt.shortname]
				assert.Equal(s.T(), port.Code, "100")
			},
		},
	}

	for _, tt := range tests {
		db := inmemory.NewPortUpsertInMemoryStorage(tt.storage)
		svc := portupsert.New(db)
		err := svc.UpsertPort(tt.shortname, tt.portreq)
		if err != nil {
			assert.Error(s.T(), err)
			return
		}
		tt.after(tt)
	}
}

func (s *UpsertPortSuite) TestUpsertPortFile() {
	type testcase struct {
		name     string
		portfile io.Reader
		storage  domain.PortMap
		before   func(testcase)
		after    func(testcase)
	}

	emptyStorage := domain.PortMap{}

	sampleInput := bytes.NewBufferString(`{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  },
  "AEAUH": {
    "name": "Abu Dhabi",
    "coordinates": [
      54.37,
      24.47
    ],
    "city": "Abu Dhabi",
    "province": "Abu Z¸aby [Abu Dhabi]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAUH"
    ],
    "code": "52001"
  }`)

	tests := []testcase{
		{
			name:     "insert port file",
			storage:  emptyStorage,
			portfile: sampleInput,
			before: func(tt testcase) {
				assert.Equal(s.T(), len(tt.storage), 0)
			},
			after: func(tt testcase) {
				assert.Equal(s.T(), len(tt.storage), 2)
			},
		},
	}

	for _, tt := range tests {
		tt.before(tt)
		db := inmemory.NewPortUpsertInMemoryStorage(tt.storage)
		svc := portupsert.New(db)
		err := svc.UpsertPortFile(tt.portfile)
		if err != nil {
			assert.Error(s.T(), err)
			return
		}
		tt.after(tt)
	}
}

func TestUpsertPortSuite(t *testing.T) {
	suite.Run(t, new(UpsertPortSuite))
}
