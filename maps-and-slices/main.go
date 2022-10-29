package main

import "log"

func main() {

	myMap := make(map[string]int)

	myMap["dog"] = 1
	myMap["other-dog"] = 2
	myMap["dog"] = 3
	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])
}
