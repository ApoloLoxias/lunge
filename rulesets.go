package main

import "errors"

type ruleset struct {
	initiative   initiativeEnum
	approach     approachEnum
	exchangeType exchangeTypeEnum
}

type initiativeEnum string

const (
	initiativeFIX initiativeEnum = "fixed initiative"
	initiativeALT initiativeEnum = "alternating initiative"
)

type approachEnum string

const (
	initiativeBidding approachEnum = "Bid for initiative on approach"
	noAproach         approachEnum = "Go directly to exchange phase"
)

type exchangeTypeEnum string

const (
	exchangeFencing   exchangeTypeEnum = "fencing exchange phase"
	exchangingAuction exchangeTypeEnum = "auction exchange phase"
)

/***********************************/
/* RULESET-SPECIFIC STATE ELEMENTS */
/***********************************/

func (s *gameState) scorePointer() *scoreStateElement {
	return s.facultativeState[score].(*scoreStateElement)
}

type scoreStateElement struct {
	p1 uint
	p2 uint
}

func (s scoreStateElement) facultativeStateElementFunc() error {
	return errors.New("Uncallable function called @./rulesets.go oIas6YUh")
}
