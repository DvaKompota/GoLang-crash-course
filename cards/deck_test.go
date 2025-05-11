package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
	if d[0] != "Two of Spades" {
		t.Errorf("Expected first card of Two of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Ace of Clubs" {
		t.Errorf("Expected last card of  Ace of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in loaded deck, but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}

func TestDeal(t *testing.T) {
	deck := newDeck()
	hand := deck.deal(5)
	if len(hand) != 5 {
		t.Errorf("Expected hand size of 5, but got %v", len(hand))
	}
	if len(deck) != 47 {
		t.Errorf("Expected deck size of 47, but got %v", len(deck))
	}
}

func TestShuffle(t *testing.T) {
	cards := newDeck()
	// Create a copy of the original deck to compare after shuffling
	originalDeck := make(deck, len(cards))
	copy(originalDeck, cards)
	cards.shuffle()
	different := false
	for i := range cards {
		if cards[i] != originalDeck[i] {
			different = true
			break
		}
	}
	if !different {
		t.Errorf("Expected deck to be shuffled, but it was not")
	}
}