package main

//micro kernel pattern
//ao iniciar, kernel instancia todos os plugins registrados
//cada plugin registra ao kernel quais mensagens se interessa

// Message formato da message
type Message interface {
	GetMessageID() MessageID
}

//Plugin indica formato do plugin
type Plugin interface {
	Init(k *Kernel)
	ReceiveMessage(msg Message)
}

//Kernel registra plugins, mensagen e observadores. Despacha mensagens a todos os interessados
type Kernel struct {
	Plugins []Plugin

	//cada mensagem, tem lista de plugins interessados\
	MessageObservers map[MessageID]map[Plugin]bool
}

//Init em todos os Plugins
func (k *Kernel) Init() {
	//Init em todos os plugins
	for _, v := range k.Plugins {
		v.Init(k)
	}
}

//Register plugin informa qual message deseja ser notificado
func (k *Kernel) Register(plugin Plugin, msg MessageID) {
	if k.MessageObservers == nil {
		k.MessageObservers = make(map[MessageID]map[Plugin]bool)
	}

	if _, ok := k.MessageObservers[msg]; ok == false {
		k.MessageObservers[msg] = map[Plugin]bool{}
	}
	//procurar se plugin já está presente no map
	if _, ok := k.MessageObservers[msg][plugin]; ok == false {
		//adicionar plugin ao map
		k.MessageObservers[msg][plugin] = true
	}
}

//PostMessage Plugin solicita ao Kernel envio de sua message
func (k *Kernel) PostMessage(msg Message) {
	//Kernel deve notificar todos os interessado da message postada

	for plugin := range k.MessageObservers[msg.GetMessageID()] {
		(plugin).ReceiveMessage(msg)
	}
}
