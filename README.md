# port-svc

port-svc receives ports information as file and upserts it (insert or update) in a datastore. Current implementation is using an inmemory datastore backed by `map[string]domain.Port` (check [./domain](domain) for more domain definitions).

It also exposes a http server listening on a default port 8000. This server handles incoming requests for upserting ports:

```
POST /ports
```

Using curl for testing:

```
curl -svf -X POST -d "$(cat ports.json)" http://localhost:8000/ports
```

Using multipart upload:

```
curl -vf -F "uploadFile=@\"ports.json\"" http://localhost:8000/ports

```

## Running

Using docker-compose:

```
docker compose up port-svc
```

## Building

`make build`

or

`go build -o port-svc ./cmd/port-svc/*.go`

## Testing

`make test`

or

`go test -race ./...`

## Configuration

Environment variables (and `.env` file) are supported.

Example:

```
SERVER_ADDRESS=127.0.0.1:8000
```

## Architecture overview

```
.
├── adapter        # adapter implementations
│   ├── http       # http adapters
│   └── inmemory   # in memory adapters
├── cmd            # binary entrypoints
│   └── port-svc   # main binary
├── domain         # domain definitions
└── usecase        # usecase implementations
    └── portupsert # main implementation

9 directories
```

## Next steps

- Validation of the ports file json schema
- Limit configuration for the ports file size
- HTTP adapter: add logging, metrics and health check - (maybe move from net/http)
- Add other data storages for actual persistence of the data
- Github actions pipeline
- Improve command-line ux: use cobra, config file, add help/version subcommand, etc
