package models

// LeafMap ...
type LeafMap struct {
	Leaves map[string]*Leaf
}

// Put ..
func (l *LeafMap) Put(leaf *Leaf) {
	l.Leaves[leaf.ID] = leaf
}

// Remove ..
func (l *LeafMap) Remove(ID string) {
	if _, ok := l.Leaves[ID]; ok {
		delete(l.Leaves, ID)
	}
}

// Has ...
func (l *LeafMap) Has(ID string) bool {
	_, ok := l.Leaves[ID]
	return ok
}

// Get ...
func (l *LeafMap) Get(ID string) *Leaf {
	if leaf, ok := l.Leaves[ID]; ok {
		return leaf
	}
	return nil
}

// Len ...
func (l *LeafMap) Len() int {
	return len(l.Leaves)
}
