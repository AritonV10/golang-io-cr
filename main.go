package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	
	cards := createDeck() // create the deck 
	fmt.Println("Shuffeled Deck: ") 
	cards.shuffleDeck() // shuffle the cards
	cards.printDeck() // print the shuffeled cards
	
	fmt.Println("\nHand: ") // print the hand
	hand := cards.dealCards(5) // deal 5 cards
	hand.printHand() // print the hand
	fmt.Println("\nAfter hand:")
	cards.printDeck() // print the deck again without those 5 cards

	fmt.Println("\nSaving the deck...")
	byt := stringToByte(hand) // save the hand and convert the cards from string to byte
	saveHand(byt) // save to file

	test := readHand("hand.txt") // read from file and convert from byte to string
	fmt.Println("\nReading File...")
	test.printHand() // print the hand 
}
func stringToByte(hand []string) []byte{
	c, delimiter:= " ", ", " // an option is to use the Join function
	for card := range hand{
		if card != len(hand)-1 {
			c += hand[card] + delimiter
		} else { c += hand[card] }
		
	}
	byt := make([]byte, 1024)
	byt = []byte(c)
	return byt
}
func bytesToString(bytes []byte) []string {
	newString := []string{}
	c := string(bytes[:])
	char := ""
	idxOfComma := strings.LastIndex(string(c), ",")
	for idx := 0; idx < len(c); idx++{ // an option would be strings.Split() function
		cChar := string(c[idx])
		if(cChar != ","){
			char += string(c[idx])
		} else {	
			newString = append(newString, char)
			char = ""
		}
		if(idx+1==len(c)){
			newString = append(newString, string(c[idxOfComma+1:len(c)])) // append the last substring without a comma
		}
	} 
	
	return newString
}
func saveHand(data []byte){
	err := ioutil.WriteFile("hand.txt", data, 0644)
	checkForError(err)
}
func readHand(fileName string) deck {
	byt, err := ioutil.ReadFile(fileName)
	checkForError(err)
	newString := bytesToString(byt)
	return newString
}
func checkForError(e error){
	if(e != nil){
		panic(e)
	}
}

