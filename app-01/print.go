package main

import (
	"fmt"
	"strconv"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func intArrayToStringArray(a []int) []string {
	var r []string
	for _, value := range a {
		r = append(r, strconv.Itoa(value))
	}
	return r
}

func intArrayToStringArray2(a []int) *[]string {
	var r []string
	for index := range a {
		r = append(r, strconv.Itoa(a[index]))
	}
	return &r
}
