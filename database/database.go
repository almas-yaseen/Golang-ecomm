package database

import (
	"fmt"
	"ginapp/config"
	"ginapp/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)

	db, dberr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if dberr != nil {
		return nil, fmt.Errorf("faild to connect to database:%w", dberr)
	}
	DB = db

	DB.AutoMigrate(&domain.User{})
	return DB, nil

}
