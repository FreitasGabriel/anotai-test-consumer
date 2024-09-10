package model

type CatalogPayload struct {
	Owner   string            `json:"owner"`
	Catalog []CatalogCategory `json:"catalog"`
}

type CatalogCategory struct {
	CategoryTitle       string        `json:"category_title"`
	CategoryDescription string        `json:"category_description"`
	Itens               []CatalogItem `json:"itens"`
}

type CatalogItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
