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
	inputService      *InputServicesDTO
	responseServices1 string
	responseServices2 string
	responseServices3 string
}

func (i *ContextChainLink) GetInputService() *InputServicesDTO {
	return i.inputService
}

func (i *ContextChainLink) SetResponseService1(response string) {
	i.responseServices1 = response
}

func (i *ContextChainLink) GetResponseService1() string {
	return i.responseServices1
}

func (i *ContextChainLink) SetResponseService2(response string) {
	i.responseServices2 = response
}

func (i *ContextChainLink) GetResponseService2() string {
	return i.responseServices2
}

func (i *ContextChainLink) SetResponseService3(response string) {
	i.responseServices3 = response
}

func (i *ContextChainLink) GetResponseService3() string {
	return i.responseServices3
}
