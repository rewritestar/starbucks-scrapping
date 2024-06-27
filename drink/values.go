package drink

const (
	SiteUrl          = "https://www.starbucks.co.kr/menu/drink_list.do"
	ChromeDriverPath = "./chromedriver"
	OutputPath       = "./output/drink.csv"
)

var DrinkNutrientClassMap = map[string]string{
	"Kcal":        ".kcal dd",
	"SatFat":      ".sat_FAT dd",
	"Protein":     ".protein dd",
	"Fat":         ".fat dd",
	"TransFat":    ".trans_FAT dd",
	"Sodium":      ".sodium dd",
	"Sugars":      ".sugars dd",
	"Caffeine":    ".caffeine dd",
	"Cholesterol": ".cholesterol dd",
	"Chabo":       ".chabo dd",
}
