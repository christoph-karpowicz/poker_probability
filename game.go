package main

const cardsInDeck = 52

var cardSuits []string = []string{"spades", "diamonds", "clubs", "hearts"}
var cardValues []string = []string{"ace", "king", "queen", "jack", "10", "9", "8", "7", "6", "5", "4", "3", "2"}
var hands []string = []string{"royalFlush", "straightFlush", "fourOfAkind", "fullHouse", "flush", "straight", "threeOfAkind", "twoPairs", "pair", "highCard"}
var handsFullNames []string = []string{"royal flush", "straight flush", "four of a kind", "full house", "flush", "straight", "three of a kind", "two pairs", "pair", "high card"}

func findCardValuesIndex(cardValue string) int {
	for i, val := range cardValues {
		if cardValue == val {
			return i
		}
	}

	return 0
}

func findHandIndex(hand string) int {
	for i, hnd := range hands {
		if hand == hnd {
			return i
		}
	}

	return 0
}

// get3CardCombination finds all possible combinations of given
// number of community cards dealt at the table.
func get3CardCombination(numberOfCards int, cards []*card) [][]*card {
	allCombinations := make([][]*card, 0)
	tmp := make([]*card, numberOfCards)

	cardCombinations(numberOfCards, cards, 0, tmp, &allCombinations, 0)

	return allCombinations
}

// cardCombinations recursively adds subsequent card combinations
// from a temp slice to a received slice of all combinations.
func cardCombinations(numberOfCards int, cards []*card, index int, tmp []*card, allCombinations *[][]*card, i int) {
	if index == numberOfCards {
		tmpCpy := make([]*card, len(tmp))

		copy(tmpCpy, tmp)
		*allCombinations = append(*allCombinations, tmpCpy)
		return
	}

	if i >= len(cards) {
		return
	}

	tmp[index] = cards[i]
	cardCombinations(numberOfCards, cards, index+1, tmp, allCombinations, i+1)

	cardCombinations(numberOfCards, cards, index, tmp, allCombinations, i+1)
}
