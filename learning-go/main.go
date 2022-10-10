package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	var whatToSay string
	var i int = 8
	whatToSay = "Goodbye world"
	fmt.Println(whatToSay)
	fmt.Println("i is set to", i)

	whatWasSaid, whenWasSaid := saySomething()
	fmt.Println("The function returned", whatWasSaid, whenWasSaid)
}

func saySomething() (string, string) {
	return "I am saying something", "now"
}
