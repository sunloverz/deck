package deck

import (
	"testing"
	"fmt"
)

func ExampleCard() {
	fmt.Println(Card{Rank:Ace, Suit:Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Diamond})
	fmt.Println(Card{Rank: Jack, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Diamonds
	// Jack of Clubs
	// Joker
}

func TestNew(t * testing.T){
  cards := New()
  if len(cards) != 52 {
	  t.Error("Wrong number of cards in a new deck.")
  }
}
