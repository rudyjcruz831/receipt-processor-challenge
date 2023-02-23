package model

type Receipt struct {
	ReceiptID    string `json:"receipt_id"`
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchase_date"`
	PurchaseTime string `json:"purchase_time"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ItemID           string `json:"item_id"`
	ShortDescription string `json:"short_description"`
	Price            string `json:"price"`
}
