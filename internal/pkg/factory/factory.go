package factory

import (
	"PortsServer/internal/pkg/database"
	ps "PortsServer/internal/pkg/portserver"
)

func NewServer() *ps.PortServer {
	return &ps.PortServer{Db: newDB()}
}

func newDB() database.Repository {
	return database.DB{}
}
