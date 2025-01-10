package user

/* type Stock struct {
	ID         int       `json:"id" gorm:"column:Id"`
	StockName  string    `json:"stock_name" gorm:"column:Stock_Name"`
	StockPrice float64   `json:"stock_price" gorm:"column:Stock_Price"`
	Segment    string    `json:"segment" gorm:"column:Segment"`
	ISIN       string    `json:"isin" gorm:"column:ISIN"`
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

func GetStocks(w http.ResponseWriter, r *http.Request) {

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "ADMIN, CLIENT, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(w).Header().Set("Content-Type", "application/json")

	log.Println("GetStocks-(+)")

	var lStockResp StockResp

	// var lStock Stock

	if r.Method == http.MethodGet {
		lStockResp.Status = common.SuccessCode

		lGormDb, lErr := gormdb.GormDBConnection()

		lSql, _ := lGormDb.DB()

		defer lSql.Close()

		if lErr != nil {
			log.Println("SGS-001", lErr.Error())
			lStockResp.Status = common.ErrorCode
			lStockResp.ErrMsg = lErr.Error()

		} else {

			lResult := lGormDb.Table("st_918_Stock_Table").Find(&lStockResp.StockDetailsArr)

			if lResult.Error != nil {
				log.Println("SGS-002", lResult.Error)
				lStockResp.ErrMsg = lResult.Error.Error()
				lStockResp.Status = common.ErrorCode
			} else {
				log.Println("Successfully Get")
				lStockResp.Status = common.SuccessCode
				fmt.Printf("%+v", lStockResp)
			}
		}

	} else {
		log.Println("SGS-", "Invalid Method")
		lStockResp.ErrMsg = "Invalid Method"
		lStockResp.Status = common.ErrorCode
	}

	lData, lErr := json.Marshal(lStockResp)

	if lErr != nil {
		log.Println("SGS-", lErr.Error())
		lStockResp.ErrMsg = lErr.Error()
		lStockResp.Status = common.SuccessCode
	} else {
		fmt.Fprintf(w, string(lData))
	}

	log.Println("GetStocks-(-)")

}
*/
