package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("hello, world\n")
	rand.Seed(time.Now().Unix())
	fmt.Printf("the time is %d ",rand.Intn(100))
}