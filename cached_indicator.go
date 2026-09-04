package techan

import "github.com/sdcoffey/big"

type resultCache []*big.Decimal

type cachedIndicator interface {
	Indicator
	cache() resultCache
	setCache(cache resultCache)
}

type indicatorDependencies interface {
	dependencies() []Indicator
}

// CacheResetter describes indicators whose cached values can be invalidated after
// their underlying data changes.
type CacheResetter interface {
	ResetCacheFrom(index int)
}

// ResetCacheFrom invalidates cached values from index onward when the indicator
// supports cache resets or uses the built-in result cache, including inputs of
// built-in composites. It returns true if any cache was reset. Reset each root
// used by the caller; parents outside that root's input graph are not visited.
func ResetCacheFrom(indicator Indicator, index int) bool {
	if resetter, ok := indicator.(CacheResetter); ok {
		resetter.ResetCacheFrom(index)
		return true
	}

	reset := false
	if cached, ok := indicator.(cachedIndicator); ok {
		resetResultCache(cached, index)
		reset = true
	}
	if composite, ok := indicator.(indicatorDependencies); ok {
		for _, input := range composite.dependencies() {
			reset = ResetCacheFrom(input, index) || reset
		}
	}
	return reset
}

func cacheResult(indicator cachedIndicator, index int, val big.Decimal) {
	if index < len(indicator.cache()) {
		indicator.cache()[index] = &val
	} else if index == len(indicator.cache()) {
		indicator.setCache(append(indicator.cache(), &val))
	} else {
		expandResultCache(indicator, index+1)
		cacheResult(indicator, index, val)
	}
}

func expandResultCache(indicator cachedIndicator, newSize int) {
	sizeDiff := newSize - len(indicator.cache())

	expansion := make([]*big.Decimal, sizeDiff)
	indicator.setCache(append(indicator.cache(), expansion...))
}

func resetResultCache(indicator cachedIndicator, index int) {
	cache := indicator.cache()
	clear(cache[min(max(index, 0), len(cache)):])
}
