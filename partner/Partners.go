package partner

type Partner struct {
	SizeSlab  string
	MinSize   int
	MaxSize   int
	MinCost   int
	CostPerGB int
	PartnerId string
}

type Partners map[string][]Partner
