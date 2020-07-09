package main

type hand []*card

func (h *hand) areOfSameSuit() bool {
	var firstSuit string

	for i, crd := range *h {
		if i == 0 {
			firstSuit = crd.suit
			continue
		}

		if crd.suit != firstSuit {
			return false
		}
	}

	return true
}
