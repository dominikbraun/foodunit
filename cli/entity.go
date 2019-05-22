package main

import "fmt"

// Offer represents someone's suggestion to order food from a certain
// supplier. An Offer therefore holds the corresponding supplier's ID.
type Offer struct {
	ID         string `json:"id"`
	SupplierID string `json:"supplier_id"`
}

// Returns a formatted String representing an Offer instance.
func (o Offer) String() string {
	str := fmt.Sprintf("#%s\tSupplier: %s", o.ID, o.SupplierID)
	return str
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

func (s Supplier) String() string {
	str := fmt.Sprintf("#%s\tName: %s\n", s.ID, s.Name)
	str += fmt.Sprintf("\tAdress: %s\n", s.Address)
	str += fmt.Sprintf("\tPhone: %s\n", s.Phone)

	opened := map[string]string{
		"Mon": s.Mon,
		"Tue": s.Tue,
		"Wed": s.Wed,
		"Thu": s.Thu,
		"Fri": s.Fri,
	}
	for day, time := range opened {
		str += fmt.Sprintf("\t%s: %s\n", day, time)
	}
	return str
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

// Returns a formatted String representing an Cat instance.
func (c Cat) String() string {
	str := fmt.Sprintf("#%s\tName: %s\n", c.ID, c.Name)

	for _, d := range c.Dishes {
		str += fmt.Sprintf("\t%s\n", d)
	}
	return str
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

// Returns a formatted String representing an Dish instance.
func (d Dish) String() string {
	str := fmt.Sprintf("#%s\tName: %s\tPrice: %s", d.ID, d.Name, d.Price)
	return str
}
