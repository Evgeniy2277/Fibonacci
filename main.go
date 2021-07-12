package main

import (
	"fmt"
	"github.com/Evgeniy2277/fibonacci/cache"
	"github.com/Evgeniy2277/fibonacci/closure"
	"github.com/Evgeniy2277/fibonacci/recursion"
	"github.com/Evgeniy2277/fibonacci/web"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fibImplementations := []web.Fiber{recursion.NewRecursion(), cache.NewCache(), closure.NewClosure()}
	fiber := fibImplementations[rand.Intn(len(fibImplementations))]
	fmt.Printf("Starting %s\n", fiber.Name())
	web.Serve(fiber)
}
