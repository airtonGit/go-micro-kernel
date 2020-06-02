package main

//micro kernel pattern
//ao iniciar, kernel instancia todos os plugins registrados
//cada plugin registra ao kernel quais mensagens se interessa
// Message formato da message
type Message interface {
}

//Plugin indica formato do plugin
type Plugin interface {
	init()
	receiveMessage(msg Message)
}

type kernel struct {
	Plugins []Plugin

	//cada mensagem, tem lista de plugins interessados\
	MessageObservers map[Message]map[Plugin]bool
}

func (k *kernel) register(plugin Plugin, msg Message) {
	if _, ok := k.MessageObservers[plugin]; ok == false {
		k.MessageObservers[msg] = []Plugin{}
	}
	//procurar se plugin já está presente na lista
	
}
