package main

import "errors"

type gameState struct {
	p1 fencer
	p2 fencer

	kind gameStateTypeEnum

	rules            ruleset
	facultativeState map[facultativeStateElementEnum]facultativeStateElement

	parent *gameState
}

var gameStateERR = gameState{ //would be const if golang allowed fconst structs
	p1:     fencerERR,
	p2:     fencerERR,
	kind:   stateERR,
	parent: nil,
}

type facultativeStateElementEnum string

const score facultativeStateElementEnum = "score"

type facultativeStateElement interface{ facultativeStateElementFunc() error }

type gameStateTypeEnum string

const (
	stateERR      gameStateTypeEnum = "zero-value gameStateType due to error retur"
	stateEXCHANGE gameStateTypeEnum = "exchanging blows"
	stateOOM      gameStateTypeEnum = "out of measure"
	stateWIN1     gameStateTypeEnum = "p1 is victorious"
	stateWIN2     gameStateTypeEnum = "p2 is victorious"
)

type move struct {
	kind moveTypeEnum
	bid1 int
	bid2 int
}

type moveTypeEnum string

const (
	moveTypeERR      moveTypeEnum = "zero-value moveType due to error return"
	moveTypeONEATK   moveTypeEnum = "p1 attacks p2"
	moveTypeTWOATK   moveTypeEnum = "p2 attacks p1"
	moveTypeAPPROACH moveTypeEnum = "players approach one another"
)

type fencer struct {
	balance int
	role    roleEnum
}

var fencerERR = fencer{
	balance: 0,
	role:    roleERR,
}

type roleEnum string

const (
	roleERR roleEnum = "zero-value role"
	roleATK roleEnum = "attacker"
	roleDEF roleEnum = "defender"
)

type hitEnum string

const (
	hitERR hitEnum = "zero-value hit due to error return"
	hitONE hitEnum = "p1 hit p2"
	hitTWO hitEnum = "p2 hit p1"
	hitMIS hitEnum = "no hit"
	hitDIS hitEnum = "disengage"
)

type lungeEnum string

const (
	lungeERR  lungeEnum = "zero-value lunge"
	lungeONE  lungeEnum = "Player 1 lungers forward and seizes the initiative"
	lungeTWO  lungeEnum = "Player 2 lunges forward and seizes the initiative"
	lungeBOTH lungeEnum = "Both players try to lunge together"
	lungeNO   lungeEnum = "Players wait for an opportunity to lunge"
)

func getNextState(oldStPointer *gameState, mv move) (newSt gameState, err error) {
	switch oldStPointer.kind {
	case stateEXCHANGE:
		return getNextStateExchange(oldStPointer, mv)
	case stateOOM:
		return getNextStateOOM(oldStPointer, mv)
	default:
		return gameStateERR, errors.New("unimplemented state or broken core game logic @./engine.go fd774fae")
	}
}

func getNextStateOOM(oldStPointer *gameState, mv move) (newSt gameState, err error) {
	oldSt := *oldStPointer

	legalityErr := checkMoveLegalityOOM(oldSt, mv)
	if legalityErr != nil {
		return oldSt, legalityErr
	}

	switch oldSt.rules.approach {
	case noAproach:
		newSt = gameStateERR
		err = errors.New("Trying to play OOM phase on a ruleset with no OOM phase")
		return newSt, err
	case initiativeBidding:
		return getNextStateOOMInitiativeBidding(oldStPointer, mv)
	}

	return gameStateERR, errors.New("Broken core game logic @./engine.go d51932ab")
}

func getNextStateOOMInitiativeBidding(oldStPointer *gameState, mv move) (newSt gameState, err error) {
	oldSt := *oldStPointer

	lunge, lungeError := func(oldSt gameState, mv move) (lungeEnum, error) {
		b1 := mv.bid1
		b2 := mv.bid2

		if b1 == 0 && b2 == 0 {
			return lungeNO, nil
		}
		if b1 > b2 {
			return lungeONE, nil
		}
		if b2 > b1 {
			return lungeTWO, nil
		}
		if b1 == b2 { // implied b1, b2 != 0
			return lungeBOTH, nil
		}

		return lungeERR, errors.New("Broken core game logic @./engine hJKg68a6")
	}(oldSt, mv)
	if lungeError != nil {
		return gameStateERR, lungeError
	}

	newSt = oldSt
	newSt.parent = oldStPointer

	switch lunge {
	case lungeONE:
		newSt.p1.role = roleATK
		newSt.p2.role = roleDEF
		newSt.kind = stateEXCHANGE
		newSt.p1.balance -= mv.bid1
		newSt.p2.balance -= mv.bid2
	case lungeTWO:
		newSt.p1.role = roleDEF
		newSt.p2.role = roleATK
		newSt.kind = stateEXCHANGE
		newSt.p1.balance -= mv.bid1
		newSt.p2.balance -= mv.bid2
	case lungeBOTH, lungeNO:
		return
	case lungeERR:
		return gameStateERR, errors.New("Invalid lunge calculated")
	default:
		return gameStateERR, errors.New("Broken core game logic @./engine.go uIy78t")
	}

	return newSt, nil
}

