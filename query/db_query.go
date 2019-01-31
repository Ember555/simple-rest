package query

import (
	"database/sql"
	"log"

	"simple-rest/models"
)

func initDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(mariadb:3306)/wallet_localnetwork")
	log.Println("Connect database")
	if err != nil {
		panic(err)
	}
	return db
}

func QueryAll() *[]models.DBdata {
	db := initDB()
	defer db.Close()

	arrDdata := make([]models.DBdata, 0)
	results, err := db.Query("SELECT participant FROM network_users;")
	log.Println("Searching")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var dbData models.DBdata
		err = results.Scan(&dbData.User)
		if err != nil {
			panic(err.Error())
		}
		arrDdata = append(arrDdata, models.DBdata{User: dbData.User})
	}
	log.Println("SearchParticipant Success")
	return &arrDdata
}
