package domain

type IContextChainLink interface {
	GetInputService() *InputServicesDTO
	SetResponseService1(response string)
	GetResponseService1() string
	SetResponseService2(response string)
	GetResponseService2() string
	SetResponseService3(response string)
	GetResponseService3() string
}

type ContextChainLink struct {
	InputService      *InputServicesDTO
	ResponseServices1 string
	ResponseServices2 string
	ResponseServices3 string
}

func (i *ContextChainLink) GetInputService() *InputServicesDTO {
	return i.InputService
}

func (i *ContextChainLink) SetResponseService1(response string) {
	i.ResponseServices1 = response
}

func (i *ContextChainLink) GetResponseService1() string {
	return i.ResponseServices1
}

func (i *ContextChainLink) SetResponseService2(response string) {
	i.ResponseServices2 = response
}

func (i *ContextChainLink) GetResponseService2() string {
	return i.ResponseServices2
}

func (i *ContextChainLink) SetResponseService3(response string) {
	i.ResponseServices3 = response
}

func (i *ContextChainLink) GetResponseService3() string {
	return i.ResponseServices3
}
