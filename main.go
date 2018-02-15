package main

import (
	"path"
	"runtime"
	"os/exec"
	"os/user"
	"os"
	"io"
	"path/filepath"
	"log"
	"fmt"
	"strings"
	"errors"

)

func getCurrentExecPath() string {
	fpath, _ := exec.LookPath(os.Args[0])
	log.Print("filepath",fpath)
	dir := filepath.Dir(fpath)
	log.Print("dir",dir)
	return dir
}

func CopyFile(src string, dest string) (int64, error) {
	srcFile,err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return -1, errors.New("open src" + src + " fail")
	}
	defer srcFile.Close()

	destFile,err := os.Create(dest)
	if err != nil {
		fmt.Println(err.Error())
		return -1, errors.New("create file " + dest + " fail")		
	}
	defer destFile.Close()
	return io.Copy(destFile, srcFile)
}

func setGOROOTEnv(gopath string){
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

func main() {
	goos := runtime.GOOS
	currentUser,_ := user.Current()
	goroot := os.Getenv("GOROOT")
	thisFile := os.Args[0]
	_, thisFileName := filepath.Split(thisFile)
	log.Printf("os:%s\n", goos)
	log.Printf("home:%s\n", currentUser.HomeDir)
	log.Printf("GOROOT:%s\n", goroot)
	if goos == "windows" {
		// no parameter
		if len(os.Args) == 1 {
			
			
			fmt.Printf("will you copy %s to %s:", thisFile, goroot)
			var s string
			fmt.Scanf("%s", &s)
			if !strings.EqualFold("y", s) {
				return
			}
			CopyFile(thisFile, path.Join(goroot,"bin", thisFileName))
			return 
		}
		if len(os.Args) == 3 {
			if os.Args[1] == "get" {
				pwd,_ := os.Getwd()
				if goroot != pwd {
					setGOROOTEnv(pwd)
				}

				os.Setenv("GOPATH",pwd)
				cmd := exec.Command("go", "get", os.Args[2])
				stdout, err := cmd.StdoutPipe()
				if err != nil {
					log.Fatal(err)
				}
				defer stdout.Close()
				if err := cmd.Start(); err != nil {
					log.Fatal(err)
				}
			}
		}
		


	}
	
}