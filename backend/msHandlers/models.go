package msHandlers

type PricingPerItem struct {
	PriceID   string `json:"price_id"`
	ItemPrice int    `json:"item_price"`
	PriceName string `json:"price_name"`
}
