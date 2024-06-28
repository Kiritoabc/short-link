package midware

import (
	"fmt"
	"github.com/Kiritoabc/short-link/gateway/cmd/pkg/config"
	"math/rand"
	"time"
)

var weightPoolServer = &WeightPoolServer{}

type WeightPoolServer struct {
	weightNodes []WeightNode
}

type WeightNode struct {
	Node   string
	Weight int
}

func InitWeightPoolServer() {
	// init
	weightFlags := config.WeightFlags
	for _, wflag := range weightFlags {
		WeightIps = append(WeightIps, WeightNode{
			Node:   wflag.Value,
			Weight: wflag.Weight,
		})
	}
	weightPoolServer.addWeightNodes(WeightIps)
}

func (w *WeightPoolServer) addWeightNode(wNode WeightNode) {
	w.weightNodes = append(w.weightNodes, wNode)
}

func (w *WeightPoolServer) addWeightNodes(wNodes []WeightNode) {
	w.weightNodes = append(w.weightNodes, wNodes...)
}

func (w *WeightPoolServer) getNode() string {
	totalWeight := w.getTotalWeight()
	// get node
	curWeight := 0
	randWeight := w.getRandomWeight(totalWeight)
	for _, wNode := range w.weightNodes {
		curWeight += wNode.Weight
		if curWeight > randWeight {
			return wNode.Node
		}
	}
	//
	fmt.Println("err...")
	return ""
}

func (w *WeightPoolServer) getTotalWeight() int {
	totalWeight := 0
	for _, wNode := range w.weightNodes {
		totalWeight += wNode.Weight
	}
	return totalWeight
}

func (w *WeightPoolServer) getRandomWeight(totalWeight int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(totalWeight)
}
