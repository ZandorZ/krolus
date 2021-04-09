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

	if production {
		usr, err := user.Current()
		if err != nil {
			panic(err)
		}
		basePath = usr.HomeDir + "/.krolus"
	} else {
		// ex, err := os.Executable()
		// if err != nil {
		// 	panic(err)
		// }
		basePath = "./db"
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
	// defer bh.CloseDB()

	ob := feed.NewObserver()

	//if options.Tor {
	// torClient := feed.NewTorClient()
	// defer torClient.Close()
	//}

	httpClient := feed.NewGenericClient()
	defer httpClient.CloseIdleConnections()

	agg := feed.NewAggregator(
		feed.NewParser(
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
	treeState, err := treex.NewState(models.NewNode(".", "."), filePersist)
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
