package main //nome do pacote

import "fmt" //formatação

func main() {

	var options int = 1

	for options != 0 {
		fmt.Println("** Agenda **\n")

		fmt.Println("Menu\n")
		fmt.Println("1 - Adicionar um contato")
		fmt.Println("2 - Mostrar todos os contatos")
		fmt.Println("3 - Remover um contato")
		fmt.Println("0 - Sair")
		fmt.Scanln(&options)
		fmt.Println(options)

	}

}
