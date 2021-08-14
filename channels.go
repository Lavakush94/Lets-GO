package main

import "log"

const NUM_POOL = 10

func CalVal(intChan chan int) {
	randomNumber := Rand(NUM_POOL)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalVal(intChan)

	num := <-intChan

	log.Println(num)
}
