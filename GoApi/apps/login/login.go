package login

import (
	"encoding/json"
	"flattrade/apps/DBConnection/gormdb"
	common "flattrade/common"
	"flattrade/genpkg"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type Login struct {
	UserId   string `json:"user_id" gorm:"column:userid"`
	Password string `json:"password" gorm:"column:password"`
}

type LoginResponse struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMsg"`
}

/* =================================== login for CLIENT, ADMIN, USER(employee) =============================================

		* USER - U
		* ADMIN - A
		* CLIENT - C

   =========================================================================================================================*/

func LoginAll(w http.ResponseWriter, r *http.Request) {

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Content-Type", "application/json")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "ROLE, USERROLE, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	log.Println("LoginUser-(+)")

	var lLoginResponse LoginResponse
	var lUser Login

	lRole := r.Header.Get("ROLE")

	lRole = strings.ToUpper(lRole)

	if isRoleValid(lRole) {
		lLoginResponse.ErrMsg = "Provide a role in header"
		lLoginResponse.Status = common.ErrorCode

	} else {

		lLoginResponse.Status = common.SuccessCode

		if r.Method == http.MethodPost {

			lBody, lErr := io.ReadAll(r.Body)

			if lErr != nil {

				log.Println("LULU-001", lErr.Error())
				lLoginResponse.ErrMsg = lErr.Error()
				lLoginResponse.Status = common.ErrorCode

			} else {

				lErr := json.Unmarshal(lBody, &lUser)

				lUserPassword := lUser.Password

				if lErr != nil {

					log.Println("LULU-002", lErr.Error())
					lLoginResponse.ErrMsg = lErr.Error()
					lLoginResponse.Status = common.ErrorCode

				} else {
					if validate(&lLoginResponse, &lUser) {

						checkUser(&lLoginResponse, &lUser, lUserPassword, lRole, r)

					}

				}
			}

		} else {
			lLoginResponse.Status = common.ErrorCode
			lLoginResponse.ErrMsg = "Invalid Method"
		}
	}

	lData, lErr := json.Marshal(lLoginResponse)

	if lErr != nil {
		log.Println("")
		lLoginResponse.Status = common.ErrorCode
		lLoginResponse.ErrMsg = "Server Error"
	} else {
		fmt.Fprintf(w, string(lData))
	}

	log.Println("LoginUser-(-)")

}

func isRoleValid(lRole string) bool {

	if lRole != "C" && lRole != "A" && lRole != "U" {
		return true
	}
	return false
}

/* ========================================== Validate the userid and password =================================================== */

func validate(lLoginResponse *LoginResponse, lUser *Login) bool {

	log.Println("validate-(+)")
	if lUser.Password == "" && lUser.UserId == "" {
		lLoginResponse.ErrMsg = "user_id & password field is required"
		lLoginResponse.Status = common.ErrorCode
		return false
	} else if lUser.UserId == "" {
		lLoginResponse.ErrMsg = "userid field is required"
		lLoginResponse.Status = common.ErrorCode
		return false
	} else if lUser.Password == "" {
		lLoginResponse.ErrMsg = "password field is required"
		lLoginResponse.Status = common.ErrorCode
		return false
	}

	log.Println("validate-(-)")
	return true

}

/* ========================================== For check in db user is Available or Not ============================================ */

func checkUser(lLoginResponse *LoginResponse, lUser *Login, lUserPassword string, lRole string, r *http.Request) {

	log.Println("checkUser-(+)")

	if lRole == "A" {

		checkAdmin(lLoginResponse, lUser, lRole)

	} else {

		lGormDB, lErr := gormdb.GormDBConnection()

		lSql, _ := lGormDB.DB()

		defer lSql.Close()

		if lErr != nil {

			log.Println("LCU-001", lErr.Error())
			lLoginResponse.ErrMsg = lErr.Error()
			lLoginResponse.Status = common.ErrorCode

		} else {

			fmt.Println(lLoginResponse)
			fmt.Println(lUserPassword)

			var lTableName string

			var lResult *gorm.DB

			if lRole == "C" {

				lTableName = "st_918_client_table"

				lResult = lGormDB.Table(lTableName).Select("password").Where("unique_id=?", lUser.UserId).Find(&lUser.Password)

			} else if lRole == "U" {

				/*  ==============USER-ROLES=============
					1. BO  -  Back Officer
					2. B   - Biller
					3. APR - Approver
				    ===================================== */

				lUserRole := r.Header.Get("USERROLE")
				lUserRole = strings.ToUpper(lUserRole)

				if isValidUserRole(lUserRole) {

					lLoginResponse.ErrMsg = "Provide a UserRole in header - B|BO|APR"
					lLoginResponse.Status = common.ErrorCode
					return

				} else {

					lTableName = "st_918_config_users_table"

					lResult = lGormDB.Table(lTableName).Select("password").Where("user_id=?", lUser.UserId).Where("role=?", lUserRole).Find(&lUser.Password)

				}

			}

			if lResult != nil && lResult.Error != nil || lResult.RowsAffected == 0 {
				log.Println("LCU-002", lResult.Error)
				lLoginResponse.Status = common.ErrorCode
				lLoginResponse.ErrMsg = "Invalid User Id"

			} else {

				if lUser.Password != lUserPassword {

					lLoginResponse.Status = common.ErrorCode
					lLoginResponse.ErrMsg = "Password is Incorrect"

				} else {

					lLoginResponse.Status = common.SuccessCode
					lLoginResponse.ErrMsg = ""
					fmt.Println("Login Successfully")
				}
			}
		}

	}

	log.Println("checkUser-(-)")

}

/* =============   IF ROLE IS USER -- then SUB-ROLES are B, BO, APR   =========================== */

func isValidUserRole(lUserRole string) bool {

	log.Println("isValidUserRole-(-)")

	if lUserRole != "BO" && lUserRole != "B" && lUserRole != "APR" {
		return true
	}
	log.Println("isValidUserRole-(-)")
	return false
}

/* ======================= FOR check Admin id & Password in toml file =============================== */

func checkAdmin(lLoginResponse *LoginResponse, lUser *Login, lRole string) {

	log.Println("checkAdmin-(+)")
	config := genpkg.ReadTomlConfig("./toml/dbconfig.toml")
	AdminId := fmt.Sprintf("%v", config.(map[string]interface{})["AdminId"])
	AdminPassword := fmt.Sprintf("%v", config.(map[string]interface{})["AdminPassword"])

	if lUser.UserId != AdminId {

		lLoginResponse.Status = common.ErrorCode
		lLoginResponse.ErrMsg = "Admin Not Found"

	} else if lUser.Password != AdminPassword {
		lLoginResponse.Status = common.ErrorCode
		lLoginResponse.ErrMsg = "Password Incorrect"
	} else {
		lLoginResponse.Status = common.SuccessCode
	}

	log.Println("checkAdmin-(-)")
}
