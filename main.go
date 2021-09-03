package main

import "fmt"

func mp() {
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
	fmt.Println("key", key, "value", value)
	}
}

func array() {
	arr := []string{"one", "two", "three"}
	for key, value := range arr {
		fmt.Println("index", key, "value", value)
	}
}

func main() {
	mp()
}