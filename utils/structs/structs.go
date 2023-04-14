package structs

type Coordinates struct {
	Type        string     `json:"type"`
	Coordinates [2]float32 `json:"coordinates"`
}

type Address struct {
	Building string      `json:"building"`
	Coord    Coordinates `json:"coord"`
	Street   string      `json:"street"`
	Zipcode  string      `json:"zipcode"`
}

type Grade struct {
	Date struct {
		Date int64 `json:"$date"`
	} `json:"date"`
	Grade string `json:"grade"`
	Score int    `json:"score"`
}

type Restaurant struct {
	Address      Address `json:"address"`
	Borough      string  `json:"borough"`
	Cuisine      string  `json:"cuisine"`
	Grades       []Grade `json:"grades"`
	Name         string  `json:"name"`
	RestaurantID string  `json:"restaurant_id"`
}
