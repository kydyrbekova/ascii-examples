package main

import (
	"io/ioutil"
	"testing"
)

func TestReadFile(t *testing.T) {
	var tests = []struct {
		input    []string
		output   string
		expected string
	}{
		{[]string{"First\nTest", "shadow", "--output=test00.txt"},
			"test00.txt",
			"exp00.txt",
		},
		{[]string{"hello", "standard", "--output=test01.txt"},
			"test01.txt",
			"exp01.txt",
		},
		{[]string{"123 -> #$%", "standard", "--output=test02.txt"},
			"test02.txt",
			"exp02.txt",
		},
		{[]string{"432 -> #$%&@", "shadow", "--output=test03.txt"},
			"test03.txt",
			"exp03.txt",
		},
		{
			[]string{"There", "shadow", "--output=test04.txt"},
			"test04.txt",
			"exp04.txt",
		},
		{[]string{"123 -> \"#$%@", "thinkertoy", "--output=test05.txt"},
			"test05.txt",
			"exp05.txt",
		},
		{[]string{"2 you", "thinkertoy", "--output=test06.txt"},
			"test06.txt",
			"exp06.txt",
		},
		{[]string{"Testing long output!", "standard", "--output=test07.txt"},
			"test07.txt",
			"exp07.txt",
		},
	}

	for _, test := range tests {
		Output(test.input)

		dataoutput, err := ioutil.ReadFile(test.output)
		if err != nil {
			t.Fatal("Could not open output file")
		}
		dataexp, errexp := ioutil.ReadFile(test.expected)
		if errexp != nil {
			t.Fatal("Could not open expected file")
		}
		if string(dataoutput) != string(dataexp) {
			t.Fatal("Output do not match expected")
		}
	}

}
