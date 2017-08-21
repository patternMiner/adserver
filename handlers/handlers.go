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
	Creative string
}

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

func InfoHandler(w http.ResponseWriter, req *http.Request) {
	setAccessControlResponseHeaders(w, req)

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

// Set access control response headers, only when the origin is known.
func setAccessControlResponseHeaders (w http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
}
