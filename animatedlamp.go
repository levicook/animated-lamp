package animatedlamp

import "sync/atomic"

var nodeCounter uint64 = 10000

func NextNodeID() NodeID {
	next := atomic.AddUint64(&nodeCounter, 1)
	return NodeID(next)
}

type NodeID uint64

type Node struct {
	NodeID NodeID
}

type Network struct {
	Nodes map[NodeID]*Node
}
