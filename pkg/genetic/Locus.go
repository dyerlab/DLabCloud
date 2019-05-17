package genetic

import "fmt"

// Locus structure collects genotypes
type Locus struct {
	Name              string      `json:"name"`
	Genotypes         []Genotype  `json:"genotypes"`
	AlleleFrequencies Frequencies `json:"Frequencies"`
}

// AddGenotype is the entry for adding genetic information to this locus
func (l *Locus) AddGenotype(g Genotype) {
	l.AlleleFrequencies.AddGenotype(g)
	l.Genotypes = append(l.Genotypes, g)

	fmt.Printf("ADded locus %s now has %d\n", g, len(l.Genotypes))
}

// Count returns the number of genotypes at the locus
func (l Locus) Count() int {
	return len(l.Genotypes)
}
