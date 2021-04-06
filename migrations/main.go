package main

import (
	"fmt"
	"krolus/app"
	"krolus/data"
	"krolus/data/sqte"
	"krolus/models"
	"krolus/treex"
	treexModels "krolus/treex/models"
	"krolus/treex/persistence"
	"os"
	"path"

	"github.com/gilliek/go-opml/opml"
	"github.com/hokaccha/go-prettyjson"
)

var manager *data.Manager
var basePath string
var treeState *treex.State
var filePersist persistence.Persister

func init() {
	basePath = app.GetPath(false)
	manager = sqte.NewManager(basePath + "/dev.db")

	var err error
	filePersist, err = persistence.NewFile(basePath + "/tree.x_")
	if err != nil {
		panic(err)
	}
	treeState, err = treex.NewState(treexModels.NewNode("Subscriptions", "My Subscriptions"), filePersist)
	if err != nil {
		panic(err)
	}
}

func pretty(data interface{}) {
	s, _ := prettyjson.Marshal(data)
	fmt.Println(string(s))
}

func traverseOPML(outlines []opml.Outline, parent *treexModels.Node) {

	for _, outline := range outlines {
		//folder
		if outline.XMLURL == "" {
			node := treexModels.NewNode(outline.Title, outline.Text)
			parent.AddNode(node)
			if len(outline.Outlines) > 0 {
				traverseOPML(outline.Outlines, node)
			}
		} else {
			leaf := treexModels.NewLeaf(outline.Title, outline.Text)
			parent.AddLeaf(leaf)
			if err := manager.Subscription.Add(&models.SubscriptionModel{
				ID:          leaf.ID,
				Title:       leaf.Label,
				XURL:        outline.XMLURL,
				Description: outline.Text,
				URL:         outline.HTMLURL,
			}); err != nil {
				panic(err)
			}
		}
	}
}

func traverseTreex(doc *opml.Outline, parent *treexModels.Node) {
	doc.Title = parent.Label
	doc.Text = parent.Description
	treeState.LoadNode(parent.ID)

	doc.Outlines = make([]opml.Outline, len(parent.Nodes)+len(parent.Leaves))
	for i, node := range parent.Nodes {
		traverseTreex(&doc.Outlines[i], node)
	}

	start := len(parent.Nodes)
	for i, leaf := range parent.Leaves {
		doc.Outlines[start+i].Title = leaf.Label
		doc.Outlines[start+i].Text = leaf.Description
		sub, err := manager.Subscription.Get(leaf.ID)
		if err != nil {
			panic(err)
		}
		doc.Outlines[start+i].HTMLURL = sub.URL
		doc.Outlines[start+i].XMLURL = sub.XURL
		doc.Outlines[start+i].Type = "rss"
	}
}

func exportOPML(fileName string) error {
	doc := &opml.OPML{
		Version: "2.0",
	}
	doc.Body.Outlines = make([]opml.Outline, 1)

	traverseTreex(&doc.Body.Outlines[0], treeState.Root)

	xml, err := doc.XML()
	if err != nil {
		return err
	}

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(xml)
	if err != nil {
		return err
	}
	return nil
}

func importOPML(file string) error {
	doc, err := opml.NewOPMLFromFile(file)
	if err != nil {
		return err
	}
	traverseOPML(doc.Outlines(), treeState.Root)

	return filePersist.Save(*treeState.Root)
}

// main function
func main() {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	// if err := exportOPML(path.Dir(ex) + "/mine.xml"); err != nil {
	// 	panic(err)
	// }

	if err := importOPML(path.Dir(ex) + "/mine.xml"); err != nil {
		panic(err)
	}

}
