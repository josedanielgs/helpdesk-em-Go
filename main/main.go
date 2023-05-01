package main

import (
	"bufio"
	"os"
	conv "strconv"

	//c "dit/paraestudo/src/helpdesk/controller"
	"fmt"
	m "helpdesk/model"
	t "time"
)

// lista de "objetos" de usuários para definir grupos de acesso
var admins = []m.User{{User: "admin.admin", Password: "123"}}
var users = []m.User{{User: "jose.teste", Password: "123"}}
var usuarioAtual m.User
var chamados = []m.Chamado{}
var sigla = "HP-"

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
					if !seExisteChamadosAbertos {
						fmt.Println("Não há chamados abertos")
					} else {
						fmt.Println("--> Todos os chamados")
						for index := range chamados {
							fmt.Println("Codigo: ", listaChamadosAbertos[index].Codigo,
								"\nTitulo: ", listaChamadosAbertos[index].Titulo,
								"\nDescricao: ", listaChamadosAbertos[index].Descricao,
								"\nCriado por: ", listaChamadosAbertos[index].CreatedBy.User,
								"\n Data de criação: ", listaChamadosAbertos[index].CreatedAt,
								"\n-")
						}

					}
				} else if opcao == 2 {
					seExisteChamadosAbertos, listaChamadosAbertos := chamadosAbertos()
					if seExisteChamadosAbertos {
						var codigo string
						fmt.Println("Digite o codigo do chamado que deseja solucionar:")
						fmt.Scan(&codigo)
						for indice := range listaChamadosAbertos {
							if listaChamadosAbertos[indice].Codigo == codigo {
								fmt.Println("deseja solucionar o seguinte chamado?:", listaChamadosAbertos[indice], "--\n(sim/não):")
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
				fmt.Println("Digite a opção que deseja:")
				fmt.Scan(&opcao)
				if opcao == 2 {
					abrir := ""
					if abrir != "sim" {
						var chamado *m.Chamado
						chamado = new(m.Chamado)
						scanner := bufio.NewReader(os.Stdin)

						fmt.Println("Digite o titulo do chamado")
						//fmt.Scan(&chamado.Titulo)
						titulo, _ := scanner.ReadString('\n')
						chamado.Titulo = titulo

						//bufio.NewReader(os.Stdin).Reset(os.Stdin)

						fmt.Println("Digite a descrição do chamado")
						//fmt.Scan(&chamado.Descricao)
						descricao, _ := scanner.ReadString('\n')
						chamado.Descricao = descricao

						fmt.Println("\n --> Deseja abrir esse chamado?: \nTitulo:", chamado.Titulo, ", \nDescricao: '", chamado.Descricao, "'\n --\n(sim/não):")
						fmt.Scan(&abrir)
						if abrir == "sim" {
							chamado.Status = "aberto"
							chamado.Codigo = gerarCodigo()
							agora := t.Now().Local()
							dataFormatada := agora.Format("02/01/2006 15:04:05")
							chamado.CreatedAt = dataFormatada
							chamado.CreatedBy = usuarioAtual
							chamados = append(chamados, *chamado)
							fmt.Println("Chamado aberto com sucesso")
						} else {
							fmt.Println("Deseja parar de abrir chamado?")
							fmt.Scan(&abrir)
						}
					}
					//append(chamados, c.)
				} else if opcao == 3 {
					logaut = 1
					usuarioAtual = m.User{}
					break
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
			usuarioAtual = admins[indice]
			return "admin"
		}
	}
	for indice := range users {
		if writeUser == users[indice].User && writePassword == users[indice].Password {
			fmt.Println("Logado com sucesso")
			usuarioAtual = users[indice]
			return "user"
		}
	}
	fmt.Println("Login ou senha incorretor")
	return "null"

}

func chamadosAbertos() (bool, []m.Chamado) {
	var chamadosAbertos = []m.Chamado{}
	if len(chamados) == 0 {
		return false, chamados
	} else {
		for indice := range chamados {
			if chamados[indice].Status == "aberto" {
				chamadosAbertos = append(chamadosAbertos, chamados[indice])
			}
		}
		return true, chamadosAbertos
	}
}

func gerarCodigo() string {
	quantidade := (len(chamados) + 1)
	ano := t.Now().Year()
	return sigla + "00" + conv.Itoa(quantidade) + "-" + conv.Itoa(ano)
}

//func findByCod(cod string) (bool, m.Chamado) {
//	for indice := range chamados {
//		if chamados[indice].Codigo == cod {
//			return true, chamados[indice]
//		}
//	}
//	return false, _
//}
