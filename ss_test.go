package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowestManager(t *testing.T) {
	// n = 2
	assert.Equal(t, "Fred", lowestManager(2, "Fred", "George", []string{"Fred George"}))
	assert.Equal(t, "George", lowestManager(2, "Fred", "George", []string{"George Fred"}))

	// n = 3
	//	Peter
	//		Hilary
	//		John

	assert.Equal(t,
		"Peter",
		lowestManager(3,
			"Hilary", "John",
			[]string{
				"Peter Hilary",
				"Peter John"}))

	// n = 4
	//	Sarah
	//		Claudiu
	//			Simon
	//		Paul

	assert.Equal(t,
		"Claudiu",
		lowestManager(4,
			"Simon", "Claudiu",
			[]string{
				"Sarah Claudiu",
				"Sarah Paul",
				"Claudiu Simon"}))

	// n = 5
	//	June
	//		Alex
	//		Qing
	//			Paul
	//			Gareth

	assert.Equal(t,
		"June",
		lowestManager(5,
			"Gareth", "Alex",
			[]string{
				"June Alex",
				"June Qing",
				"Qing Paul",
				"Qing Gareth"}))

	// n = 6
	//	Sarah
	//		Fred
	//			Hilary
	//			Jenny
	//				James
	//		Paul
	assert.Equal(t,
		"Fred",
		lowestManager(6,
			"Hilary", "James",
			[]string{
				"Sarah Fred",
				"Sarah Paul",
				"Fred Hilary",
				"Fred Jenny",
				"Jenny James"}))

	// ===================================
	// My examples
	// ===================================

	// Simple
	assert.Equal(t,
		"Deneris",
		lowestManager(
			3,
			"Tirion", "Mormon",
			[]string{
				"Deneris Tirion",
				"Deneris Mormon"}))

	// n = 4
	assert.Equal(t,
		"Dani",
		lowestManager(4,
			"Stefan", "Koceto",
			[]string{
				"Max Dani",
				"Max Kiril",
				"Dani Stefan",
				"Dani Koceto"}))

	// Indirect manager
	assert.Equal(t,
		"Didi",
		lowestManager(4,
			"Stefan", "Koceto",
			[]string{
				"Max Didi",
				"Didi Koceto",
				"Didi Dani",
				"Dani Stefan"}))

	assert.Equal(t,
		"Stefan",
		lowestManager(3,
			"Stefan", "Koce",
			[]string{
				"Max Stefan",
				"Max Koce",
				"Stefan Koce"}))

	assert.Equal(t,
		"Stefan",
		lowestManager(5,
			"Stefan", "Koce",
			[]string{
				"Max Stefan",
				"Max Koce",
				"Stefan Mimi",
				"Mimi Koce"}))

	assert.Equal(t,
		"Stefan",
		lowestManager(5,
			"Stefan", "Koce",
			[]string{
				"Max Stefan",
				"Max Mimi",
				"Stefan Mimi",
				"Mimi Koce"}))

	assert.Equal(t,
		"Queen",
		lowestManager(4,
			"Pawn1", "Pawn2",
			[]string{
				"Queen Bishop1",
				"Queen Bishop2",
				"Bishop1 Pawn1",
				"Bishop1 Pawn2",
				"Bishop2 Pawn1",
				"Bishop2 Pawn2",
			}))

	assert.Equal(t,
		"Rook",
		lowestManager(5,
			"Pawn1", "Pawn2",
			[]string{
				"Queen Bishop1",
				"Queen Bishop2",
				"Bishop1 Rook",
				"Rook Pawn1",
				"Rook Pawn2",
				"Bishop2 Pawn1",
				"Bishop2 Pawn2",
			}))

	assert.Equal(t,
		"m",
		lowestManager(9,
			"a", "b",
			[]string{
				"r x",
				"x l",
				"x z",
				"r y",
				"y z",
				"l p",
				"l m",
				"z m",
				"z q",
				"m a",
				"m q",
				"q b",
			}))
}
