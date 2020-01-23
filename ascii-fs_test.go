package main

import (
	"io/ioutil"
	"testing"
	student ".."
)

func TestReadFile(t *testing.T) {
	var tests = []struct {
		input []string
		output string
		expected string
	} {
		{"First\nTest" shadow --output=test00.txt, test00.txt, exp00.txt },
		{"hello" standard --output=test01.txt, test01.txt, exp01.txt},
		{"123 -> #$%" standard --output=test02.txt, test02.txt, exp02.txt},
		{"432 -> #$%&@" shadow --output=test03.txt, test03.txt, exp03.txt},
		{"There" shadow --output=test04.txt, test04.txt, exp04.txt},
		{"123 -> \"#$%@" thinkertoy --output=test05.txt, test05.txt, exp05.txt},
		{"2 you" thinkertoy --output=test06.txt, test06.txt, exp06.txt},
		{"Testing long output!" standard --output=test07.txt, test07.txt, exp07.txt},
	}

		for _, test:= range tests {
			Output(test.input)
			dataoutput, err := ioutill.ReadFile(test.output)
			if err!=nil {
				t.Fatal("Could not open output file")
			}
			dataexp, errexp := ioutill.ReadFile(test.expected)
			if errexp!=nil {
				t.Fatal("Could not open expected file")
			}
			if string(dataoutput) != dataexp {
				t.Fatal("Output do not match expected")
			}
		}


}
