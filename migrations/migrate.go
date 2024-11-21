package migrations

import (
	"github.com/gustavNdamukong/library-inventory-golang/initializers"
	"github.com/gustavNdamukong/library-inventory-golang/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	initializers.ConnectToDB()
}

func main() {
	DB.AutoMigrate(&models.Bindas{})
}
