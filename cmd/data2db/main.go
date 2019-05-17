package main

import (
	"fmt"

	"github.com/dyerlab/DLabCloud/pkg/genetic"
)



func main() {
	fmt.Println("data2db")


	alleles := [][]string{{"A", "B"}, {"B", "A"}, {"A", "A"}}

	fmt.Printf("this is the number %d\n", 5)

	loc := genetic.Genotype{Alleles: []string{"A","B"} }

	fmt.Println(loc)


	fmt.Println(alleles)
}