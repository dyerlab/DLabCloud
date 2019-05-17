package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// Individual type has strata records for individuals
type Individual struct {
	ID         int
	Population string
	Country    string
	Region     string
	Sex        string
}

func main() {
	fmt.Println("data2db")

	csvFile, err := os.Open("./data/chr2.umich.phased.ordered.snp")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = ' '
	reader.FieldsPerRecord = -1 // header and all rows must have the same number of fields
	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, _ := sql.Open("sqlite3", "./data/chr2.db")

	// Make the SNP table info rs 0, loc1 3, loc2 4
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS snps (rs TEXT PRIMARY KEY, loc1 INTEGER, loc2 INTEGER)")
	statement.Exec()
	statement, _ = db.Prepare("INSERT INTO snps (rs, loc1, loc2) VALUES (?,?,?)")
	fmt.Println("Making SNP db")
	for i := 0; i < len(csvData[0]); i++ {
		rs := csvData[0][i]
		l1, _ := strconv.Atoi(csvData[3][i])
		l2, _ := strconv.Atoi(csvData[4][i])
		statement.Exec(rs, l1, l2)
		fmt.Print(".")
		if i > 0 && i%100 == 0 {
			fmt.Printf(" [%d / %d snp desc]\n", i, len(csvData[0]))
		}
	}
	fmt.Printf(" [%d / %d snp desc]\n", len(csvData[0]), len(csvData[0]))
	statement.Close()
	fmt.Println(" ")

	// Make the individual stuff
	fmt.Println("Creating Individual Strata Data")
	statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS individuals (id INTEGER, population TEXT, country TEXT, region TEXT, sex TEXT)")
	statement.Exec()
	statement, _ = db.Prepare("INSERT INTO individuals (id, population, country, region, sex) VALUES (?,?,?,?,?)")
	for i := 5; i < (len(csvData) - 1); i += 2 {

		if csvData[i][0] != csvData[(i + 1)][0] {
			panic("individual syncing is out of whack.  ID's do not match on subsequent line")
		}

		statement.Exec(i, csvData[i][2], csvData[i][3], csvData[i][4], csvData[i][6])
		fmt.Print(".")
		if (i-4)%100 == 0 && i != 4 {
			fmt.Printf(" [%d / %d inds]\n", (i - 4), len(csvData))
		}
	}
	fmt.Printf(" [%d / %d inds]\n", len(csvData), len(csvData))
	statement.Close()
	fmt.Println(" ")

	// Make the loci
	fmt.Println("Creating Locus Table")
	statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS loci (rs TEXT PRIMARY KEY, fA Float64, fT Float64, fG Float64, fC Float64, Ae Float64, Ho Float64, He Float64, genotypes TEXT)")
	statement.Exec()
	statement, _ = db.Prepare("INSERT INTO loci (rs, fA, fT, fG, fC, Ae, Ho, He, genotypes) VALUES(?,?,?,?,?,?,?,?,?)")

	for j := 7; j < len(csvData[5]); j++ {
		var genos []string

		freq := map[string]float64{
			"A": 0.0,
			"T": 0.0,
			"G": 0.0,
			"C": 0.0,
		}

		Ho := 0.0
		for i := 5; i < (len(csvData) - 1); i += 2 {
			freq[csvData[i][j]]++
			freq[csvData[(i + 1)][j]]++
			if csvData[i][j] != csvData[(i + 1)][j] {
				Ho++
			}
			loc := csvData[i][j] + ":" + csvData[(i + 1)][j]
			genos = append(genos, loc)
		}

		N := freq["A"] + freq["T"] + freq["G"] + freq["C"]
		freq["A"] /= N
		freq["T"] /= N
		freq["G"] /= N
		freq["C"] /= N

		Ae := 0.0
		if freq["A"] > 0 {
			Ae += (freq["A"] * freq["A"])
		}
		if freq["T"] > 0 {
			Ae += (freq["T"] * freq["T"])
		}
		if freq["C"] > 0 {
			Ae += (freq["C"] * freq["C"])
		}
		if freq["G"] > 0 {
			Ae += (freq["G"] * freq["G"])
		}

		Ae = 1.0 / Ae

		Ho = Ho / (N / 2.0)
		He := 1.0 - (freq["A"]*freq["A"] + freq["T"]*freq["T"] + freq["G"]*freq["G"] + freq["C"]*freq["C"])

		genotypes := strings.Join(genos[:], ",")
		statement.Exec(csvData[0][j-7], freq["A"], freq["T"], freq["G"], freq["C"], Ae, Ho, He, genotypes)

		fmt.Print(".")
		if (j-7)%100 == 0 && (j-7) != 0 {
			fmt.Printf("%d / %d loci]\n", (j - 7), len(csvData[5])-7)
		}
	}
	fmt.Printf("%d / %d loci]\n", len(csvData[5])-7, len(csvData[5])-7)
	statement.Close()
	fmt.Println("\nFinished")
}
