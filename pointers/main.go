package main

import "fmt"

type Person struct {
	id   int
	name string
}

func main() {

	fmt.Println("INIT")
	a := 100
	aPointer := &a
	report(a, aPointer)

	fmt.Println("MODIFY A")
	a = 102
	report(a, &a) // &a == aPointer

	fmt.Println("INIT B BY POINTER")
	bPointer := new(int)
	report(*bPointer, bPointer)

	fmt.Println("MODIFY B BY POINTER")
	*bPointer = 12
	report(*bPointer, bPointer)

	fmt.Println("ACCESS TO POINTER BY &*BPOINTER")
	report(*bPointer, &*bPointer)

	fmt.Println("COPY B TO C VIA POINTER")
	c := *bPointer
	report(c, &c)

	fmt.Println("PERSON")
	myPerson := &Person{
		id: 1,
	}
	fmt.Printf("    myPerson: %v\n", *myPerson)
	myPerson.SetTemporaryName("Joe")
	fmt.Printf("    myPerson: %v\n", *myPerson)
	myPerson.SetName("John")
	fmt.Printf("    myPerson: %v\n", *myPerson)

}

func (p *Person) SetName(n string) {
	p.name = n
}

func (p Person) SetTemporaryName(n string) {
	p.name = n
}

func report(v int, vPointer *int) {
	fmt.Println("    value =", v)
	fmt.Println("    pointer =", vPointer)
}
