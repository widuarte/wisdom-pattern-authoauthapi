package adapter

import (
	"fmt"
	"wisdom-pattern-authoauthapi/domain"
)

type IAdapter interface {
	CalculateResponse(context domain.IContextChainLink) *domain.OutputServicesDTO
}

// CreateAdapter Factory method
func CreateAdapter(input *domain.InputServicesDTO) IAdapter {
	return &Adapter{}
}

type Adapter struct {
}

func (*Adapter) CalculateResponse(context domain.IContextChainLink) *domain.OutputServicesDTO {
	return &domain.OutputServicesDTO{
		ResponseServices: fmt.Sprintf(
			"IN: %v\nR1: %s\nR2: %s\nR3: %s",
			context.GetInputService(),
			context.GetResponseService1(),
			context.GetResponseService2(),
			context.GetResponseService3(),
		),
	}
}
