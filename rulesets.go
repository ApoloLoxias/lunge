package main

type ruleset struct {
	initiative initiativeEnum
	approach   approachEnum
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
