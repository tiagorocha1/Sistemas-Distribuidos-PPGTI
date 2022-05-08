package main

import (
	"fmt"
	"net/rpc"
	"os"
	"strconv"
)

func main() {

	for {

		client, err := rpc.Dial("tcp", ":5000")
		if err != nil {
			fmt.Print("dialing:", err)
		}

		exibeMenu()

		comando := lerComando()

		opcao, _ := strconv.Atoi(comando)

		switch opcao {
		case 1:
			fmt.Println("ID da Lista:")
			id := lerComando()
			fmt.Println("Elemento:")
			elemento := lerComando()
			addItem(id, elemento, *client)
		case 2:
			fmt.Println("ID da Lista:")
			id := lerComando()
			removeItem(id, *client)
		case 3:
			fmt.Println("ID da Lista:")
			id := lerComando()
			fmt.Println("Posicao:")
			position := lerComando()
			getPosition(id, position, *client)
		case 4:
			fmt.Println("ID da Lista:")
			id := lerComando()
			sizeList(id, *client)
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando desconhecido")
			os.Exit(-1)
		}

	}

}

func exibeMenu() {
	fmt.Println("1- Adicionar item")
	fmt.Println("2- Remover item")
	fmt.Println("3- Pegar item na posição")
	fmt.Println("4- Tamanho da Lista")
	fmt.Println("0- Sair")
}

func lerComando() string {
	var comando string
	fmt.Scan(&comando)
	fmt.Println("")
	return comando
}

func addItem(id string, elemento string, client rpc.Client) {

	var args []string
	args = append(args, id)
	args = append(args, elemento)
	var reply bool
	err := client.Call("RemoteList.Append", args, &reply)
	if err != nil {
		fmt.Println(err)
	}

}

func removeItem(id string, client rpc.Client) {

	var args []string
	args = append(args, id)
	var reply int
	err := client.Call("RemoteList.Remove", args, &reply)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Elemento removido: ", reply)
}

func getPosition(id string, position string, client rpc.Client) {
	var args []string
	args = append(args, id)
	args = append(args, position)
	var reply int
	err := client.Call("RemoteList.Get", args, &reply)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Item retornado: ", reply)
}

func sizeList(id string, client rpc.Client) {
	var args []string
	args = append(args, id)
	var reply int
	err := client.Call("RemoteList.Size", args, &reply)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Tamanho da Lista: ", reply)
}
