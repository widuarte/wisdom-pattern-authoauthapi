package services

import "wisdom-pattern-authoauthapi/domain"

func SearchService1(context domain.IContextChainLink) {
	context.SetResponseService1("Respuesta servicio 1")
}

func SearchService2(context domain.IContextChainLink) {
	context.SetResponseService1("Respuesta servicio 2")
}

func SearchService3(context domain.IContextChainLink) {
	context.SetResponseService1("Respuesta servicio 3")
}
