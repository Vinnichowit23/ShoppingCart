package model1

type Users struct {
	User_id    int     `json:"user_id"`
	Username   string  `json:"username"`
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	Created_at []uint8 `json:"created_at"`
}

// type Products struct {
// 	product_id     int    `json:"product_id"`
// 	name           string `json:"name"`
// 	price          int    `json:"price"`
// 	stock_quantity int    `json:"stock_quantity"`
// 	created_at     int    `json:"created_at"`
// 	updated_at     int    `json:"updated_at"`
// }

// type Shoppingcart struct {
// 	cart_id    int `json:"cart_id"`
// 	user_id    int `json:"user_id"`
// 	created_at int `json:"created_at"`
// }
// type Orderitems struct {
// 	item_id     int `json:"item_id"`
// 	order_id    int `json:"order_id"`
// 	product_id  int `json:"product_id"`
// 	quantity    int `json:"quantity"`
// 	item_price  int `json:"item_price"`
// 	total_price int `json:"total_price"`
// }

// type Orders struct {
// 	order_id     int `json:"order_id"`
// 	user_id      int `json:"user_id"`
// 	total_amount int `json:"total_amount"`
// 	order_date   int `json:"order_date"`
// }
