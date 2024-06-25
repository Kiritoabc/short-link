package midware

import (
	"math/rand"
	"time"
)

var randServer = &RandServer{}

type RandServer struct {
	nodes []string
}

func InitRandServer() {
	randServer.addNodes(Ip)
}

func (rs *RandServer) addNode(node string) {
	rs.nodes = append(rs.nodes, node)
}

func (rs *RandServer) addNodes(nodes []string) {
	rs.nodes = append(rs.nodes, nodes...)
}

func (rs *RandServer) getNode() string {
	rand.Seed(time.Now().UnixNano())
	// get rand index
	randIndex := rand.Intn(len(rs.nodes))

	return rs.nodes[randIndex]
}

func (rs *RandServer) removeNode(node string) {
	if ok, index := rs.getIndexForKey(node); ok {
		rs.nodes = append(rs.nodes[:index], rs.nodes[index+1:]...)
	}
}

func (rs *RandServer) getIndexForKey(key string) (bool, int) {
	for index, node := range rs.nodes {
		if node == key {
			return true, index
		}
	}
	return false, -1
}
