package main

import (
	"fmt"
)

//SaidaPadraoMessage transporta texto para saida padrao
type SaidaPadraoMessage struct {
	Texto string
}

//GetMessageID pegar o const da Message
func (m SaidaPadraoMessage) GetMessageID() MessageID {
	return SaidaPadraoMessageID
}

//SaidaPadraoPlugin Printar na tela por message
type SaidaPadraoPlugin struct{}

//Init registrar plugin e mensagens
func (p *SaidaPadraoPlugin) Init(kernel *Kernel) {
	//registrar no kernel mensagens
	kernel.Register(p, SaidaPadraoMessageID)
}

//ReceiveMessage notificar quando chega message
func (p *SaidaPadraoPlugin) ReceiveMessage(msg Message) {
	msgSaida := msg.(SaidaPadraoMessage)
	fmt.Println(msgSaida.Texto)
}
