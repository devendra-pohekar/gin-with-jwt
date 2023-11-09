package response

type AddProductStruct struct {
	ProductName string `json:"product_name"`
}

type UpdateProduct struct {
	ProductName string `json:"product_name"`
}

type FetchProduct struct {
	ProductID int `json:"product_id"`
}
