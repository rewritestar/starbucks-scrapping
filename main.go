package main

import "starbucks/menu/scrapping/drink"

func main() {
	drink.ScrapAndWriteCsv()
	drink.WriteDatabase()
}
