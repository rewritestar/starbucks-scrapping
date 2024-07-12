package drink

type DrinkList []*Drink

type Drink struct {
	NameKR string
	NameEN string
	ImgUrl string
	DrinkNutrient
	Price      int
	Likes      int64
	IsExistent bool
	Cate       uint
}

type DrinkNutrient struct {
	Kcal        int64
	SatFat      int64
	Protein     int64
	Fat         int64
	TransFat    int64
	Sodium      int64
	Sugars      int64
	Caffeine    int64
	Cholesterol int64
	Chabo       int64
}

type DrinkStringList []*DrinkString

type DrinkString struct {
	NameKR      string
	NameEN      string
	ImgUrl      string
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
	Cate        string
}
