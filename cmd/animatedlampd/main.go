package main

import (
	"animatedlamp"
	"log"
)

func main() {
	for i := 0; i < 2000; i++ {
		log.Println(animatedlamp.NextNodeID())
	}
}
