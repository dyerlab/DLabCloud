package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/nickng/bibtex"
)

func main() {
	fmt.Println("bbitex2db")

	example := "./cmd/bibtex2db/DyerCitations.bib"

	bibFile, err := ioutil.ReadFile(example)
	if err != nil {
		log.Fatalf("Error reading bibfile: %v", err.Error())
	}

	// var bib *bibtex.BibTex

	bib, _ := bibtex.Parse(bytes.NewBuffer(bibFile))

	fmt.Printf("Parsed and formatted: %v", *bib)
}
