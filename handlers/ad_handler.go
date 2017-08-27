package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/patternMiner/adserver/context"
)

// Return the matching ads based on the requested adunit, size, and targeting parameters
func AdHandler (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	setAccessControlResponseHeaders(w, req)

	query := req.URL.Query()
	adUnit := query["adunit"]
	adSet := context.AdSetByAdUnit(adUnit[0])
	var ads []Ad
	for adId := range adSet {
		record := context.AdsMap[adId]
		ads = append(ads, Ad{Id: record[0], Url: record[1], Width: record[2], Height: record[3],
			Description: record[4]})
	}

	data, err := json.Marshal(Response{Items: ads})
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, string(data))
}
