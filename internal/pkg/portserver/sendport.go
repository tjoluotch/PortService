package portserver

import (
	pb "PortsServer/internal/pkg/portsprotobuf"
	"context"
	"errors"
	"log"
)

const errMsg = "failed to save port"

func (s *PortServer) SendPort(ctx context.Context, port *pb.Port) (*pb.Result, error) {
	log.Println(port)
	// save to DB
	if err := s.Db.CreatePort(port); err != nil {
		return &pb.Result{Result: errMsg}, errors.New(errMsg)
	}
	log.Println("saved port in database")
	result := &pb.Result{Result: "port received & saved in DB"}
	return result, nil
}
