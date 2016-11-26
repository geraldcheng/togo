package main

import (
	"fmt"
	"greetings"
	"time"
)

func main() {
	tnow := time.Now()
	fmt.Println(tnow)
	if greetings.IsMorning(tnow) {
		fmt.Printf("Good Morning\n")
	} else if greetings.IsAfternoon(tnow) {
		fmt.Printf("Good Afternoon\n")
	} else if greetings.IsEvening(tnow) {
		fmt.Printf("Good Evening\n")
	}
}
