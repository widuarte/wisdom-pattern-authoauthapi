package main

import (
	"fmt"
	"wisdom-pattern-authoauthapi/domain"
	"wisdom-pattern-authoauthapi/strategy"
)

func main() {
	in := &domain.InputServicesDTO{
		GrandType: domain.REFRESH_TOKEN,
	}

	s := strategy.CreateStrategy(in)
	out := s.Process(in)

	fmt.Println(out.ResponseServices)
}
