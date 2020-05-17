package main

import (
	"fmt"

	"./poker"
)

func main() {
	fmt.Println("Generowanie losowych rozdań w pokera i empiryczne wyznaczenie prawdopodobieństwa wszystkich konfiguracji.")

	stol10 := poker.NowyStol(10)
	stol10.RozdajNrazy(1000)
}
