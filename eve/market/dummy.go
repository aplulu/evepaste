package market

import (
	"fareastdominions.com/evepaste/eve/entity"
	"fareastdominions.com/evepaste/eve/repository"
)

type Dummy struct {

}

func (s *Dummy) GetMarketPrice(typeIds []int32, systemId int) (map[int32]entity.Price, error) {
	prices := make(map[int32]entity.Price)

	for _, id := range typeIds {
		t := repository.GetTypeByID(id)
		if t != nil && t.MarketGroupId != nil &&  *t.MarketGroupId == 0 {
			prices[id] = entity.Price{}
		}
	}

	return prices, nil
}
