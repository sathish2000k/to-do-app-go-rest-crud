package config

import (
	"fmt"
	"log"
	models "to-do-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/go-redis/redis/v8"
)

var DB *gorm.DB
var RedisClient *redis.Client

func ConnectDB() {
	config := loadConfig()

	dsn := fmt.Sprintf("host = %s user =%s password = %s dbname = %s port = %d sslmode = disable", config.Database.Host, config.Database.User, config.Database.Password, config.Database.Database, config.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	} else {
		log.Println("Database connected successfully")
	}

	db.AutoMigrate(&models.ToDo{})
	db.AutoMigrate(&models.User{})
	DB = db
}

func ConnectRedis() {
	config := loadConfig()
	log.Printf("Connecting to Redis %s:%d\n", config.Redis.Host, config.Redis.Port)
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
	})

	if redisClient == nil {
		log.Fatalf("Error connecting to Redis")
	} else {
		log.Println("Redis connected successfully")
	}

	RedisClient = redisClient
}