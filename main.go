package main

import (
	"cryptopals/challenges/set1"
	"cryptopals/challenges/set2"
	"cryptopals/utilities/english"
)

var standard english.Standard     // Standard English Frequency checker implementation
var chiSquared english.ChiSquared // ChiSquared English Frequency implementation

func main() {
	//set1.Challenge1()
	//set1.Challenge2()
	set1.Challenge3(chiSquared)
	//set1.Challenge4("texts/set1/4.txt", standard)
	//set1.Challenge5()
	set1.Challenge6("texts/set1/6.txt", chiSquared)
	//set1.Challenge7("texts/set1/7.txt")
	//set1.Challenge8("texts/set1/8.txt")
	//set2.Challenge1()
	//set2.Challenge2()
	set2.Challenge3()
}
