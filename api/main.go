package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"

	"github.com/julienschmidt/httprouter"
	"github.com/condrowiyono/sakudompet/database"
	"github.com/condrowiyono/sakudompet/handler"
	sd "github.com/condrowiyono/sakudompet"
)


// our main function
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbOpt := database.Option{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		Database: os.Getenv("MYSQL_DATABASE"),
		Charset:  os.Getenv("MYSQL_CHARSET"),
	}
	mysql, err := database.NewMySQL(dbOpt)
	
	
	if err != nil {
		log.Fatal("Error loading mysql")
	}

	sakudompet := sd.NewSakuDompet(mysql)
	sdHandler := handler.NewHandler(sakudompet)


	router := httprouter.New()

	router.GET("/healthz", sdHandler.Healthz)
	router.GET("/debits",  sdHandler.FindAllDebits)

	http.ListenAndServe(":3000", router)
	
}

