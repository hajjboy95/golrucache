package main

import (
	"fmt"
	"github.com/hajjboy95/lru_cache/lru_cache"
)

func main() {
	var cache lru_cache.Cache = lru_cache.New(4)
	val := cache.Get("val")
	fmt.Println("yiked " , val)

	cache.Put("val1", "hi1")
	cache.Put("val2", "hi2")
	cache.Put("val3", "hi3")
	cache.Put("val4", "hi4")
	cache.Put("val5", "hi5")
	cache.Put("val6", "hi6")

	cache.Print()
	a := cache.Get("val4")
	fmt.Println("val4 = " , a)

	a = cache.Get("val6")
	fmt.Println("val6 = " , a)

	a = cache.Get("val5")
	fmt.Println("val5 = " , a)
	cache.Print()
}

