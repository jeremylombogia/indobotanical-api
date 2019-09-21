package transaction

type Payload struct {
	Data struct {
		Products []struct {
			ProductID string `json:"productId"`
			Amount    int    `json:"amount"`
		} `json:"products"`
		PromoCode string `json:"promoCode"`
	} `json:"data"`
}
