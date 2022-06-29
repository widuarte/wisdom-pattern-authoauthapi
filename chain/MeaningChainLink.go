package chain

import "wisdom-pattern-authoauthapi/domain"

type Execute func(context domain.IContextChainLink) error

type IChainLink interface {
	Execute(context domain.IContextChainLink)
	SetNext(next IChainLink) IChainLink
}

// chainLinkAbstract Definición basica para flujos simple.
// actua como estructura abstracta
type chainLinkAbstract struct {
	next IChainLink
}

func (i *chainLinkAbstract) SetNext(next IChainLink) IChainLink {
	i.next = next
	return next
}

func (i *chainLinkAbstract) processNext(context domain.IContextChainLink) {
	if i.next != nil {
		i.next.Execute(context)
	}
}

// SimpleChainLink Ejecuta una tarea y continua
type SimpleChainLink struct {
	chainLinkAbstract
	ExecuteMethod Execute
}

func (i *SimpleChainLink) Execute(context domain.IContextChainLink) {
	i.ExecuteMethod(context)
	i.processNext(context)
}

// ListChainLink Ejecuta una lista de tareas secuencial
type ListChainLink struct {
	chainLinkAbstract
	ExecutesMethods []Execute
}

func (i *ListChainLink) Execute(context domain.IContextChainLink) {
	for _, e := range i.ExecutesMethods {
		if err := e(context); err != nil {
			// Si se presenta un error no deja seguir la cadena
			return
		}
	}
	i.processNext(context)
}
