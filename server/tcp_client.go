package main

import (
	"bufio"

	"fmt"

	"net"

	"os"

	"time"
)

var host string
var port string

type Clienter struct {
	client net.Conn

	isAlive bool

	SendStr chan string

	RecvStr chan string
}

func (c *Clienter) Connect(ip, port string) bool {

	if c.isAlive {
		return true
	} else {

		var err error

		c.client, err = net.Dial("tcp", ip+":"+port)

		if err != nil {
			fmt.Printf("Failure to connet:%s\n", err.Error())
			return false

		}

		c.isAlive = true

	}

	return true

}

func (c *Clienter) Echo() {

	line := <-c.SendStr

	c.client.Write([]byte(line))

	buf := make([]byte, 1024)

	n, err := c.client.Read(buf)

	if err != nil {
		fmt.Printf("echo error,%s\n", err)
		c.RecvStr <- string("Server close...")

		c.client.Close()

		c.isAlive = false

		return

	}

	time.Sleep(1 * time.Second)

	c.RecvStr <- string(buf[0:n])

}

func Work(tc *Clienter) {

	if !tc.isAlive {

		if tc.Connect(host, port) {
			tc.Echo()

		} else {

			<-tc.SendStr

			tc.RecvStr <- string("Server close...")

		}

	} else {

		tc.Echo()

	}

}

func main() {
	if os.Args == nil || len(os.Args) < 3 {
		fmt.Println("usage:command ip port")
		return
	} else {
		host = os.Args[1]
		port = os.Args[2]
	}

	var tc Clienter
	tc.SendStr = make(chan string)
	tc.RecvStr = make(chan string)
	if !tc.Connect(host, port) {
		return
	}

	r := bufio.NewReader(os.Stdin)

	for {

		switch line, ok := r.ReadString('\n'); true {

		case ok != nil:

			fmt.Printf("bye bye!\n")

			return

		default:

			go Work(&tc)

			tc.SendStr <- line

			s := <-tc.RecvStr

			fmt.Printf("back:%s\n", s)

		}

	}
}
