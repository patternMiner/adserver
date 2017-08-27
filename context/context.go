/**
 * Created by jbisa on 8/20/17.
 */

package context

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
)

// Fetches data records from the given file path
func fetch(path string) (records [][]string, err error) {
	fp, _ := filepath.Abs(path)
	fmt.Printf("%s: %s\n", path, fp)
	fh, err := os.Open(fp)
	if err != nil {
		return
	}
	defer fh.Close()
	data := csv.NewReader(bufio.NewReader(fh))
	records, err = data.ReadAll()
	return
}

type DataFetcherTask struct {
	path string
}

var wg sync.WaitGroup

func (t DataFetcherTask) Run() {
	defer wg.Done()
	records, err := fetch(t.path)
	if err != nil {
		log.Printf("Error fetching %s: %s\n", t.path, err)
		return
	}
	switch t.path {
	case ads_data:
		for _, record := range records[1:] {
			id := record[0]
			AdsMap[id] = record
		}
		break
	case adunit_ads_data:
		for _, record := range records[1:] {
			adunit_id := record[0]
			ad_id := record[1]
			AdUnitAdsMap.stringSet(adunit_id).append(ad_id)
		}
		break
	}
}

// Initializes the context by fetching all data records into various maps.
func InitContext() error {
	async.StartTaskDispatcher(2)
	wg.Add(len(data_files))
	for _, path := range data_files {
		async.TaskQueue <- DataFetcherTask{path}
	}
	wg.Wait()
	return nil
}

// Returns the adSet for the given adUnit
func AdSetByAdUnit(adUnit string) StringSet {
	return AdUnitAdsMap[adUnit]
}
