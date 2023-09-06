package config

import (
	"log"

	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
)

var (
	retryConnection int = 0
)

const (
	HOST     string = "simasbanking.oracle.database"
	PORT     int    = 1521
	SERVICE  string = "ORCL"
	USERNAME string = "SIMASBANKINGPROD"
	PASSWORD string = "simasbanking"
)

func SetupDb() (*gorm.DB, error) {
	log.Printf("connecting to db with username || password => %s||%s host => %s:%d/%s ", USERNAME, PASSWORD, HOST, PORT, SERVICE)
	url := oracle.BuildUrl(HOST, PORT, SERVICE, USERNAME, PASSWORD, nil)
	log.Println("result build url : ", url)
	db, err := gorm.Open(oracle.Open(url), &gorm.Config{})
	if err != nil {
		if retryConnection < 3 {
			retryConnection++
			log.Println("retry connection at : ", retryConnection)
			return SetupDb()
		}
		log.Println("max attempt to connect ", retryConnection, " stopping...")
		log.Fatal("error : ", err)
		return nil, err
		// panic(err)
		// panic error or log error info
	}
	log.Println("successfully connected to db.")
	return db, nil
}
