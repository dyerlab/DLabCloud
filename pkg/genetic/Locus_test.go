package genetic

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"testing"
)

func TestLocus(t *testing.T) {
	var loc Locus

	if loc.Name != "" {
		t.Error("non-empty name")
	}
	if loc.Count() != 0 {
		t.Error("non-empty genotypes")
	}

	het := []string{"A", "B"}
	hom := []string{"A", "A"}
	for i := 0; i < 100; i++ {
		var g Genotype

		if i%2 == 0 {
			g.Alleles = het
		} else {
			g.Alleles = hom
		}

		loc.AddGenotype(g)
	}
	if loc.Count() != 100 {
		t.Errorf("Wrong number of genotypes at locus expected 100 fuond %d", loc.Count())
	}

	if loc.AlleleFrequencies.GetFrequency("C") != 0.0 {
		t.Error("wrong allele frequency for missing allele")
	}

	f := loc.AlleleFrequencies.GetFrequency("A")
	if f != 0.75 {
		t.Errorf("Error in freq(A), should be 0.75, got %f", f)
	}

	fmt.Println(loc.AlleleFrequencies)

	allalleles := loc.AlleleFrequencies.Alleles()
	sort.Strings(allalleles)
	exp := []string{"A", "B"}

	if !reflect.DeepEqual(allalleles, exp) {
		t.Error("Alleles not returning right alleles")
	}
	loc.AlleleFrequencies.Ae()

	ae := loc.AlleleFrequencies.Ae()
	aeExp := 1.0 / (0.25*0.25 + 0.75*0.75)
	if math.Abs(ae-aeExp) > 0.000001 {
		t.Errorf("Value was: %f, expected %f, who are different as %v", ae, aeExp, ae-aeExp)
	}

}
