package context

// Returns the adSet for the given adUnit
func AdSetByAdUnit(adUnit string) StringSet {
	return AdUnitAdsMap[adUnit]
}

