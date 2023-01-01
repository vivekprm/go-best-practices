package main

import (
	"fmt"
	"net"
	"time"
)

func broadcastMsg(msg string, addrs []string) error {
	errc := make(chan error, len(addrs))
	for _, addr := range addrs {
		go func(addr string) {
			errc <- sendMsg(msg, addr)
			fmt.Println("done")
		}(addr)
	}

	for range addrs {
		if err := <-errc; err != nil {
			return err
		}
	}
	return nil
}

func sendMsg(msg, addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = fmt.Fprint(conn, msg)
	return err
}

func main() {
	addr := []string{"localhost:8080", "http://google.com"}
	err := broadcastMsg("hi", addr)

	time.Sleep(time.Second)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("everything went fine!")
}
