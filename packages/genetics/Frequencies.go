package genetic

import (
	"fmt"
	"sort"
	"strconv"
)

// Frequencies a struct for allele frequences at a locus
type Frequencies struct {
	Counts map[string]float64 `json:"counts"`
	N      float64            `json:"n"`
	NDip   float64            `json:"ndip"`
	NHet   float64            `json:"nhet"`
}

// AddGenotype is method for adding data to the struct
func (f *Frequencies) AddGenotype(g Genotype) {

	fmt.Printf("Adding genotype %s\n", g)

	if f.Counts == nil {
		f.Counts = make(map[string]float64)
	}

	for _, v := range g.Alleles {
		fmt.Printf("adding %s\n", v)
		f.Counts[v] += 1.0
	}

	f.N += float64(g.Ploidy())
	if g.Ploidy() == 2 {
		f.NDip += 1.0
		if g.IsHeterozygote() {
			f.NHet += 1.0
		}
	}

}

// GetFrequency returns specific allele frequencies
func (f Frequencies) GetFrequency(allele string) float64 {
	if v, found := f.Counts[allele]; found {
		return v / f.N
	}
	return 0.0
}

// Alleles returns all the alleles entered
func (f Frequencies) Alleles() []string {
	var ret []string
	for k := range f.Counts {
		ret = append(ret, k)
	}
	sort.Strings(ret)
	return ret
}

// String overload
func (f Frequencies) String() string {
	ret := "Allele Frequencies:\n"
	for k := range f.Counts {
		freq := f.GetFrequency(k)
		ret += k + " : " + strconv.FormatFloat(freq, 'f', -1, 64) + "\n"
	}
	return ret
}

// Ae is the effective number of alleles
func (f Frequencies) Ae() float64 {
	ret := 0.0

	for k := range f.Counts {
		p := f.GetFrequency(k)
		ret += p * p
	}

	if ret > 0 {
		ret = 1.0 / ret
	}

	return ret

}
