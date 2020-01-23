package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	arguments := os.Args[1:]
	n := 0
	if len(arguments) == 3 {
		FileName := arguments[1]
		switch arguments[1] {

		case "standard":
			FileName = "standart.txt"
		case "shadow":
			FileName = "shadow.txt"
		case "thinkertoy":
			FileName = "thinkertoy.txt"
		}
		if len(arguments[2]) >= 8 {
			if arguments[2][0:8] == "--align=" {
				if arguments[2][8:] == "left" {
					n = 1
				} else if arguments[2][8:] == "center" {
					n = 2
				} else if arguments[2][8:] == "right" {
					n = 3
				} else if arguments[2][8:] == "justify" {
					n = 4
				}
			} else {
				fmt.Println("Flag for alignment should be --align=")
				os.Exit(2)
			}
		} else {
			fmt.Println("Determine your alignment from: right, left, center, justify !!!")
			os.Exit(2)
		}
		if n == 0 {
			fmt.Println("Determine your alignment from: right, left, center, justify !!!")
			os.Exit(2)
		}
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
		var ArrayofText []string
		letters := ""
		for t, letter := range arguments[0] {

			if letter == '\n' {
				index++
				ArrayofText = append(ArrayofText, letters)
				letters = ""
				continue
			} else {
				letters = letters + string(letter)
				if t == len(arguments[0])-1 {
					ArrayofText = append(ArrayofText, letters)
				}
			}
			if letter > 31 && letter < 127 { // check for handling the character
				result[index] = append(result[index], ToLinesArray(symbols[letter]))
			} else {
				result[index] = append(result[index], ToLinesArray(symbols[32]))
			}
		}

		text := ""
		for _, bigLine := range result { //the whole text (picture)
			for i := 0; i < 8; i++ {
				for _, line := range bigLine {
					text = text + line[i]
				}
			}
		}
		terminalW, err := TerminalWidth()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		pic_len := len(text) / 8
		for ll, bigLine := range result { //printing result line by line
			for k := 0; k < 8; k++ {
				if n == 2 {
					for c := 0; c < (terminalW-pic_len)/2; c++ { //center
						fmt.Print(" ")
					}
				}
				if n == 3 {
					for r := 0; r < (terminalW - pic_len); r++ { //right
						fmt.Print(" ")
					}
				}
				for jj, line := range bigLine {
					fmt.Print(line[k])
					if n == 4 && ArrayofText[ll][jj] == ' ' {
						for i := 0; i < (terminalW-pic_len)/strings.Count(ArrayofText[ll], " "); i++ { // justify
							fmt.Print(" ")
						}
					}
				}
				fmt.Print("\n")
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
		fmt.Print(err.Error()) //if error? print error and close programm
		fmt.Println(" ", path)
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

func TerminalWidth() (int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdout
	out, _ := cmd.Output()
	return strconv.Atoi(string(out[3 : len(out)-1]))
}
