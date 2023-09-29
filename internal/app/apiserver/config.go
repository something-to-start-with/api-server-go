package apiserver

import (
	"github.com/something-to-start-with/api-server-go/internal/pkg/db/postgres"
)

type Config struct {
	Server   *ServerConfig
	Postgres *postgres.Config
}

type ServerConfig struct {
	Port string
}
