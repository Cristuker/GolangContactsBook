package main //Definindo para o pacote main

import (
	"fmt"     //Biblioteca para formatar a saida do terminal
	"strconv" //Biblioteca para conversão, usado para converter inteiro em string
)

type Contact struct { //Struct da mesma forma que é usada no C
	name     string
	phone    int
	age      int
	isActive bool
}

func main() {
	//A atribuição no Golang é feito por meio do := que é usado para declarar variaveis de formas curtas
	contacts := [50]Contact{}
	var addContacts int = 0 // O = também é usado para atribuições mais extensas var NomedaVariavel tipo da variavel
	var options int = 1

	//Em go os parenteses não são usados para definiri uma expressão como nas outras linaguagens
	for options != 0 { //No golang não existe while ou do while e sim o for acaba sendo adaptado para a forma que deseja usar
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

				if contacts[i].isActive == true { //Em if os parenteses não aparecem também, apenas as chaves
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
