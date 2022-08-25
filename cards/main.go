package main
import "fmt"

func main(){
	// card := "Ace of Spades"
	// card := newCard()
	// fmt.Println(card)

	cards := []string{"Ace of Diamonds" ,newCard()}
	fmt.Println(cards)

	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}


// package main
// import "fmt"

// // deckSize := 52
// var deckSize int

// func main(){
//     // var card string = "Ace of Spades"
//     // card := "Ace of Spades"
//     // card = "Five of Diamonds"
//     // fmt.Println(card)
//     // deckSize := 52
    
//     deckSize = 52
    
//     fmt.Println(deckSize)
// }