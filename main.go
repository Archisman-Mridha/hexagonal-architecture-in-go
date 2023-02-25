package main

import (
	gRPC "hexagonalArchitecture/adapters/inbound/grpc"
	application "hexagonalArchitecture/application"
	businessLogic "hexagonalArchitecture/business-logic"
)

func main( ) {

	// business logic layer
	businessLogicLayer := businessLogic.CreateBusinessLogicLayer( )

	// application layer
	applicationLayer := application.CreateApplicationLayer(businessLogicLayer)

	// gRPC server as an inbound adapter which acts as the presentation layer
	gRPCAdapter := gRPC.CreateGRPCAdapter(applicationLayer)
	gRPCAdapter.StartServer( )
	defer gRPCAdapter.StopServer( )
}