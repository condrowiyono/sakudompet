package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"

	"github.com/julienschmidt/httprouter"
	"github.com/condrowiyono/sakudompet/database"
	"github.com/condrowiyono/sakudompet/handler"
	"github.com/condrowiyono/sakudompet/middleware"
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
	
	//DDEBITS
	router.GET("/debits", middleware.Gateway("",sdHandler.BasicAuth(sdHandler.FindAllDebits)))
	router.POST("/debits",  middleware.Gateway("",sdHandler.BasicAuth(sdHandler.CreateDebit)))
	router.GET("/debits/:id",  middleware.Gateway("",sdHandler.BasicAuth(sdHandler.FindDebit)))
	router.PUT("/debits/:id",  middleware.Gateway("",sdHandler.BasicAuth(sdHandler.PutDebit))) //Entire Resource
	router.DELETE("/debits/:id",  middleware.Gateway("",sdHandler.BasicAuth(sdHandler.DeleteDebit))) //Delete

	//PASSES
	router.GET("/passes", middleware.Gateway("",sdHandler.BasicAuth(sdHandler.GetPasses)))
	router.GET("/passes/:id", middleware.Gateway("",sdHandler.BasicAuth(sdHandler.FindPass)))
	router.DELETE("/passes/:id",  middleware.Gateway("",sdHandler.BasicAuth(sdHandler.DeletePass))) 
	
	// router.GET("/create-pass",  sdHandler.BasicAuth(sdHandler.CreatePass))


	http.ListenAndServe(":3000", router)
	
}