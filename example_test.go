package stlmap

import (
	"fmt"
)

func Example() {
	mp := New()
	mp.Set("key", 123)
	fmt.Println(mp.Get("key"))
	mp.Delete("key")
	fmt.Println(mp.Get("key"))
	// Output:
	// 123
	// <nil>
}
