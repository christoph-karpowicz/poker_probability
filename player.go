package main

type player struct {
	flop [2]*card
}

// checkCombination looks for hands in all possible combinations
// in a given hand.
func (p *player) checkCombination(combinations [][]*card) []string {
	var allHands []string

	for _, comb := range combinations {
		if len(comb) == 3 {
			cardsWithFlop := newPlayersCardsWholeHand(p.flop, comb)
			allHands = append(allHands, cardsWithFlop.findHands())
		} else {
			cardsWith1stFlopCard := newPlayersCards1Hand(p.flop[0], comb)
			allHands = append(allHands, cardsWith1stFlopCard.findHands())

			cardsWith2ndFlopCard := newPlayersCards1Hand(p.flop[1], comb)
			allHands = append(allHands, cardsWith2ndFlopCard.findHands())
		}
	}

	return allHands
}

// findHands looks for hands in all combinations of
// players hand and all community cards.
func (p *player) findHands(combinationsOf3Cards [][]*card, combinationsOf4Cards [][]*card) string {
	var allHands []string
	var allCombinations [][]*card = combinationsOf3Cards

	allCombinations = append(allCombinations, combinationsOf4Cards...)

	allHands = p.checkCombination(allCombinations)

	if combinationsOf3Cards == nil && combinationsOf4Cards == nil {
		cardsWithFlopOnly := newPlayersCardsOnlyFlop(p.flop)
		handName := cardsWithFlopOnly.findHandsOnlyFlop()

		allHands = append(allHands, handName)
	}

	var bestHandsIndex int

	for i, ukl := range allHands {
		if i == 0 {
			bestHandsIndex = findHandIndex(ukl)
		} else {
			if currentHandsIndex := findHandIndex(ukl); bestHandsIndex > currentHandsIndex {
				bestHandsIndex = currentHandsIndex
			}
		}

	}

	return hands[bestHandsIndex]
}
