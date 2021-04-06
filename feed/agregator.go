package feed

import (
	"krolus/data"
	"krolus/models"
	"time"

	"github.com/wailsapp/wails/lib/logger"
)

// Aggregator ...
type Aggregator struct {
	*Looper
	*Pool

	feedParser Parser
	ob         Observer
	dataMng    *data.Manager
	logger     *logger.CustomLogger
	subInfos   *SubInfos

	interval   time.Duration
	maxWorkers int
	batchSize  int
}

// NewAggregator ...
func NewAggregator(parser Parser, interval time.Duration, maxWorkers int, ob Observer, man *data.Manager, logger *logger.CustomLogger) *Aggregator {

	// TODO: create options?
	agg := &Aggregator{
		interval:   interval,
		batchSize:  maxWorkers * 4,
		feedParser: parser,
		maxWorkers: maxWorkers,
		ob:         ob,
		dataMng:    man,
		logger:     logger,
		subInfos:   NewSubInfos(),
	}

	agg.Looper = NewLooper(interval, agg.eachLoop)
	agg.Pool = NewPool(maxWorkers)

	return agg
}

// eachLoop ...
func (a *Aggregator) eachLoop(now time.Time) {
	a.logger.Warnf("Starting loop at %s", now.Format(time.RFC822))

	if err := a.dataMng.Subscription.ForEachOlderThan(a.interval, a.eachSub); err != nil {
		a.logger.Errorf("Error loading subs: %v", err)
	}

	// rest
	if a.Pool.Size() > 0 {
		a.Pool.Run()
		a.saveBatchItems(nil)
	}

	a.logger.Warnf("Finish loop at %s", now.Format(time.RFC822))
}

// CheckSub ...
func (a *Aggregator) CheckSub(sub *models.SubscriptionModel) error {

	//TODO: fix this /////////////////////////////////////////////////////////////////////////
	a.logger.Infof("Checking: %s, since: %s", sub.Title, sub.LastUpdate)
	items, err := a.feedParser.Parse(sub)
	if err != nil {
		a.logger.Errorf("Error parsing sub: %s, %v", sub.Title, err)
	}

	a.logger.Debugf("Total items: %d", len(items))
	if len(items) > 0 {
		a.logger.Infof("%d new items in %s, since %s", len(items), sub.Title, sub.LastUpdate)
		a.subInfos.Put(sub, &items)
	}

	a.saveBatchItems(nil)
	/////////////////////////////////////////////////////////////////////////////////////////

	return nil
}

// eachSub ...
func (a *Aggregator) eachSub(sub *models.SubscriptionModel, tx interface{}) error {

	a.Pool.Add(func() {
		a.logger.Infof("Checking: %s, since: %s", sub.Title, sub.LastUpdate)
		items, err := a.feedParser.Parse(sub)
		if err != nil {
			a.logger.Errorf("Error parsing sub: %s, %v", sub.Title, err)
		}
		a.logger.Debugf("Total items: %d", len(items))
		if len(items) > 0 {
			a.logger.Infof("%d new items in %s, since %s", len(items), sub.Title, sub.LastUpdate)
			a.subInfos.Put(sub, &items)
		}
	})

	// group by batchsize
	if a.Pool.Size()%a.batchSize == 0 {
		a.logger.Infof("Batching: ")
		a.Pool.Run()
		a.saveBatchItems(tx)
	}

	return nil
}

// saveBatchItems
func (a *Aggregator) saveBatchItems(tx interface{}) {

	if a.subInfos.Len() > 0 {
		a.logger.Debugf("Adding batch ...")
		err := a.dataMng.Item.AddInBatch(a.subInfos.Infos(), tx)
		if err != nil {
			a.logger.Errorf("Error adding items into batch: %v", err)
		}
		if err == nil {
			a.ob.Publish(a.subInfos.Infos())
			a.subInfos.Reset()
		}
	}
}

// Start ....
func (a *Aggregator) Start(flag bool) {
	go a.Looper.Start(flag)
}
