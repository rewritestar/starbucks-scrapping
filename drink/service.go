package drink

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"starbucks/menu/scrapping/data"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/fedesog/webdriver"
)

const ETCCategoryID = 10

func ScrapAndWriteCsv() {
	// 크롤링 & csv 파일 작성
	session, err := setUpSession()
	if err != nil {
		handleError(err)
	}

	err = writeCsv(session)
	if err != nil {
		handleError(err)
	}
}

func setUpSession() (*webdriver.Session, error) {
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
	w.Write([]string{"제품명(한글)", "제품명(영어)", "이미지 URL", "칼로리(kcal)", "포화지방(g)", "단백질(g)", "지방(g)", "트랜스 지방(g)", "나트륨(mg)", "당류(g)", "카페인 함량(mg)", "콜레스테롤(mg)", "탄수화물(g)", "카테고리"})

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
			drink.Cate,
		})
	}

	defer file.Close()
	defer w.Flush()

	return nil
}

func retrieveData(session *webdriver.Session) (DrinkStringList, error) {
	drinkList := DrinkStringList{}

	list, err := session.FindElements(webdriver.ClassName, "menuDataSet")
	if err != nil {
		return nil, err
	}
	for _, item := range list {
		drink := DrinkString{}

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
		})
		drink.Cate = html.Find(".cate").Text()
		drinkList = append(drinkList, &drink)

		session.Back()
	}
	return drinkList, nil
}

func WriteDatabase() {
	// csv 파일을 데이터 베이스에 저장
	err := writeDatabase()
	if err != nil {
		handleError(err)
	}
}

func writeDatabase() error {
	db, gormDB, err := setUpDatabase()
	if err != nil {
		return err
	}

	drinkList, err := readCsv()
	if err != nil {
		return err
	}

	if err = createTable(db); err != nil {
		return err
	}

	err = gormDB.Create(drinkList).
		Error

	if err != nil {
		return err
	}

	defer db.Close()
	return nil
}

func setUpDatabase() (*sql.DB, *gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, nil, err
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPW := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPW, dbHost, dbPort, dbName))
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return db, gormDB, nil
}

func readCsv() (DrinkList, error) {
	file, _ := os.Open(OutputPath)

	r := csv.NewReader(bufio.NewReader(file))

	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	drinkList := DrinkList{}
	for i, row := range rows {
		if i == 0 {
			continue
		}
		drink := Drink{}
		drink.NameKR = row[0]
		drink.NameEN = row[1]
		drink.ImgUrl = row[2]
		drink.Kcal, err = stringToInt64(row[3])
		drink.SatFat, err = stringToInt64(row[4])
		drink.Protein, err = stringToInt64(row[5])
		drink.Fat, err = stringToInt64(row[6])
		drink.TransFat, err = stringToInt64(row[7])
		drink.Sodium, err = stringToInt64(row[8])
		drink.Sugars, err = stringToInt64(row[9])
		drink.Caffeine, err = stringToInt64(row[10])
		drink.Cholesterol, err = stringToInt64(row[11])
		drink.Chabo, err = stringToInt64(row[12])
		drink.Cate = getCategory(row[13])
		drink.Price = data.InitDrinkData[i-1].Price
		drink.Likes = data.InitDrinkData[i-1].Likes
		drink.IsExistent = data.InitDrinkData[i-1].IsExistent
		if err != nil {
			return nil, err
		}
		drinkList = append(drinkList, &drink)
	}

	defer file.Close()

	return drinkList, nil
}

func createTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS drinks")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE TABLE drinks (\n    id INT NOT NULL AUTO_INCREMENT,\n    name_kr VARCHAR(255),\n    name_en VARCHAR(500),\n    img_url VARCHAR(500),\n    kcal INT,\n    sat_fat INT,\n    protein INT,\n    fat INT,\n    trans_fat INT,\n    sodium INT,\n    sugars INT,\n    caffeine INT,\n    cholesterol INT,\n    chabo INT,\n    price INT,\n    likes BIGINT,\n    is_existent BOOLEAN,\n    cate INT,\n    PRIMARY KEY (id)\n)")
	if err != nil {
		return err
	}
	return nil
}

func handleError(err error) {
	log.Fatalln(err)
}

func stringToInt64(str string) (int64, error) {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func getCategory(name string) uint {
	for _, category := range data.DrinkCategories {
		if category.Name == name {
			return category.ID
		}
	}
	return ETCCategoryID
}
