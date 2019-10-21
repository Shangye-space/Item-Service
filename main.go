package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/gorilla/mux"
	"github.com/johan-lejdung/go-microservice-api-guide/rest-api/app"
)

type Items struct {
	ItemID      int
	ItemName    string
	Quantity    int
	Description string
	Price       float32
	Discount    int
	InSale      bool
	Category    string
	SubCategory *string
	AddedTime   time.Time
	RemovedTime *time.Time
}

func main() {
	db, er := database.CreateDatabase()
	if er != nil {
		log.Fatal("Database connection failed: %s", er.Error())
	} else {
		rows, err := db.Query("SELECT * FROM Items;")

		if err != nil {
			panic(err)
		}

		defer rows.Close()
		for rows.Next() {
			items := Items{}
			err = rows.Scan(&items.ItemID, &items.ItemName, &items.Quantity, &items.Description, &items.Price, &items.Discount, &items.InSale, &items.Category, &items.SubCategory, &items.AddedTime, &items.RemovedTime)
			if err != nil {
				panic(err)
			}
			fmt.Println(items)
		}
		err = rows.Err()
		if err != nil {
			panic(err)
		}
	}

	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: db,
	}

	app.SetupRouter()

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
