package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	X int
	Y int
}

var v1 = new(Vertex)
var (
	a1 = [10]int{11, 222, 3333}
	a2 = []int{11, 222, 3333}
	a4 = make([]int, 10)

	//
	v2 = make([]Vertex, 1)
)

var a3 []int

func main() {
	// Appending to a slick

	//
	println("Convert Array Type from Integer to String:", strings.Join(intArrayToStringArray(a2), ","))
	println("Convert Array Type from Integer to String(2):", strings.Join(*intArrayToStringArray2(a2), ","))

	printSlice(a4)
	fmt.Println(v2)

	//
	v := Vertex{1, 2}
	fmt.Println(v.X, v1)

	//
	fmt.Println(a1[0], a1[9], a1)
	fmt.Println(a2[0], a2[len(a2)-1], a2)

	//
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println("1:4数组的长度", s, "len", len(s))
	fmt.Println("1:数组的长度", primes[1:])
	fmt.Println(":3数组的长度", primes[:3])

	//

	customObject := []struct {
		id   int
		name string
	}{
		{id: 1, name: "aa"},
		{id: 2, name: "bb"},
	}
	fmt.Println(customObject)

	//
	printSlice(s)
	printSlice(primes[2:])
	printSlice(primes[:3])

	// var a3 []int
	printSlice(a3)
	if a3 == nil {
		fmt.Print("Is null")
		a3 := []int{111, 111}
		a3[0] = 2
		// a3[1] = 2
		printSlice(a3)
	}
}
