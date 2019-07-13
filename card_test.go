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

func TestSort(t *testing.T) {
	cards := New()
	cards = Sort(cards)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
	// fmt.Println(cards)
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(4))
	// fmt.Println(len(cards))
	count := 0
	for _, c:= range cards {
		if c.Suit == Joker {
			count++
		}
	}

	if count != 4 {
		t.Error("Expected 4 jokers, received:", count)
	}
}

// func TestShuffle(t *testing.T) {
// 	cards := New(Shuffle)
// 	fmt.Println(cards)
// }