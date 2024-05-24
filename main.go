package main

import (
	"fmt"
	"heroku/config"
	"heroku/handler"
)

func main() {
	router, server := config.InitServer()

	db := config.Connect()
	defer db.Close()

	DB := handler.HandlerDB{DB: db}
	router.GET("/reports", DB.GetReport)
	router.GET("/report/:id", DB.GetReportById)
	router.POST("/report", DB.CreateReport)
	router.PUT("/report/:id", DB.UpdateReport)
	router.DELETE("/report/:id", DB.DeleteReport)

	fmt.Println("localhost running in port 8080")
	panic(server.ListenAndServe())
}
