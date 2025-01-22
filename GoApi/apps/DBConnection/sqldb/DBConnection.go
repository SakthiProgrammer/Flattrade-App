package sqldb

import (
	"database/sql"
	genpkg "flattrade/genpkg"
	"fmt"
	"log"
	"strconv"
)

func DBConnection() (*sql.DB, error) {
	log.Println("DBConnection-(+)")

	config := genpkg.ReadTomlConfig("./toml/dbconfig.toml")
	lUserName := fmt.Sprintf("%v", config.(map[string]interface{})["DBUser"])
	lPassword := fmt.Sprintf("%v", config.(map[string]interface{})["DBPass"])
	lHostName := fmt.Sprintf("%v", config.(map[string]interface{})["DBHost"])
	lPort, _ := strconv.Atoi(fmt.Sprintf("%v", config.(map[string]interface{})["DBPort"]))
	lDBName := fmt.Sprintf("%v", config.(map[string]interface{})["DBName"])

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", lUserName, lPassword, lHostName, lPort, lDBName)
	lDb, lErr := sql.Open("mysql", connString)

	if lErr != nil {
		log.Println("Open Connection Failed: ", lErr.Error())
		return lDb, lErr
	}

	log.Println("DBConnection-(-)")
	return lDb, nil
}
