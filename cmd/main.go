package main

import (
	"PortsServer/internal/pkg/factory"
	pb "PortsServer/internal/pkg/portsprotobuf"
	"PortsServer/internal/pkg/serverutils"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	portServer := factory.NewServer()
	listener, err := net.Listen(serverutils.NETWORK, serverutils.GetAddress())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	chanExit := serverutils.InitChannel()

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	log.Println("GRPC server startup...")
	log.Printf("--------------------\n")
	pb.RegisterPortsResolverServer(grpcServer, portServer)

	go func() {
		for {
			select {
			case <-chanExit:
				log.Println("---------------------")
				log.Println("shutting down grpc server connection")
				log.Println("---------------------")
				grpcServer.GracefulStop()
				break
			default:
				continue
			}
		}
	}()
	grpcServer.Serve(listener)
}
