package common

import (
	"flattrade/genpkg"
	"fmt"
	"log"
	"strconv"
)

const (
	SuccessCode = "S"
	ErrorCode   = "E"
	Pending     = "Pending"
	Approved    = "Approved"
	Rejected    = "Rejected"
	BackOfficer = "BO"
	Biller      = "B"
	Approver    = "APPR"
)

func ValidateUserRole(lRole string) bool {
	log.Println("ValidateUserRole-(+)")
	if lRole == "" || lRole != BackOfficer || lRole != Biller || lRole != Approver {
		return false
	}
	log.Println("ValidateUserRole-(+)")
	return true
}

func GenerateUniqueID(clientId int, role string) string {

	log.Println("generateUniqueID-(+)")

	var lUniqueId = ""

	lUniqueIdFormatInTomal := ""

	switch role {
	case "C":
		lUniqueIdFormatInTomal = "UniqueId"
	case "BO":
		lUniqueIdFormatInTomal = "UserUniqueId"
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
