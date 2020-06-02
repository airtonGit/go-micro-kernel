package main

//MessageID identificador mensagens
type MessageID int

const (
	//SaidaPadraoMessageID printar texto na tela
	SaidaPadraoMessageID MessageID = iota
)

//SaidaPadraoMessage Carrega Texto para printar

func main() {
	saidaPradrao := &SaidaPadraoPlugin{}

	kernel := Kernel{
		Plugins: []Plugin{saidaPradrao}}

	kernel.Init()

	olaMundo := SaidaPadraoMessage{
		Texto: "Olá Mundo!",
	}

	kernel.PostMessage(olaMundo)
}
