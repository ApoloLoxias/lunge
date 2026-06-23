//go:build scratch
// +build scratch

package main

func main() {
	screen := screen{
		lines:   []string{"line 1", "line 2", "", "line 4"},
		borders: [3]string{"------", "|", "------"},
	}
	screen.draw()

}
