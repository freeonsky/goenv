package main

import (
	"runtime"
	"os/exec"
	"os/user"
	"os"
	"log"
	"fmt"
)

func main() {
	goos := runtime.GOOS
	currentUser,_ := user.Current()
	log.Printf("os:%s\n", goos)
	log.Printf("home:%s\n", currentUser.HomeDir)
	
	if goos == "windows" {
		gopath := currentUser.HomeDir
		if len(os.Args) == 2 {
			gopath = os.Args[1]
		}
		log.Println("gopath:", gopath)

		cmd := exec.Command("setx", "GOPATH", gopath)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		defer stdout.Close()
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("set GOPATH=", gopath)
		}

	}
	
}