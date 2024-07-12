package drink_category

import (
	"fmt"
	"log"
	"os"
	"starbucks/menu/scrapping/data"

	"database/sql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Run() {
	if err := writeDrinkCategoryDatabase(); err != nil {
		log.Fatalln(err)
	}
}

func writeDrinkCategoryDatabase() error {
	db, gormDB, err := setUpDatabase()
	if err != nil {
		return err
	}

	if err = createTable(db); err != nil {
		return err
	}

	if err = insertCategory(gormDB); err != nil {
		return err
	}

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

func createTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS drink_category")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE TABLE drink_category (\n    id INT,\n    name VARCHAR(255),\n    code VARCHAR(255),\n    primary key (id)\n)")
	if err != nil {
		return err
	}
	return nil
}

func insertCategory(db *gorm.DB) error {
	result := data.DrinkCategories

	err := db.Model(data.DrinkCategory{}).Create(&result).
		Error
	if err != nil {
		return err
	}
	return nil
}
