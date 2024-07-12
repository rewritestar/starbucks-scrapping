package data

type DrinkCategoryList []*DrinkCategory

type DrinkCategory struct {
	ID   uint
	Name string
	Code string
}

func (d DrinkCategory) TableName() string {
	return "drink_category"
}

var DrinkCategories = DrinkCategoryList{
	&DrinkCategory{
		ID:   1,
		Name: "콜드 브루",
		Code: "COLD-BREW",
	},
	&DrinkCategory{
		ID:   2,
		Name: "브루드 커피",
		Code: "BREWED",
	},
	&DrinkCategory{
		ID:   3,
		Name: "에스프레소",
		Code: "ESPRESSO",
	},
	&DrinkCategory{
		ID:   4,
		Name: "프라푸치노",
		Code: "FRAPPUCCINO",
	},
	&DrinkCategory{
		ID:   5,
		Name: "블렌디드",
		Code: "BLENDED",
	},
	&DrinkCategory{
		ID:   6,
		Name: "스타벅스 리프레셔",
		Code: "REFRESHER",
	},
	&DrinkCategory{
		ID:   7,
		Name: "스타벅스 피지오",
		Code: "FIZZIO",
	},
	&DrinkCategory{
		ID:   8,
		Name: "티",
		Code: "TEA",
	},
	&DrinkCategory{
		ID:   9,
		Name: "스타벅스 주스(병음료)",
		Code: "BOTTLE",
	},
	&DrinkCategory{
		ID:   10,
		Name: "기타",
		Code: "ETC",
	},
}
