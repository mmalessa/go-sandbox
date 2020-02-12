package main

import "fmt"

type Vehicle interface {
	Drive()
	Break()
	Turn()
}

type Car struct{}

func (c Car) Drive() {
	fmt.Println("Car is driving")
}

func (c Car) Break() {
	fmt.Println("Car is breaking")
}

func (c Car) Turn() {
	fmt.Println("Car is turning")
}

type Bicycle struct{}

func (b Bicycle) Drive() {
	fmt.Println("Bicycle is driving")
}

func (b Bicycle) Break() {
	fmt.Println("Bicycle is breaking")
}

func (b Bicycle) Turn() {
	fmt.Println("Bicycle is turning")
}

func main() {
	c := Car{}
	b := Bicycle{}

	DriveVehicle(c)
	DriveVehicle(b)

	// extremity
	var d interface{} = 5
	var e interface{} = 4.9999
	var f interface{} = "some string"
	g := []interface{}{1, "string", false, true, 1.99, []int{1, 2, 3}}

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}

func DriveVehicle(v Vehicle) {
	v.Drive()
	v.Turn()
	v.Break()
}
