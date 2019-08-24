package transaction

type Payload struct {
	Data struct {
		ProductID string `json:"productId"`
		Amount    int    `json:"amount"`
	} `json:"data"`
}
