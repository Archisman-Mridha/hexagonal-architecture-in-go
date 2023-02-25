package server

import (
	"context"

	protoGenerated "hexagonalArchitecture/generated/proto"
)

func(adapter GRPCAdapter) Add(
	ctx context.Context, operands *protoGenerated.Operands) (*protoGenerated.OperationResponse, error) {

	return adapter.applicationLayer.Add(operands)
}

func(adapter GRPCAdapter) Subtract(
	ctx context.Context, operands *protoGenerated.Operands) (*protoGenerated.OperationResponse, error) {

	return adapter.applicationLayer.Subtract(operands)
}

func(adapter GRPCAdapter) Multiply(
	ctx context.Context, operands *protoGenerated.Operands) (*protoGenerated.OperationResponse, error) {

	return adapter.applicationLayer.Multiply(operands)
}

func(adapter GRPCAdapter) Divide(
	ctx context.Context, operands *protoGenerated.Operands) (*protoGenerated.OperationResponse, error) {

	return adapter.applicationLayer.Divide(operands)
}