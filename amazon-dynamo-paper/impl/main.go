package main

import (
	"fmt"
	"crypto/md5"
	"encoding/binary"
)

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

func KeyToPosition(key []byte) uint64 {
	hash := md5.Sum(key)
	return binary.BigEndian.Uint64(hash[:8])
}

func main() {
	Put([]byte("hello"), nil, []byte("world"))
	_, val, ok := Get([]byte("hello"))
	if ok {
		fmt.Println("got:", string(val))
	} else {
		fmt.Println("not found")
	}

	fmt.Println("position(hello):", KeyToPosition([]byte("hello")))
}
