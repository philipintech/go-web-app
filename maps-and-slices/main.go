package main

import "log"

func main() {
	var mySlice []string
	mySlice = append(mySlice, "Fish")
	mySlice = append(mySlice, "Squid")
	mySlice = append(mySlice, "battery")

	log.Println(mySlice)
}

//Maps are immutable and are fast (performant)
//Maps are programmatically built into the system not sorted (it does not stay in sorted way, must be accessed via key lookup myMap["myKey"])
//unlike the values in Maps, the values within the slices remain in the same order when accessing the slice
