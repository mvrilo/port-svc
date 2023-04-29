package domain

type Config struct {
	ServerAddress string `envconfig:"server_address"`
}
