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

func TestFilterCard(t *testing.T) {
	filter := func(card Card) bool {
	  return card.Rank == Two || card.Rank == Three	
	}
	cards := New(FilterCards(filter))
    for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all twos and threes to be filtered out.")
		}
	}
}

func TestDeck(t *testing.T) {
   cards := New(Deck(3))

   if len(cards) != 52 * 3 {
	   t.Errorf("Expected %d cards, received %d cards.", 52*3, len(cards))
   }
}