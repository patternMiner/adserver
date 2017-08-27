package context

import (
	"path/filepath"
	"fmt"
	"os"
	"encoding/csv"
	"bufio"
	"log"
)

type DataFetcherTask struct {
	path string
}

func (t DataFetcherTask) Run() {
	defer func() {
		log.Printf("Done fetching: %s\n", t.path)
		wg.Done()
	}()
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
