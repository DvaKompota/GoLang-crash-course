package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	hand1 := cards.deal(5)
	hand2 := cards.deal(5)
	fmt.Println("\nHand 1:")
	hand1.print()
	fmt.Println("\nHand 2:")
	hand2.print()
	fmt.Println("\nRemaining cards in deck:")
	cards.print()
}
