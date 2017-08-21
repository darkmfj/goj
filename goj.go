package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	goPath := os.Getenv("GOPATH")
	packageName := os.Args[1]

	run, err := checkRun()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	targetPath := filepath.Join(goPath, "src", packageName)

	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		fmt.Printf("package does not exist, maybe try go get %s", packageName)
	} else {
		cmd := exec.Command(run, targetPath)
		err := cmd.Start()
		if err != nil {
			log.Fatal("Error:", err.Error())
		}
	}
}

// check run code-insiders or code
func checkRun() (string, error) {
	// code-insiders will open first
	_, err := exec.LookPath("code-insiders")
	if err == nil {
		return "code-insiders", nil
	}

	// code
	_, err = exec.LookPath("code")
	if err == nil {
		return "code", nil
	} else {
		var errorMsg = errors.New("make sure you have installed vscode")
		return "", errorMsg
	}
}
