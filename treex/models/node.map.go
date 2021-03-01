package models

// NodeMap ...
type NodeMap struct {
	Nodes map[string]*Node
}

// Put ..
func (n *NodeMap) Put(node *Node) {
	n.Nodes[node.ID] = node
}

// Remove ..
func (n *NodeMap) Remove(ID string) {
	if _, ok := n.Nodes[ID]; ok {
		delete(n.Nodes, ID)
	}
}

// Has ...
func (n *NodeMap) Has(ID string) bool {
	_, ok := n.Nodes[ID]
	return ok
}

// Get ...
func (n *NodeMap) Get(ID string) *Node {
	if node, ok := n.Nodes[ID]; ok {
		return node
	}
	return nil
}

// Len ...
func (n *NodeMap) Len() int {
	return len(n.Nodes)
}
