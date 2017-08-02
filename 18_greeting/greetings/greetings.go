package greetings

import (
	"fmt"
	"time"
)

func GoodDay() {
	fmt.Println("Good day!")
}

func GoodNight() {
	fmt.Println("Good night!")
}

func IsAM() bool {
	return time.Now().Hour() < 12
}

func IsAfternoon() bool {
	t := time.Now().Hour()
	return t > 12 && t < 18
}

func IsEvening() bool {
	t := time.Now().Hour()
	return t > 18 && t <= 23

}
