package connection

import (
	"log"

	wrk "webapiingo/worker"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DataMigration() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:12345@tcp(127.0.0.1:3306)/workerdbs?parseTime=True"), &gorm.Config{})
	if err != nil {
		log.Fatal("error in connection : ", err)
	}
	db.AutoMigrate(&wrk.WorkerInfo{})

	return db
}
