package main

import (
	"fmt"

	"github.com/willfantom/go-eightball"
)

func main() {
	fmt.Printf("Positive: %s\n", eightball.GetPositive().Message)
	fmt.Printf("Neutral: %s\n", eightball.GetNeutral().Message)
	fmt.Printf("Negative: %s\n", eightball.GetNegative().Message)
	fmt.Printf("Random: %+v\n", eightball.GetRandom())
}
