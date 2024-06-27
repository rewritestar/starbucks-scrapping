package drink

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/fedesog/webdriver"
)

func Run() {
	settion, err := setUp()
	if err != nil {
		handleError(err)
	}

	err = writeCsv(settion)
	if err != nil {
		handleError(err)
	}
}

func setUp() (*webdriver.Session, error) {
	chromeDriver := webdriver.NewChromeDriver(ChromeDriverPath)
	err := chromeDriver.Start()
	if err != nil {
		return nil, err
	}

	desired := webdriver.Capabilities{"Platform": "Windows"}
	required := webdriver.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)
	if err != nil {
		return nil, err
	}

	err = session.Url(SiteUrl)
	if err != nil {
		return nil, err
	}
	time.Sleep(3 * time.Second)
	return session, nil
}

func writeCsv(session *webdriver.Session) error {
	file, err := os.Create(OutputPath)
	if err != nil {
		return err
	}

	w := csv.NewWriter(bufio.NewWriter(file))
	w.Write([]string{"제품명(한글)", "제품명(영어)", "이미지 URL", "칼로리(kcal)", "포화지방(g)", "단백질(g)", "지방(g)", "트랜스 지방(g)", "나트륨(mg)", "당류(g)", "카페인 함량(mg)", "콜레스테롤(mg)", "탄수화물(g)"})

	drinkList, err := retrieveData(session)
	if err != nil {
		return err
	}

	for _, drink := range drinkList {
		w.Write([]string{
			drink.NameKR,
			drink.NameEN,
			drink.ImgUrl,
			drink.Kcal,
			drink.SatFat,
			drink.Protein,
			drink.Fat,
			drink.TransFat,
			drink.Sodium,
			drink.Sugars,
			drink.Caffeine,
			drink.Cholesterol,
			drink.Chabo,
		})
	}

	defer file.Close()
	defer w.Flush()

	return nil
}

func retrieveData(session *webdriver.Session) (DrinkList, error) {
	drinkList := DrinkList{}

	list, err := session.FindElements(webdriver.ClassName, "menuDataSet")
	if err != nil {
		return nil, err
	}
	for _, item := range list {
		drink := Drink{}

		got, err := item.FindElement(webdriver.TagName, "a")
		if err != nil {
			return nil, err
		}
		img, err := got.FindElement(webdriver.TagName, "img")
		if err != nil {
			return nil, err
		}
		imgUrl, err := img.GetAttribute("src")
		if err != nil {
			return nil, err
		}
		drink.ImgUrl = imgUrl

		got.Click()
		time.Sleep(1 * time.Second)

		src, err := session.Source()
		if err != nil {
			return nil, err
		}
		html, err := goquery.NewDocumentFromReader(strings.NewReader(src))
		if err != nil {
			return nil, err
		}

		html.Find(".product_view_detail").Each(func(i int, s *goquery.Selection) {
			drink.NameKR = s.Find("h4").Nodes[0].FirstChild.Data
			drink.NameEN = s.Find("h4 span").Text()

			info := s.Find(".product_info_content")
			drink.Kcal = info.Find(DrinkNutrientClassMap["Kcal"]).Text()
			drink.SatFat = info.Find(DrinkNutrientClassMap["SatFat"]).Text()
			drink.Protein = info.Find(DrinkNutrientClassMap["Protein"]).Text()
			drink.Fat = info.Find(DrinkNutrientClassMap["Fat"]).Text()
			drink.TransFat = info.Find(DrinkNutrientClassMap["TransFat"]).Text()
			drink.Sodium = info.Find(DrinkNutrientClassMap["Sodium"]).Text()
			drink.Sugars = info.Find(DrinkNutrientClassMap["Sugars"]).Text()
			drink.Caffeine = info.Find(DrinkNutrientClassMap["Caffeine"]).Text()
			drink.Cholesterol = info.Find(DrinkNutrientClassMap["Cholesterol"]).Text()
			drink.Chabo = info.Find(DrinkNutrientClassMap["Chabo"]).Text()

			drinkList = append(drinkList, &drink)
		})
		session.Back()
	}
	return drinkList, nil
}

func handleError(err error) {
	log.Fatal(err)
}
