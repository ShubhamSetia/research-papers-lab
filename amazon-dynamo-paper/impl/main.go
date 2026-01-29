package main

import "fmt"

type item struct {
	Context []byte
	Value   []byte
}

var storage = make(map[string]item)

func Put(key, context, value []byte) {
	storage[string(key)] = item{
		Context: context,
		Value:   value,
	}
}

func Get(key []byte) (context, value []byte, ok bool) {
	it, ok := storage[string(key)]
	if !ok {
		return nil, nil, false
	}
	return it.Context, it.Value, true
}

func main() {
	Put([]byte("hello"), nil, []byte("world"))
	_, val, ok := Get([]byte("hello"))
	if ok {
		fmt.Println("got:", string(val))
	} else {
		fmt.Println("not found")
	}
}