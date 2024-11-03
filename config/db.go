package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/inventori-app-jeff/internal/model"
	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", Cfg.Database.Host, Cfg.Database.Username, Cfg.Database.Password, Cfg.Database.Dbname, Cfg.Database.Port)
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	}

	once.Do(func() {
		DB = Db
		fmt.Println("Successfully Connected To Database!")
	})
}

func SyncDB() {
	if err := DB.AutoMigrate(&model.User{}); err != nil {
		fmt.Print(err.Error())
	}
}

func UserSeeder(count int) []*model.User {
	users := []*model.User{}
	for i := 0; i < count; i++ {
		users = append(users, GenerateUser())
	}

	return users
}

func GenerateUser() *model.User {
	return &model.User{
		Username:    faker.Username(),
		Email:       faker.Email(),
		Password:    "admin",
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		PhoneNumber: faker.Phonenumber(),
	}
}
