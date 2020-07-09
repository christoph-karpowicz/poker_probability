package main

import (
	"math/rand"
	"time"
)

type deck struct {
	cards []*card
}

func newDeck() *deck {
	newDeck := deck{
		cards: make([]*card, 0),
	}

	for _, flush := range cardSuits {
		for _, cardValue := range cardValues {
			newDeck.cards = append(newDeck.cards, &card{flush, cardValue})
		}
	}

	return &newDeck
}

// returnCard returns a card to the deck.
func (t *deck) returnCard(krt *card) {
	t.cards = append(t.cards, krt)
}

// drawLastCard draws one card from the top of the deck.
func (t *deck) drawLastCard() *card {
	card := t.cards[len(t.cards)-1]
	t.cards = t.cards[:len(t.cards)-1]

	return card
}

// splitAndSwap splits the deck in two and
// swaps the two parts.
func (t *deck) splitAndSwap() {
	numberOfCardsInPart := rand.Intn(30-20) + 20

	part1 := t.cards[0:numberOfCardsInPart]
	part2 := t.cards[numberOfCardsInPart:len(t.cards)]

	t.cards = append(part2, part1...)
}

func (t *deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(t.cards), func(i, j int) { t.cards[i], t.cards[j] = t.cards[j], t.cards[i] })
}
