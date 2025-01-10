package stock

import (
	"encoding/json"
	"flattrade/apps/DBConnection/gormdb"
	"flattrade/common"
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
	type Stock struct {
		ID         int     `json:"id" gorm:"column:Id"`
		StockName  string  `json:"stock_name" gorm:"column:Stock_Name"`
		StockPrice float64 `json:"stock_price" gorm:"column:Stock_Price"`
		Segment    string  `json:"segment" gorm:"column:Segment"`
		ISIN       string  `json:"isin" gorm:"column:ISIN"`
		CreatedBy  string    `json:"created_by" gorm:"column:Created_By"`
		CreatedAt  time.Time `json:"created_at" gorm:"column:Created_At"`
		UpdatedBy  string    `json:"updated_by" gorm:"column:Updated_By"`
		UpdatedAt  time.Time `json:"updated_at" gorm:"column:Updated_At"`
	}

	type StockResp struct {
		StockDetailsArr []Stock `json:"stock_details"`
		ErrMsg          string  `json:"errMsg"`
		Status          string  `json:"status"`
	}
*/
func CreateStock(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "PRODUCTS, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")

	log.Println("CreateStock - (+)")
	var lStock Stock
	var lStockResp StockResp
	lStockResp.Status = common.SuccessCode

	if r.Method == http.MethodPost {

		lData, lErr := io.ReadAll(r.Body)

		if lErr != nil {
			log.Println("SCS-001", lErr.Error())
			lStockResp.ErrMsg = lErr.Error()
			lStockResp.Status = common.ErrorCode
		} else {
			lErr = json.Unmarshal(lData, &lStock)

			if lErr != nil {
				log.Println("SCS-002", lErr.Error())
				lStockResp.ErrMsg = lErr.Error()
				lStockResp.Status = common.ErrorCode
			} else {

				if lErr != nil {
					log.Println("SCS-003", lErr.Error())
					lStockResp.ErrMsg = lErr.Error()
					lStockResp.Status = common.ErrorCode

				} else {
					createStocksInDB(&lStockResp, &lStock)
				}

			}

		}

	} else {
		lStockResp.ErrMsg = "Invalid Method"
		lStockResp.Status = common.ErrorCode
	}

	lData, lErr := json.Marshal(lStockResp)

	if lErr != nil {
		log.Println("SCS-006", lErr.Error())
		lStockResp.ErrMsg = lErr.Error()
		lStockResp.Status = common.ErrorCode

	} else {
		fmt.Fprint(w, string(lData))
	}

	log.Println("CreateStock - (-)")

}

func createStocksInDB(lStockResp *StockResp, lStock *Stock) {

	log.Println("CreateStocksInDB-(+)")

	lGormDB, lErr := gormdb.GormDBConnection()

	lSql, _ := lGormDB.DB()

	defer lSql.Close()

	if lErr != nil {
		log.Println("SCS-004", lErr.Error())
		lStockResp.ErrMsg = lErr.Error()
		lStockResp.Status = common.ErrorCode
	} else {

		lResult := lGormDB.Table("st_918_Stock_Table").Create(&lStock)

		if lResult.Error != nil {
			log.Println("SCS-005", lResult.Error)
			lStockResp.ErrMsg = lErr.Error()
			lStockResp.Status = common.ErrorCode

		} else {
			fmt.Println(lResult.RowsAffected)
			lStockResp.StockDetailsArr = append(lStockResp.StockDetailsArr, *lStock)
		}
	}

	log.Println("CreateStocksInDB-(-)")
}
