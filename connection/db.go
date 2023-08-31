package connection

import (
	// "fmt"
	"link-aja/models"
	"log"

	// "os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB {
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	dsn := "root:@tcp(localhost:3306)/test_link_aja"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(
		&models.Customer{},
		&models.Account{},
	)

	return db
}
