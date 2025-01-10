package stock

import (
	"encoding/json"
	"flattrade/common"
	"fmt"
	"io"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Stock, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")

	log.Println("UpdateStock- (+)")

	var lStock Stock
	var lStockResp StockResp

	lStockResp.Status = common.SuccessCode
	if r.Method == http.MethodPost {

		lData, lErr := io.ReadAll(r.Body)

		if lErr != nil {
			log.Println("SUS-001", lErr.Error())
			lStockResp.ErrMsg = lErr.Error()
			lStockResp.Status = common.ErrorCode

		} else {
			lErr = json.Unmarshal(lData, &lStock)

			if lErr != nil {

				log.Println("SUS-002", lErr.Error())
				lStockResp.ErrMsg = lErr.Error()
				lStockResp.Status = common.ErrorCode
			} else {

				if lErr != nil {
					log.Println("SUS-003", lErr.Error())
					lStockResp.ErrMsg = lErr.Error()
					lStockResp.Status = common.ErrorCode

				} else {

					lDsn := "ST918:Best@123@tcp(192.168.2.5)/training"
					lGormDB, lErr := gorm.Open(mysql.Open(lDsn), &gorm.Config{})

					if lErr != nil {
						log.Println("SUS-004", lErr.Error())
						lStockResp.ErrMsg = lErr.Error()
						lStockResp.Status = common.ErrorCode

					} else {

						lResult := lGormDB.Table("st_918_Stock_Table").Where("stock_id=?", lStock.ID).UpdateColumns(&lStock)

						if lResult.Error != nil {
							log.Println("SUS-005", lResult.Error)
							lStockResp.ErrMsg = lErr.Error()
							lStockResp.Status = common.ErrorCode

						} else {

							fmt.Println(lResult.RowsAffected)
							lStockResp.StockDetailsArr = append(lStockResp.StockDetailsArr, lStock)
						}
					}

				}

			}

		}

	} else {
		// w.Header().Set(http.S)
		lStockResp.ErrMsg = "Method Not Allowed!"
		lStockResp.Status = common.ErrorCode
	}

	lData, lErr := json.Marshal(lStockResp)

	if lErr != nil {
		log.Println("SUS-006", lErr.Error())
		lStockResp.ErrMsg = lErr.Error()
		lStockResp.Status = common.ErrorCode

	} else {
		fmt.Fprint(w, string(lData))
	}

	log.Println("UpdateStock- (-)")
}
