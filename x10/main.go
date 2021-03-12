package main

import (
	"fmt"
	"strconv"
)

//TotalCowsAdd
func TotalCowsAdd(cowA, cowB int32) int32 {
	// returns the total number of Cows in Both Fields
	tca := int(cowA + cowB)

	fmt.Println("======  TotalCowsAdd = ", strconv.Itoa(tca))
	return cowA + cowB
}

type Adder struct{ id int32 }

//AddPtr go:noinline
func (adder *Adder) AddPtr(cowA, cowB int32) int32 {

	tca := cowA + cowB

	fmt.Println("======  Addptr = ", strconv.Itoa(int(tca)))

	return tca

}

//AddVall go:noinline
func (adder Adder) AddVal(cowA, cowB int32) int32 {
	tca := cowA + cowB

	fmt.Println("======  AddVal = ", strconv.Itoa(int(tca)))
	return cowA + cowB
}

func HowManyCowsDoIHave() (int, error) {
	//  get cowA and cowB from external service
	//  This would make a call to external but I will just hard code for now.
	cowA := int32(10)
	cowB := int32(20)
	systemUp := true // status of external system

	adder := Adder{id: 6754}

	if systemUp {
		return int(adder.AddVal(cowA, cowB)), nil
	} else {
		return 0, fmt.Errorf("system Down")
	}

}

func main() {

	// So there are Two Fields with cows
	// I just want to know how many cows total

	TotalCowsAdd(10, 32) // direct call of top-level function

	adder := Adder{id: 6754}
	adder.AddPtr(10, 32) // direct call of method with pointer receiver
	adder.AddVal(10, 32) // direct call of method with value receiver

	(&adder).AddVal(10, 32) // implicit dereferencing

	//===============================================
	//  this is a way to abstract away the cowcount.
	//  The service would have a MAP or a DB in the background
	//  This just returns all the cows in all the fields.

	currentCowCount, err := HowManyCowsDoIHave()
	if err == nil {
		fmt.Println("I have...", currentCowCount)
	} else {
		fmt.Println("problem with cowcount", err)
	}

}
