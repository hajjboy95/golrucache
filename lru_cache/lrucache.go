package lru_cache

import (
	"container/list"
	"fmt"
)

type Cache interface {
	Get(k interface{}) *list.Element
	Put(k , v interface{})
	Print()
}

type LruCache struct {
	entries map[interface{}]*list.Element
	list    *list.List
	size    int
}

func New(size int) *LruCache {
	emptyMap := make(map[interface{}]*list.Element)
	emptyList := list.New()
	return &LruCache{ entries: emptyMap, list: emptyList, size: size}
}

func (lc *LruCache) Size() int {
	return lc.size
}

func (lc *LruCache) Get(k interface{}) *list.Element {
	if elem, ok := lc.entries[k]; ok {
		lc.list.MoveToFront(elem) // needs fixing
		return elem
	}
	return nil
}

func (lc *LruCache) Put(k, v interface{}) {
	if elem, ok := lc.entries[k]; ok {
		elem.Value = v
		lc.list.MoveToFront(elem)
	} else {
		if lc.list.Len() >= lc.size {
			e := lc.list.Back()
			lc.list.Remove(e)
		}
		newElem := list.Element{Value: v}
		lc.entries[k] = &newElem // add value to entries
		lc.list.PushFront(&newElem)
	}
}

func (lc *LruCache) Print() {
	i := 0
	for e := lc.list.Front(); e != nil; e = e.Next() {
		fmt.Printf("i = %d, e = %s \n",i, e.Value)
		i += 1
	}
}



