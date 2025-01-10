package client

/*
import "net/http"

type Client struct{

}

func GetClients(w http.ResponseWriter, r *http.Request){

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "ADMIN, CLIENT, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(w).Header().Set("Content-Type", "application/json")

	log.Println("GetClients-(+)")

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

	log.Println("GetClients-(-)")

} */
