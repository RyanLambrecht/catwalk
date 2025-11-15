package factory

import "github.com/RyanLambrecht/catwalk/building"

func (s *NodeStack) isEmpty() bool {
	return len(s.nodes) == 0
}

func (s *NodeStack) Push(node *building.Node) {
	s.nodes = append(s.nodes, node)
}

func (s *NodeStack) Peek() *building.Node {
	if s.isEmpty() {
		return nil
	}
	return s.nodes[0]
}

func (s *NodeStack) Pop() *building.Node {
	if s.isEmpty() {
		return nil
	}

	lastIndex := len(s.nodes) - 1
	top := s.nodes[lastIndex]
	s.nodes = s.nodes[:lastIndex]

	return top
}

func (s *NodeStack) Size() int {
	return len(s.nodes)
}

func (s *NodeQueue) isEmpty() bool {
	return len(s.nodes) == 0
}

func (s *NodeQueue) Enqueue(node *building.Node) {
	s.nodes = append(s.nodes, node)
}

func (s *NodeQueue) Peek() *building.Node {
	if s.isEmpty() {
		return nil
	}
	return s.nodes[len(s.nodes)-1]
}

func (s *NodeQueue) Dequeue() *building.Node {
	if s.isEmpty() {
		return nil
	}

	first := s.nodes[0]
	s.nodes = s.nodes[1:]

	return first
}

func (s *NodeQueue) Size() int {
	return len(s.nodes)
}
