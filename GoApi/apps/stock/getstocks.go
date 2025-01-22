package trade

import (
	"encoding/json"
	"log"
	"net/http"
)

type GetTradeRec struct {
	TradeID             uint    `json:"trade_id"`
	TradeType           string  `json:"trade_type"`
	TotalPrice          float64 `json:"total_price"`
	TradeDate           string  `json:"trade_date"`
	BackOfficerApproval string  `json:"back_officer_approval"`
	BillerApproval      string  `json:"biller_approval"`
	ApproverApproval    string  `json:"approver_approval"`
}

type GetBankRec struct {
	BankId   uint   `json:"bank_id"`
	BankName string `json:"bank_name"`
	IFSCCode string `json:"ifsc_code"`
	Address  string `json:"address"`
}

type GetClientRec struct {
	ClientID       uint          `json:"client_id"`
	FirstName      string        `json:"first_name"`
	LastName       string        `json:"last_name"`
	Email          string        `json:"email"`
	PhoneNumber    string        `json:"phone_number"`
	PanNumber      string        `json:"pan_number"`
	NomineeName    string        `json:"nominee_name"`
	UniqueId       string        `json:"unique_id"`
	BankAccount    string        `json:"bank_account"`
	KycIsCompleted string        `json:"kyc_is_completed"`
	TradesArr      []GetTradeRec `json:"trade_details"`
	BankRec        GetBankRec    `json:"bank_detail"`
}

func GetClientFullDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "USERID, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")

	log.Println("GetClientFullDetails-(+)")

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	clientID := 3 // Replace with dynamic value if required
	sqldb
	if err != nil {
		log.Println("Error connecting to database:", err)
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var clientRec GetClientRec

	// Query for client details
	clientQuery := `
		SELECT client_id, first_name, last_name, email, phone_number, pan_number,
		       nominee_name, unique_id, bank_account, kyc_is_completed
		FROM client_table
		WHERE client_id = ?
	`

	err = db.QueryRow(clientQuery, clientID).Scan(
		&clientRec.ClientID, &clientRec.FirstName, &clientRec.LastName, &clientRec.Email,
		&clientRec.PhoneNumber, &clientRec.PanNumber, &clientRec.NomineeName, &clientRec.UniqueId,
		&clientRec.BankAccount, &clientRec.KycIsCompleted,
	)
	if err != nil {
		log.Println("Error fetching client details:", err)
		http.Error(w, "Error fetching client details", http.StatusInternalServerError)
		return
	}

	// Query for bank details
	bankQuery := `
		SELECT bank_id, bank_name, ifsc_code, address
		FROM bank_table
		WHERE bank_id = (
			SELECT bank_id FROM client_table WHERE client_id = ?
		)
	`

	err = db.QueryRow(bankQuery, clientID).Scan(
		&clientRec.BankRec.BankId, &clientRec.BankRec.BankName,
		&clientRec.BankRec.IFSCCode, &clientRec.BankRec.Address,
	)
	if err != nil {
		log.Println("Error fetching bank details:", err)
		http.Error(w, "Error fetching bank details", http.StatusInternalServerError)
		return
	}

	// Query for trade details
	tradeQuery := `
		SELECT trade_id, trade_type, total_price, trade_date,
		       back_officer_approval, biller_approval, approver_approval
		FROM trade_table
		WHERE client_id = ?
	`

	rows, err := db.Query(tradeQuery, clientID)
	if err != nil {
		log.Println("Error fetching trade details:", err)
		http.Error(w, "Error fetching trade details", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var trade GetTradeRec
		err := rows.Scan(
			&trade.TradeID, &trade.TradeType, &trade.TotalPrice, &trade.TradeDate,
			&trade.BackOfficerApproval, &trade.BillerApproval, &trade.ApproverApproval,
		)
		if err != nil {
			log.Println("Error scanning trade record:", err)
			http.Error(w, "Error processing trade records", http.StatusInternalServerError)
			return
		}
		clientRec.TradesArr = append(clientRec.TradesArr, trade)
	}

	// Send the response
	response, err := json.Marshal(clientRec)
	if err != nil {
		log.Println("Error marshalling response:", err)
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}

	w.Write(response)
	log.Println("GetClientFullDetails-(-)")
}
