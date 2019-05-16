package genetic

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"testing"
)

func TestFrequencies(t *testing.T) {

	var freqs Frequencies

	var vals = freqs.Alleles()

	if len(vals) != 0 {
		t.Error("Empty frequency object not zero alleles")
	}

	alleles := [][]string{{"A", "B"}, {"B", "A"}, {"A", "A"}}
	var loc1 Genotype
	var loc2 Genotype
	var loc3 Genotype
	loc1.Alleles = alleles[0][0:2]
	loc2.Alleles = alleles[1][0:2]
	loc3.Alleles = alleles[2][0:2]

	freqs.AddGenotype(loc1)
	freqs.AddGenotype(loc2)
	freqs.AddGenotype(loc3)

	vals = freqs.Alleles()
	if len(vals) != 2 {
		t.Errorf("Freq should have two alleles but had %d instead", len(vals))
	}

	if freqs.GetFrequency("Z") != 0.0 {
		t.Error("Bad frequency for missing allele.")
	}
	if freqs.GetFrequency("A") != 4.0/6.0 {
		t.Error("Bad frequency for present allele.")
	}

	allalleles := freqs.Alleles()
	sort.Strings(allalleles)
	exp := []string{"A", "B"}

	if !reflect.DeepEqual(allalleles, exp) {
		t.Error("Alleles not returning right alleles")
	}

	fmt.Println(freqs)

	ae := freqs.Ae()
	aeExp := 1.0 / ((1.0/3.0)*(1.0/3.0) + (2.0/3.0)*(2.0/3.0))
	if math.Abs(ae-aeExp) > 0.000001 {
		t.Errorf("Value was: %f, expected %f, who are different as %v", ae, aeExp, ae-aeExp)
	}

}
