package conf

import (
	. "challenge2019/partner"
	"sort"
)

func SortPartner(partner map[string][]Partner) map[string][]Partner {
	for _, v := range partner {
		sort.Sort(sortByCost(v))
	}

	return partner
}

type sortByCost []Partner

func (partners sortByCost) Len() int {
	return len(partners)
}
func (partners sortByCost) Swap(i, j int) {
	partners[i], partners[j] = partners[j], partners[i]
}
func (partners sortByCost) Less(i, j int) bool {
	if partners[i].CostPerGB < partners[j].CostPerGB {
		return true
	}
	if partners[i].CostPerGB > partners[j].CostPerGB {
		return false
	}
	return partners[i].CostPerGB < partners[j].CostPerGB
}
