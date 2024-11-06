package gacha

import (
	"math/rand"
	"time"
)

type Item struct {
	ID     string
	Name   string
	Weight int64
}

func Execute(items []*Item, n int) ([]*Item, error) {
	result := make([]*Item, 0, n)
	for i := 0; i < n; i++ {
		item := weightedRandom(items)
		if item != nil {
			result = append(result, item)
		}
	}

	return result, nil
}

func weightedRandom(items []*Item) *Item {
	if len(items) == 0 {
		return nil
	}

	var totalWeight int64
	for _, item := range items {
		totalWeight += item.Weight
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(totalWeight)
	var weightSum int64
	for _, item := range items {
		weightSum += item.Weight
		if random < weightSum {
			return item
		}
	}

	return nil
}
