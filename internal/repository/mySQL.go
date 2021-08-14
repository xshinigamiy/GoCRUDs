package repository

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var (
	MySQLRepoVar *MySQLRepo
)

type MySQLRepo struct {
	db *gorm.DB
}

func GetMySQLRepo () *MySQLRepo {
	if MySQLRepoVar != nil {
		return MySQLRepoVar
	}
	InitMySQLDBRepo()
	return MySQLRepoVar
}

func InitMySQLDBRepo() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/person?charset=utf8&parseTime=True")
	if err != nil {
		log.Error("Couldn't get DB connection")
	} else {
		log.Info("DB connection initialized successfully")
	}
	MySQLRepoVar.db = db
}


