package handler

type AddItemToCartRequest struct {
	UserID   string `json:"userId"`
	ItemID   string `json:"itemId"`
	Quantity int    `json:"quantity"`
}
