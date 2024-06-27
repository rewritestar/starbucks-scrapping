package drink

type DrinkList []*Drink

type Drink struct {
	NameKR string
	NameEN string
	ImgUrl string
	DrinkNutrient
}

type DrinkNutrient struct {
	Kcal        string
	SatFat      string
	Protein     string
	Fat         string
	TransFat    string
	Sodium      string
	Sugars      string
	Caffeine    string
	Cholesterol string
	Chabo       string
}
