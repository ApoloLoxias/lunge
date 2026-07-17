package engine

import (
	"fmt"
	"strings"
)

type GameState struct {
	Players []Player
	Phase   Phase
	Pending []Move
}
type GameStateInfo struct {
	PlayerInfo
	Phase   bool
	Pending []MoveInfo
}

type Player struct {
	ID      int
	Name    Optional[string]
	Balance Optional[int]
	Role    Optional[Role]
	Score   Optional[int]
}
type PlayerInfo struct {
	ID      bool
	Name    []Optional[bool]
	Balance []Optional[bool]
	Role    []Optional[bool]
	Score   []Optional[bool]
}

/**/

func (g *GameState) NewPlayer(p Player) {
	np := p
	ID := len(g.Players)
	np.ID = ID
	g.Players = append(g.Players, np)
}

func (g GameState) String() string {
	var b strings.Builder

	b.WriteString("[")
	for i, p := range g.Players {
		if i != 0 {
			b.WriteString("; ")
		}
		b.WriteString(fmt.Sprintf("P%d: %s", i, p.String()))
	}
	b.WriteString("]")

	result := fmt.Sprintf(
		"GameState{Players: %v; Phase: %v; Rules: %v}",
		b.String(), g.Phase,
	)
	return result
}

func (p Player) String() string {
	var b strings.Builder
	b.WriteString("Player{")

	var strings []string

	if p.Name.OK {
		strings = append(strings, fmt.Sprintf("Name: %s", p.Name.Value))
	}
	if p.Balance.OK {
		strings = append(strings, fmt.Sprintf("Balance: %v", p.Balance.Value))
	}
	if p.Role.OK {
		strings = append(strings, (fmt.Sprintf("Role: %v", p.Role.Value)))
	}
	if p.Score.OK {
		strings = append(strings, (fmt.Sprintf("Score: %v", p.Score.Value)))
	}

	for i, s := range strings {
		if i != 0 {
			b.WriteString("; ")
		}
		b.WriteString(s)
	}

	b.WriteString("}")

	return b.String()
}
