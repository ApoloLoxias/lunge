package main

import "fmt"
import "strings"

/*
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
*/

type screen struct {
	header   string
	players  []string
	position []string
	prompt   string
	height   int
	length   int
}

func (s screen) draw() {
	fmt.Println(s.topString())
	fmt.Println(s.headerString())
	fmt.Println(s.middleString())
	fmt.Println(s.playersString())
	fmt.Println()
	fmt.Println(s.positionString())
	fmt.Println()
	fmt.Println(s.bottomString())
}

func (s screen) topString() string {
	var b strings.Builder
	b.WriteString("╭")
	i := 0
	for i < s.length {
		b.WriteString("─")
		i++
	}
	b.WriteString("╮")
	return b.String()
}

func (s screen) bottomString() string {
	var b strings.Builder
	b.WriteString("╰")
	i := 0
	for i < s.length {
		b.WriteString("─")
		i++
	}
	b.WriteString("╯")
	return b.String()
}

func (s screen) middleString() string {
	var b strings.Builder
	b.WriteString("├")
	i := 0
	for i < s.length {
		b.WriteString("─")
		i++
	}
	b.WriteString("┤")

	return b.String()
}

func (s screen) headerString() string {
	var b strings.Builder
	b.WriteString("|")
	i := 0
	for i < (s.length/2 - len([]rune(s.header))/2 - 1) {
		b.WriteString(" ")
		i++
	}
	b.WriteString(s.header)
	i = 0
	for i < (s.length/2 - len([]rune(s.header))/2 - 1) {
		b.WriteString(" ")
		i++
	}
	if len([]rune(s.header))%2 == 1 {
		b.WriteString(" ")
	}
	b.WriteString("|")

	return b.String()
}

func (s screen) playersString() string {
	var b strings.Builder

	for i := 0; i < len(s.players); i++ {
		b.WriteString("| ")
		b.WriteString(s.players[i])
		j := 0
		for j < (s.length - len([]rune(s.players[i])) - 1) {
			b.WriteString(" ")
			j++
		}
		b.WriteString("|\n")
	}

	return b.String()
}

func (s screen) positionString() string {
	var b strings.Builder

	for _, position := range s.position {
		b.WriteString("|")
		b.WriteString(position)
		for j := 0; j < (s.length - len(position)); j++ {
			b.WriteString(" ")
		}
		b.WriteString("|\n")
	}

	return b.String()
}
