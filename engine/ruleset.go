package engine

type Move struct {
	Actor int
	Bid   Optional[int]
}
type MoveInfo struct {
	Actor bool
	Bid   Optional[bool]
}

type Ruleset struct {
	PhaseRules   []PhaseRules
	Phases       map[Phase]int
	InitialState GameState
}

type PhaseRules struct {
	Name       string
	Disclose   func(GameState) ([]GameStateInfo, error)
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

type Comparison struct {
	Delta Optional[[][]int]
}

/* Specific Ruleset Functions*/

/* */

/* Disclosures */

func DiscloseStandard(g GameState) ([]GameStateInfo, error) {
	var info GameStateInfo
	num_players := len(g.Players)
	TRUE := Optional[bool]{Value: true, OK: true}
	FALSE := Optional[bool]{Value: false, OK: true}

	info.Phase = true
	info.Name = make([]Optional[bool], num_players)
	for i := range num_players {
		info.Name[i] = TRUE
		info.Balance[i] = TRUE
		info.Role[i] = TRUE
		info.Score[i] = TRUE
		info.Pending[i].Bid = FALSE
		info.Pending[i].Actor = true
	}

	var result []GameStateInfo
	for len(result) < num_players {
		result = append(result, info)
	}

	for i, r := range result {
		for j := range r.Pending {
			if i == j {
				result[i].Pending[j].Bid = TRUE
			}
		}
	}

	return result, nil
}

/* Comparisons */

func CompareSimpleBid(g GameState) Comparison {
	SortedMoves := SortMoves(g.Pending)
	result := make([][]int, len(SortedMoves))
	for i := range result {
		result[i] = make([]int, len(SortedMoves))
	}

	for i, move := range SortedMoves {
		for j := range len(result) {
			result[i][j] = move.Bid.Value
		}
	}

	for i := range result {
		for j := range result {
			result[i][j] = result[i][j] - result[j][j]
		}
	}

	return Comparison{Delta: Optional[[][]int]{OK: true, Value: result}}
}

func SortMoves(mm []Move) []Move {
	var j int
	for i, m := range mm {
		j = i
		for j > 0 && m.Actor < mm[j].Actor {
			mm[j+1] = mm[j]
			j = j - 1
		}
		mm[j+1] = m
	}
	return mm
}
