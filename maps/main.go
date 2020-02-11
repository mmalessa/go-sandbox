package main

import "fmt"

func main() {
	var Map1 map[string]int
	Map2 := make(map[string]int, 3)
	Map3 := map[string]int{
		"Zosia":  22,
		"Krysia": 33,
		"Basia":  44,
	}
	fmt.Println("INIT")
	report(Map1, Map2, Map3)

	fmt.Println("APPEND TO MAP1")
	// panic!
	// Map1["Jan"] = 11
	Map1 = make(map[string]int, 0)
	Map1["Jan"] = 11
	Map1["Joseph"] = 70
	report(Map1)

	fmt.Println("SOME TESTS")
	if _, ok := Map3["Stefania"]; !ok {
		fmt.Println("    Map3: Stefania not found!")
	}

	fmt.Println("CONTENT OF MAP3")
	for k, v := range Map3 {
		fmt.Printf("    key: %v, \tvalue: %d\n", k, v)
	}
}

func report(maps ...map[string]int) {
	for idx, mymap := range maps {
		fmt.Printf("    Arg %d\t", idx)
		fmt.Println("content: ", mymap)
	}
}
