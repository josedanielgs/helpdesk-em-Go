package controller

import (
	"fmt"
	m "helpdesk/model"
)

type Chamado struct {
	Codigo        string
	Titulo        string
	Descricao     string
	CreatedAt     string
	DataInicio    string
	DataFim       string
	CreatedBy     m.User
	Classificacao string
	Status        string
}

func (m *Chamado) Criar(chamdos []Chamado) Chamado {
	var codigo string
	var titulo string
	var descricao string

	fmt.Scan("digite o codigo:", &codigo)
	fmt.Scan("digite o titulo:", &titulo)
	fmt.Scan("digite a descrição:", &descricao)
	return Chamado{Codigo: codigo, Titulo: titulo, Descricao: descricao}
}
