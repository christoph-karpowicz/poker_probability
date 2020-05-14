package main

import (
	"fmt"
)

type talia struct {
	karty karta[]
}

type karta struct {
	kolor  string
	figura string
}

func main() {
	fmt.Println("Generowanie losowych rozdań w pokera i empiryczne wyznaczenie prawdopodobieństwa wszystkich konfiguracji.")
}
