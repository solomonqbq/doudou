package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

func Echo(c net.Conn) {
	defer c.Close()
	for {
		line, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Printf("Failure to read:%s\n", err.Error())
			return
		}
		_, err = c.Write([]byte(line))
		if err != nil {
			fmt.Printf("Failure to write: %s\n", err.Error())
			return
		}
	}
}

func main() {
	port := "3306"
	flag.StringVar(&port, "p", "", "port to listen")
	flag.Parse()
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("Failure to listen: %s\n", err.Error())
	}
	fmt.Printf("Server is listen on %s ...\n", port)
	for {
		if c, err := l.Accept(); err == nil {
			go Echo(c) //new thread
		}
	}
}
