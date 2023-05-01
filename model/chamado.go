package model

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
}
