package main

import (
	"fmt"
	"os"

	"golang.design/x/clipboard"
)

func getdir(){
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Print("could not copy directory to clipboard")
	}

	// Create OS copy command
	err = clipboard.Init()
	if err != nil {
		fmt.Print("could not access sytem clipboard")
	}
	clipboard.Write(clipboard.FmtText, []byte(cwd))

	return

}

func main(){
	getdir()
}
