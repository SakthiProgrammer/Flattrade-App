package main

import (
	"flattrade/apps/bank"
	"flattrade/apps/brokerage"
	"flattrade/apps/client"
	login "flattrade/apps/login"
	"flattrade/apps/stock"
	"flattrade/apps/user"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	log.Println("Server Started")
	f, err := os.OpenFile("./log/logfile"+time.Now().Format("02012006.15.04.05.000000000")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer f.Close()

	log.SetOutput(f)

	http.HandleFunc("/login", login.LoginAll)
	http.HandleFunc("/registerclient", login.RegisterClient)
	http.HandleFunc("/getbanks", bank.GetBanks)
	http.HandleFunc("/createbank", bank.CreateBank)
	http.HandleFunc("/updatebank", bank.UpdateBank)
	http.HandleFunc("/createcharge", brokerage.CreateChager)
	http.HandleFunc("/getcharges", brokerage.GetCharges)
	http.HandleFunc("/updatecharge", brokerage.UpdateChage)
	http.HandleFunc("/getclients", client.GetClients)
	http.HandleFunc("/updateclient", client.UpdateClient)
	http.HandleFunc("/createstock", stock.CreateStock)
	http.HandleFunc("/getstock", stock.GetStocks)
	http.HandleFunc("/updatestock", stock.UpdateStock)
	http.HandleFunc("/createuser", user.CreateUser)
	http.HandleFunc("/getuser", user.GetUsers)
	http.HandleFunc("/updateuser", user.UpdateUser)

	http.ListenAndServe(":29091", nil)
}
