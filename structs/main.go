package main

import "fmt"

type Car struct {
	brand string
	model string
	year  int
	Engine
}

type Engine struct {
	name     string
	power    int
	capacity float64
}

func main() {
	vv := Engine{
		name:     "Volkswagen",
		power:    92,
		capacity: 1.991,
	}

	myCar := Car{
		brand:  "Ford",
		model:  "Galaxy",
		year:   2000,
		Engine: vv,
	}

	mySecondCar := Car{
		brand: "Mini",
		model: "One",
		year:  2001,
		Engine: Engine{
			name:     "Unknown",
			power:    88,
			capacity: 1.4,
		},
	}

	fmt.Printf("MyCar: %+v\n", myCar)
	fmt.Printf("MyCar2: %+v\n", mySecondCar)
}
