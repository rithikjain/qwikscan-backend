package entities

import "github.com/jinzhu/gorm"

type Cart struct {
	gorm.Model
	UUID     string `json:"id"`
	CartName string `json:"cart_name"`
	UserID   string `json:"user_id"`
}

type CartItem struct {
	gorm.Model
	UUID         string `json:"id"`
	CartID       string `json:"cart_id"`
	ItemName     string `json:"item_name"`
	ItemPrice    int    `json:"item_price"`
	ItemQuantity int    `json:"item_quantity"`
	ItemImageUrl string `json:"item_image_url"`
}
