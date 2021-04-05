package main

import (
	"fmt"
	"krolus/data"
	"krolus/data/sqte"
	_models "krolus/models"
	"krolus/treex/models"
	"krolus/treex/persistence"
	"log"
	"os"
	"path"

	"github.com/gilliek/go-opml/opml"
	"github.com/hokaccha/go-prettyjson"
)

var manager *data.Manager

var pathDB = "./db"

func init() {
	manager = sqte.NewManager(pathDB + "/dev.db")
}

func pretty(data interface{}) {
	s, _ := prettyjson.Marshal(data)
	fmt.Println(string(s))
}

func traverse(outlines []opml.Outline, parent *models.Node) {

	for _, outline := range outlines {
		//folder
		if outline.XMLURL == "" {
			node := models.NewNode(outline.Title, outline.Text)
			parent.AddNode(node)
			if len(outline.Outlines) > 0 {
				traverse(outline.Outlines, node)
			}
		} else {
			leaf := models.NewLeaf(outline.Title, outline.Text)
			parent.AddLeaf(leaf)
			if err := manager.Subscription.Add(&_models.SubscriptionModel{
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

func importFile(name string, parent *models.Node) error {

	doc, err := opml.NewOPMLFromFile(name)
	if err != nil {
		log.Fatal(err)
	}

	traverse(doc.Outlines(), parent)

	return nil
}

// main function
func main() {

	// defer db.CloseDB()

	filePersist, err := persistence.NewFile(pathDB + "/tree.x_")
	if err != nil {
		panic(err)
	}

	root := models.NewNode("Subscriptions", "My subscriptions")

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	importFile(path.Dir(ex)+"/test.xml", root)

	if err := filePersist.Save(*root); err != nil {
		panic(err)
	}

	pretty(root)
}
