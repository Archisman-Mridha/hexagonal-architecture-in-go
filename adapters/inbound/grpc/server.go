package server

import (
	"log"
	"net"

	"google.golang.org/grpc"

	protoGenerated "hexagonalArchitecture/generated/proto"
	"hexagonalArchitecture/ports"
)

type GRPCAdapter struct {
	protoGenerated.UnimplementedMathServer
	applicationLayer ports.ApplicationLayerPort

	serverInstance *grpc.Server
}

func(adapter GRPCAdapter) StartServer( ) {

	var (
		tcpListener net.Listener
		error error
	)
	if tcpListener, error= net.Listen("tcp", "0.0.0.0:4000");
		error != nil {
			log.Fatalf("failed to listen on port 4000 : %s", error.Error( ))}

	adapter.serverInstance= grpc.NewServer( )

	implementedMathGRPCService := adapter
	protoGenerated.RegisterMathServer(adapter.serverInstance, implementedMathGRPCService)

	if error= adapter.serverInstance.Serve(tcpListener);
		error != nil {
			log.Fatalf("error occured in gRPC server : %s", error.Error( ))}
}

func(adapter GRPCAdapter) StopServer( ) {
	if adapter.serverInstance != nil {
		adapter.StopServer( )

		log.Println("stopped gRPC server")
	}
}

func CreateGRPCAdapter(applicationLayer ports.ApplicationLayerPort) *GRPCAdapter {
	return &GRPCAdapter{
		applicationLayer: applicationLayer,
		serverInstance: nil,
	}
}