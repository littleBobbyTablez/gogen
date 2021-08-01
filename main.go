package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Print("Invalid input: 2 arguments needed")
		fmt.Print("gogen [tempolate] [typename]")
		os.Exit(1)
	}
	templateFileName := os.Args[1]
	typeName := os.Args[2]

	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		fmt.Printf("Error Reading file: %s: %s", templateFileName, err)
		os.Exit(1)
	}
	outName := fmt.Sprintf("gogen_%s_gen.go", typeName)
	fd, err := os.OpenFile(outName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Error creating destination file: %s", err)
		os.Exit(1)
	}
	defer func(fd *os.File) {
		err := fd.Close()
		if err != nil {
			os.Exit(0)
		}
	}(fd)

	data := struct {
		T string
	}{
		typeName,
	}
	err = t.Execute(fd, data)
	if err != nil {
		fmt.Printf("Error executing Template: %s", err)
	}
}
