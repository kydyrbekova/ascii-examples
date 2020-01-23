package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("writing to files...")
}

func Output(arguments []string) {
	//	arguments := os.Args[1:]
	fmt.Println(arguments[2])
	FileName := "standart.txt"
	if len(arguments) == 3 {
		switch arguments[1] {
		case "standart":
			FileName = "standart.txt"
		case "shadow":
			FileName = "shadow.txt"
		case "thinkertoy":
			FileName = "thinkertoy.txt"
		}

		fileNAme := arguments[2][9:]
		FileW, err := os.Create(fileNAme)
		if err != nil {
			fmt.Println(err.Error()) //if error? print error and close programm
			os.Exit(1)
		}

		defer FileW.Close()

		symbols := GetMapOfSymbols(FileName)
		lineIndex := 1
		arguments[0] = strings.Replace(arguments[0], "\\n", "\n", -1)
		for _, char := range arguments[0] {
			if char == '\n' {
				lineIndex++
			}
		}
		var result [][][]string //slice of slices to contain result string
		result = make([][][]string, lineIndex)
		index := 0
		for _, letter := range arguments[0] {
			if letter == '\n' { //
				index++
				continue
			}
			if letter > 31 && letter < 127 { // check for handling the character
				result[index] = append(result[index], ToLinesArray(symbols[letter]))
			} else {
				result[index] = append(result[index], ToLinesArray(symbols[32]))
			}
		}
		for _, bigLine := range result { //printing result line by line
			for i := 0; i < 8; i++ {
				for _, line := range bigLine {
					fmt.Fprint(FileW, line[i])
				}
				fmt.Fprintln(FileW)
			}
		}
		os.Exit(0)

	} else {
		fmt.Println("Number of arguments should be 3")
	}

}

func ToLinesArray(s string) []string { // funcion to convert string to string slice line by line
	var result []string
	temp := ""
	for _, letter := range s {
		if letter == '\n' {
			result = append(result, temp)
			temp = ""
			continue
		}
		temp += string(letter)
	}
	return result
}

func GetMapOfSymbols(path string) map[rune]string { // function, to create dictionary of strings by runes
	symbols := make(map[rune]string)   //dictionary to contain representation of runes
	body, err := ioutil.ReadFile(path) // reading standart.txt into []byte
	if err != nil {
		fmt.Println(err.Error()) //if error? print error and close programm
		os.Exit(1)
	}
	lineCounter := 0
	var keysCounter rune = ' '
	temp := "" //temp string for symbol containing

	for k, letter := range string(body) {
		if k == 0 {
			continue //the first /n dont need to be read into temp
		}
		if letter == '\n' {
			lineCounter++
		}
		if lineCounter > 8 {
			symbols[keysCounter] = temp
			keysCounter++
			lineCounter = 0
			temp = ""
			continue
		}
		temp += string(letter)
		if k == (len(body) - 1) { // if last char Add line? and add to map
			temp += "\n"
			symbols[keysCounter] = temp
			keysCounter++
			lineCounter = 0
			temp = ""
			continue
		}
	}
	return symbols
}
