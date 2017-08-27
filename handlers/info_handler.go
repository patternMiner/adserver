package handlers

import (
	"net/http"
	"fmt"
	"github.com/patternMiner/adserver/context"
)

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

