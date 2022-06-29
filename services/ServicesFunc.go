package services

import "wisdom-pattern-authoauthapi/domain"

func SearchService1(context domain.IContextChainLink) error {
	context.SetResponseService1("Respuesta servicio 1")
	return nil
}

func SearchService2(context domain.IContextChainLink) error {
	context.SetResponseService1("Respuesta servicio 2")
	return nil
}

func SearchService3(context domain.IContextChainLink) error {
	context.SetResponseService1("Respuesta servicio 3")
	return nil
}
