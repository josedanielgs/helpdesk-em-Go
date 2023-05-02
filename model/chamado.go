package model

import "fmt"

type Chamado struct {
	Codigo        string
	Titulo        string
	Descricao     string
	CreatedAt     string
	DataInicio    string
	DataFim       string
	CreatedBy     User
	Classificacao string
	Status        string
	Solucao       string
	DataSolucao   string
	Responsavel   User
}

func (m *Chamado) PrintChamado() {
	fmt.Println("Codigo: ", m.Codigo,
		"\nTitulo: ", m.Titulo,
		"\nDescricao: ", m.Descricao,
		"\nCriado por: ", m.CreatedBy.User,
		"\nData de criação: ", m.CreatedAt,
		"\n-")
}

func (m *Chamado) PrintChamadoComAtribuicao() {
	fmt.Println("Codigo: ", m.Codigo,
		"\nTitulo: ", m.Titulo,
		"\nDescricao: ", m.Descricao,
		"\nCriado por: ", m.CreatedBy.User,
		"\nData de criação: ", m.CreatedAt,
		"\nResponsavel: ", m.Responsavel.User,
		"\nClassificação: ", m.Classificacao,
		"\nData de inicio: ", m.DataInicio,
		"\nData fim: ", m.DataFim,
		"\n-")
}
