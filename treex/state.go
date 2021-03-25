package treex

import (
	"fmt"
	dataModels "krolus/models"
	"krolus/treex/models"
	"krolus/treex/persistence"
	"os"
	"sync"
)

// State ...
type State struct {
	lock      sync.RWMutex
	Root      *models.Node
	MapNodes  *models.NodeMap // flat map of nodes
	MapLeaves *models.LeafMap // flat map of leaves
	persister persistence.Persister
}

// NewState ...
func NewState(root *models.Node, saver persistence.Persister) (*State, error) {

	err := saver.Load(root)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	//bootstrap mapNodes
	mapNodes := &models.NodeMap{Nodes: make(map[string]*models.Node)}
	root.EachNode(func(node *models.Node) {
		mapNodes.Put(node)
	})

	//bootstrap mapLeaves
	mapLeaves := &models.LeafMap{Leaves: make(map[string]*models.Leaf)}
	root.EachLeaf(func(leaf *models.Leaf) {
		mapLeaves.Put(leaf)
	})

	return &State{
		Root:      root,
		MapNodes:  mapNodes,
		MapLeaves: mapLeaves,
		persister: saver,
	}, nil
}

// UpdateCounters ...
func (s *State) UpdateCounters(subInfos dataModels.SubscriptionItemsMap) {
	s.lock.Lock()
	defer s.lock.Unlock()

	for sub, items := range subInfos {
		if leaf := s.MapLeaves.Get(sub.ID); leaf != nil {
			leaf.ItemsCount += len(*items)
			leaf.NewItemsCount += len(*items)
		}
	}
}

// UnLoadAll ...
func (s *State) UnLoadAll() {
	s.lock.RLock()
	defer s.lock.RUnlock()
	for _, n := range s.MapNodes.Nodes {
		//except root
		if s.Root.ID != n.ID {
			n.UnLoadChildren()
		}
	}
}

// LoadNode ...
func (s *State) LoadNode(ID string) error {
	s.lock.RLock()
	defer s.lock.RUnlock()

	node := s.MapNodes.Get(ID)
	if node == nil {
		return fmt.Errorf("node not found: %s", ID)
	}
	node.LoadChildren()

	return nil
}

// UnLoadNode ...
func (s *State) UnLoadNode(ID string) error {
	s.lock.RLock()
	defer s.lock.RUnlock()

	node := s.MapNodes.Get(ID)
	if node == nil {
		return fmt.Errorf("node not found: %s", ID)
	}
	node.UnLoadChildren()

	return nil
}

// MoveNode ...
func (s *State) MoveNode(nodeID, parentID string) error {

	s.lock.Lock()
	defer s.lock.Unlock()

	node := s.MapNodes.Get(nodeID)
	if node == nil {
		return fmt.Errorf("node %s not found", nodeID)
	}

	//same parent
	if node.ParentID == parentID {
		return fmt.Errorf("same parent")
	}

	parent := s.MapNodes.Get(parentID)
	if parent == nil {
		return fmt.Errorf("parent %s not found", parentID)
	}

	//is descendent
	if node.IsDescendent(parent) {
		return fmt.Errorf("error moving. New parent '%s' is descendent of '%s'", parent.Label, node.Label)
	}

	oldParent := s.MapNodes.Get(node.ParentID)
	if oldParent == nil {
		return fmt.Errorf("old parent  %s not found", node.ParentID)
	}

	//remove from old parent
	oldParent.RemoveNode(node.ID)

	//add to new Parent
	parent.AddNode(node)

	// saving
	if err := s.Save(); err != nil {
		return err
	}

	return nil
}

