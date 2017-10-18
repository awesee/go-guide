package main

import "fmt"

func main() {
	// make(map[key-type]val-type)
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2
	m["k3"] = 3
	fmt.Println("map: ", m)
	fmt.Println("len:", len(m))

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	delete(m, "k2")
	fmt.Println("map:", m)
	// The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	v2, prs := m["k2"]
	fmt.Println("v2:", v2)
	fmt.Println("prs:", prs)
	// You can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
