package main

import (
	"math/rand"

	gotea "github.com/jpincas/go-tea"
)

// Types

type Deck []Card

// Message generators

func FlipAllBack(_ gotea.MessageArguments) gotea.Message {
	return gotea.NewMsg(flipAllBack, nil)
}

func RemoveMatches(_ gotea.MessageArguments) gotea.Message {
	return gotea.NewMsg(removeMatches, nil)
}

// Messages

func flipAllBack(_ gotea.MessageArguments, s *gotea.Session) {
	s.State.(Model).Deck.flipAllBack()
}

func removeMatches(_ gotea.MessageArguments, s *gotea.Session) {
	s.State.(Model).Deck.removeMatches()
}

// Actions

func newDeck(n int) (deck Deck) {
	cardValues := append([]int{}, append(rand.Perm(n), rand.Perm(n)...)...)
	for _, v := range cardValues {
		deck = append(deck, Card{v, false, false})
	}
	return
}

func (deck Deck) flipAllBack() {
	for i := range deck {
		deck[i].Flipped = false
	}
}

func (deck Deck) flipCard(i int) {
	deck[i].Flipped = !deck[i].Flipped
}

func (deck Deck) onlyFlipped() (flippedCards Deck) {
	// filter to only flipped cards
	for _, card := range deck {
		if card.Flipped {
			flippedCards = append(flippedCards, card)
		}
	}
	return
}

func (deck Deck) numberFlippedCards() int {
	return len(deck.onlyFlipped())
}

func (deck Deck) hasFoundMatch() bool {
	// if there are exactly 2 cards flipped
	// and their valus match
	flippedCards := deck.onlyFlipped()
	if len(flippedCards) == 2 {
		if flippedCards[0].Value == flippedCards[1].Value {
			return true
		}
	}
	return false
}

func (deck Deck) removeMatches() {
	for i, card := range deck {
		if card.Flipped {
			deck[i].Matched = true
			deck[i].Flipped = false
		}
	}
}
