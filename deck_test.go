package main

import (
	"os"
	"testing"
)

// What should we even test? What do we care about?
// I.e. newDeck() should have:
//		- X amount of cards
//		- first card is Ace of Spades
//		- last card is what we'd expect of playing deck

// Capitalized function indicates ???
// testing.T is our testing "object"
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 { // this should change to 52 in final version
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {
	// Remove the testing deck file
	os.Remove("_decktesting")

	// Calling our init & save function
	deck := newDeck()
	deck.saveToFile("_decktesting")

	// Grab the test deck, load it
	loadedDeck := newDeckFromFile("_decktesting")

	// The actual "check/test"
	// TODO: should be 52 cards not 16
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	// Making sure we clean up!
	os.Remove("_decktesting")
}
