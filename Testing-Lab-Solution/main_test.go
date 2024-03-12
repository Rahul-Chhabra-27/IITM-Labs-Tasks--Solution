package main

import (

	"testing"
)
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}
func TestPrimeDecomposition(t *testing.T) {
	number := 17;
	primes := primeDecomposition(number);

	if len(primes) == 0 {
		t.Errorf("# of prime multiples cannot be zero");
	}

	if number % 2 == 0 && primes[0] != 2 {
		t.Errorf("Function not working properly");
	}
	
	if len(primes) == 1 && number != primes[0] {
		t.Errorf("Function not working properly");
	}
}