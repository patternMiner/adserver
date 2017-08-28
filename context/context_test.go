package context

import (
	"testing"
)

func TestContext(t *testing.T) {
	InitContext()
	adSet := AdSetByAdUnit("ad-unit-1")
	adCount := len(adSet)
	if adCount != 1 {
		t.Errorf("Expected adSet of length %d, but go %d", 1, adCount)
	}
	var record StringSlice
	for adId := range adSet {
		record = AdsMap[adId]
	}
	expectedUrl := "https://www.youtube.com/embed/B9XveQRpoeg"
	if record[1] !=  expectedUrl {
		t.Errorf("Expected url %s, but go %s", expectedUrl, record[1])
	}
}
