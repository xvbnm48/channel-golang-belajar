package main

import (
	"channel/helpers"
	"log"
)

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)

	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)

}
