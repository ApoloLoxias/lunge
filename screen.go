package main

import "fmt"
import "strings"
import "unicode/utf8"

type screen struct {
	lines   []string
	borders [2]string
	length  int
	height  int
}

func (s screen) draw() {
	fmt.Println()
	fmt.Println(strings.Repeat(s.borders[0], s.length+2))
	for _, line := range s.lines {
		filler := strings.Repeat(" ", s.length-utf8.RuneCountInString(line))
		fmt.Println(s.borders[1] + line + filler + s.borders[1])
	}
	filler := s.borders[1] + strings.Repeat(" ", s.length) + s.borders[1] + "\n"
	fmt.Print(strings.Repeat(filler, s.height))
	fmt.Println(strings.Repeat(s.borders[0], s.length+2))
}

//type line string
