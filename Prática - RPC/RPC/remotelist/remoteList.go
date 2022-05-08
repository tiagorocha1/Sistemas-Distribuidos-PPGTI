package remotelist

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type RemoteList struct {
	mu      sync.Mutex
	MapList map[string][]int
}

func (l *RemoteList) Append(args []string, reply *bool) error {
	fmt.Println("Adicionando item...")
	l.mu.Lock()
	defer l.mu.Unlock()
	defer save(l)

	id := args[0]
	value, _ := strconv.Atoi(args[1])

	l.MapList[id] = append(l.MapList[id], value)
	fmt.Println(l.MapList)

	*reply = true
	return nil
}

func (l *RemoteList) Remove(args []string, reply *int) error {
	fmt.Println("Removendo item...")
	l.mu.Lock()
	defer l.mu.Unlock()
	id := args[0]

	if len(l.MapList[id]) > 0 {

		l.MapList[id] = l.MapList[id][:len(l.MapList[id])-1]
		fmt.Println(l.MapList)
	} else {
		return errors.New("empty list")
	}

	return nil
}

func (l *RemoteList) Get(args []string, reply *int) error {
	fmt.Println("Pegando item...")
	l.mu.Lock()
	defer l.mu.Unlock()
	id := args[0]
	position, _ := strconv.Atoi(args[1])

	if len(l.MapList[id]) > 0 {
		for i := 0; i < len(l.MapList[id]); i++ {
			if position == i {
				*reply = l.MapList[id][i]
				fmt.Println("Valor retornado: ", l.MapList[id][i])
			}
		}

		fmt.Println(l.MapList)
	} else {
		return errors.New("empty list")
	}

	return nil
}

func (l *RemoteList) Size(args []string, reply *int) error {
	fmt.Println("Tamanho da lista...")
	l.mu.Lock()
	defer l.mu.Unlock()
	id := args[0]

	if len(l.MapList[id]) > 0 {
		*reply = len(l.MapList[id])
		fmt.Println("Tamanho retornado: ", len(l.MapList[id]))
	} else {
		return errors.New("empty list")
	}

	return nil
}

func NewRemoteList() *RemoteList {
	return new(RemoteList)
}

func save(l *RemoteList) {
	arquivo, err := os.OpenFile("maplist.txt", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println(err)
	}

	jsonMap, _ := json.Marshal(l.MapList)
	arquivo.WriteString(string(jsonMap))

	arquivo.Close()
}
