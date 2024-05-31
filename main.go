package main

import (
	"fmt"
	"log"
	"net/http"

	"fizzbuzz/handler"
	logger "fizzbuzz/log"
	"fizzbuzz/usecase"
	"os"
	"os/signal"
	"syscall"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("== Start HTTP Server ====")
	router := httprouter.New()

	fizzbuzzRoutes(router)
	logger.GenerateLog()
	logger.CommonLog.Println("Server is Running")
	server := http.Server{Addr: ":8080", Handler: router}
	c := make(chan os.Signal, 1)
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	logger.CommonLog.Print("Shuting down on progress...")
	fmt.Println("Start Shutting Down... Please wait for 10 seconds")
	usecase.StopGraceful()

	fmt.Println("Finished Shutting Down...")
	logger.CommonLog.Print("fully shuting down")

}

func fizzbuzzRoutes(router *httprouter.Router) {
	router.GET("/range-fizzbuzz", handler.FizzBuzz)

}
