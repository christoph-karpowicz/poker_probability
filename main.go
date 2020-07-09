package main

func main() {
	// Creates a new table with 10 players.
	table10 := newTable(10)
	table10.dealNTimes(10000)
	table10.calculateProbabilities()
}
