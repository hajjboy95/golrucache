package main

import (
	"fmt"
	"github.com/hajjboy95/lru_cache/lru_cache"
)

func main() {
	var cache  = lru_cache.New(4)
	val := cache.Get("val")
	fmt.Println("yiked " , val)

	cache.Put("val1", "hi1")
	cache.Put("val2", "hi2")
	cache.Put("val3", "hi3")
	cache.Put("val4", "hi4")
	cache.Put("val5", "hi5")
	cache.Put("val6", "hi6")

	cache.PrintAll()
	a := cache.Get("val4")
	fmt.Println(a)
	cache.PrintAll()
}

