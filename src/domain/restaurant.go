package domain

type RestaurantInfoShort struct {
	Areacode int32 `json:"areacode,omitempty"`
	ID int32	`json:"id,omitempty"`
	Name string	`name:"title,omitempty"`
}

type RestaurantInfoLong struct {
	Address string
	Zip string
	City string
	Description string
	Email string
	Name string `json:"restaurant,omitempty"`
}