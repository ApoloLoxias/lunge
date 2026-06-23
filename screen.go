package main

import "fmt"

type screen struct {
	lines   []string
	borders [3]string
}

func (s screen) draw() {
	fmt.Println()
	fmt.Println(s.borders[0])
	for _, line := range s.lines {
		fmt.Println(s.borders[1] + line + s.borders[1])
	}
	fmt.Println(s.borders[2])
}

//type line string
