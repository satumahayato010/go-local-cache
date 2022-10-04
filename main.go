package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := sync.Map{}

	m.Store("key 1", "value 1")
	m.Store("key 2", "value 2")
	m.Store("key 3", "value 3")
	m.Store("key 4", "value 4")
	m.Store("key 5", "value 5")

	LOS1, ok := m.LoadOrStore("key 6", "value 6")
	if ok {
		fmt.Printf("key 6: %v \n", LOS1)
	}
	fmt.Printf("key 6: %v \n", LOS1)

	<-time.After(5 * time.Millisecond)

	m.Range(func(key interface{}, value interface{}) bool {
		fmt.Printf("key: %v(Type: %T) -> Value: %v(Type: %T\n", key, key, value, value)
		return true
	})

	<-time.After(5 * time.Millisecond)
	fmt.Println("-------------------------------------------------------")

	m.Delete("key 5")

	m.Range(func(key interface{}, value interface{}) bool {
		fmt.Printf("key: %v(Type: %T) -> Value: %v(Type: %T\n", key, key, value, value)
		return true
	})
}
