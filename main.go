package main

import (
	db "krolus/data/bh"
	"krolus/feed"
	"krolus/store"
	"krolus/treex"
	"krolus/treex/models"
	"krolus/treex/persistence"
	"os/user"
	"time"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

const (
	production = false
	numWorkers = 3
	interval   = 30 * time.Minute
)

var pathDB = ".krolus"

func init() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	//TODO: Fix this, use system paths
	pathDB = usr.HomeDir + "/" + pathDB
	if !production {
		pathDB += "/dev"
	} else {
		pathDB += "/db"
	}
}

func main() {

	man := db.NewManager(pathDB)
	defer db.CloseDB()

	ob := feed.NewObserver()

	// torClient := feed.NewTorClient()
	// defer torClient.Close()

	agg := feed.NewAggregator(
		feed.NewParser(
			feed.NewRequester(feed.NewGenericClient()),
		),
		interval,
		numWorkers,
		ob,
		man,
		&wails.CustomLogger{},
	)
	agg.Start(true)

	// treex
	filePersist, err := persistence.NewFile(pathDB + "/tree.x_")
	if err != nil {
		panic(err)
	}
	treeState, err := treex.NewState(models.NewNode(".", "."), filePersist)
	if err != nil {
		panic(err)
	}

	js := mewn.String("./frontend/dist/frontend/main.js")
	css := mewn.String("./frontend/dist/frontend/styles.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:            1024,
		Height:           768,
		Title:            "krolus",
		JS:               js,
		CSS:              css,
		Resizable:        true,
		DisableInspector: false,
	})
	app.Bind(store.NewMediaStore(man))
	app.Bind(store.NewItemStore(man, treeState))
	app.Bind(store.NewTreeStore(agg, man, treeState, ob.Add("tree")))
	app.Bind(store.NewFeedStore(man, treeState, ob.Add("feed")))
	if err := app.Run(); err != nil {
		panic(err)
	}
}
