// Ftoc prints two Farhenheit-to-Celsius conversions.
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g degrees F = %g degrees C \n", freezingF, fToC(freezingF)) //"32 degrees F = 0 degrees C"
	fmt.Printf("%g degrees F = %g degrees C \n", boilingF, fToC(boilingF))   //"212 degrees F = 100 degrees C"

}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9

}
