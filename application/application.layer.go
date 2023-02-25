package application

import (
	businessLogic "hexagonalArchitecture/business-logic"
	protoGenerated "hexagonalArchitecture/generated/proto"
)

type ApplicationLayer struct {
	businessLogicLayer *businessLogic.BusinessLogicLayer
}

func(application ApplicationLayer) Add(
	operands *protoGenerated.Operands) (*protoGenerated.OperationResponse, error) {

	return application.businessLogicLayer.Add(operands)
}

func(application ApplicationLayer) Subtract(
	operands *protoGenerated.Operands) (*protoGenerated.OperationResponse, error) {

	return application.businessLogicLayer.Subtract(operands)
}

func(application ApplicationLayer) Multiply(
	operands *protoGenerated.Operands) (*protoGenerated.OperationResponse, error) {

	return application.businessLogicLayer.Multiply(operands)
}

func(application ApplicationLayer) Divide(
	operands *protoGenerated.Operands) (*protoGenerated.OperationResponse, error) {

	return application.businessLogicLayer.Divide(operands)
}

func CreateApplicationLayer(businessLogicLayer *businessLogic.BusinessLogicLayer) ApplicationLayer {
	return ApplicationLayer{ businessLogicLayer }
}