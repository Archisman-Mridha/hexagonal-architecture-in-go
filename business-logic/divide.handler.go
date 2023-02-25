package business

import (
	"errors"

	protoGenerated "hexagonalArchitecture/generated/proto"
)

func(*BusinessLogicLayer) Divide(
	operands *protoGenerated.Operands) (*protoGenerated.OperationResponse, error) {

	if operands.B == 0 {
		return nil, errors.New("can't divide by 0")}

	answer := operands.A / operands.B
	return &protoGenerated.OperationResponse{ Answer: answer }, nil
}