package store

import (
	"krolus/data"
	"krolus/feed"
	m "krolus/models"
	"krolus/treex"
	"krolus/treex/models"

	"github.com/mitchellh/mapstructure"
	"github.com/wailsapp/wails"
)

// TreeStore ...
type TreeStore struct {
	store      *wails.Store
	treeState  *treex.State
	logger     *wails.CustomLogger
	obs        feed.Observable
	aggregator *feed.Aggregator
	manager    *data.Manager
	isFlat     bool
}

// NewTreeStore ...
func NewTreeStore(agg *feed.Aggregator, man *data.Manager, treeState *treex.State, obs feed.Observable) ITreeStore {
	return &TreeStore{
		treeState:  treeState,
		obs:        obs,
		manager:    man,
		aggregator: agg,
	}
}

// WailsInit is called when the component is being initialised
func (t *TreeStore) WailsInit(runtime *wails.Runtime) error {
	t.logger = runtime.Log.New("TreeStore")
	t.store = runtime.Store.New("TreeState", &models.Node{})
	go t.listenSubNews()
	return nil
}

// listen for new items
func (t *TreeStore) listenSubNews() {

	for info := range t.obs {
		t.logger.Warnf("Receiving from observable: %v", info)
		// Update tree state
		t.treeState.UpdateCounters(info)

		if err := t.treeState.Save(); err != nil {
			t.logger.Errorf("Error saving treex: %v", err)
		}

		if t.isFlat {
			t.store.Set(t.treeState.GetFavorites())
		} else {
			t.store.Set(t.treeState.Root)
		}
	}

}

// LoadNode  ...
func (t *TreeStore) LoadNode(ID string) error {
	if ID == "" {
		return t.store.Set(t.treeState.Root)
	}
	if err := t.treeState.LoadNode(ID); err != nil {
		return err
	}
	return t.store.Set(t.treeState.Root)
}

// LoadAncestors ...
func (t *TreeStore) LoadAncestors(ID string, isLeaf bool) error {
	t.treeState.LoadAncestors(ID, isLeaf)
	return t.store.Set(t.treeState.Root)
}

// UnLoadNode  ...
func (t *TreeStore) UnLoadNode(ID string) error {
	if err := t.treeState.UnLoadNode(ID); err != nil {
		return err
	}
	return t.store.Set(t.treeState.Root)
}

// UnLoadAll ...
func (t *TreeStore) UnLoadAll() error {
	t.treeState.UnLoadAll()
	return t.store.Set(t.treeState.Root)
}

// FilterFavorites ...
func (t *TreeStore) FilterFavorites(enabled bool) error {
	t.isFlat = enabled
	if enabled {
		return t.store.Set(t.treeState.GetFavorites())
	}
	return t.store.Set(t.treeState.Root)
}

// ToggleLeafFavorite ...
func (t *TreeStore) ToggleLeafFavorite(leafID string) error {
	return t.treeState.TogleLeafFavorite(leafID)
}

// MoveLeaf ...
func (t *TreeStore) MoveLeaf(leafID, nodeID string) error {
	if err := t.treeState.MoveLeaf(leafID, nodeID); err != nil {
		return err
	}
	return t.store.Set(t.treeState.Root)
}

// MoveNode ...
func (t *TreeStore) MoveNode(nodeID, toID string) error {
	if err := t.treeState.MoveNode(nodeID, toID); err != nil {
		return err
	}
	return t.store.Set(t.treeState.Root)
}

// AddNode ....
func (t *TreeStore) AddNode(_node map[string]interface{}, toID string) error {
	node := models.NewNode(_node["label"].(string), _node["description"].(string))
	if err := t.treeState.AddNode(node, toID); err != nil {
		return err
	}
	return t.store.Set(t.treeState.Root)
}

// AddSubscription ....
func (t *TreeStore) AddSubscription(_sub map[string]interface{}, toID string) error {

	sub := &m.SubscriptionModel{}
	if err := mapstructure.Decode(_sub, sub); err != nil {
		return err
	}

	leaf := models.NewLeaf(sub.Title, sub.Description)
	sub.ID = leaf.ID
	if err := t.manager.Subscription.Add(sub); err != nil {
		return err
	}

	//TODO: Feed.selectedLeaves is outdated if new sub is in selected
	if err := t.aggregator.CheckSub(sub); err != nil {
		return err
	}

	if sub.Provider != "generic" {
		leaf.Icon = sub.Provider
	}
	if err := t.treeState.AddLeaf(leaf, toID); err != nil {
		return err
	}
	if err := t.store.Set(t.treeState.Root); err != nil {
		return err
	}

	return nil
}

// EditNode ...
func (t *TreeStore) EditNode(_node map[string]interface{}) error {

	//TODO: create function
	node := &models.Node{
		Leaf: &models.Leaf{
			ID:          _node["id"].(string),
			ParentID:    _node["parent"].(string),
			Label:       _node["label"].(string),
			Description: _node["description"].(string),
			Icon:        _node["icon"].(string),
			Color:       _node["color"].(string),
		},
		Nodes:  models.NodeCollection{},
		Leaves: models.LeafCollection{},
	}

	if err := t.treeState.EditNode(node); err != nil {
		return err
	}

	return t.store.Set(t.treeState.Root)
}

// RemoveSubscription ...
func (t *TreeStore) RemoveSubscription(ID string) error {

	//TODO: refresh feed pagination if sub in selectedLeaves

	if err := t.manager.Subscription.Remove(ID); err != nil {
		return err
	}

	if err := t.treeState.RemoveLeaf(ID); err != nil {
		return err
	}
	return t.store.Set(t.treeState.Root)
}

// EditSubscription ...
func (t *TreeStore) EditSubscription(_sub map[string]interface{}) error {

	//TODO: refresh feed pagination if sub in selectedLeaves

	sub := &m.SubscriptionModel{}
	if err := mapstructure.Decode(_sub, sub); err != nil {
		return err
	}

	if err := t.manager.Subscription.Update(sub); err != nil {
		return err
	}

	//TODO: fix this
	leaf := &models.Leaf{
		ID:          sub.ID,
		Label:       sub.Title,
		Description: sub.Description,
	}

	if err := t.treeState.EditLeaf(leaf); err != nil {
		return err
	}

	return t.store.Set(t.treeState.Root)
}
