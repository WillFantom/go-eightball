# 8Ball Go Package

A super simple library for getting a magic 8-ball message in a golang program offline

--- 

### Why?

 1. I needed it for a thing

### Usage

```bash
go get "github.com/willfantom/go-eightball"
```

```golang
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

```

Output:
```plain
Positive: Without a doubt
Neutral: Reply hazy, try again
Negative: Don’t count on it
Random: {Message:Concentrate and ask again ResponseType:Neutral}
```