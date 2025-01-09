package gormdb

import (
	"flattrade/genpkg"
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormDBConnection() (*gorm.DB, error) {

	config := genpkg.ReadTomlConfig("./toml/dbconfig.toml")
	lUserName := fmt.Sprintf("%v", config.(map[string]interface{})["DBUser"])
	lPassword := fmt.Sprintf("%v", config.(map[string]interface{})["DBPass"])
	lHostName := fmt.Sprintf("%v", config.(map[string]interface{})["DBHost"])
	lPort, _ := strconv.Atoi(fmt.Sprintf("%v", config.(map[string]interface{})["DBPort"]))
	lDBName := fmt.Sprintf("%v", config.(map[string]interface{})["DBName"])

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", lUserName, lPassword, lHostName, lPort, lDBName)

	gormDB, lErr := gorm.Open(mysql.Open(connString), &gorm.Config{})

	if lErr != nil {
		log.Println("GGDBC-001", lErr.Error())
		return gormDB, lErr
	}

	/* sqlDB, err := gormDB.DB()
	if err != nil {
		log.Println("GGDBC-002", err.Error())
		return nil, err
	}

	// Set connection pool settings
	sqlDB.SetMaxOpenConns(10)                  // Max open connections
	sqlDB.SetMaxIdleConns(5)                   // Max idle connections
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Max connection lifetime */
	log.Println("Database connection established successfully.")
	return gormDB, nil
}
