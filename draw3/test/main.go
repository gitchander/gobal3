package main

import (
	"log"
)

func main() {
	checkError(exampleDrawSamples())
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
