package models

// Node ....
type Node struct {
	*Leaf
	NodesCount int             `json:"nodes_count" msgpack:"n"`
	Nodes      NodeCollection  `json:"-" msgpack:"s"`
	MemNodes   *NodeCollection `json:"children" msgpack:"-"`
	Leaves     LeafCollection  `json:"-" msgpack:"l"`
	MemLeaves  *LeafCollection `json:"leaves" msgpack:"-"`
}

// NewNode ...
func NewNode(label, description string) *Node {
	return &Node{
		Leaf:   NewLeaf(label, description),
		Nodes:  NodeCollection{},
		Leaves: LeafCollection{},
	}
}

// AddNode ...
func (n *Node) AddNode(node *Node) {

	n.lock.Lock()
	defer n.lock.Unlock()

	node.ParentID = n.ID
	n.Nodes = append(n.Nodes, node)
	n.NodesCount = n.Nodes.Len() + n.Leaves.Len()
	n.Nodes.Sort()
}

// AddLeaf ...
func (n *Node) AddLeaf(leaf *Leaf) {
	n.lock.Lock()
	defer n.lock.Unlock()

	leaf.ParentID = n.ID
	n.Leaves = append(n.Leaves, leaf)
	n.NodesCount = n.Nodes.Len() + n.Leaves.Len()
	n.Leaves.Sort()
}

// RemoveNode ...
func (n *Node) RemoveNode(ID string) {
	n.lock.Lock()
	defer n.lock.Unlock()

	index := -1
	for i, node := range n.Nodes {
		if node.ID == ID {
			index = i
			break
		}
	}
	if index > -1 {
		n.Nodes = append(n.Nodes[:index], n.Nodes[index+1:]...)
		n.NodesCount = n.Nodes.Len() + n.Leaves.Len()
	}
}

// RemoveLeaf ...
func (n *Node) RemoveLeaf(ID string) {
	n.lock.Lock()
	defer n.lock.Unlock()

	index := -1
	for i, leaf := range n.Leaves {
		if leaf.ID == ID {
			index = i
			break
		}
	}
	if index > -1 {
		n.Leaves = append(n.Leaves[:index], n.Leaves[index+1:]...)
		n.NodesCount = n.Nodes.Len() + n.Leaves.Len()
	}
}

// EachNode ...
func (n *Node) EachNode(eachFunc func(node *Node)) {

	stack := []*Node{n}
	for len(stack) > 0 {
		// pop
		i := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		eachFunc(i)

		for _, c := range i.Nodes {
			stack = append(stack, c)
		}
	}
}

// EachLeaf ...
func (n *Node) EachLeaf(eachFunc func(leaf *Leaf)) {

	stack := []*Node{n}
	for len(stack) > 0 {
		// pop
		i := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, leaf := range i.Leaves {
			eachFunc(leaf)
		}

		for _, c := range i.Nodes {
			stack = append(stack, c)
		}
	}
}

// FindNode ...
func (n *Node) FindNode(findFunc func(node *Node) bool) *Node {
	n.lock.RLock()
	defer n.lock.RUnlock()

	stack := []*Node{n}
	for len(stack) > 0 {
		// pop
		i := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if findFunc(i) {
			return i
		}

		for _, c := range i.Nodes {
			stack = append(stack, c)
		}
	}
	return nil
}

// FindLeafByID ...
func (n *Node) FindLeafByID(ID string) *Leaf {
	return n.FindLeaf(func(leaf *Leaf) bool {
		return leaf.ID == ID
	})
}

// FindLeaf ...
func (n *Node) FindLeaf(findFunc func(leaf *Leaf) bool) *Leaf {

	n.lock.RLock()
	defer n.lock.RUnlock()
	stack := []*Node{n}
	for len(stack) > 0 {
		// pop
		i := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, leaf := range i.Leaves {
			if findFunc(leaf) {
				return leaf
			}
		}

		for _, c := range i.Nodes {
			stack = append(stack, c)
		}
	}
	return nil
}

// LoadChildren loads nodes and leaves
func (n *Node) LoadChildren() {
	n.MemLeaves = &n.Leaves
	n.MemNodes = &n.Nodes
}

// UnLoadChildren unloads nodes and leaves
func (n *Node) UnLoadChildren() {
	n.MemLeaves = nil
	n.MemNodes = nil
}

// IsDescendent check if node is descendent
func (n *Node) IsDescendent(descendent *Node) bool {
	return n.FindNode(func(node *Node) bool {
		return descendent.ParentID == node.ID
	}) != nil
}

// DescendentLeaves ...
func (n *Node) DescendentLeaves() []string {
	var leavesIDs []string
	n.EachLeaf(func(leaf *Leaf) {
		leavesIDs = append(leavesIDs, leaf.ID)
	})
	return leavesIDs
}

// IsDescendentLeaf check if leaf is descendent
func (n *Node) IsDescendentLeaf(descendent *Leaf) bool {
	return n.FindLeaf(func(leaf *Leaf) bool {
		return descendent.ID == leaf.ID
	}) != nil
}
