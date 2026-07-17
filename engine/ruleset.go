package engine

type Move struct {
	Bid Optional[int]
}
type MoveInfo struct {
	Bid Optional[bool]
}

type Ruleset struct {
	PhaseRules   []PhaseRules
	Phases       map[Phase]int
	InitialState GameState
}

type PhaseRules struct {
	Name       Optional[string]
	Disclose   func(GameState) ([]PlayerView, error)
	Compare    func([]Move) (Comparison, error)
	Resolve    func(GameState, Comparison) (GameState, error)
	Validation func(GameState, Move) error
}

type Role string

const (
	roleAttacker Role = "attacker"
	roleDefender Role = "defender"
)

type Phase string

const (
	phaseApproach Phase = "approach"
	phaseExchange Phase = "exchange"
	phaseEnd      Phase = "concluded game"
)

type PlayerView struct {
	GameStateInfo
}

/* Specific Ruleset Functions*/

/* */

/* Disclosures */

func DiscloseStandard(g GameState) ([]PlayerView, error) {
	var view PlayerView
	num_players := len(g.Players)
	TRUE := Optional[bool]{Value: true, OK: true}
	FALSE := Optional[bool]{Value: false, OK: true}

	view.Phase = true
	for i := range num_players {
		view.Name[i] = TRUE
		view.Balance[i] = TRUE
		view.Role[i] = TRUE
		view.Score[i] = TRUE
		view.Pending[i].Bid = FALSE
	}

	var result []PlayerView
	for len(result) < num_players {
		result = append(result, view)
	}

	for i, r := range result {
		for j := range r.Pending {
			if i == j {
				r.Pending[j].Bid = TRUE
			}
		}
	}

	return result, nil
}
