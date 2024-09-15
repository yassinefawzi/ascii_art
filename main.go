// "\n"

package main

import (
	"fmt"
	"os"
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
	ret = append(ret, 0)
	for i := 0; i < len(line); i++ {
		if i+1 < len(line) && line[i] == '\\' {
			if line[i+1] == 'n' {
				ret[j]++
			}
			i++
			if i+1 < len(line) && line[i+1] != '\\' {
				ret = append(ret, 0)
				j++
			}
		}
	}
	return ret
}

func print_art(file []string, splitted_line []string, lines_count []int) {
	holder := 0
	i := 0
	for ; i < len(splitted_line); i++ {
		for j := 0; j < 8; j++ {
			for k := 0; k < len(splitted_line[i]); k++ {
				holder = (int(splitted_line[i][k])-32)*9 + j
				fmt.Printf("%s", file[holder])
			}
			fmt.Println()
		}
		if len(lines_count) > 0 && i+1 < len(splitted_line) && lines_count[i] > 1 {
			fmt.Println()
		}
	}
	i--
	if i > 0 && i < len(lines_count) {
		for ; lines_count[i] > 0; lines_count[i]-- {
			fmt.Println()
		}
	}

}

func cleaned_split(s []string, lines_count []int) ([]string, []int) {
	var ret []string
	if len(s) == 2 && s[0] == "" {
		return nil, nil
	}
	i := 0
	if ; i < len(s) && s[i] == "" {
		if lines_count[0] > 1 {
			fmt.Println()
		}
		if len(lines_count) > 1 {
		lines_count = lines_count[1:]
		}
	}
	for ; i < len(s); i++ {
		if s[i] != "" {
			ret = append(ret, s[i])
		}
	}
	return ret, lines_count
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
	splitted_line := strings.Split(line, "\\n")
	lines_count := count_next_line(line)
	splitted_line, lines_count = cleaned_split(splitted_line, lines_count)
	if splitted_line == nil {
		fmt.Println()
		return
	}
	print_art(file[1:], splitted_line, lines_count)
}
