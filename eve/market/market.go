package market

import (
	"net/http"
	"fareastdominions.com/evepaste/eve/entity"
	"golang.org/x/net/context"
)

type MarketPriceProvider interface {
	GetMarketPrice(typeIds []int32, systemId int) (map[int32]entity.Price, error)
}

/**
 * Get Market prices from Providers
 */
func GetMarketPrice(typeIds []int32, systemId int, c context.Context, transport http.RoundTripper) []entity.Price {
	results := make([]entity.Price, 0)
	prices := make(map[int32]entity.Price)

	providers := []MarketPriceProvider{
		&Dummy{},
		&EVECentralMemcache{
			Context: c,
		},
		&EVECentral{
			Context: c,
			Transport: transport,
		},
	}

	for _, provider := range providers {
		if len(typeIds) == 0 {
			break
		}

		processedPrices, err := provider.GetMarketPrice(typeIds, systemId)
		if err != nil {
			continue
		}

		typeIds = removeProcessedTypeIds(typeIds, processedPrices)
		for id, price := range processedPrices {
			prices[id] = price
		}
	}

	for _, p := range prices {
		results = append(results, p)
	}

	return results
}

/**
 * Remove processed Type IDs
 */
func removeProcessedTypeIds(typeIds []int32, processed map[int32]entity.Price) []int32 {
	results := make([]int32, 0)

	for _, id := range typeIds {
		if _, ok := processed[id]; !ok {
			results = append(results, id)
		}
	}

	return results
}