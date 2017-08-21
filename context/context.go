/**
 * Created by jbisa on 8/20/17.
 */

package context

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
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

// Initializes the context by fetching all data records into various maps.
func InitContext() error {
	for _, path := range data_files {
		records, err := fetch(path)
		if err != nil {
			return err
		}
		switch path {
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
	return nil
}

// Returns the adSet for the given adUnit
func AdSetByAdUnit(adUnit string) StringSet {
	return AdUnitAdsMap[adUnit]
}
