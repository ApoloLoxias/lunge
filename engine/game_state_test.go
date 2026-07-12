package engine

import "testing"
import "fmt"

func TestPlayerString(t *testing.T) {
	testCases := []struct {
		name          string
		p             Player
		expectedValue string
	}{
		{
			name:          "empty",
			p:             Player{},
			expectedValue: "Player{}",
		},
		{
			name: "uninitialized values",
			p: Player{
				Name:    Optional[string]{Value: "Heracles"},
				Balance: Optional[int]{Value: 7},
				Role:    Optional[Role]{Value: roleAttacker},
				Score:   Optional[int]{Value: 12},
			},
			expectedValue: "Player{}",
		},
		{
			name:          "name only",
			p:             Player{Name: Optional[string]{Value: "Sample Name", OK: true}},
			expectedValue: "Player{Name: Sample Name}",
		},
		{
			name:          "Balance only",
			p:             Player{Balance: Optional[int]{Value: 3, OK: true}},
			expectedValue: "Player{Balance: 3}",
		},
		{
			name:          "Role only",
			p:             Player{Role: Optional[Role]{Value: roleAttacker, OK: true}},
			expectedValue: "Player{Role: attacker}",
		},
		{
			name:          "Score only",
			p:             Player{Score: Optional[int]{Value: 0, OK: true}},
			expectedValue: "Player{Score: 0}",
		},
		{
			name: "All fields",
			p: Player{
				Name:    Optional[string]{Value: "Fiori", OK: true},
				Balance: Optional[int]{Value: 99, OK: true},
				Role:    Optional[Role]{Value: roleDefender, OK: true},
				Score:   Optional[int]{Value: 1, OK: true},
			},
			expectedValue: "Player{Name: Fiori; Balance: 99; Role: defender; Score: 1}",
		},
	}

	for _, testCase := range testCases {
		got := testCase.p.String()
		if got != testCase.expectedValue {
			t.Errorf(
				"Wrong Player string. Expected: %s; Actual: %s",
				testCase.expectedValue, got,
			)
		}
	}
}

func TestGameStateString(t *testing.T) {
	players := []Player{
		{},
		{
			Name:    Optional[string]{Value: "Ephialtes", OK: false},
			Balance: Optional[int]{Value: 3, OK: true},
		},
	}

	gameState := GameState{
		Players: players,
		Phase:   phaseApproach,
		Rules:   Ruleset{},
	}

	want := fmt.Sprintf(
		"GameState{Players: [P0: %s; P1: %s]; Phase: %s; Rules: %s}",
		players[0].String(), players[1].String(),
		"approach",
		"ruleset",
	)

	if want != gameState.String() {
		t.Errorf(
			"Wrong GameState string. Expected '%s'; got '%s'",
			want, gameState.String(),
		)
	}
}
