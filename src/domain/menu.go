package domain

type Menu struct {
	Data []MenuItem `json:"data,omitempty"`
	Date string `json:"date,omitempty"`
}

type MenuItem struct {
	Ingredients string `json:"ingredients,omitempty"`
	Name string `json:"name,omitempty"`
	NameEN string `json:"name_en,omitempty"`
	NameSV string `json:"name_sv,omitempty"`
	Nutrition string `json:"nutritionn,omitempty"`
	Price struct {
		Name string `json:"name,omitempty"`
		Value struct {
			Contract string `json:"contract,omitempty"`
			Graduate string `json:"graduate,omitempty"`
			GraduateHYY string `json:"graduate_hyy,omitempty"`
			Normal string `json:"normal,omitempty"`
			Student string `json:"student,omitempty"`
			StudentHYY string `json:"student_hyy,omitempty"`
		} `json:"value,omitempty"`
	} `json:"price,omitempty"`
	SKU string `json:"sku,omitempty"`
}