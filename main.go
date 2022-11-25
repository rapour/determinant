package main

import (
	"example/det/matrix"
	"fmt"
	"log"
)

func main() {

	// sample matrix
	matrix := matrix.Matrix{
		{3, 6, 2, 4},
		{7, 1, 5, 3},
		{9, 9, 1, 2},
		{4, 6, 3, 2}}

	det, err := matrix.Det()
	if err != nil {
		log.Fatalf("Error in calculating the determinant: %v", err)
	}

	fmt.Printf("The determinant is: %f", det)

}
