// "\n"

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read_file() []string {
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Ascii file not found")
		return nil
	}
	ret := strings.Split(string(file), "\n")
	return ret
}

func count_next_line(line string) []int {
	var ret []int
	j := 0
	ret = append(ret, -1)
	for i := 0; i < len(line); i++ {
		if i+2 < len(line) && line[i] == '\\' {
			if line[i+1] == 'n' {
				if len(ret) == 0 {
					ret[j]++
				}
				i += 2
			}
			if line[i] != '\\' {
				ret = append(ret, 0)
			}
		}
	}
	return ret
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Argument Error")
		return
	}
	file := read_file()
	line := os.Args[1]
	if len(line) < 1 {
		return
	}
	splitted_line := strings.Split(line, "\n")
	lines_count := count_next_line(line)
	print(file, splitted_line, lines_count)
}