// Package model babab
//go:generate stringer -type=HeroType -trimprefix=HeroType -linecomment
package model

type HeroType int

const (
	HeroTypeStrength     HeroType = iota + 1 // STRENGTH
	HeroTypeAgility                          // AGILITY
	HeroTypeIntelligence                     // INTELLIGENCE
)
