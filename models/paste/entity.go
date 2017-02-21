package paste

import (
	"encoding/json"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"fareastdominions.com/evepaste/eve/market"
	"fareastdominions.com/evepaste/eve/entity"
	"google.golang.org/appengine/urlfetch"
	"fareastdominions.com/evepaste/utils"
	"github.com/mjibson/goon"
	"fareastdominions.com/evepaste/models/user"
)

type Paste struct {
	ID int64 `datastore:"-" goon:"id"`
	EncodedID string `datastore:"-"`
	Text string `datastore:",noindex"`
	Type string
	User user.User
	MarketSystemID int
	ScanSystemID int
	ItemCount int `datastore:",noindex"`
	TotalBuyPrice float64 `datastore:",noindex"`
	TotalSellPrice float64 `datastore:",noindex"`
	TotalVolume float64 `datastore:",noindex"`
	TotalQuantity int `datastore:",noindex"`
	Items []entity.Item `datastore:"-"`
	Prices []entity.Price `datastore:"-"`
	StoreItems []byte `datastore:"Items,noindex"`
	StorePrices []byte `datastore:"Prices,noindex"`
	PriceTable map[int32]entity.Price `datastore:"-"`
	IPAddr string
	Created time.Time
}

/**
 * Validates
 */
func (p *Paste) IsValid() bool {
	return p.Type != "" && p.Text != "" && len(p.Items) > 0
}

/**
 * Generate key
 */
func (p *Paste) key(c context.Context) *datastore.Key {
	if p.ID == 0 {
		p.Created = time.Now()
		return datastore.NewIncompleteKey(c, "Paste", nil)
	}
	return datastore.NewKey(c, "Paste", "", p.ID, nil)
}

/**
 * Save entity
 */
func (p *Paste) Save(c context.Context) (*Paste, error) {
	p.calcTotal()
	err := p.serialize()
	if err != nil {
		return nil, err
	}
	g := goon.FromContext(c)
	k, err := g.Put(p)
	if err != nil {
		return nil, err
	}
	p.ID = k.IntID()
	return p, nil
}

/**
 * Fetching Market Price
 */
func (p *Paste) FetchMarketPrice(c context.Context) {
	typeIds := p.GetUniqueTypeIds()

	p.Prices = market.GetMarketPrice(typeIds, p.MarketSystemID, c, &urlfetch.Transport{Context: c})
}

/**
 * Build unique TypeID List
 */
func (p *Paste) GetUniqueTypeIds() []int32 {
	sets := make(map[int32]struct{})
	for _, i := range p.Items {
		sets[i.TypeId] = struct{}{}
	}

	ids := make([]int32, 0)
	for i, _ := range sets {
		ids = append(ids, i)
	}
	return ids
}

/**
 * Calculating total
 */
func (p *Paste) calcTotal() {
	p.ItemCount = 0
	p.TotalQuantity = 0
	p.TotalVolume = 0
	p.TotalBuyPrice = 0
	p.TotalSellPrice = 0

	// build price table
	p.buildPriceTable()

	for _, item := range p.Items {
		quantity := float64(item.Quantity)
		p.TotalBuyPrice += item.Prices.Buy.Price * quantity
		p.TotalSellPrice += item.Prices.Sell.Price * quantity
		p.TotalQuantity += item.Quantity
		p.TotalVolume += item.Volume * quantity
		p.ItemCount++
	}
}

/**
 * Serialize entity
 */
func (p *Paste) serialize() error {
	// marshal items
	bytes, err := json.Marshal(p.Items)
	if err != nil {
		return err
	}
	p.StoreItems = bytes

	// marshal prices
	bytes, err = json.Marshal(p.Prices)
	if err != nil {
		return err
	}
	p.StorePrices = bytes

	return nil
}

/**
 * Unserialize entity
 */
func (p *Paste) unserialize() error {
	p.EncodedID = utils.EncodeBase62(p.ID)

	// unmarshal items
	p.Items = make([]entity.Item, 0)
	err := json.Unmarshal(p.StoreItems, &p.Items)
	if err != nil {
		return err
	}

	// unmarshal price
	p.Prices = make([]entity.Price, 0)
	err = json.Unmarshal(p.StorePrices, &p.Prices)
	if err != nil {
		return err
	}

	p.buildPriceTable()

	return nil
}

/**
 * Build Price Table
 */
func (p *Paste) buildPriceTable() {
	p.PriceTable = make(map[int32]entity.Price)
	for _, price := range p.Prices {
		p.PriceTable[price.TypeId] = price
	}

	// mapping to items
	for i, item := range p.Items {
		if _, ok := p.PriceTable[item.TypeId]; ok {
			p.Items[i].Prices = p.PriceTable[item.TypeId]
		}
	}
}

/**
 * Export EFT
 */
func (p *Paste) exportEFT(lang string) {
}

func (p *Paste) GetGroupedItems() ([]entity.ItemGroup) {
	groupMap := make(map[string]*entity.ItemGroup)
	items := make([]entity.Item, 0)

	for _, i := range p.Items {
		if i.ItemGroupKey != "" {
			g, ok := groupMap[i.ItemGroupKey]
			if !ok {
				g = entity.NewItemGroup(i.ItemGroupKey, i.ItemGroupName)
				groupMap[i.ItemGroupKey] = g
			}

			g.Items = append(g.Items, i)
			quantity := float64(i.Quantity)
			g.TotalBuyPrice += i.Prices.Buy.Price * quantity
			g.TotalSellPrice += i.Prices.Sell.Price * quantity
			g.TotalVolume += i.Volume * quantity
		} else {
			items = append(items, i)
		}
	}

	groups := make([]entity.ItemGroup, 0)
	for _, g := range groupMap {
		groups = append(groups, *g)
	}

	if len(items) > 0 {
		g := entity.NewItemGroup("other", "Other")
		for _, i := range items {
			g.Items = append(g.Items, i)
			quantity := float64(i.Quantity)
			g.TotalBuyPrice += i.Prices.Buy.Price * quantity
			g.TotalSellPrice += i.Prices.Sell.Price * quantity
			g.TotalVolume += i.Volume * quantity
		}
		groups = append(groups, *g)
	}

	return groups
}