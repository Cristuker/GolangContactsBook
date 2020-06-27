package main //nome do pacote

import (
	"fmt"
	"strconv"
) //formatação

type Contact struct {
	name  string
	phone int
	age   int
}

func main() {

	contacts := [50]Contact{}
	var options int = 1

	for options != 0 {
		fmt.Println("** Agenda **\n")

		fmt.Println("Menu\n")
		fmt.Println("1 - Adicionar um contato")
		fmt.Println("2 - Mostrar todos os contatos")
		fmt.Println("3 - Remover um contato")
		fmt.Println("0 - Sair")
		fmt.Print("-> ")
		fmt.Scanln(&options)

		switch options {
		case 1:
			var name string
			var phone int
			var age int
			var position int

			fmt.Println("Qual a posição desse contato ?")
			fmt.Scanln(&position)

			fmt.Println("Nome:")
			fmt.Scanln(&name)

			fmt.Println("Telefone:")
			fmt.Scanln(&phone)

			fmt.Println("Idade: ")
			fmt.Scanln(&age)

			contacts[position].name = name
			contacts[position].phone = phone
			contacts[position].age = age

			fmt.Println("Contao adicionado:\n")
			fmt.Println("Nome: " + name + "\n Idade: " + strconv.Itoa(age) + "\n Telefone: " + strconv.Itoa(phone))

		case 2:
			fmt.Println("Mostrar todos os contatos")
		case 3:
			fmt.Println("Remover um contato")
		default:
			fmt.Println("Insira uma opção válida\n")
		}

	}
}

/*
pegar dados do menu - ok
switch para as operacações
adicionar no array
tirar do array
mostrar o conteudo do array
*/
