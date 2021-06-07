package main

import (
	"flag"
	"fmt"
	"krolus/app"
	"krolus/data"
	"krolus/data/sqte"
	"krolus/feed/providers"
	"krolus/models"
	"krolus/treex"
	treexModels "krolus/treex/models"
	"krolus/treex/persistence"
	"log"
	"os"
	"path"

	"github.com/gilliek/go-opml/opml"
	"github.com/hokaccha/go-prettyjson"
)

var manager *data.Manager
var basePath string
var treeState *treex.State
var filePersist persistence.Persister
var opmlFile string
var production bool
var export bool
var proxy *providers.Proxy

func init() {

	flag.StringVar(&opmlFile, "file", "test.xml", "opml filename")
	flag.BoolVar(&production, "production", false, "production mode")
	flag.BoolVar(&export, "export", false, "export mode")
	flag.Parse()

	basePath = app.GetPath(production)
	manager = sqte.NewManager(basePath + "/mine.db")

	var err error
	filePersist, err = persistence.NewFile(basePath + "/tree.x_")
	if err != nil {
		panic(err)
	}
	treeState, err = treex.NewState(treexModels.NewNode("Subscriptions", "My Subscriptions"), filePersist)
	if err != nil {
		panic(err)
	}

	proxy = providers.NewProxy()
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
			sub := &models.SubscriptionModel{
				ID:          leaf.ID,
				Title:       leaf.Label,
				XURL:        outline.XMLURL,
				Description: outline.Text,
				URL:         outline.HTMLURL,
				Provider:    proxy.GetRegister().GetRegisterByURL(outline.XMLURL).Name,
			}
			if err := manager.Subscription.Add(sub); err != nil {
				panic(err)
			}
			if sub.Provider != "generic" {
				leaf.Icon = sub.Provider
			}
			parent.AddLeaf(leaf)
		}
	}
}

func traverseTreex(outlines *[]opml.Outline, parent *treexModels.Node) {

	treeState.LoadNode(parent.ID)

	tempOuts := make([]opml.Outline, len(parent.Nodes)+len(parent.Leaves))

	for i, leaf := range parent.Leaves {
		tempOuts[i].Title = leaf.Label
		tempOuts[i].Text = leaf.Description
		sub, err := manager.Subscription.Get(leaf.ID)
		if err != nil {
			log.Printf("%s", err)
		}
		tempOuts[i].HTMLURL = sub.URL
		tempOuts[i].XMLURL = sub.XURL
		tempOuts[i].Type = "rss"
	}
	start := len(parent.Leaves)

	for i, node := range parent.Nodes {
		tempOuts[start+i].Title = node.Label
		tempOuts[start+i].Text = node.Description
		traverseTreex(&tempOuts[start+i].Outlines, node)
	}
	*outlines = tempOuts
}

func exportOPML(fileName string) error {
	doc := &opml.OPML{
		Version: "1.0",
	}
	doc.Head.Title = treeState.Root.Label

	traverseTreex(&doc.Body.Outlines, treeState.Root)

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
	xmlFile := path.Dir(ex) + "/" + opmlFile

	if export {
		if err := exportOPML(xmlFile); err != nil {
			panic(err)
		}

	} else {
		if err := importOPML(xmlFile); err != nil {
			panic(err)
		}
	}

}
