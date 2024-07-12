package main

import "starbucks/menu/scrapping/drink"

func main() {
	// csv 파일 취득
	// drink.ScrapAndWriteCsv()

	// drink_category DB 생성
	// drink_category.Run()

	//csv 파일 기반으로 drink DB 생성
	drink.WriteDatabase()
}
