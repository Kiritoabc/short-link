package config

type WeightFlag struct {
	Name        string
	Value       string
	Weight      int
	Description string
}

func NewWeightFlag(name, value, description string, weight int) *WeightFlag {
	return &WeightFlag{
		Name:        name,
		Value:       value,
		Weight:      weight,
		Description: description,
	}
}

var WeightFlags = []*WeightFlag{
	Node1,
	Node2,
	Node3,
}

var (
	// Node1 WeightConfig is the config of weight
	Node1 = NewWeightFlag("node1", "127.0.0.1:8081", "node1", 1)
	// Node2 WeightConfig is the config of weight
	Node2 = NewWeightFlag("node2", "127.0.0.1:8082", "node2", 2)
	// Node3 WeightConfig is the config of weight
	Node3 = NewWeightFlag("node3", "127.0.0.1:8083", "node3", 3)
)
