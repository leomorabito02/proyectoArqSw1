package db

import (
	"../proyectoArqSw1/model"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
 	"github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "arqsw"
	DBUser := "root" //a chequear
	DBPass := ""
	//DBPass := os.Getenv("MVC_DB_PASS")
	DBHost := "root"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	usuarioClient.Db = db
	reservaClient.Db = db
	hotelClient.Db = db
}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&model.Usuario{})
	db.AutoMigrate(&model.Reserva{})
	db.AutoMigrate(&model.Hotel{})

	log.Info("Finishing Migration Database Tables :)")
}
