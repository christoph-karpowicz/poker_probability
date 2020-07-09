package main

import (
	"fmt"
	"strconv"
	"strings"
)

type table struct {
	players               []*player
	deck                  *deck
	handCounter           int
	dealCounter           int
	handCounterPreFlop    map[string]int
	handCounterPostFlop   map[string]int
	handCounterAfterTurn  map[string]int
	handCounterAfterRiver map[string]int
}

// newTable creates a new table with a
// given number of players.
func newTable(numberOfPlayers int) *table {
	newDeck := newDeck()
	players := make([]*player, 0)
	handCounterPreFlop := make(map[string]int)
	handCounterPostFlop := make(map[string]int)
	handCounterAfterTurn := make(map[string]int)
	handCounterAfterRiver := make(map[string]int)

	for i := 1; i <= numberOfPlayers; i++ {
		players = append(players, &player{})
	}

	for _, hand := range hands {
		handCounterPreFlop[hand] = 0
		handCounterPostFlop[hand] = 0
		handCounterAfterTurn[hand] = 0
		handCounterAfterRiver[hand] = 0
	}

	newTable := table{
		deck:                  newDeck,
		players:               players,
		handCounterPreFlop:    handCounterPreFlop,
		handCounterPostFlop:   handCounterPostFlop,
		handCounterAfterTurn:  handCounterAfterTurn,
		handCounterAfterRiver: handCounterAfterRiver,
	}

	return &newTable
}

// calculateProbabilities calculates and prints the probabilities
// of hand values after N deals using the counters from all stages
// of the game.
func (t *table) calculateProbabilities() {
	t.calculateProbabilitiesInGameStage("pre flop", 2, t.handCounterPreFlop)
	t.calculateProbabilitiesInGameStage("post flop", 5, t.handCounterPostFlop)
	t.calculateProbabilitiesInGameStage("after turn", 6, t.handCounterAfterTurn)
	t.calculateProbabilitiesInGameStage("after river", 7, t.handCounterAfterRiver)
}

func (t *table) calculateProbabilitiesInGameStage(gameStageName string, numberOfCards int, counter map[string]int) {
	fmt.Println()
	fmt.Println("Probabilities of hand values " + gameStageName + ":")
	fmt.Println(strings.Repeat("_", 66))
	fmt.Printf("%-30s| %-34s|\n", "Hand value name", "Probability (%)")
	fmt.Println(strings.Repeat("-", 67))

	for i, hand := range hands {
		if numberOfCards == 2 && i != 8 && i != 9 {
			continue
		}

		nFound := counter[hand]
		var probability float64 = (float64(nFound) / float64(t.handCounter)) * float64(100)

		fmt.Printf("%-30s| %-34.3f|\n", handsFullNames[i], probability)
	}
	fmt.Println(strings.Repeat("-", 67))

	fmt.Printf("%-42s: %v\n", "Number of cards available for one player", strconv.FormatInt(int64(numberOfCards), 10))
	fmt.Printf("%-42s: %v\n", "Number of players", strconv.FormatInt(int64(len(t.players)), 10))
	fmt.Printf("%-42s: %v\n", "Number of deals", strconv.FormatInt(int64(t.dealCounter), 10))
	fmt.Printf("%-42s: %v\n", "Number of hands", strconv.FormatInt(int64(t.handCounter), 10))
}

func (t *table) deal() {
	t.deck.shuffle()
	t.deck.splitAndSwap()

	deal := newDeal(t)
	deal.dealFlops()
	deal.findHands(t.handCounterPreFlop)

	deal.flop()
	deal.findHands(t.handCounterPostFlop)

	deal.turn()
	deal.findHands(t.handCounterAfterTurn)

	deal.river()
	deal.findHands(t.handCounterAfterRiver)

	t.gatherCards(deal)

	t.dealCounter++
}

func (t *table) dealNTimes(numberOfDeals int) {
	for i := 0; i < numberOfDeals; i++ {
		t.deal()
		if i%1000 == 0 {
			fmt.Println(strconv.FormatInt(int64(i), 10) + " hands dealt.")
		}
	}
}

// gatherCards gathers dealt cards back to
// the deck.
func (t *table) gatherCards(roz *deal) {
	for _, crd := range roz.communityCards {
		t.deck.returnCard(crd)
	}

	for _, player := range t.players {
		t.deck.returnCard(player.flop[0])
		t.deck.returnCard(player.flop[1])
	}
}
