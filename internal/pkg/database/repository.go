package database

import (
	pb "PortsServer/internal/pkg/portsprotobuf"
	"errors"
)

type Repository interface {
	CreatePort(port *pb.Port) error
	GetAllPorts() (*[]pb.Port, error)
}

type DB map[string]pb.Port

func (db DB) CreatePort(port *pb.Port) error {
	// since a map can only have unique keys no UpdatePort function needed
	// however if this project were to be implemented with an actual DB then this logic would be necessary
	db[port.Name] = *port
	return nil
}

func (db DB) GetAllPorts() (*[]pb.Port, error) {
	var ports []pb.Port
	if len(db) <= 0 {
		return nil, errors.New("no ports found in database")
	}
	for _, port := range db {
		ports = append(ports, port)
	}
	return &ports, nil
}
