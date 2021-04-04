package main

import "fmt"

func main() {
	cards := []string{"Ace Of Dimaonds", newCard()}
	cards = append(cards, "Four Of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard () string {
	return "Five Of Diamonds"
}

func estimatePi() float64 {
    return 3.14
}