package ports

import protoGenerated "hexagonalArchitecture/generated/proto"

type ApplicationLayerPort interface {

	Add(operands *protoGenerated.Operands) (*protoGenerated.OperationResponse, error)
	Subtract(operands *protoGenerated.Operands) (*protoGenerated.OperationResponse, error)
	Multiply(operands *protoGenerated.Operands) (*protoGenerated.OperationResponse, error)
	Divide(operands *protoGenerated.Operands) (*protoGenerated.OperationResponse, error)
}