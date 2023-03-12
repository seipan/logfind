package a

import (
	"fmt"
	"log"
)

func f() {
	log.Println("hi im mmm")
	fmt.Print("log err")
	log.Fatal() //nocheck:thislog
}
