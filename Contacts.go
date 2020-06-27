package main //nome do pacote

import (
	"fmt"
	"strconv"
) //formatação

type Contact struct {
	name     string
	phone    int
	age      int
	isActive bool
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
			fmt.Print("-> ")
			fmt.Scanln(&name)

			fmt.Println("Telefone:")
			fmt.Print("-> ")
			fmt.Scanln(&phone)

			fmt.Println("Idade: ")
			fmt.Print("-> ")
			fmt.Scanln(&age)

			contacts[addContacts].name = name
			contacts[addContacts].phone = phone
			contacts[addContacts].age = age
			contacts[addContacts].isActive = true

			fmt.Println("Contao adicionado:\n")
			fmt.Println("Nome: " + name + "\nIdade: " + strconv.Itoa(age) + "\nTelefone: " + strconv.Itoa(phone))
			addContacts++

		case 2:

			fmt.Println("Contatos cadastrados:\n")

			for i := 0; i < addContacts; i++ {

				if contacts[i].isActive == true {
					fmt.Println("Nº " + strconv.Itoa(i))
					fmt.Println("Nome: " + contacts[i].name + "\nIdade: " + strconv.Itoa(contacts[i].age) + "\nTelefone: " + strconv.Itoa(contacts[i].phone) + "\n")
				}

			}

		case 3:

			var removeIndex int

			fmt.Println("Digite o número do contato que deseja remover")
			fmt.Print("-> ")
			fmt.Scanln(&removeIndex)
			contacts[removeIndex].isActive = false
			var showLater = contacts[removeIndex]

			fmt.Println("\n\nContato removido com sucesso!\n")
			fmt.Println("Nº " + strconv.Itoa(removeIndex))
			fmt.Println("Nome: " + showLater.name + "\nIdade: " + strconv.Itoa(showLater.age) + "\nTelefone: " + strconv.Itoa(showLater.phone) + "\n")

		default:
			fmt.Println("Insira uma opção válida\n")
		}

	}
}

/*
pegar dados do menu - ok
switch para as operacações - ok
adicionar no array - ok
tirar do array
mostrar o conteudo do array - ok
*/
