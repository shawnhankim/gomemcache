package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	mc := memcache.New("127.0.0.1:11211", "127.0.0.1:11212")
	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})

	it, err := mc.Get("foo")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("1. ", string(it.Value))

	err = mc.Delete("foo")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("2. the key of foo is deleted")
	it, err = mc.Get("foo")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("3. ", it)
}
