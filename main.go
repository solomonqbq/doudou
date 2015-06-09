package main

import (
	"fmt"
	"os/exec"
	//	"github.com/dmotylev/goproperties"
	//	"jcloud/doudou/db"
	//	"os"
	//	"path/filepath"
)

func main() {
	argv := []string{"a"}
	c := exec.Command("ip", argv...)
	d, _ := c.Output()
	fmt.Println(string(d)) //因为装的git bash所以可以用ls -a
	/*
	 *	.
	 *	..
	 *	command.go
	 *	lookpath.go
	 */
}
