package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("INIT")
	var Slice1 []int
	Slice2 := make([]int, 6)
	Slice3 := []int{100, 221, 31, 4}
	report(Slice1, Slice2, Slice3)

	fmt.Println("AFTER APPEND")
	Slice1 = append(Slice1, 5, 6, 7)
	Slice2 = append(Slice2, 9)
	Slice3 = append(Slice3, 1, 99, 900)
	report(Slice1, Slice2, Slice3)

	fmt.Println("COPY OF SLICE3")
	CopyOfSlice3 := make([]int, len(Slice3))
	copy(CopyOfSlice3, Slice3)
	report(Slice3, CopyOfSlice3)

	fmt.Println("SORT OF SLICE3")
	sort.Ints(Slice3)
	report(Slice3)

	fmt.Println("DELETE LAST ELEMENT FROM SLICE3")
	deleted := Slice3[len(Slice3)-1]
	fmt.Printf("    deleted: %d\n", deleted)
	report(Slice3[:len(Slice3)-1])
}

func report(slices ...[]int) {
	for idx, slice := range slices {
		fmt.Printf("    Arg %d\t", idx)
		fmt.Printf("len: %d\t", len(slice))
		fmt.Printf("cap: %d\t", cap(slice))
		fmt.Printf("content: %d\n", slice)
	}
}
