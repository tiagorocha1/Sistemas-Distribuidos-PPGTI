package main

import (
	"fmt"
	"math/rand"
	"net/rpc"
	"strconv"
)

func main() {

	for i := 0; i < 100; i++ {
		done := make(chan string)
		go teste(100, "50", done)
		fmt.Println(i, <-done)

	}

}

func teste(loop int, id string, done chan string) {
	client, err := rpc.Dial("tcp", ":5000")
	if err != nil {
		fmt.Print("dialing:", err)
	}

	for i := 0; i < loop; i++ {
		randomNumber := rand.Int()
		randomNumberS := strconv.FormatInt(int64(randomNumber), 10)

		addItem(id, randomNumberS, *client)
	}

	done <- "Terminei"
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
