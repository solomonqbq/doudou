package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
	"time"
)

func loop(startport, endport int, inport chan int) {
	for i := startport; i <= endport; i++ {
		inport <- i
	}
}

func scanner(inport, outport, out chan int, ip net.IP, endport int) {
	for {
		in := <-inport
		//fmt.Println(in)
		tcpaddr := &net.TCPAddr{IP: ip, Port: in}
		conn, err := net.DialTCP("tcp", nil, tcpaddr)
		fmt.Println(err)
		if err != nil {
			outport <- 0
		} else {
			outport <- in
		}
		if conn != nil {
			conn.Close()
		}
		if in == endport {
			out <- in
		}
	}
}

func main() {
	starttime := time.Now().Unix()
	runtime.GOMAXPROCS(4)
	inport := make(chan int)
	outport := make(chan int)
	out := make(chan int)
	collect := []int{}
	if len(os.Args) != 4 {
		fmt.Println("Usage: scanner.exe IP startport endport")
		fmt.Println("Endport must be larger than startport")
		os.Exit(0)
	}
	ip := net.ParseIP(os.Args[1])
	if os.Args[3] < os.Args[2] {
		fmt.Println("Usage: scanner IP startport endport")
		fmt.Println("Endport must be larger than startport")
		os.Exit(0)
	}
	startport, _ := strconv.Atoi(os.Args[2])
	endport, _ := strconv.Atoi(os.Args[3])
	go loop(startport, endport, inport)
	for {
		select {
		case <-out:
			fmt.Println(collect)
			endtime := time.Now().Unix()
			fmt.Println("The scan process has spent ", endtime-starttime, "second")
			os.Exit(0)
		default:
			go scanner(inport, outport, out, ip, endport)

			port := <-outport

			if port != 0 {
				collect = append(collect, port)
			}
		}
	}
}

//该代码片段来自于: http://www.sharejs.com/codes/go/8939
