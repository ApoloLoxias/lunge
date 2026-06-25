//go:build scratch
// +build scratch

package main

func main() {
	screen := screen{
		lines:   []string{"line 1", "line     2", "", "line 4   "},
		borders: [2]string{"-", "|"},
		height:  10,
		length:  20,
	}
	screen.draw()

}
