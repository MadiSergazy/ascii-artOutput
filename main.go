package main

import (
	"Mado/Checker"
	"Mado/FileReadWork"
	"Mado/PrintArt"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	massivString, errF := FileReadWork.FileReadWork("standard.txt") // reading file
	if errF != nil {
		fmt.Println(errF)
		os.Exit(0)
	}
	args := os.Args[1:]

	fileName := args[1]

	switch fileName {
	case "standard":
		fileName = fileName + ".txt"
	case "shadow":
		fileName = fileName + ".txt"
	case "thinkertoy":
		fileName = fileName + ".txt"
	default:
		fmt.Printf("Incorect args %v\n ddddddddd", fileName)
		return
	}

	if len(args) == 1 {
		// fmt.Println("EMPRY ARGS")
		os.Exit(0)
	}
	if len(args) != 3 {
		fmt.Println("Incorrect input")
		os.Exit(0)
	}

	err := Checker.Check(args[1]) // check correct args
	if err != nil {
		return
	}

	Checker.FileCheker(fileName)

	text := PrintArt.PrintArt(massivString, args[0]) // print and check err

	FileNameToWrite := strings.Split(args[2], "=")
	if FileNameToWrite[0] != "--output" {
		fmt.Println("Incorrect args in writeFile")
		return
	}

	fileExtensionWrite := filepath.Ext(FileNameToWrite[1])
	if fileExtensionWrite != ".txt" { // проверка формата файла
		fmt.Println("INCORRECT FILEANME for write")
		return
	}

	errNewLine, count := Checker.ChekcOnlyN(args[0]) // if we have only \n
	if !errNewLine {
		str := ""
		for i := 0; i < count; i++ {
			str += "\n"
		}
		text = str
	}

	f, err := os.Create(FileNameToWrite[1])
	if err != nil {
		fmt.Println("ERROR WHEN CREATE FILE")
		return

	}

	defer f.Close()

	_, err2 := f.WriteString(text)

	if err2 != nil {
		fmt.Println("ERROR WHEN write FILE")
		return

	}

}
