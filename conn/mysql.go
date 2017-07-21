package conn

import (
	"fmt"
	"os"

	"gitlab.com/guszak/test/models"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// InitDb init db connection
func InitDb() *gorm.DB {

	env := os.Getenv("ENV_VAR_FROM_SYSTEM")
	if env != "1" {
		err := godotenv.Load("config/.env")
		if err != nil {
			panic(err)
		}
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME"))
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	db.LogMode(true)

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Product{}, &models.Company{})

	return db
}
