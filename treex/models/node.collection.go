package models

import (
	"sort"
	"strings"
)

// NodeCollection ...
type NodeCollection []*Node

func (n NodeCollection) Len() int {
	return len(n)
}

func (n NodeCollection) Less(i, j int) bool {
	return strings.ToUpper(n[i].Label) < strings.ToUpper(n[j].Label)
}

func (n NodeCollection) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

// Sort ...
func (n NodeCollection) Sort() {
	sort.Sort(n)
}

// Get ...
func (n NodeCollection) Get(ID string) *Node {
	for _, node := range n {
		if node.ID == ID {
			return node
		}
	}
	return nil
}
