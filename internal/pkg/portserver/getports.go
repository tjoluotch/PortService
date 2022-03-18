package portserver

import (
	pb "PortsServer/internal/pkg/portsprotobuf"
	"errors"
	"log"
)

func (s *PortServer) GetPorts(req *pb.Request, stream pb.PortsResolver_GetPortsServer) error {
	if req.Request == "" {
		return errors.New("empty request sent")
	}

	ports, err := s.Db.GetAllPorts()
	if err != nil {
		return err
	}
	for _, port := range *ports {
		if err := stream.Send(&port); err != nil {
			return err
		}
	}
	log.Println("finished sending all ports")
	return nil
}
