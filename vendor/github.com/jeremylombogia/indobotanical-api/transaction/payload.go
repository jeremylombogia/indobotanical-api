package transaction

type Payload struct {
	Data struct {
		ProductID string `json:"productId"`
		Amount    int    `json:"amount"`
		PromoCode string `json:"promoCode"`
	} `json:"data"`
}
