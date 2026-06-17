package main

type ruleset struct {
	initiative initiativeEnum
}

type initiativeEnum string

const (
	initiativeFIX initiativeEnum = "fixed initiative"
	initiativeALT initiativeEnum = "alternating initiative"
)
