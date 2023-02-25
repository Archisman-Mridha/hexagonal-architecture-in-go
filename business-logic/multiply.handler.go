package business

import protoGenerated "hexagonalArchitecture/generated/proto"

func(*BusinessLogicLayer) Multiply(
	operands *protoGenerated.Operands) (*protoGenerated.OperationResponse, error) {

	answer := operands.A * operands.B
	return &protoGenerated.OperationResponse{ Answer: answer }, nil
}