func getNextStateExchange(oldStPointer *gameState, mv move) (newSt gameState, err error) {
	oldSt := *oldStPointer

	legalityErr := checkMoveLegalityExchange(oldSt, mv)
	if legalityErr != nil {
		return oldSt, legalityErr
	}

	hit, hitError := strike(oldSt, mv)
	if hitError != nil {
		return oldSt, hitError
	}

	var newStType gameStateTypeEnum
	switch hit {
	case hitONE:
		newStType = stateWIN1
	case hitTWO:
		newStType = stateWIN2
	case hitMIS:
		newStType = stateEXCHANGE
	case hitDIS:
		newStType = stateOOM
	default:
		newStType = stateERR
	}
	if newStType == stateERR {
		newSt = gameStateERR
		newSt.parent = oldStPointer
		return newSt, errors.New("broken core game logic @./engine.go 3c6a8ee1")
	}

	newP1, newP2, balanceERR := updatePlayersExchange(oldSt, mv)
	if balanceERR != nil {
		newSt = gameStateERR
		newSt.parent = oldStPointer
		return newSt, errors.New("broken core game loigic @./engine.go d4fcc62c")
	}

	newState := gameState{
		p1:     newP1,
		p2:     newP2,
		kind:   newStType,
		rules:  oldSt.rules,
		parent: oldStPointer,
	}

	return newState, nil
}

func updatePlayersExchange(st gameState, mv move) (fencer, fencer, error) {
	p1 := st.p1
	p2 := st.p2
	b1 := p1.balance - mv.bid1
	b2 := p2.balance - mv.bid2

	if b1 < 0 || b2 < 0 {
		return fencerERR, fencerERR, errors.New("can't update a player into negative balance")
	}

	p1.balance = b1
	p2.balance = b2

	if st.rules.initiative == initiativeALT {
		p1.role, p2.role = p2.role, p1.role
	}

	return p1, p2, nil
}

func checkMoveLegalityExchange(st gameState, mv move) error {
	if st.p1.balance == 0 || st.p2.balance == 0 {
		return errors.New("An out of balance player is trying to make a move")
	}

	if mv.bid1 > st.p1.balance || mv.bid2 > st.p2.balance {
		return errors.New("Bids exceed balance")
	}

	if st.p1.role == st.p2.role {
		return errors.New("Invalid role combination")
	}

	if mv.kind == moveTypeONEATK && st.p1.role != roleATK {
		return errors.New("invalid move type for given roles")
	}

	if mv.kind == moveTypeTWOATK && st.p2.role != roleATK {
		return errors.New("invalid move type for given roles")
	}

	if mv.kind == moveTypeERR {
		return errors.New("malformed move type")
	}

	if st.kind == stateERR {
		return errors.New("malformed gameStateType")
	}

	if st.kind == stateWIN1 || st.kind == stateWIN2 {
		return errors.New("game ended")
	}

	if st.kind == stateOOM {
		return errors.New("unimplemented gameStateType: out-of-measure")
	}

	return nil
}

func checkMoveLegalityOOM(st gameState, mv move) error {
	if st.p1.balance == 0 || st.p2.balance == 0 {
		return errors.New("An out of balance player is trying to make a move")
	}

	if mv.bid1 > st.p1.balance || mv.bid2 > st.p2.balance {
		return errors.New("Bids exceed balance")
	}

	if mv.kind != moveTypeAPPROACH {
		return errors.New("invalid stateType-moveKind pair")
	}

	if st.kind != stateOOM {
		return errors.New("wtf jk@s98Y")
	}

	return nil
}

func strike(st gameState, mv move) (hit hitEnum, err error) {
	switch st.rules.exchangeType {
	case exchangeFencing:
		return strikeClassic(st, mv)
	case exchangingAuction:
		return strikeAuction(st, mv)
	default:
		return hitERR, errors.New("Broken logic @ ./engine.go nsA19Tugy")
	}
}

func strikeClassic(st gameState, mv move) (hit hitEnum, err error) {
	if mv.bid1 > st.p1.balance || mv.bid2 > st.p2.balance {
		return hitERR, errors.New("illegal strike")
	}
	if st.p1.balance == 0 || st.p2.balance == 0 {
		return hitERR, errors.New("illegal strike")
	}

	if mv.kind == moveTypeONEATK {
		hit, err = resolveStrike(st.p1, st.p2, mv.bid1, mv.bid2)
		if err != nil {
			return hitERR, err
		}
	}
	if mv.kind == moveTypeTWOATK {
		hit, err = resolveStrike(st.p2, st.p1, mv.bid2, mv.bid1)
		if err != nil {
			return hitERR, err
		}
		if hit == hitONE {
			hit = hitTWO
		} else if hit == hitTWO {
			hit = hitONE
		}
	}

	return hit, nil
}

func resolveStrike(a, d fencer, atk, def int) (hitEnum, error) {
	if atk == 0 && def == 0 {
		return hitDIS, nil
	}
	if atk > def {
		return hitONE, nil
	}
	if (a.balance-atk) == 0 && (d.balance-def) == 0 {
		return hitDIS, nil
	}
	if (a.balance - atk) == 0 { //defender implied to remain balanced
		return hitTWO, nil
	}
	if (d.balance - def) == 0 { //attacker implied to remain blaanced
		return hitONE, nil
	}
	if def >= atk && a.balance > atk && d.balance > def {
		return hitMIS, nil
	}

	return hitERR, errors.New("Core game logic broken @./engine.go 6e5a77f1")
}
