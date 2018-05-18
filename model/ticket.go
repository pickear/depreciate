package model

type RuYueResponse struct {
	Success bool `json:"success"`
	Code int `json:"code"`
	Data Data `json:"data"`
}

type Data struct {
	Pct string `json:"pct"`
	Product Product `json:"Product"`
	Items []Item `json:"items"`
}

type Product struct {
	Pnm string `json:"pnm"`
	Rdc string `json:"rdc"`
}

type Item struct {
	Date string `json:"date"`
	Priceyj string `json:"priceyj"`
	Clsinf []Clsinf `json:"clsinf"`
}

type Clsinf struct {
	Clstm string `json:"clstm"`
	Seats string `json:"seats"`
} 