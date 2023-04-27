package main

import (
	//c "dit/paraestudo/src/helpdesk/controller"
	"fmt"
	m "helpdesk/model"
)

// lista de "objetos" de usuários para definir grupos de acesso
var admins = []m.User{{User: "admin.admin", Password: "123"}}
var users = []m.User{{User: "jose.teste", Password: "123"}}
var chamados = []m.Chamado{}

func main() {

	login := 0

	for login < 1 {
		fmt.Println("por favor faça login:")
		grupoUsuario := validacao()
		if grupoUsuario == "admin" {
			fmt.Println("1. ver chamados abertos\n2. solucionar chamado\n3. direcionar chamado")
			var opcao int
			fmt.Scan("\ngigite a opção que deseja:", &opcao)
			if opcao == 1 {
				chamadosAbertos := []m.Chamado{}
				for indice := range chamados {
					if chamados[indice].Status == "aberto" {
						chamadosAbertos = append(chamadosAbertos, chamados[indice])
					}
				}
				fmt.Println("Todos os chamados abertos são: ", chamadosAbertos)
			}
		} else if grupoUsuario == "user" {
			fmt.Println("1. ver seus chamados abertos\n2. abrir chamado")
			var opcao int
			fmt.Scan("\ngigite a opção que deseja:", &opcao)
			if opcao == 2 {
				//append(chamados, c.)
			}
		}

	}

}

func validacao() string {
	var writeUser string
	var writePassword string
	fmt.Print("login:")
	fmt.Scan(&writeUser)
	fmt.Print("password:")
	fmt.Scan(&writePassword)
	for indice := range admins {
		if writeUser == admins[indice].User && writePassword == admins[indice].Password {
			fmt.Println("Lofin suceful")
			return "admin"
		}
	}
	for indice := range users {
		if writeUser == users[indice].User && writePassword == users[indice].Password {
			fmt.Println("Logado com sucesso")
			return "user"
		}
	}
	fmt.Println("Login ou senha incorretor")
	return "null"

}
