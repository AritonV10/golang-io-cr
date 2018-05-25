package main

import (
	"fmt"
	"math/rand"
	"time"
)
type deck []string
// create the deck
func createDeck() deck {
	return createDeckHelperFunction()
}
func (d *deck) dealCards(numberOfCardsToDeal int) deck {
	return (*d).dealRandomCards(numberOfCardsToDeal)
}
// shuffle the deck
func (d *deck) shuffleDeck()  {
	for i := 0; i < len(*d); i++{
		for j := 0; j < len(*d); j++{
			randomNumber := randomInteger(len(*d))
			(*d).swapCards(j, randomNumber)
		}
	}
}
// print the deck
func (d deck) printDeck(){
	for _, card := range d{
		fmt.Println(card)
	}
}
func (h deck) printHand(){
	for _, handCards := range h{
		fmt.Println(handCards)
	}
}
/* utils */
func createDeckHelperFunction() deck {
    d := deck{}
	f := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	s := []string{"Ace of", "Two of", "Three of", "Four of", "Five of", 
				  "Six of", "Seven of", "Eight of", "Nine of", "Ten of", "Jack of", "King of", "Queen of"}
	for _, ftype := range f{
		for _, stype := range s{
			d = append(d, stype + " " + ftype)
		}
	}
	return d
}
func (d *deck) swapCards(firstPosition int, secondPosition int) {
	temp := (*d)[firstPosition]
	(*d)[firstPosition] = (*d)[secondPosition]
	(*d)[secondPosition] = temp
}
func randomInteger(upTo int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(upTo-1)
}
func (d *deck) dealRandomCards(numberOfCards int) deck{
	newDeck := deck{}
	
	for i := 0; i < numberOfCards; i++{
		randomNumber := randomInteger(len(*d))
		if(!newDeck.containsCard((*d)[randomNumber])){
			newDeck = append(newDeck, (*d)[randomNumber])
			(*d).swapCards(randomNumber, len(*d)-1)
			*d = (*d)[:(len(*d)-1)]
		} else { i-- }	
	}
	return newDeck
}
func (h deck) containsCard(cardName string) bool{
	for _, card := range h{
		if(card==cardName) { return true }
	}
	return false
}
