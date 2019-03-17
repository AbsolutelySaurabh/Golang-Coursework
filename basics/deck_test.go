package main

import (
	"os"
	"testing"
)

func testNewDeck(t *testing.T) {

	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Exapected ength of deck 16 but it is %v", len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Exapected ength of deck 16 but it is %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
