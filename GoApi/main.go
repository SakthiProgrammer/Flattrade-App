package main

import (
	"flattrade/apps/bank"
	login "flattrade/apps/login"
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
	http.HandleFunc("/registeruser", login.RegisterClient)
	http.HandleFunc("/getbanks", bank.GetBanks)

	http.ListenAndServe(":29091", nil)
}
