package test

import (
	"flattrade/genpkg"
	"fmt"
	"log"
	"strconv"
)

func GenerateUniqueID(clientId int, role string) string {

	log.Println("generateUniqueID-(+)")

	var lUniqueId = ""

	lUniqueIdFormatInTomal := ""

	switch role {
	case "C":
		lUniqueIdFormatInTomal = "UniqueId"
	case "BO":
		lUniqueIdFormatInTomal = "BillerUniqueId"
	case "B":
		lUniqueIdFormatInTomal = "BackOfficerUniqueId"
	case "APPR":
		lUniqueIdFormatInTomal = "ApproverUniqueId"
	}

	config := genpkg.ReadTomlConfig("./toml/dbconfig.toml")
	lUniqueIdFormat := fmt.Sprintf("%v", config.(map[string]interface{})[lUniqueIdFormatInTomal])
	lUniqueId = lUniqueIdFormat
	lUniqueId = lUniqueId + strconv.Itoa(clientId+1)

	log.Println("generateUniqueID-(-)")
	return lUniqueId

}
