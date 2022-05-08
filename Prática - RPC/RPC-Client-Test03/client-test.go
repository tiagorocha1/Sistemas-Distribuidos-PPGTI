package main

import (
	"fmt"
	"math/rand"
	"net/rpc"
	"strconv"
)

func main() {

	done1 := make(chan int)
	done2 := make(chan int)
	done3 := make(chan int)
	done4 := make(chan int)
	done5 := make(chan int)

	id1 := 10
	id2 := 11
	id3 := 12
	id4 := 13
	id5 := 14

	for i := 0; i < 10; i++ {

		go teste(id1, "50", done1)
		go teste(id2, "50", done2)
		go teste(id3, "50", done3)
		go teste(id4, "50", done4)
		go teste(id5, "50", done5)

	}

	thread01Terminou := <-done1
	thread02Terminou := <-done2
	thread03Terminou := <-done3
	thread04Terminou := <-done4
	thread05Terminou := <-done5

	fim := 9
	for thread01Terminou != fim && thread02Terminou != fim && thread03Terminou != fim && thread04Terminou != fim && thread05Terminou != fim {
		fmt.Println("thread 01 terminou?", thread01Terminou)
		fmt.Println("thread 02 terminou?", thread02Terminou)
		fmt.Println("thread 03 terminou?", thread03Terminou)
		fmt.Println("thread 04 terminou?", thread04Terminou)
		fmt.Println("thread 05 terminou?", thread05Terminou)
		fmt.Println("==================================")

		thread01Terminou = <-done1
		thread02Terminou = <-done2
		thread03Terminou = <-done3
		thread04Terminou = <-done4
		thread05Terminou = <-done5
	}

	fmt.Println("Fim")
}

func teste(loop int, id string, done chan int) {
	client, err := rpc.Dial("tcp", ":5000")
	if err != nil {
		fmt.Print("dialing:", err)
	}

	for i := 0; i < loop; i++ {
		randomNumber := rand.Int()
		randomNumberS := strconv.FormatInt(int64(randomNumber), 10)

		addItem(id, randomNumberS, *client)
		done <- i
	}

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
