package tests

import (
	// "fmt"
	"link-aja/models"
	// "os"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestConnection(t *testing.T) {
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER_TEST"), os.Getenv("DB_PASSWORD_TEST"), os.Getenv("DB_HOST_TEST"), os.Getenv("DB_PORT_TEST"), os.Getenv("DB_NAME_TEST"))
	dsn := "root:@tcp(localhost:3306)/testing_link_aja"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(
		&models.Customer{},
		&models.Account{},
	)
}
