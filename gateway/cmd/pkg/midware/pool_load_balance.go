package midware

import "sync/atomic"

var poolServer = &PoolServer{}

type PoolServer struct {
	nodes []string // 	节点
	cur   int64
}

func InitPoolServer() {
	poolServer.addNodes(Ip)
}

func (ps *PoolServer) addNode(node string) {
	ps.nodes = append(ps.nodes, node)
}

func (ps *PoolServer) addNodes(nodes []string) {
	ps.nodes = append(ps.nodes, nodes...)
}

func (ps *PoolServer) getNode() string {
	if ps.cur >= int64(len(ps.nodes)) {
		atomic.StoreInt64(&ps.cur, 0)
	}
	node := ps.nodes[ps.cur]
	atomic.AddInt64(&ps.cur, 1)
	return node
}

func (ps *PoolServer) removeNode(node string) {
	if ok, index := ps.getIndexForKey(node); ok {
		ps.nodes = append(ps.nodes[:index], ps.nodes[index+1:]...)
	}
}

func (ps *PoolServer) getIndexForKey(node string) (bool, int) {
	for index, n := range ps.nodes {
		if n == node {
			return true, index
		}
	}
	return false, -1
}
