package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const (
	L     = "/"
	R     = "\\"
	CLEAR = "\033[H\033[2J"
)

var (
	MOD       = flag.Int("m", 4, "How common the left variant?")
	SPEED     = flag.Int("s", 1, "How fast the 10print gonna be?")
	SPEED_MOD = 100
)

func main() {
	flag.Parse()
	fmt.Print(CLEAR)

	speed := time.Duration((SPEED_MOD / *SPEED))
	sleep := speed * time.Millisecond

	for {
		num := rand.Intn(100)

		if (num % *MOD) == 0 {
			fmt.Print(L)
		} else {
			fmt.Print(R)
		}

		time.Sleep(sleep)
	}
}
