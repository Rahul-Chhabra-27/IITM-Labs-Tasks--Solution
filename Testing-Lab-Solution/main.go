package main

import "fmt"

// creates a slice of string...
type deck []string

func newDeck() deck {
	// deck(slice) of cards
	cards := deck{}
	// slice of suits....
	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	// slice of value of cards...
	cardsValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardsSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Prime Decompotion
func primeDecomposition(number int) [] int {
	primeMultiples := []int {};
	num := 2;

	for {
		if(number % num == 0) {
			primeMultiples = append(primeMultiples, num);
			number /= num;
		} else {
			num++;
		}

		if number == 1 {
			return primeMultiples;
		}
	}
	return primeMultiples;
}
func main() {
	var cards deck = newDeck()
	fmt.Printf("%v \n", cards)

	var primes [] int = primeDecomposition(100);
	fmt.Printf("Primes are %v \n", primes);
}
