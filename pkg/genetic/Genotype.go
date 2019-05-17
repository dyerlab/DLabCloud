package genetic

import (
	"strings"
)

// Genotype is a light weight structore for a diploid locus
type Genotype struct {
	Alleles []string `json:"Alleles"`
}

// Ploidy returns the number of alleles in the genotype
func (g Genotype) Ploidy() int {
	return len(g.Alleles)
}

// IsHeterozygote returns bool indicating both diploid and unequal alleles
func (g Genotype) IsHeterozygote() bool {
	var vals = make(map[string]int)
	for _, v := range g.Alleles {
		vals[v]++
	}
	return g.Ploidy() > 1 && len(vals) > 1
}

func (g Genotype) String() string {
	return strings.Join(g.Alleles, ":")
}
