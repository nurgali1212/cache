package main

import (
	"fmt"

	"github.com/nurgali1212/cache.git/memory"
)

func main() {
	cache := memory.New()

	cache.Set("userId", 42)
	userId, _ := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId, _ = cache.Get("userId")

	fmt.Println(userId)

}
