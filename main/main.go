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
		for logaut != 1 && grupoUsuario != "null" {
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
						fmt.Println("--> Todos os chamados abertos")
						for index := range chamados {
							listaChamadosAbertos[index].PrintChamado()
						}

					}
				} else if opcao == 2 {
					//seExisteChamadosAbertos, listaChamadosAbertos := chamadosAbertos()
					if len(chamados) >= 1 {
						var codigo string
						fmt.Println("Digite o codigo do chamado que deseja solucionar:")
						fmt.Scan(&codigo)
						seExiste, chamado := findByCodigo(&chamados, codigo)
						if seExiste {
							fmt.Println("--> deseja solucionar o seguinte chamado?:")
							chamado.PrintChamado()
							fmt.Println("(sim/não):")
							var resp string
							fmt.Scan(&resp)
							if resp == "sim" {
								fmt.Println("Escreva a solução:")
								//precisa mudar o método de input
								fmt.Scan(&(*chamado).Solucao)
								agora := t.Now().Local()
								dataFormatada := agora.Format("02/01/2006 15:04:05")
								(*chamado).DataSolucao = dataFormatada
								(*chamado).Status = "solucionado"
								fmt.Println("Chamado '", (*chamado).Titulo, "' solucionado")
							} else {
								break
							}
						}

					} else {
						fmt.Println("Não há chamados para ser fechados")
					}
				} else if opcao == 3 {
					//depois terei que transforar a função de chamados abertos em ponteiro
					var chamadosAbertos []m.Chamado
					for indice := range chamados {
						if chamados[indice].Status == "aberto" {
							chamadosAbertos = append(chamadosAbertos, chamados[indice])
						}
					}
					if len(chamadosAbertos) == 0 {
						fmt.Println("Não há chamados abertos")
					} else {
						//for indice := range chamadosAbertos {
						//	chamadosAbertos[indice].PrintChamado()
						//}
						seCodigoExiste := false
						if !seCodigoExiste {
							fmt.Println("Selecione o codigo do chamado que quer fazer a atribuição do responsável\n:")
							var codigo string
							fmt.Scan(&codigo)
							seExiste, chamado := findByCodigo(&chamados, codigo)
							if !seExiste {
								fmt.Println("Não existe chamado com esse codigo")
							} else {
								fmt.Println("Digite a data de inicio do chamado no sequinte formato, DD/MM/AAAA :")
								fmt.Scan(&(*chamado).DataInicio)

								fmt.Println("Digite a data de fim do chamado no sequinte formato, DD/MM/AAAA :")
								fmt.Scan(&(*chamado).DataFim)

								fmt.Println("Digite a classificação do chamado :")
								fmt.Scan(&(*chamado).Classificacao)
								bufio.NewReader(os.Stdin).Reset(os.Stdin)
								//fazer validações
								fmt.Println("Digite o responsável pelo chamado :")
								var user string
								fmt.Scan(&user)
								responsavel := m.User{}
								_, responsavel = findUserByName(user)
								(*chamado).Responsavel = responsavel

								fmt.Println("--> chamado atribuido")
								for indice := range chamados {
									if chamados[indice].Responsavel.User != "" {
										chamados[indice].PrintChamadoComAtribuicao()
									}
								}
								fmt.Println("---")
							}
						}

					}

					fmt.Println("Escreva o codigo do chamado que você quer atribuir:")
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

						fmt.Println("\n--> Chamado:\nTitulo:", chamado.Titulo, "\nDescricao: '", chamado.Descricao, "'\nDeseja abrir esse chamado?:--\n(sim/não):")
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
	}
	if len(chamadosAbertos) > 0 {
		return true, chamadosAbertos
	} else {
		return false, chamados
	}
}

func gerarCodigo() string {
	quantidade := (len(chamados) + 1)
	ano := t.Now().Year()
	return sigla + "00" + conv.Itoa(quantidade) + "-" + conv.Itoa(ano)
}

func findByCodigo(chamados *[]m.Chamado, codigo string) (bool, *m.Chamado) {
	for indice := range *chamados {
		if (*chamados)[indice].Codigo == codigo {
			return true, &(*chamados)[indice]
		}
	}
	vazio := m.Chamado{}
	return false, &vazio
}

func findUserByName(name string) (bool, m.User) {
	for indice := range admins {
		if name == admins[indice].User {
			//fmt.Println("Lofin suceful")
			return true, admins[indice]
		}
	}
	for indice := range users {
		if name == users[indice].User {
			//fmt.Println("Logado com sucesso")
			return true, users[indice]
		}
	}
	return false, m.User{}
}

//func findByCod(cod string) (bool, m.Chamado) {
//	for indice := range chamados {
//		if chamados[indice].Codigo == cod {
//			return true, chamados[indice]
//		}
//	}
//	return false, _
//}
