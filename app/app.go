package app

import (
	"krolus/data/sqte"
	"krolus/feed"
	"krolus/media"
	"krolus/store"
	"krolus/treex"
	"krolus/treex/models"
	"krolus/treex/persistence"
	"os"
	"os/user"

	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/lib/logger"
)

func GetPath(production bool) string {

	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	var basePath string

	if production {
		basePath = usr.HomeDir + "/.krolus"
	} else {
		basePath = os.TempDir() + "/.krolus"
	}

	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		err := os.Mkdir(basePath, os.FileMode(0755))
		if err != nil {
			panic(err)
		}
	}

	return basePath
}

type KrolusApp struct {
	options Options
}

func (k *KrolusApp) Start(options Options) {

	k.options = options
	basePath := GetPath(options.Production)
	man := sqte.NewManager(basePath+"/mine.db", !options.Production)
	ob := feed.NewObserver()
	myLogger := logger.NewCustomLogger("Krolus")

	//if options.Tor {
	// torClient := feed.NewTorClient()
	// defer torClient.Close()
	//}
	httpClient := feed.NewGenericClient()
	defer httpClient.CloseIdleConnections()

	req := feed.NewRequester(httpClient)

	agg := feed.NewAggregator(
		feed.NewChecker(req),
		options.Interval,
		options.Workers,
		ob,
		man,
		myLogger,
	)
	agg.Start(options.CheckAtStart)

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
		DisableInspector: options.Production,
	})
	appW.Bind(store.NewMediaStore(man, media.NewDownloader(req)))
	appW.Bind(store.NewItemStore(man, treeState))
	appW.Bind(store.NewTreeStore(agg, man, treeState, ob.Add("tree")))
	appW.Bind(store.NewFeedStore(man, treeState, ob.Add("feed"), req))
	if err := appW.Run(); err != nil {
		myLogger.Fatalf("App error: %e", err)
	}

}
