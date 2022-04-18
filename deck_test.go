package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()
	first := d[0]
	last := d[len(d)-1]

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}

	if first != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %s", first)
	}

	if last != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %s", last)
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {

	file := "_decktesting"

	os.Remove(file)

	newDeck := newDeck()
	newDeck.saveToFile(file)

	loadedDeck := newDeckFromFile(file)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %d", len(loadedDeck))
	}

	os.Remove(file)
}
