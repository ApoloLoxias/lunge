//go:build scratch
// +build scratch

package main

func main() {
	screen := screen{
		header:   "+++++++ 󰞇 LUNGE! 󰞇 ++++++",
		players:  []string{"󰓥 P1  ●●●●        4 Balance", "󰦝 P2  ●●●●●       5 Balance"},
		position: []string{" Initiative: P1", " Range: 󰚌 At Measure 󰚌"},
		height:   10,
		length:   40,
	}
	screen.draw()

}
