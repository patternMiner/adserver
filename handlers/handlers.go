/**
 * Created by jbisa on 8/20/17.
 */

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/patternMiner/adserver/context"
)

type Response struct {
	Items interface{}
	Err error
}

type Ad struct {
	Id string
	Url string
	Width string
	Height string
	Description string
}

func DefaultHandler (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintln(w, "/info\n\tShow context\n")
	fmt.Fprintln(w, "/ad?adunit=ad-unit-1&width=560&height=315")
}

func AdHandler (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
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

func InfoHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintln(w, "Ads")
	for id, ad := range context.AdsMap {
		fmt.Fprintf(w, "%2s: %s\n", id, ad)
	}
	fmt.Fprintln(w)

	fmt.Fprintln(w, "Ads by adUnit")
	for adUnit, ads := range context.AdUnitAdsMap {
		fmt.Fprintf(w, "%s: %s\n", adUnit, ads)
	}
	fmt.Fprintln(w)
}
