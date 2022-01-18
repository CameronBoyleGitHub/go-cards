package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// The arg before the name is the "receiver". The word "deck"
// enforces type and 'd' is the named copy being passed.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// No receiver, it makes the function potentially ambiguous,
// the receiver would imply a modification of the underlying
// slice and not so much the creation of two, new slices.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// Reverting back to "parent type" with the cast of 'd'
	// Join takes a string slice and joins elements sequentially
	// using a delimiter as second arg.
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Doing some more conversion on the byte-slice
	// .Split() will divide up our casted-to-string byte-slice
	// by the delimiter and save them sequentially as elements
	s := strings.Split(string(bs), ",")
	// Type conversion from slice of strings --> deck
	return deck(s)
}

// Takes the deck we're working on, randomizes the elements.
func (d deck) shuffle() {
	// Initiating random seed -- will use time to give us our
	// value of 0-63; time.UnixnNano() based on the epoch
	source := rand.NewSource(time.Now().UnixNano())
	// Creating a random source "object"
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// One-line swap
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
