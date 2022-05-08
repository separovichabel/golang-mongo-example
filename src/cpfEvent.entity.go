package main

type CPF string

type CPFEventType int

const (
	UltimaConsultaBureau CPFEventType = iota
	Movimentacao
	UltimaCompraCartao
)

func (c CPFEventType) String() string {
	return [...]string{"UltimaConsultaBureau", "Movimentacao", "UltimaCompraCartao"}[c-1]
}
