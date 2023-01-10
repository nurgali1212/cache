Cache

go get -u -v  https://github.com/nurgali1212/cache.git

Methods
Set(key string, value any) error
Get(key string) (any, error)
Delete(key string) error
  
  
  
  
Examole

package main

import (
	"fmt"
	"github.com/OrxanKerimov/Cache/cache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId, _ := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId, _ = cache.Get("userId")

	fmt.Println(userId)
}