//MoveLeaf ...
func (s *State) MoveLeaf(leafID, parentID string) error {

	s.lock.Lock()
	defer s.lock.Unlock()

	leaf := s.Root.FindLeafByID(leafID)
	if leaf == nil {
		return fmt.Errorf("leaf %s not found", leafID)
	}

	//same parent
	if leaf.ParentID == parentID {
		return fmt.Errorf("same parent")
	}

	parent := s.MapNodes.Get(parentID)
	if parent == nil {
		return fmt.Errorf("parent %s not found", parentID)
	}

	oldParent := s.MapNodes.Get(leaf.ParentID)
	if oldParent == nil {
		return fmt.Errorf("old parent  %s not found", leaf.ParentID)
	}

	//remove from old parent
	oldParent.RemoveLeaf(leaf.ID)
	//add to new Parent
	parent.AddLeaf(leaf)
	s.MapLeaves.Put(leaf)

	// saving
	if err := s.Save(); err != nil {
		return err
	}

	return nil
}

// AddNode ...
func (s *State) AddNode(node *models.Node, parentID string) error {

	s.lock.Lock()
	defer s.lock.Unlock()

	parent := s.MapNodes.Get(parentID)
	if parent == nil {
		return fmt.Errorf("parent: %s not found", parentID)
	}
	parent.AddNode(node)
	s.MapNodes.Put(node)

	// saving
	if err := s.Save(); err != nil {
		return err
	}

	return nil
}

// AddLeaf ...
func (s *State) AddLeaf(leaf *models.Leaf, parentID string) error {

	s.lock.Lock()
	defer s.lock.Unlock()

	parent := s.MapNodes.Get(parentID)
	if parent == nil {
		return fmt.Errorf("parent: %s not found", parentID)
	}
	parent.AddLeaf(leaf)
	s.MapLeaves.Put(leaf)

	// saving
	if err := s.Save(); err != nil {
		return err
	}

	return nil
}

// EditNode ...
func (s *State) EditNode(_node *models.Node) error {

	s.lock.Lock()
	defer s.lock.Unlock()

	node := s.MapNodes.Get(_node.ID)
	if node == nil {
		return fmt.Errorf("node: %s not found", _node.ID)
	}

	//label changed?
	if node.Label != _node.Label {
		node.Label = _node.Label

		//find parent
		if parent := s.MapNodes.Get(node.ParentID); parent != nil {
			parent.Nodes.Sort()
			parent.MemNodes.Sort()
		}
	}

	node.Color = _node.Color
	node.Label = _node.Label
	node.Icon = _node.Icon
	node.Description = _node.Description

	// saving
	if err := s.Save(); err != nil {
		return err
	}

	return nil
}

// RemoveLeaf ...
func (s *State) RemoveLeaf(id string) error {

	s.lock.Lock()
	defer s.lock.Unlock()

	leaf := s.MapLeaves.Get(id)
	if leaf == nil {
		return fmt.Errorf("leaf: %s not found", id)
	}

	parent := s.MapNodes.Get(leaf.ParentID)
	if parent == nil {
		return fmt.Errorf("parent of Leaf: %s not found", id)
	}

	parent.RemoveLeaf(id)
	s.MapLeaves.Remove(id)

	// saving
	if err := s.Save(); err != nil {
		return err
	}

	return nil
}

// EditLeaf ...
func (s *State) EditLeaf(_leaf *models.Leaf) error {

	s.lock.Lock()
	defer s.lock.Unlock()

	leaf := s.MapLeaves.Get(_leaf.ID)
	if leaf == nil {
		return fmt.Errorf("leaf: %s not found", _leaf.ID)
	}

	//label changed?
	if leaf.Label != _leaf.Label {
		leaf.Label = _leaf.Label

		//find parent
		if parent := s.MapNodes.Get(leaf.ParentID); parent != nil {
			parent.Leaves.Sort()
			parent.MemLeaves.Sort()
		}
	}

	leaf.Color = _leaf.Color
	leaf.Icon = _leaf.Icon
	leaf.Description = _leaf.Description

	// saving
	if err := s.Save(); err != nil {
		return err
	}

	return nil
}

//Save persists Root
func (s *State) Save() error {
	return s.persister.Save(*s.Root)
}
