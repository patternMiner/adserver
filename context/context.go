/**
 * Created by jbisa on 8/20/17.
 */

package context

import (
	"github.com/patternMiner/async"
	"sync"
)

const (
	ads_data = "data/ads.csv"
	adunit_ads_data = "data/adunit_ads.csv"
)

var (
	data_files = []string {ads_data, adunit_ads_data}

	// dictionary of ads by id
	AdsMap = make(StringSliceMap)

	// dictionary of ads by adunit
	AdUnitAdsMap = make(StringSetMap)

	// data fetcher task synchronization lock
	wg sync.WaitGroup
)

// Initializes the context by fetching all data records into various maps.
func InitContext() error {
	async.StartDispatcher(2)
	wg.Add(len(data_files))
	for _, path := range data_files {
		async.TaskQueue <- DataFetcherTask{path}
	}
	wg.Wait()
	return nil
}
