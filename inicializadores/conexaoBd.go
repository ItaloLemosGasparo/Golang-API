package inicializadores

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var BD *gorm.DB

func ConectarAoBD() {
	var err error

	dsn := os.Getenv("BD_URL")
	BD, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Falha ao conectar a Base de Dados")
		return
	}
}
