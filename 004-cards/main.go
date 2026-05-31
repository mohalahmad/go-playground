package main

// -----------------------------------------------------------------------------
// Evolution 1: plain variable
// -----------------------------------------------------------------------------
// var card string = "Ace of Spades"

// -----------------------------------------------------------------------------
// Evolution 2: short declaration (type inferred)
// -----------------------------------------------------------------------------
// card := "Ace of Spades"
// card = "Five of Diamonds"  // updating after the fact

// -----------------------------------------------------------------------------
// Evolution 3: delegate to a function
// -----------------------------------------------------------------------------
// card := newCard()

// -----------------------------------------------------------------------------
// Evolution 4: slice of cards
// -----------------------------------------------------------------------------
// cards := []string{"Ace of Spades", newCard()}
// cards = append(cards, "Six of Spades")

// -----------------------------------------------------------------------------
// Evolution 5: custom type (deck) + constructor
// -----------------------------------------------------------------------------
// cards := deck{"Ace of Spades", newCard()}

// -----------------------------------------------------------------------------
// Evolution 6: full deck + constructor
// -----------------------------------------------------------------------------
// cards := newDeck()
// func newCard() string {
// 	return "Five of Diamonds"
// }
// -----------------------------------------------------------------------------
// Evolution 7: printing the deck — direct, then loop, then method receiver
// -----------------------------------------------------------------------------
// fmt.Println(cards)                        // raw: prints the whole slice at once
// for i, card := range cards {              // manual: iterate and print each card
// 	fmt.Println(i, card)
// }
// cards.print()                             // clean: method receiver on deck type

// -------------------------------------------------------------------------

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
