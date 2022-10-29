package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {

	myMap := make(map[string]User)

	me := User{
		FirstName: "Philip",
		LastName:  "Cheung",
	}

	myMap["me"] = me
	log.Println(myMap["me"].FirstName, myMap["me"].LastName)

}

//Maps are immutable and are fast (performant)
//Maps are programmatically built into the system not sorted (it does not stay in sorted way, must be accessed via key lookup myMap["myKey"])
