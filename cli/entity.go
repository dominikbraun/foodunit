package main

type Offer struct {
	ID         string `json:"id"`
	SupplierID string `json:"supplier_id"`
}

type Order struct {
	Email     string `json:"email"`
	Positions []struct {
		DishID string `json:"dish_id"`
		Name   string `json:"name"`
	} `json:"positions"`
	Total float64 `json:"total"`
}

type Supplier struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Mon     string `json:"mon"`
	Tue     string `json:"tue"`
	Wed     string `json:"wed"`
	Thu     string `json:"thu"`
	Fri     string `json:"fri"`
}

type Cat struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Img    string `json:"img"`
	Dishes []Dish `json:"dishes"`
}

type Dish struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Price string `json:"price"`
}
