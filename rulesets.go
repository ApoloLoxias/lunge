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

/********************************
## RUSET SPECIFC STATE ELEMENTS #
*********************************/

func (g gameState) scoreValue() scoreStateElement {
	return g.facultativeState[score].(scoreStateElement)
}

func (g gameState) numGoodsToBeAuctionedValue() numGoodsToBeAuctionedStateElement {
	return g.facultativeState[numGoodsToBeAuctioned].(numGoodsToBeAuctionedStateElement)
}

type scoreStateElement map[fencer]int

func (s scoreStateElement) facultativeStateElementFunc() error {
	return errors.New("Uncallable func called @./rulesets vfrtY79iKm")
}

type numGoodsToBeAuctionedStateElement int

func (n numGoodsToBeAuctionedStateElement) facultativeStateElementFunc() error {
	return errors.New("Uncallable func called @./rulesets vfRtY79iKm")
}
