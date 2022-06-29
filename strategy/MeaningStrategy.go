package strategy

import (
	"wisdom-pattern-authoauthapi/adapter"
	chainlink "wisdom-pattern-authoauthapi/chain"
	"wisdom-pattern-authoauthapi/domain"
	"wisdom-pattern-authoauthapi/services"
)

type IStrategy interface {
	Process(input *domain.InputServicesDTO) *domain.OutputServicesDTO
}

func CreateStrategy(input *domain.InputServicesDTO) IStrategy {
	adapter := adapter.CreateAdapter(input)
	switch input.GrandType {
	case domain.AUTHORIZATION_CODE:
		return CreateAuthorizationCodeStrategy(adapter)
	case domain.CLIENT_CREDENTIAL:
		return CreateClientCredentialStrategy(adapter)
	case domain.REFRESH_TOKEN:
		return CreateRefreshTokenStrategy(adapter)
	default:
		return nil
	}
}

func CreateAuthorizationCodeStrategy(adapter adapter.IAdapter) IStrategy {
	s := &abstractStrategy{
		adapter: adapter,
	}
	s.firstChainLink = &chainlink.SimpleChainLink{
		ExecuteMethod: services.SearchService1,
	}
	s.firstChainLink.SetNext(&chainlink.SimpleChainLink{
		ExecuteMethod: services.SearchService2,
	}).SetNext(&chainlink.SimpleChainLink{
		ExecuteMethod: services.SearchService3,
	})
	return s
}

func CreateRefreshTokenStrategy(adapter adapter.IAdapter) IStrategy {
	s := &abstractStrategy{
		adapter: adapter,
	}
	s.firstChainLink = &chainlink.ListChainLink{
		ExecutesMethods: []chainlink.Execute{
			services.SearchService2,
			services.SearchService3,
		},
	}
	s.firstChainLink.SetNext(&chainlink.SimpleChainLink{
		ExecuteMethod: services.SearchService1,
	})
	return s
}

func CreateClientCredentialStrategy(adapter adapter.IAdapter) IStrategy {
	s := &abstractStrategy{
		adapter: adapter,
	}
	s.firstChainLink = &chainlink.SimpleChainLink{
		ExecuteMethod: services.SearchService1,
	}
	s.firstChainLink.SetNext(&chainlink.ListChainLink{
		ExecutesMethods: []chainlink.Execute{
			services.SearchService2,
			services.SearchService3,
		},
	})
	return s
}

type abstractStrategy struct {
	adapter        adapter.IAdapter
	firstChainLink chainlink.IChainLink
}

func (i *abstractStrategy) Process(input *domain.InputServicesDTO) *domain.OutputServicesDTO {
	context := makeContext(input)
	i.firstChainLink.Execute(context)
	return i.adapter.CalculateResponse(context)
}

func makeContext(input *domain.InputServicesDTO) domain.IContextChainLink {
	context := &domain.ContextChainLink{
		InputService: input,
	}
	return context
}
