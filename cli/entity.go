package main

// Offer represents someone's suggestion to order food from a certain
// supplier. An Offer therefore holds the corresponding supplier's ID.
type Offer struct {
	ID         string `json:"id"`
	SupplierID string `json:"supplier_id"`
}

// Order represents a collection of positions that are associated with
// a given e-mail address. Multiple Order instances eventually form the
// actual order which will be sent to the supplier.
type Order struct {
	Email     string `json:"email"`
	Positions []struct {
		DishID string `json:"dish_id"`
		Name   string `json:"name"`
	} `json:"positions"`
	Total float64 `json:"total"`
}

// Supplier represents a restaurant, takeaway or even supermarket. Any
// offer applies to one particular supplier, and the several dishes
// that can be ordered depict the supplier's menu.
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

// Cat represents the category a dish is associated to. While it may
// be used by multiple suppliers at once, the individual dishes belong
// to only one category and one supplier. When receiving a supplier's
// menu, the API will return a list of categories, each containing the
// dishes that are associated with respective category and supplier.
type Cat struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Img    string `json:"img"`
	Dishes []Dish `json:"dishes"`
}

// Dish represents a menu item and is therefore associated with a certain
// supplier. However, this information won't be depicted in the entity
// because the API actually returns the menu as a list of all categories,
// holding a list of corresponding dishes - the supplier ID isn't needed.
type Dish struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Price string `json:"price"`
}
