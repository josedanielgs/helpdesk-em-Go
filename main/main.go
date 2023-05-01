package main

import (
	//c "dit/paraestudo/src/helpdesk/controller"
	"fmt"
	m "helpdesk/model"
	t "time"
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
		logaut := 0
		for logaut != 1 {
			if grupoUsuario == "admin" {
				fmt.Println("1. ver chamados abertos\n2. solucionar chamado\n3. direcionar chamado\n4. logaut")
				var opcao int
				fmt.Println("\ngigite a opção que deseja:")
				fmt.Scan(&opcao)
				if opcao == 1 {
					seExisteChamadosAbertos, listaChamadosAbertos := chamadosAbertos()
					if seExisteChamadosAbertos {
						fmt.Println("Não há chamados abertos")
					} else {
						fmt.Println("Todos os chamados abertos são: ", listaChamadosAbertos)
					}
				} else if opcao == 2 {
					seExisteChamadosAbertos, listaChamadosAbertos := chamadosAbertos()
					if seExisteChamadosAbertos {
						var codigo string
						fmt.Println("Digite o codigo do chamado que deseja solucionar:")
						fmt.Scan(&codigo)
						for indice := range listaChamadosAbertos {
							if listaChamadosAbertos[indice].Codigo == codigo {
								fmt.Println("deseja solucionar o seguinte chamado?:", listaChamadosAbertos[indice])
								var resp string
								fmt.Scan(&resp)
								if resp == "sim" {
									fmt.Println("Escreva a solução:")
									fmt.Scan(&chamados[indice].Solucao)
									agora := t.Now().Local()
									dataFormatada := agora.Format("01/01/2000")
									chamados[indice].DataSolucao = dataFormatada
									chamados[indice].Status = "solucionado"
									fmt.Println("Chamado '", chamados[indice].Titulo, "' solucionado")
								} else {
									break
								}
							}
						}
					}
				}
			} else if grupoUsuario == "user" {
				fmt.Println("1. ver seus chamados abertos\n2. abrir chamado\n3. logaut")
				var opcao int
				fmt.Scan("\ngigite a opção que deseja:", &opcao)
				if opcao == 2 {
					abrir := ""
					if abrir != "sim" {
						var chamado *m.Chamado
						chamado = new(m.Chamado)

						fmt.Println("Digite o titulo do chamado")
						fmt.Scan(&chamado.Titulo)

						fmt.Println("Digite a descrição do chamado")
						fmt.Scan(&chamado.Descricao)

						fmt.Println("Deseja abrir esse chamado?: \nTitulo:", chamado.Titulo, ", \nDescricao: '", chamado.Descricao, "'")
						fmt.Scan(&abrir)
						if abrir == "sim" {
							chamados = append(chamados, *chamado)
							fmt.Println("Chamado aberto com sucesso")
						} else {
							fmt.Println("Deseja parar de abrir chamado?")
							fmt.Scan(&abrir)
						}
					}
					//append(chamados, c.)
				}
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

func chamadosAbertos() (bool, []m.Chamado) {
	chamadosAbertos := []m.Chamado{}
	if len(chamadosAbertos) == 0 {
		return false, chamadosAbertos
	} else {
		for indice := range chamados {
			if chamados[indice].Status == "aberto" {
				chamadosAbertos = append(chamadosAbertos, chamados[indice])
			}
		}
		return true, chamadosAbertos
	}
}

func findByCod(cod string) (bool, m.Chamado) {
	for indice := range chamados {
		if chamados[indice].Codigo == cod {
			return true, chamados[indice]
		}
	}
	return false, _
}
