package login

import (
	"encoding/json"
	"flattrade/apps/DBConnection/gormdb"
	common "flattrade/common"
	"fmt"
	"io"
	"log"
	"net/http"
)

type UserLogin struct {
	UserId   string `json:"user_id" gorm:"userid"`
	Password string `json:"password" gorm:"password"`
}

type UserLoginResponse struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMsg"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "CLIENT, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(w).Header().Set("Content-Type", "application/json")

	log.Println("LoginUser-(+)")

	var lUserLoginResponse UserLoginResponse
	var lUser UserLogin

	lUserLoginResponse.Status = common.SuccessCode

	if r.Method == http.MethodPost {

		lBody, lErr := io.ReadAll(r.Body)

		if lErr != nil {

			log.Println("LULU-001", lErr.Error())
			lUserLoginResponse.ErrMsg = lErr.Error()
			lUserLoginResponse.Status = common.ErrorCode

		} else {

			lErr := json.Unmarshal(lBody, &lUser)

			lUserPassword := lUser.Password

			if lErr != nil {

				log.Println("LULU-002", lErr.Error())
				lUserLoginResponse.ErrMsg = lErr.Error()
				lUserLoginResponse.Status = common.ErrorCode

			} else {

				if validate(&lUserLoginResponse, &lUser) {
					checkUser(&lUserLoginResponse, &lUser, lUserPassword)
				}

			}
		}

	} else {
		lUserLoginResponse.Status = common.ErrorCode
		lUserLoginResponse.ErrMsg = "Invalid Method"
	}

	lData, lErr := json.Marshal(lUserLoginResponse)

	if lErr != nil {
		log.Println("")
		lUserLoginResponse.Status = common.ErrorCode
		lUserLoginResponse.ErrMsg = "Server Error"
	} else {
		fmt.Fprintf(w, string(lData))
	}

	log.Println("LoginUser-(-)")

}

/* ========================================== Validate the userid and password =================================================== */

func validate(lUserLoginResponse *UserLoginResponse, lUser *UserLogin) bool {

	log.Println("validate-(+)")
	if lUser.Password == "" && lUser.UserId == "" {
		lUserLoginResponse.ErrMsg = "user_id & password field is required"
		lUserLoginResponse.Status = common.ErrorCode
		return false
	} else if lUser.UserId == "" {
		lUserLoginResponse.ErrMsg = "userid field is required"
		lUserLoginResponse.Status = common.ErrorCode
		return false
	} else if lUser.Password == "" {
		lUserLoginResponse.ErrMsg = "password field is required"
		lUserLoginResponse.Status = common.ErrorCode
		return false
	}

	log.Println("validate-(-)")
	return true

}

/* ========================================== For check in db user is Available or Not ============================================ */

func checkUser(lUserLoginResponse *UserLoginResponse, lUser *UserLogin, lUserPassword string) {

	log.Println("checkUser-(+)")

	lGormDB, lErr := gormdb.GormDBConnection()

	if lErr != nil {

		log.Println("LCU-001", lErr.Error())
		lUserLoginResponse.ErrMsg = lErr.Error()
		lUserLoginResponse.Status = common.ErrorCode

	} else {

		fmt.Println(lUserLoginResponse)
		fmt.Println(lUserPassword)
		lResult := lGormDB.Table("st_918_emp").Select("password").Where("userid=?", lUser.UserId).Find(&lUser.Password)

		if lResult.Error != nil || lResult.RowsAffected == 0 {
			log.Println("LCU-002", lResult.Error.Error())
			lUserLoginResponse.Status = common.ErrorCode
			lUserLoginResponse.ErrMsg = "Invalid User Id"

		} else {

			if lUser.Password != lUserPassword {

				lUserLoginResponse.Status = common.ErrorCode
				lUserLoginResponse.ErrMsg = "Password is Incorrect"

			} else {

				lUserLoginResponse.Status = common.SuccessCode
				lUserLoginResponse.ErrMsg = ""
				fmt.Println("Login Successfully")
			}
		}
	}

	log.Println("checkUser-(-)")

}
