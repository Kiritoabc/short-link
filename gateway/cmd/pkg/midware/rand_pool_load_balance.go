package midware

import (
	"math/rand"
	"sync/atomic"
)

var randPoolServer = &RandPoolServer{}

type RandPoolServer struct {
	nodes     []string
	randIndex int64
}

func InitRandPoolServer() {
	// add nodes
	randPoolServer.addNodes(Ip)
	// init rand index
	randPoolServer.initRandIndex()
}

func (r *RandPoolServer) addNode(node string) {
	r.nodes = append(r.nodes, node)
}

func (r *RandPoolServer) addNodes(nodes []string) {
	r.nodes = append(r.nodes, nodes...)
}

func (r *RandPoolServer) getNode() string {
	if r.randIndex >= int64(len(r.nodes)) {
		atomic.StoreInt64(&r.randIndex, 0)
	}
	node := r.nodes[r.randIndex]
	atomic.AddInt64(&r.randIndex, 1)
	return node
}

func (r *RandPoolServer) removeNode(node string) {
	if ok, index := r.getIndexForKey(node); ok {
		r.nodes = append(r.nodes[:index], r.nodes[index+1:]...)
	}
}

func (r *RandPoolServer) getIndexForKey(key string) (bool, int) {
	for index, node := range r.nodes {
		if node == key {
			return true, index
		}
	}
	return false, 0
}

func (r *RandPoolServer) initRandIndex() {
	r.randIndex = rand.Int63n(int64(len(r.nodes)))
}
