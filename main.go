package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fileName, project := getArgs()
	var exists bool = fileExists(fileName)
	if !exists {
		result := fmt.Sprintf("El archivo %s no existe", fileName)
		fmt.Println(result)
		return
	}

	executeBat(fileName, project)
}

func setParams() (fileName string, project string) {
	fmt.Print("Indica la ubicación del archivo: ")
	fmt.Scanln(&fileName)
	fmt.Print("Indica el nombre del proyecto: ")
	fmt.Scanln(&project)

	return fileName, project
}

func getArgs() (fileName string, project string) {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 2 {
		fileName, project = setParams()
	} else {
		fileName = argsWithoutProg[0]
		project = argsWithoutProg[1]
	}
	return fileName, project
}

func executeBat(fileName string, project string) {
	cmd := exec.Command(fileName, project)
	// create a pipe for the output of the script
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		return
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("\t > %s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		return
	}
}

func fileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}
