package genetic

import (
	"fmt"
	"testing"
)

/*
	Testing for all the units can proceed like this
	go test -coverprofile cp.out
	go tool cover -html=cp.out
*/
func TestGenotype(t *testing.T) {
	var geno Genotype

	if geno.IsHeterozygote() {
		t.Error("Empty genotype is heterozygote")
	}
	if geno.Ploidy() != 0 {
		t.Error("Expected zero ploidy got something else.")
	}
	geno.Alleles = []string{"A", "B"}
	if !geno.IsHeterozygote() {
		t.Error("Type II error")
	}

	if geno.Ploidy() != 2 {
		t.Error("Wrong ploidy")
	}

	fmt.Println(geno)

}
