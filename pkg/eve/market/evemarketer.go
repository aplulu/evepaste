package market

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/evepaste/evepaste/pkg/eve/entity"
	"golang.org/x/net/context"
	"google.golang.org/appengine/memcache"
)

type EVEMarketerResult struct {
	Buy  EVEMarketerPriceItem `json:"buy"`
	Sell EVEMarketerPriceItem `json:"sell"`
}

func (r *EVEMarketerResult) TypeId() int32 {
	return r.Sell.ForQuery.Types[0]
}

func (r *EVEMarketerResult) ToPrice(isUniverse bool) entity.Price {
	return entity.Price{
		TypeId: r.TypeId(),
		Buy: r.Buy.ToPriceItem(isUniverse, "buy"),
		Sell: r.Sell.ToPriceItem(isUniverse, "sell"),
	}
}


type EVEMarketerPriceItem struct {
	ForQuery EVEMarketerForQuery `json:"forQuery"`
	WAvg float64 `json:"wavg"`
	Avg float64 `json:"avg"`
	Median float64 `json:"median"`
	Max float64 `json:"max"`
	Min float64 `json:"min"`
	FivePercent float64 `json:"fivePercent"`
}

func (i *EVEMarketerPriceItem) ToPriceItem(isUniverse bool, t string) entity.PriceItem {
	price := i.FivePercent
	if !isUniverse {
		if t == "buy" {
			price = i.Max
		} else if t == "sell" {
			price = i.Min
		}
	}

	return entity.PriceItem{
		Avg: i.Avg,
		Median: i.Median,
		Max: i.Max,
		Min: i.Min,
		Price: price,
	}
}

type EVEMarketerForQuery struct {
	Types []int32 `json:"types"`
}


/**
 * EVEMarketer Market Price Provider
 */
type EVEMarketer struct {
	Transport http.RoundTripper
	Context context.Context
}

func (s *EVEMarketer) GetMarketPrice(typeIds []int32, systemId int) (map[int32]entity.Price, error) {
	prices := make(map[int32]entity.Price)
	client := &http.Client{
		Transport: s.Transport,
	}

	batchSize := 190
	l := len(typeIds)
	for i := 0; i < l; i += batchSize {
		params := []string{}
		end := i + batchSize
		if end > l {
			end = l
		}

		for j := i; j < end; j++ {
			params = append(params, "typeid=" + strconv.Itoa(int(typeIds[j])))
		}

		if systemId > 0 {
			params = append(params, "usesystem=" + strconv.Itoa(systemId))
		}


		u := "https://api.evemarketer.com/ec/marketstat/json?" + strings.Join(params, "&")

		req, err := http.NewRequest("GET", u, nil)
		if err != nil {
			return prices, err
		}
		req.Header.Set("User-Agent", "EVEPaste/1.0 (aplulu@fareastdominions.com)")
		resp, err := client.Do(req)
		if err != nil {
			return prices, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return prices, err
		}

		results := make([]EVEMarketerResult, 0)
		json.Unmarshal(body, &results)

		for _, r := range results {
			price := r.ToPrice(systemId < 1)
			prices[int32(r.TypeId())] = price

			cache, err := json.Marshal(price)
			if err == nil {
				item := &memcache.Item{
					Key: fmt.Sprintf("evemarketer_%d_%d", r.TypeId(), systemId),
					Value: cache,
					Expiration: 24 * time.Hour,
				}
				memcache.Add(s.Context, item)
			}
		}
	}

	return prices, nil
}

/**
 * EVEMarketer Cached Provider
 */
type EVEMarketerMemcache struct {
	Context context.Context
}

func (s *EVEMarketerMemcache) GetMarketPrice(typeIds []int32, systemId int) (map[int32]entity.Price, error) {
	prices := make(map[int32]entity.Price)

	for _, typeID := range typeIds {
		item, err := memcache.Get(s.Context, fmt.Sprintf("evemarketer_%d_%d", typeID, systemId))
		if err != nil {
			continue
		}

		p := entity.Price{}
		err = json.Unmarshal(item.Value, &p)
		if err != nil {
			continue
		}

		prices[typeID] = p
	}

	return prices, nil
}