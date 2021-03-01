package models

import (
	"sort"
	"strings"
)

// LeafCollection ...
type LeafCollection []*Leaf

func (l LeafCollection) Len() int {
	return len(l)
}

func (l LeafCollection) Less(i, j int) bool {
	return strings.ToUpper(l[i].Label) < strings.ToUpper(l[j].Label)
}

func (l LeafCollection) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// Sort ...
func (l LeafCollection) Sort() {
	sort.Sort(l)
}

// Get ...
func (l LeafCollection) Get(ID string) *Leaf {
	for _, leaf := range l {
		if leaf.ID == ID {
			return leaf
		}
	}
	return nil
}
