package main

import (
	"container/list"
)

func main() {
	l := list.New()
	a := &list.Element{Value: "1"}
	b := &list.Element{Value: "2"}
	c := &list.Element{Value: "3"}
	d := &list.Element{Value: "4"}
	e := &list.Element{Value: "5"}

	l.PushFront(a)
	l.PushFront(b)
	l.PushFront(c)
	l.PushFront(d)
	l.PushFront(e)

	//
	//var cache  = lru_cache.New(4)
	//
	//val := cache.Get("val")
	//fmt.Println("yiked " , val)
	//
	//cache.Put("val1", "hi1")
	//cache.Put("val2", "hi2")
	//cache.Put("val3", "hi3")
	//cache.Put("val4", "hi4")
	//cache.Put("val5", "hi5")
	//cache.Put("val6", "hi6")
	//
	//cache.PrintAll()
	//a := cache.Get("val4")
	//fmt.Println(a)
	//cache.PrintAll()
}

