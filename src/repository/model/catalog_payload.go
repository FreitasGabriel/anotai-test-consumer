package model

type CatalogAggregation struct {
	ID          string    `bson:"id,omitempty"`
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	OwnerID     string    `bson:"owner_id"`
	ProductList []Product `bson:"products"`
}

type Catalog struct {
	Owner   string     `json:"owner"`
	Catalog []Category `json:"catalog"`
}

type Category struct {
	CategoryTitle       string    `json:"category_title"`
	CategoryDescription string    `json:"category_description"`
	Itens               []Product `json:"itens"`
}

type Product struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
