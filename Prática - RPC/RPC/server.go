package main

import (
	"RPC/remotelist"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/rpc"
	"reflect"
)

func main() {
	list := new(remotelist.RemoteList)

	persistence := retrieveList()

	if persistence != nil {

		list.MapList = persistence
	} else {
		list.MapList = make(map[string][]int)
	}

	rpcs := rpc.NewServer()
	rpcs.Register(list)
	listener, err := net.Listen("tcp", ":5000")
	defer listener.Close()

	if err != nil {
		fmt.Println("listen error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err == nil {
			go rpcs.ServeConn(conn)
		} else {
			fmt.Println(err)
			break
		}
	}
}

func retrieveList() map[string][]int {

	mapTxt := read()

	var result map[string][]int

	json.Unmarshal([]byte(mapTxt), &result)

	fmt.Println(reflect.TypeOf(result))

	fmt.Println(result)

	return result
}

func read() string {

	mapData, _ := ioutil.ReadFile("maplist.txt")

	mapJson := string(mapData)

	return mapJson
}
