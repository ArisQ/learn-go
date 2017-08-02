package main

import (
	"greetings/greetings"
)

func main() {
	if greetings.IsEvening() {
		greetings.GoodNight()
	} else {
		greetings.GoodDay()
	}
}
