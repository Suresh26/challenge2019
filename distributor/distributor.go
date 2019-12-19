package distributor

import (
	"bufio"
	. "challenge2019/partner"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Distributors []Distributor

// Distributor struct for input.csv
type Distributor struct {
	DistributorId string
	Size          int
	TheatreId     string
}

// Output - Result Structure
type Output struct {
	DistributorId string
	IsValid       bool
	PartnerId     string
	TotalCost     int
	size          int
}

type partnerCapacity struct {
	DistID []string
	size   int
}

type Capacity map[string]int

type DistCapacity map[string]partnerCapacity

// IDistributor -- Interface
type IDistributor interface {
	ReadInput() Distributors
	Output1(map[string][]Partner, string, string) Distributors
	Output2(Output, map[string][]Partner) Distributors
}

// ReadInput Method for read the input.csv file
func (distributors Distributors) ReadInput(filePath string) Distributors {
	csvFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("File not found")
		os.Exit(1)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	lines, error := reader.ReadAll()
	if error != nil {
		log.Fatal(error)
	}

	for _, line := range lines {

		minCost, _ := strconv.Atoi(strings.TrimSpace(line[1]))
		distributors = append(distributors, Distributor{
			DistributorId: strings.TrimSpace(line[0]),
			Size:          minCost,
			TheatreId:     strings.TrimSpace(line[2]),
		})

	}
	return distributors
}

// Output1 Method for Find the partner for each delivery where cost of delivery is minimum
func (distributors Distributors) Output1(partners map[string][]Partner, pID string, dID string) []Output {
	outputs := []Output{}
	i := 0

Check:
	for i <= len(distributors)-1 {

		dist := distributors[i]
		if pval, Ok := partners[dist.TheatreId]; !Ok {
			fmt.Println("Theatre is not Found")
			os.Exit(1)
		} else {
			for _, p := range pval {
				if dist.Size >= p.MinSize && dist.Size <= p.MaxSize {
					if p.PartnerId == pID && dist.DistributorId == dID {
						continue
					}

					outputs = append(outputs, Output{
						DistributorId: dist.DistributorId,
						IsValid:       true,
						PartnerId:     p.PartnerId,
						TotalCost:     dist.Size * p.CostPerGB,
						size:          dist.Size,
					})
					i++
					goto Check
				}

			}
			outputs = append(outputs, Output{
				DistributorId: dist.DistributorId,
				IsValid:       false,
				PartnerId:     "",
				TotalCost:     0,
			})
			i++
		}
	}
	return outputs
}

// Output2 Method to check the capacity of partner and change the Partner
func (distributors Distributors) Output2(out []Output, partners map[string][]Partner) []Output {
	capacities := readCapacityCSV()
	distCapacities := findPartnerCapacity(out, len(capacities))
	// totalCost := getTotalCost(out)
	isExceed := false
	for K := range capacities {

		if capacities[K] < distCapacities[K].size {
			isExceed = true
			for i := range distCapacities[K].DistID {
				out = distributors.Output1(partners, K, distCapacities[K].DistID[i])
				break
			}
		}
	}
	if isExceed {
		distributors.Output2(out, partners)
	}

	return out
}

func getTotalCost(output []Output) (totalCost int) {
	for _, out := range output {
		totalCost += out.TotalCost
	}
	return
}

// func checkTheMinCostPartner(distributors Distributors, partners map[string][]Partner, totalCost int, partnerID string, distributorID string) []Output {
// 	output := distributors.Output1(partners, partnerID, distributorID)
// 	return output
// }

func findPartnerCapacity(Output []Output, capLen int) (distCapacity DistCapacity) {
	distCapacity = make(DistCapacity)
	for _, out := range Output {
		if _, Ok := distCapacity[out.PartnerId]; !Ok {
			var distID []string
			distID = append(distID, out.DistributorId)
			distCapacity[out.PartnerId] = partnerCapacity{
				DistID: distID,
				size:   out.size,
			}
		} else {
			patnerCap := partnerCapacity{}
			patnerCap.DistID = append(distCapacity[out.PartnerId].DistID, out.DistributorId)
			patnerCap.size = distCapacity[out.PartnerId].size + out.size
			distCapacity[out.PartnerId] = patnerCap
		}
	}
	return
}

func readCapacityCSV() (capacity Capacity) {
	capacity = make(Capacity)
	var csvFile *os.File
	var err error
	if len(os.Args) > 3 && os.Args[3] != "" {
		csvFile, _ = os.Open(os.Args[3])
	} else {
		csvFile, _ = os.Open("capacities.csv")
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
		capacityGB, _ := strconv.Atoi(strings.TrimSpace(line[1]))
		capacity[strings.TrimSpace(line[0])] = capacityGB

	}
	return
}

func CreateOutputCSV(output []Output, fileName string) {
	var _, err = os.Stat(fileName)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(fileName)
		if err != nil {
			return
		}
		defer file.Close()
	}
	openFile, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer openFile.Close()
	for _, out := range output {
		_, err = openFile.WriteString(fmt.Sprintf("%s,%t,%s,%d\n", out.DistributorId, out.IsValid, out.PartnerId, out.TotalCost))
	}

}
