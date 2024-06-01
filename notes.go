package main

import (
		"fmt"
		"strconv"
		// "reflect"
)


// can only declare package level variables like this, cannot assign a value to them
var z string

func gfsdgds() string {

// use '=' to assign value to variables that have already been declared
// use ':=' to both declare and initiazlie variable in one step, can only be used inside functions
// ':=' infers the type of variable from assigned value
z := "20"
// You're specifying a signed bit size so its -16 -> 15
// rather than unsined where you could go up to 31
y, err := strconv.ParseInt(z, 10, 6)

fmt.Println(y, err)
// fmt.Println()
// fmt.Printf()
return z
}


// To make a Go file executable you must declare a main package and have a main func