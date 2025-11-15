package factory

import "github.com/RyanLambrecht/catwalk/building"

// build by branch
type NodeStack struct {
	nodes []*building.Node
}

// build build multiple branches  in parralel
type NodeQueue struct {
	nodes []*building.Node
}
