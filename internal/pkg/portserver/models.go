package portserver

import (
	"PortsServer/internal/pkg/database"
	pb "PortsServer/internal/pkg/portsprotobuf"
)

type PortServer struct {
	pb.UnimplementedPortsResolverServer
	Db database.Repository
}
