package main

import "errors"

type Fencer struct {
	Balance int
	RoW     bool
}

type hitEnum string

const (
	hitONE hitEnum = "p1 hits p2"
	hitTWO hitEnum = "p2 hits p1"
	hitFAL hitEnum = "no hit"
	hitDIS hitEnum = "disengage"
)

func Strike(o, d *Fencer, atk, def int) (hitEnum, error) {
	if [2]bool{o.RoW, d.RoW} != [2]bool{true, false} {
		return hitFAL, errors.New("invalid row pairing")
	}
	if atk > o.Balance || def > d.Balance {
		return hitFAL, errors.New("invalid balance spenditure")
	}

	if atk == 0 && def == 0 {
		return hitDIS, nil
	}
	if atk > def {
		return hitONE, nil
	}

	o.Balance -= atk
	d.Balance -= def

	if o.Balance == 0 && d.Balance == 0 {
		return hitDIS, nil
	}
	if o.Balance == 0 { //d.Balance != 0 is implied
		return hitTWO, nil
	}
	if d.Balance == 0 { //o.Balance !=0 is implied
		return hitONE, nil
	}
	return hitFAL, nil
}
