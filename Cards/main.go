package main

func main() {

	cards := newDeckFromFile("my_cards")
	cards.print()
	cards.suffle()
	cards.print()
}
