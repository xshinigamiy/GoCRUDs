package repository

import (
	"GoCRUDs/config"
	"GoCRUDs/pkg/pojos"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	FirstName  string
	LastName   string
	Age        int
	Profession string
	CreatedAt  time.Time `gorm:"type:timestamp;default:current_timestamp"`
	CreatedBy  string
	UpdatedAt  time.Time `gorm:"type:timestamp;default:current_timestamp ON update current_timestamp"`
	UpdatedBy  string
	IsActive   string
}

var (
	MySQLRepo *MySQLRepository
)

type MySQLRepository struct {
	DbConnection *gorm.DB
}

func init() {
	InitMySQLDBRepo()
}

func GetMySQLRepo() *MySQLRepository {
	if MySQLRepo != nil {
		return MySQLRepo
	} else {
		InitMySQLDBRepo()
		return MySQLRepo
	}
}

func GetDbURL() string {
	return config.GetConfigs().DBConfigs.DBUser + ":" + config.GetConfigs().DBConfigs.DBPassword +
		"@tcp(" + config.GetConfigs().DBConfigs.Url + ":" + config.GetConfigs().DBConfigs.Port + ")/" +
		config.GetConfigs().DBConfigs.DatabaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func InitMySQLDBRepo() {
	log.Info("URL : ", GetDbURL())
	db, err := gorm.Open(mysql.Open(GetDbURL()), &gorm.Config{})
	if err != nil {
		log.Error("Couldn't get dbConnection connection")
		panic(err)
	} else {
		log.Info("dbConnection connection initialized successfully")
	}

	MySQLRepo = &MySQLRepository{
		DbConnection: db,
	}

	err = MySQLRepo.DbConnection.AutoMigrate(&User{})
	if err != nil {
		log.Error("error while migrating DB")
		return
	}
}

func (repo *MySQLRepository) CreateUser(user pojos.User) *gorm.DB {
	//TODO: this method does not return the ID of the inserted object
	val := repo.DbConnection.Create(user)
	log.Info(val.Statement)
	return nil
}

func (repo *MySQLRepository) GetUserById(Id string) *gorm.DB {
	//TODO: fetch user based on Id
	return nil
}
