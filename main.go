package main

import (
	"bufio"
	. "challenge2019/distributor"
	. "challenge2019/partner"
	. "challenge2019/sort"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	partnerData := make(Partners)

	//Read the Partner.CSV file
	partnerData = readTheCSVFile(partnerData)
	var distributors Distributors

	//Read the Input.csv file
	if len(os.Args) > 2 && os.Args[2] != "" {
		distributors = distributors.ReadInput(os.Args[2])
	} else {
		distributors = distributors.ReadInput("input.csv")
	}
	output := distributors.Output1(partnerData, "", "")
	CreateOutputCSV(output, "output1.csv")
	fmt.Println("OutPut1 File created sucessfully")
	output = distributors.Output2(output, partnerData)
	fmt.Println("OutPut2 File created sucessfully")
	CreateOutputCSV(output, "output2.csv")
}

func readTheCSVFile(partnerData map[string][]Partner) Partners {

	var csvFile *os.File
	var err error

	if len(os.Args) > 1 && os.Args[1] != "" {
		csvFile, err = os.Open(os.Args[1])
	} else {
		csvFile, err = os.Open("partners.csv")
	}
	if err != nil {
		fmt.Println("File not found")
		os.Exit(1)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	lines, error := reader.ReadAll()
	if error != nil {
		log.Fatal(error)
	}
	for i, line := range lines {
		if i == 0 {
			// skip header line
			continue
		}

		minCost, _ := strconv.Atoi(strings.TrimSpace(line[2]))
		costPerGB, _ := strconv.Atoi(strings.TrimSpace(line[3]))
		theatreID := strings.TrimSpace(line[0])

		sizeSlab := strings.TrimSpace(line[1])
		size := strings.Split(sizeSlab, "-")
		minSize, _ := strconv.Atoi(size[0])
		maxSize, _ := strconv.Atoi(size[1])

		partner := Partner{
			SizeSlab:  sizeSlab,
			MinSize:   minSize,
			MaxSize:   maxSize,
			MinCost:   minCost,
			CostPerGB: costPerGB,
			PartnerId: strings.TrimSpace(line[4]),
		}

		if val, ok := partnerData[theatreID]; !ok {
			var partners []Partner
			partners = append(partners, partner)
			partnerData[theatreID] = partners
		} else {
			partnerData[theatreID] = append(val, partner)
		}

	}

	//Sort the each data based on Cost per GB
	SortPartner(partnerData)
	return partnerData
}
