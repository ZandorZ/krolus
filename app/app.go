package app

import (
	"krolus/data/sqte"
	"krolus/feed"
	"krolus/store"
	"krolus/treex"
	"krolus/treex/models"
	"krolus/treex/persistence"
	"os/user"

	"github.com/wailsapp/wails"
)

func GetPath(production bool) string {

	var basePath string

	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	if production {
		basePath = usr.HomeDir + "/.krolus"
	} else {
		basePath = usr.HomeDir + "/Projects/krolus/db"
	}

	return basePath
}

type KrolusApp struct {
	options Options
}

func (k *KrolusApp) Start(options Options) {

	k.options = options
	basePath := GetPath(options.Production)
	man := sqte.NewManager(basePath + "/mine.db")
	ob := feed.NewObserver()

	//if options.Tor {
	// torClient := feed.NewTorClient()
	// defer torClient.Close()
	//}
	httpClient := feed.NewGenericClient()
	defer httpClient.CloseIdleConnections()

	agg := feed.NewAggregator(
		feed.NewChecker(
			feed.NewRequester(httpClient),
		),
		options.Interval,
		options.Workers,
		ob,
		man,
		&wails.CustomLogger{},
	)
	agg.Start(true)

	// treex
	filePersist, err := persistence.NewFile(basePath + "/tree.x_")
	if err != nil {
		panic(err)
	}
	treeState, err := treex.NewState(models.NewNode("Root", "Root folder"), filePersist)
	if err != nil {
		panic(err)
	}

	appW := wails.CreateApp(&wails.AppConfig{
		Width:            1024,
		Height:           768,
		Title:            "krolus",
		JS:               options.Js,
		CSS:              options.Css,
		Resizable:        true,
		DisableInspector: false,
	})
	appW.Bind(store.NewMediaStore(man))
	appW.Bind(store.NewItemStore(man, treeState))
	appW.Bind(store.NewTreeStore(agg, man, treeState, ob.Add("tree")))
	appW.Bind(store.NewFeedStore(man, treeState, ob.Add("feed")))
	if err := appW.Run(); err != nil {
		panic(err)
	}

}
