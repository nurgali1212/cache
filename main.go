package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nurgali1212/cache.git/memory"
)

func main() {
	cache := memory.New()

	cache.Set("userId", 42, time.Second*5)
	userId, err := cache.Get("userId")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userId)
	time.Sleep(time.Second * 6)

	cache.Delete("userId")
	userId, err = cache.Get("userId")

	if err != nil {
		log.Fatal(err)
	}

}
