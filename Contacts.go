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
	var addContacts int = 0
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

			fmt.Println("Nome:")
			fmt.Scanln(&name)

			fmt.Println("Telefone:")
			fmt.Scanln(&phone)

			fmt.Println("Idade: ")
			fmt.Scanln(&age)

			contacts[addContacts].name = name
			contacts[addContacts].phone = phone
			contacts[addContacts].age = age

			fmt.Println("Contao adicionado:\n")
			fmt.Println("Nome: " + name + "\nIdade: " + strconv.Itoa(age) + "\nTelefone: " + strconv.Itoa(phone))
			addContacts++

		case 2:

			fmt.Println("Contatos cadastrados:\n")

			for i := 0; i < addContacts; i++ {
				fmt.Println("Nome: " + contacts[i].name + "\nIdade: " + strconv.Itoa(contacts[i].age) + "\nTelefone: " + strconv.Itoa(contacts[i].phone) + "\n")
			}

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
