// Go 1.15.6

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var cells = [30000]int{0}
var ptr = 0

func main() {
	filepath := os.Args

	if len(filepath) < 2 {
		log.Fatal("main.go <.bffilepath>")

	}

	file, err := ioutil.ReadFile(filepath[1])
	if err != nil {
		log.Fatal(err)
	}

	execute(string(file))
}

func execute(c string) {

	for i := 0; i < len(c); i++ {
		switch string(c[i]) {
		case ">":
			ptr++
		case "<":
			ptr--
		case "+":
			cells[ptr]++
		case "-":
			cells[ptr]--
		case ".":
			fmt.Print(string(rune(cells[ptr])))
			// fmt.Println(cells[ptr], "ptr:", ptr)
		case ",":
			var i int
			fmt.Scanf("%d", &i)
			cells[ptr] = i
		case "[":
			for cells[ptr] != 0 {
				closeindex, err := getCloseIndex(i, c)
				if err != nil {
					log.Fatal(err)
				}
				execute(c[i+1 : closeindex])
			}
			if cells[ptr] == 0 {
				closeindex, err := getCloseIndex(i, c)
				if err != nil {
					log.Fatal(err)
				}
				i = closeindex
				break
			}

		}
	}
}

func getCloseIndex(open int, c string) (int, error) {
	// TODO:handle nested loops
	for i := open; i < len(c); i++ {
		if string(c[i]) == "]" {
			return i, nil
		}

	}
	return 0, fmt.Errorf("invalid Parentheses")
}
